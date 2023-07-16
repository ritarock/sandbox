# クラス
## クラスと継承
```typescript
// チェスのゲーム
class Game {}

// チェスの駒
class Piece {}

// 駒の位置(座標)
class Position {}

// チェスは6種類の駒がある
class King extends Piece {}
class Queen extends Piece {}
class Bishop extends Piece {}
class Knight extends Piece {}
class Rook extends Piece {}
class Pawn extends Piece {}

// 駒を表す Piece クラスに色と位置を追加する
type Color = 'Black' | 'White'
type File = 'A' | 'B' | 'C'| 'D' | 'E' | 'F' | 'G' | 'H'
type Rank = 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8

class Postion {
  constructor(
    private file: File,
    private rank: Rank
  ) {}
}

class Piece {
  protected position: Position
  constructor(
    private readonly color: Color,
    file: File,
    rank: Rank
  ) {
      this.postion = new Position(file, rank)
  }
}
```
### public
どこからでもアクセス可能
### protected
このクラスとサブクラスのインスタンスからのみアクセス可能
### private
このクラスのインスタンスからのみアクセス可能

ユーザには新しい Piece をインスタンス化してほしくない
Piece を拡張して Queen や Bishop などを作成しそれをインスタンス化してほしい
abstract を使うとそれを強制することができる
```typescript
abstract class Piece {
  constructor(
//
  )
}
```
abstract が指定されたクラスは抽象クラスと呼ばれる
抽象クラスで Piece を直接インスタンス化するとエラーとなる
```typescript
new Piece('White', 'E', 1) // error
```

# インターフェース
型エイリアスと同様にインターフェースは型に名前をつけるための方法
これを使うことでインラインで型を定義する必要がなくなる

型エイリアスをインターフェースに置き換えるのは簡単
```typescript
// 型エイリアス
type Sushi = {
  calories: number
  salty : boolean
  tasty: boolean
}
```
```typescript
interface Sushi {
  calories: number
  salty : boolean
  tasty: boolean
}
```

次に別の食べ物をモデル化
```typescript
type Cake = {
  calories: number
  sweet: boolean
  tasty: boolean
}
```

食べ物を独自の型 Food として抜き出し、それに基づいてこれまで出てきた食べ物を再定義する
```typescript
type Food = {
  calories: number
  tasty: boolean
}
type Sushi = Food & {
  salty: boolean
}
type Cake = Food & {
  sweet: boolean
}
```

今度はインターフェースで実現する
```typescript
interface Food {
  calories: number
  tasty: boolean
}
interface Sushi extends Food {
  salty: boolean
}
interface Cake extends Food {
  sweet: boolean
}
```

## 型エイリアスとインターフェースの違いは?
違いは3点ある
1つ目の違いは、型エイリアスの方が右辺に任意の型を指定できるという点で汎用的である(右辺に型の式 (&, |) を指定できる)
インターフェースの場合、右辺は形状でなければならない
例えば次のような型はインターフェースで書き直すことはできない
```typescript
type A = number
type B = A | string
```

2つめの違いは、インターフェースを拡張する場合に TypeScript は拡張元のインターフェースが拡張先のインターフェースに割当可能化どうかを確認できる
この場合エラーとなる
```typescript
interface A {
  good(x: number): string
  bad(x: number): string
}
interface B extends A {
  good(x: string | number): string
  bad(x: string): string
}
```
3つめの違いは、同じスコープ内に同じ名前のインターフェースが複数存在する場合、それらは自動的にマージされる
これは宣言のマージと呼ばれる機能

## 宣言のマージ
Userという全く同じ名前をもつ2つのインターフェースを宣言した場合、 TypeScript は自動的にそれらを結合し1つのインターフェースにまとめる
```typescript
interface User {
  name: string
}
interface User {
  age: number
}

let a: User = {
  name: 'Ashley',
  age: 30
}
```
これを型エイリアスで書き直す(これはエラー)
```typescript
type User = {
  name: string
}
type User = {
  age: number
}
```

# 実装
クラスを宣言するときに implements キーワードを使うと、そのクラスが特定のインターフェースを満たしてることを表現できる
```typescript
interface Animal {
  eat(food: string): void
  sleep(hours: number): void
}

class Cat implements Animal {
  eat(food: string) {
    console.info('Ate some', food, '. Mmm!')
  }
  sleep(hours: number) {
    console.info('Slept for', hours, 'hours')
  }
}
```
Cat クラスは Animal インターフェースをすべて満たしていないとエラーとなる

# インターフェースの実装 VS 抽象クラスの拡張
インターフェースの方が汎用的で軽量、抽象クラスは目的に特化していて機能が豊富
インターフェースは形状をモデル化するための方法
値レベルではオブジェクト、配列、クラス、クラスインスタンスを意味する
インターフェースは JavaScript を発行せずコンパイル時のみ存在する

抽象クラスがモデル化できるのはクラスのみ
抽象クラスはランタイムコードを発行する
抽象クラスはコンストラクタを持つことができ、デフォルトの実装を提供することができ、プロパティやメソッドにアクセス修飾子を設定できる

複数のクラス間で実装を共有する場合は抽象クラスを使う
このクラスは T であると表現するための軽量表現をつくときはインターフェースを使う

# クラスは値と型の両方を宣言する
typescript で表現できることの多くは、値か型のどちらか
```typescript
// 値
let a = 1999
function b() {}

// 型
type a = number
interface b {
  (): void
}
```
TypeScript では、型と値の名前空間が別々に分けられている

```typescript
class C {}
let c: C // 1
  = new c //2

enum E {F, G}
let e: E // 3
  = E.F // 4
```
1. C は、 C クラスのインスタンス型を指している
2. C は、値である C を指している
3. E は、列挙型 E の型を指している
4. E は、値である E をさしている


```typescript
type State = {
  [key: string]: string
}

class StringDatabase {
  state: State = {}
  get(key: string): string | null {
    return key in this.state ? this.state[key] : null
  }
  set(key: string, value: string): void {
    this.state[key] = value
  }
  static from(state: State) {
    let db = new StringDatabase
    for(let key in state) {
      db.set(key, state[key])
    }
    return db
  }
}
```
このクラス宣言によって下記のようなインスタンス型 StringDatabase と コンストラクタ型 typeof StringDatabase が生成される
```typescript
interface StringDatabase {
  state: State
  get(key: string): string | null
  set(key: string, value: string): void
}
interface StringDatabaseConstructor {
  new(state: State): StringDatabase
  from(state: State): StringDatabase
}
```

# ポリモーフィズム
関数や型と同様にクラスとインターフェースはじ、ジェネリック型パラメータをサポートしている
```typescript
class MyMap<K, V> { // 1
  constructor(initialkey: K, initialValue: V) { // 2
    // 処理
  }
  get(key: K): V { // 3
    // 処理
  }
  merge<K1, V1>(map: MyMap<K1, V1>): MyMap<K | K1, V | V1> { // 4
    // 処理
  }
  static of<K, V>(k: K, v: V): myMap<K, V> { // 5
    // 処理
  }
}
```
1. class を宣言するときにクラススコープのジェネリック型にバインドする
   この例だとK, V はインスタンスメソッドとインスタンスプロパティで利用できる
2. constructor の中ではジェネリック型は宣言できない
   代わりに、その宣言を class の宣言まで引き上げる
3. クラススコープのジェネリック型は、クラス内のどこでも使うことができる
4. インスタンスメソッドはクラスレベルのジェネリックにアクセスすることができ、その他の独自のジェネリックを宣言することができる
   .merge はクラスレベルのジェネリック、 K と V を使用しその他に2つの独自ジェネリック、K1, V1 を宣言する
5. 宣言メソッドは値レベルではクラスのインスタンス変数にアクセスできない
   それと同様にクラスのジェネリックにはアクセスできない
   of は1で宣言された K と V にはアクセスできないので代わりにジェネリックの K と V を宣言する

# デザインパターン
```typescript
type Shoe = {
  purpose: string
}

class BelletFlat implements Shoe {
  purpose = 'dancing'
}

class Boot implements Shoe {
  purpose = 'woodcutting'
}

class Sneaker implements Shoe {
  purpose = 'walking'
}
```
interface に書き換え
```typescript
let Shoe = {
  create(type: 'balletFlat' | 'boot' | 'sneaker'): Shoe {
    switch(type) {
      case 'balletFlat': return new BalletFlat
      case 'boot': return new Boot
      case 'sneaker': return new Sneaker
    }
  }
}
```
