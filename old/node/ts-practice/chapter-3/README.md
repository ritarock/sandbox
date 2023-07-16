# 型のはなし
## any
すべての型を持つ。あるものの型がわからないときのデフォルト型が any
最終手段であり、できるだけ回避すべき

## unknown
any と同様に任意の値を示す
その値が何であるかわからない場合は any でなく unknown を使うべき。型アサーションがないと使用できない

## bookean
true, false の2つの値を持つ

## number
整数、浮動小数点、整数、負数、無限大( Infinity )、非数( NaN )

## bigint
number 型では 2^53 までの整数を表すことができるが、 bigint 型はそれより大きい整数を表すことができる

## string
すべての文字列

## symbol
オブジェクトやマップにおいて文字列のキーの代わりとして既知のキーが適切に使われているかを確信させるときに用いる

## オブジェクト
```typescript
let a: {
    b: number // number であるプロパティb をもつ
    c?: string // string であるプロパティ c をもつ可能性がある
    [key: number]: boolean // booleanである数値プロパティを任意の数持つ
}

a = {b: 1} // ok
a = {b: 1, c: undefined } // ok
a = {b: 1, c: 'd'} // ok
a = {b: 1, 10: true} // ok
a = {b: 1, 10: true, 20: false} // ok
a = {10: true} // error
a = {b: 10, 30: 'red'} // error

let user: {
  readonly firstName: string
} = {
  firstName: 'abby'
}

user.firstName
user.firstName = 'abby with an' // error

```
### インデントシグネチャ
**[key: T]: U** この構文はインデックスシグネチャと呼ばれる
オブジェクトがより多くのキーを含む可能性があることを TypeScript に伝える
このオブジェクトは型 T のすべてのキーは、型 U の値を持たなければならない
また、インデックスシグネチャの型 T は number か string のどちらかでなければならない
インデックスシグネチャのキーには任意の言葉が使える
```typescript
let airplaneSeatingAssigments: {
  [seatNumber: string]: string
} = {
  '34D': 'Boris Cherny',
  '34E': 'Bill Gates'
}
```

## 型エイリアス
型を宣言する
```typescript
type Age = number
type Person = {
  name: string
  age: Age
}
```

## 合併型と交差型
```typescript
type Cat = {name: string, purrs: boolean}
type Dog = {name: string, barks: boolean, wags: boolean}
type CatOrDogOrBoth = Cat | Dog // 合併
type CatAndDog = Cat & Dog // 交差
```
合併型だと、どちらかのメンバーだけを持つわけでなく同時にどちらのメンバーを持つことができる
```typescript
let a = {
  name: 'Domino',
  barks: true,
  purrs: true,
  wags: true
}
```

## 配列
**T[]** もしくは **Array[]** で宣言できる

## タプル
配列のサブタイプ
固定長の配列の型付けができる
```typescript
let a: [number] = [1]
let b: [string, string, number] = ['malcolm', 'gladwell', 1963]

let trainFares: [number, number?][] = [
  [3.75],
  [8.25, 7.70],
  [10.50]
]

let moreTrainFares: ([number] | [number, number])[] = [
  // 上記と同等
]
```
読み取り専用も定義できる
```typescript
let a: readonly number[] = [1, 2, 3]
```

## null, undefined, void, never
型 | 意味
-|-
null | 値の欠如
undefined | 値がまだ割り当てられていない変数
void | return 文を持たない関数の戻り値
never | 決して戻ることのない関数の戻り値

## 列挙型
```typescript
enum Language {
    English,
    Spanish,
    Russian
}
```
明示的に値を設定することもできる
```typescript
enum Language {
    English = 0,
    Spanish = 1,
    Russian = 2
}
```
値を取得するにはドット、もしくは角括弧を使う
```typescript
let myFirstLanguage = Language.Russian
let mySecondLanguage = Language['English']
```
