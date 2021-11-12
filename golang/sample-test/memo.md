t.Helper() を使えば失敗した行が関数呼び出し側になる.

Benchmark 実行には `go test -bench=.` を実行する.

カバレッジを確認するには `go test -cover` を実行する.

スライスの test では `reflect.DeepEqual` を使って変数を比較する.
ただ `reflect.DeepEqual` は型安全ではないので注意.

マップは nil に書き込もうとするとランタイムパニックになるので
```go
var m map[string]string
```
で初期化するのではなく,下記のように初期化する.
```go
var dictonary = map[string]string{}
var dictonary = make(map[string]string)
```

標準ライブラリに `net/http/httptest` があってこれで模擬 HTTP サーバを作れる


'入力 X のとき出力 Y を期待する' というテストを作るときはテーブルベースのテストを使う.
