# 関数
## 関数の宣言
```typescript
// 名前付き関数
function greet(name: string) {
  return 'hello' + name
}

// 関数式
let greet2 = function(name: string) {
  return 'hello' + name
}

// アロー関数
let greet3 = (name: string) => {
  return 'hello' + name
}

// アロー関数省略記法
let greet4 = (name: string) => 'hello' + name

// 関数コンストラクタ
let greet5 = new Function('name', 'return "hello" + name')
```

## オプションパラメータ
```typescript
function log(message: string, userId?: string) {
  let time = new Date().toLocaleDateString()
  console.log(time, message, userId || 'Not signed in')
}

log('Page loaded') // 12:36:12 Page loaded
log('User signed in', 'da763be') // 12:36:12 User signed in da763be
```

## デフォルトパラメータ
デフォルトパラメータを与えるとオプションのアノテーションがなくなる
```typescript
function log(message: string, userId = 'Not signed in') {
  let time = new Date().toISOString()
  console.log(tmie, message, userId)
}
log('Page loaded') // 12:36:12 Page loaded
log('User signed in', 'da763be') // 12:36:12 User signed in da763be
```
デフォルトパラメータに明示的な型アノテーションを加えることも可能
```typescript
type Context = {
  appId?: string
  userId?: string
}

function log(message: string, context: Context = {}) {
  let time = new Date().toISOString()
  console.log(time, message, context.userId)
}
```

## レストパラメータ
```typescript
function sum(number: number[]): number {
  return number.reduce((total, n) => total + n, 0)
}

sum([1, 2, 3]) //6
```

## 呼び出しシグネチャ
```typescript
function add(a: number, b: number): number {
  return a + b
}
```
この関数について add は2つの number を取り、 number を返す関数でありその型を次のように表現する
```typescript
(a: number, b: number) => number
```
これを型シグネチャと呼ぶ

### 例
```typescript
type Log = (message: string, userId?: string) => void

let log: Log = (
  message,
  userId = 'Not signed in',
) => {
    let time = new Date().toISOString()
    console.log(time, message, userId)
}
```

```typescript
// 呼び出しシグネチャの省略記法
type Log = (message: string, userId?: string) => void

// 完全な呼び出しシグネチャ
type Log = {
  (message: string, userId?: string): void
}
```

## ジェネリック型
```typescript
type Filter = {
  <T>(array: T[], f: (item: T) => boolean): T[]
}
```
この関数はジェネリック型パラメータ T を使う
この型は何になるかは事前にはわからないが、呼び出すたびに TypeScript が推論を行う
