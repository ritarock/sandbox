## context パッケージ
context パッケージを見てみる.
```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

context パッケージのざっくりとした目的は
- 適切なキャンセルを行う機能を提供する
- リクエストのデータの置き場を提供する

キャンセルには 3 つ側面がある.
- ゴルーチンの親がキャンセルしたい場合
- ゴルーチンの子をキャンセルしたい場合
- ゴルーチン内のブロックしている処理がキャンセルされるように中断できる必要がある場合

## 使い方
Context の空インスタンスを作る関数は以下の 2 つ
- `func Background() Context`
  - 通常使うのはこっち.空の Context を返す
- `func TODO() Context`
  - 本番環境で使われることを想定していない.どの Context を使っていいかわからないとき,もしくは上流の実装が終わっていないときに使う

### キャンセル処理
- `context.WithCancel` を使う
```go
func context.WithCancel(parent context.Context) (ctx context.Context, cancel context.CancelFunc)
```
キャンセルされる側は `ctx.Done()` からキャンセルを受け取る.
```go
<-ctx.Done()
```

キャンセルする側は `context.WithCancel()` によって生成された cancal 関数を実行することで, キャンセルされる側の context の Done メソッドが close される.
```go
ctx, cancel := context.WithCancel(context.Background())
// 処理
cancel()
```

- `context.WithDeadline`
```go
func context.WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc)
```
キャンセルする側は `context.WithDeadline` の生成時に停止したい時刻を設定することでその時刻を超えたタイミングでキャンセルが実行される.
```go
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
```
- `context.Timeout`
```go
func context.WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
```

キャンセルする側は `context.Timeout` の生成時に停止したい時間は設定することでその時間を超えたタイミングでキャンセルが実行される.
```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
```

`context.WithDeadline` , `context.Timeout` を使ったキャンセルされる側は `context.WithCancel` 同様, `ctx.Done()` からキャンセルを受け取る.
context生成時に得られる canacl は close されたチャネルに対しては何も実行されないので,タイムアウトの処理をしていても明示的に cancel は呼ぶほうが良い.

context に対してタイムアウトが設定されているかどうかを確認するには, context の `Deadline` メソッドを実行する.
```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
}
```
設定されている場合,第 2 返り値は true で 第 1 返り値にはその時刻が設定されている.

### Err メソッド
```go
type Context interface {
	Err() error
}
```
- context がキャンセルされていない場合は `nil`
- context が明示的にキャンセルされている場合は `Canceled`
- context がタイムアウトしていた場合は `DeadlineExceeded`
```go
LOOP:
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); errors.Is(err, context.Canceled) {
				fmt.Println("Canceled")
			} else if errors.Is(err, context.DeadlineExceeded) {
				fmt.Println("DeadlineExceeded")
			}
			break LOOP
		}
	}
```

### Value メソッド
```go
func context.WithValue(parent context.Context, key interface{}, val interface{}) context.Context
```
`WithValue` を使うと context に key-value 形式でデータを保持できる.
```go
ctx, cancel := context.WithCancel(context.Background())
ctx = context.WithValue(ctx, "id", 1)
ctx = context.WithValue(ctx, "user", "abc")
```

取り出すときはアサーションして値を取り出す.
```go
id, user := ctx.Value("id").(int), ctx.Value("user").(string)
```
