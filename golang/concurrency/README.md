## 並行処理とは

### 並行処理の気をつけるポイント
#### 競合状態
データの処理に順番がある可能性がある.
処理の間に 1 時間かかっても影響がないか? と考えることで気づくことができる.

#### アトミック性
すべての処理が終了した状態か,そうでないかの 2 パターンしかない場合.
コンテキストで判断する.
`i++` はコンテキスト.中身の処理はそれぞれアトミックだが, `i++` は非アトミック.
- i の値を取得する
- i の値に 1 増やす
- i の値を保存する
アトミックな処理であれば複数のゴルーチンで安全に扱える.

#### メモリアクセス同期
排他処理が必要.クリティカルセクションと呼ばれる.
`sync.Mutex` を使って解決できるが,`Lock` のたびにプログラムが一時停止するのでパフォーマンスに影響がある.

#### デッドロック
複数の並行なプロセスが互いに処置を待ち続けていて処理が終了しない.
デッドロックは以下の条件 ( Coffman 条件 ) で起こる
- 相互排他
  - ある並行プロセスがリソースに対して排他的な権利をどの時点においても保持している
- 条件待ち
  - ある並行プロセスはリソースの保持と追加のリソース待ちを同時に行わなければならない
- 横取り不可
  - ある並行プロセスによって保持されているリソースは,そのプロセスによってのみ開放される
- 循環待ち
  - ある並行プロセス ( P1 ) は,他の連なっている並行プロセス ( P2 ) を待たなければならない.そして P2 は P1 を持っている

#### ライブロック
ライブロックが起こる原因の多くは,試行回数に上限がない.
2つの並行プロセスが互いのデッドロックを予防しようとすることでライブロックが起こる.

#### リソース枯渇
ライブロックのように,1 つ以上の貪欲な並行プロセスが他のプロセスのリソースを奪っている状態で起きる.
バランスが重要.
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops\n", count)
	}

	wg.Add(2)
	go greedWorker()
	go politeWorker()
	wg.Wait()
}
```

### sync パッケージ
使う場面の多くは, `struct` のような小さなスコープ.

#### sync.WaitGroup
ひとまとまりの並行処理があったとき処理の完了待ちに使う.
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("finish")
}
```
`wg.Add` で渡された引数の数値分カウンターを増やし, `wg.Done` でカウンターを 1 つ減らす.
`wg.Wait` はカウンターが 0 になるまで処理をブロックする.

#### Mutex
Mutual exclusion ( 相互排他 ) を表す.
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()

		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()

		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithemtic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithemtic.Add(1)
		go func() {
			defer arithemtic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithemtic.Add(1)
		go func() {
			defer arithemtic.Done()
			decrement()
		}()
	}

	arithemtic.Wait()
	fmt.Println("Finish")
}
```
`Unlock` の呼び出しは `defer` に入れたほうがいい.こうしておくと, panic になったとしても確実に呼び出せる.

### channel
ゴルーチン間の通信に使う.
チャネルの状態に対する操作結果
操作|状態|結果
-|-|-
read|nil|ブロック
read|open で空でない|値を取得
read|open で空|ブロック
read|Close|デフォルト値, false
read|書き込み専用|コンパイルエラー
write|nil|ブロック
write|open で満杯でない|値を書き込む
write|open で満杯|ブロック
write|Close|panic
write|読み込み専用|コンパイルエラー
close|nil|panic
close|open で空でない|チャネルを閉じる.読み込みはチャネルの中身がなくなるまで成功する.その後はデフォルト値を読み込む
close|open で空|チャネルを閉じる.デフォルト値を読み込む
close|closed|panic
close|読み込み専用|コンパイルエラー

チャネルを使うことで気をつけること.
- チャネルを初期化する
- 書き込みを行うか,他のゴルーチンに所有権を渡す
- チャネルを閉じる
- 上記手順をカプセル化して読み込み専用のチャネルを経由して公開する

この 4 つを行うことで下記が期待できる.
- チャネルを初期化するゴルーチンなので nil チャネルに書き込んでデッドロックを起こす危険がなくなる
- チャネルを初期化するゴルーチンなので nil チャネルを閉じることによって起こる panic の危険がなくなる
- チャネルを閉じるタイミングを決めるゴルーチンなので,閉じたチャネルに書き込んで panic になる危険がなくなる
- チャネルを閉じるタイミングを決めるゴルーチンなので,チャネルを 2 度以上閉じてしまうことによって起こる panic の危険がなくなる
- コンパイル時に型チェックを行って,チャネルに対する不適切な書き込みを防ぐ

ブロック操作の注意点
- チャネルがいつ閉じられたか把握する
- いかなる理由でもブロックする操作は責任を持って扱う

### select 文
select 文はチャネルをまとめる.
チャネルでゴルーチンをまとめて,select 文でチャネルをまとめる.

### 並行処理パターン
#### 拘束
並行処理を安全員行うパターン
- メモリを共有するための同期のプリミティブ ( sync.Mutex )
- 通信による同期 (channel)
- イミュータブルなデータ
- 拘束によって保護されたデータ

#### for-select ループ
#### ゴルーチンリークを避ける
ゴルーチンが終了するときは以下の 3 通り
- ゴルーチンが処理を完了するとき
- 回復できないエラーにより処理が続けられないとき
- 停止するように命令されたとき

#### パイプライン
何らかのデータを受け取り,何らかの処理を行い,どこかに渡すという一連の作業を行う.

下記はパイプライン処理の例.
```go
package main

import "fmt"

func main() {
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}

		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}
}
```

パイプラインをストリーム指向の処理に書き直したのが以下.
```go
package main

import "fmt"

func main() {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}
	add := func(value, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(add(multiply(v, 2), 1))
	}
}
```

もっとよく書く
```go
package main

import "fmt"

func main() {
	generator := func(done <-chan int, integers ...int) <-chan int {
		intStream := make(chan int, len(integers))
		go func() {
			defer close(intStream)
			for _, v := range integers {
				select {
				case <-done:
					return
				case intStream <- v:
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan int,
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for v := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- v * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(
		done <-chan int,
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for v := range intStream {
				select {
				case <-done:
					return
				case addedStream <- v + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan int)
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := add(done, multiply(done, intStream, 2), 1)

	for v := range pipeline {
		fmt.Println(v)
	}
}
```

#### ファンアウト / ファンインイン
ステージが以下の条件を満たせば利用できる.
- ステージが前の計算結果に依存しない
- 実行時間が長時間に及ぶ

