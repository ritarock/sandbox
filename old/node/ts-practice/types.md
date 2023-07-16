## 高度な型
- intersection types (交差型)
- type guard (型ガード)
- discriminated unions (判別されるunion型)
- type casting (型キャスト)
- function overloads (関数オーバーロード)

## 交差型

## discriminated unions (判別されるunion型)
共通のプロパティをもたせることで判別させる
```ts
interface Bird {
  type: 'bird'
  flyingSpeed: number
}

interface Horse {
  type: 'horse'
  runningSpeed: number
}

type Animal = Bird | Horse

function moveAnimal(animal: Animal) {
  // if ('flyingSpeed' in animal) {
  //   console.log(animal.flyingSpeed)
  // }
  let speed
  switch (animal.type) {
    case 'bird':
      speed = animal.flyingSpeed
      break
    case 'horse':
      speed = animal.runningSpeed
  }
}
```

## 型キャスト
前に <> で書くか、後ろに as で書く.
! は絶対nullを返さないという意味.
```ts
const userInputElement1 = <HTMLInputElement>document.getElementById("user-input")!
const userInputElement2 = document.getElementById("user-input")! as HTMLInputElement
userInputElement.value = 'hello'
userInputElement2.value = 'hello'
```

## インデックス型
```ts
interface ErrorContainer { // {email: '正しいメールアドレスではありません', username: '名前に記号を含めることはできません'}
  [prop: string]: string
}

const errorBag: ErrorContainer = {
  email: '正しいメールアドレスではありません',
  username: '名前に記号を含めることはできません'
}
```

## 関数オーバーロード
受け取る型と返す型のパターンを関数の直前に記述する
```ts
function add(a: number, b:number): number;
function add(a: string, b:string): string;
function add(a: string, b:number): string;
function add(a: number, b:string): string;
function add(a: Combinable, b:Combinable) {
  if (typeof a === 'string' || typeof b === 'string') {
    return a.toString() + b.toString()
  }
  return a + b
}

const result = add('hello', 'world')
```

## オプショナルチェイン
ネストされたオブジェクトに安全にアクセスできる
```ts
const fetchedUserData = {
  id: 'u1',
  name: 'user1',
  job: {
    title: 'Developer',
    description: 'TypeScript',
  },
}

// console.log(fetchedUserData && fetchedUserData.job && fetchedUserData.job.title)
console.log(fetchedUserData?.job?.title)
```

## null合体演算子
null か undifind のときのみ判定できる
```ts
const userInput = ''
const storedData = userInput ?? 'DEFAULT'
```
