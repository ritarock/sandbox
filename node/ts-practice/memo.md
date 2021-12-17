## class
```ts
class Person {
  id: string;
  name: string;

  constructor(id: string, n: string) {
    this.id = id;
    this.name = n;
  }

  printData() {
    console.log(`ID: ${this.id}; NAME: ${this.name}`);
  }
}

const person = new Person("1", "hoge");
person.printData(); // ID: 1; NAME: hoge
```
フィールドとコンストラクタはまとめる事ができる.やりかたは `constractor` にアクセス修飾子を書くだけ.
引数の名前はフィールド名にしないといけない.
```ts
class Person {
  constructor(public id: string, public name: string) {
  }

  printData() {
    console.log(`ID: ${this.id}; NAME: ${this.name}`);
  }
}

const person = new Person("1", "hoge");
person.printData(); // ID: 1; NAME: hoge
```

`private` は外部からアクセスできない. `readonly` は初期化後に変更ができない.
`protected` はサブクラスからのみアクセス可能.
```ts
class Person {
  constructor(private readonly id: string, public name: string) {
  }

  printData() {
    console.log(`ID: ${this.id}; NAME: ${this.name}`);
  }
}

const person = new Person("1", "hoge");
person.printData(); // ID: 1; NAME: hoge
```

### getter & setter
`()` は不要.プロパティのように実行する.
```ts
class Person {
  private _report: string;

  constructor(private id: string, public name: string) {
    this._report = "";
  }

  get report() {
    return this._report;
  }

  set report(input: string) {
    this._report = input;
  }

  printData() {
    console.log(`ID: ${this.id}; NAME: ${this.name}`);
  }
}

const person = new Person("1", "hoge");
person.printData(); // ID: 1; NAME: hoge
person.report = "report1";
console.log(person.report); // report1
```

### static メソッド & static プロパティ
static メソッド や static プロパティには this でアクセスできない.(インスタンスからアクセスできない)
```ts
class Person {
  static year = 2021;
}
console.log(Person.year); // 2021
```

### 抽象メソッド
抽象メソッドは抽象クラス内でのみ使える.抽象メソッドは関数の構造のみを定義しておく.
抽象クラスからはインスタンスを作れない.継承したサブクラスからはインスタンスを作れる.
```ts
abstract class Product {
  constructor(protected readonly id: string, public name: string) {
  }
  abstract describe(): void;
}

class Product1 extends Product {
  constructor(id: string, name: string) {
    super(id, name);
  }
  describe() {
    console.log(`ID: ${this.id}; NAME: ${this.name}`);
  }
}

const product1 = new Product1("1", "hoge");
product1.describe(); // ID: 1; NAME: hoge
```

### シングルトンパターン
オブジェクトを1つしか作らせたくない場面で使う.
```ts
class Person {
  private static instance: Person;

  static getInstance() {
    if (Person.instance) {
      return this.instance;
    }
    this.instance = new Person();
    return this.instance;
  }
}

const person = Person.getInstance();
```

## interface
interface とはオブジェクトがどんな形であるか定義するもの.
interface とカスタムタイプの使い分けは, interface はオブジェクトの構造を記述するために使う.
カスタムタイプは union 型などさまざまな型を定義できる. interface を使えばオブジェクトの構造を定義したいという意図を明確にできる.
また, interface は readonly や継承もできる.
```ts
interface Named {
  readonly name: string;
}

interface Greetable extends Named {
  hello(phrase: string): void;
}
```


### implements
`implements` を使って実装する.
抽象クラスから作るときの違いは, interface は値や実装を持たない.抽象クラスは実際の値や実装を混在させることができる.
```ts
interface PersonInterface {
  id: string;
  name: string;
  describe(): void
}

class Person implements PersonInterface {
  id: string;
  name: string;
  describe() {
    console.log(`ID: ${this.id}; NAME: ${this.name}`)
  }
}
```

## 型 いろいろ
- discriminated unions
- 型ガード
- 型キャスト
- 関数のオーバーロード


### discriminated unions
意味のある共通のプロパティをもたせて判別に使う.
```ts
interface Bird {
  flySpeed: number;
  type: "bird";
}

interface Horse {
  runSpeed: number;
  type: "horse";
}

type Animal = Bird | Horse;

function animalSpeed(animal: Animal) {
  switch (animal.type) {
    case "bird":
      console.log(animal.flySpeed);
      break;
    case "horse":
      console.log(animal.runSpeed);
  }
}
```

### 型キャスト
文字列を数値に変換するには `+` をつけるだけで良い.
```ts
const result = +a + +b;
```

前に `<>` で型を書くか,後ろに `as` を付けて型キャストできる
存在が確定している値は `!` をつけるとエラーを回避できる.
```ts
const input1 = <HTMLInputElement> document.getElementById("num")!;
const input2 = document.getElementById("num")! as HTMLInputElement;
```

### インデックス型
```ts
interface ErrorInterface {
  [prop: string]: string;
}

const errorMessage: ErrorInterface = {
  email: "hoge",
  name: "hogehoge",
  message: "fuga",
};
console.log(errorMessage); // { email: "hoge", name: "hogehoge", message: "fuga" }
```

### 関数オーバーロード
受け取る型と返す型のパターンを関数の直前に記述する.
```ts
type Input = string | number

function adder(a: number, b: number): number
function adder(a: string, b: number): string
function adder(a: number, b: string): string 
function adder(a: string, b: string): string
function adder(a: Input, b: Input) {
  if (typeof a === 'string' || typeof b === 'string') {
    return a.toString() + b.toLocaleString()
  }
  return a + b
}

console.log(adder(1, 2))
```

### オプショナルチェイン
`?` を使って安全にオブジェクトにアクセスする.
```ts
const fetched = {
  id: "id1",
  name: "hoge",
  job: {
    title: "Developer",
    desc: "TypeScript",
  },
};

console.log(fetched && fetched.job && fetched.job.title)
console.log(fetched?.job?.title);
```

### null 合体演算子
`null` か `undefined` のときのみ判定できる.
```ts
let input = "";
const inputData = input ?? "Default";

console.log(inputData);
```

## ジェネリクス
関数の後ろに `<>` を付けて表現する.
```ts
function mergeObject(objA: object, objB: object) {
  return Object.assign(objA, objB);
}

console.log(mergeObject({ id: "1" }, { name: "hoge" })); // { id: "1", name: "hoge" }
```

```ts
function mergeObject<T>(objA: T, objB: T) {
  return Object.assign(objA, objB);
}

console.log(mergeObject({ id: "1" }, { name: "hoge" })); // { id: "1", name: "hoge" }
```

### 制約をつける
extends キーを使う.
```ts
function mergeObject<T extends object, U extends object>(objA: T, objB: U) {
  return Object.assign(objA, objB);
}

console.log(mergeObject({ id: "1" }, { name: "hoge" })); // { id: "1", name: "hoge" }
```

### keyof
`keyof` を使うことでオブジェクトのキーの制約をもたせる.
```ts
function addConvert<T extends object, U extends keyof T>(obj: T, key: U) {
  return "value: " + obj[key];
}

addConvert({ name: "hoge" }, "name");
```

### ジェネリクスクラス
クラスの後ろに `<>` を付けて表現.
```ts
class DataStore<T extends string | number> {
  private data: T[] = [];

  addItem(item: T) {
    this.data.push(item);
  }
  removeItem(item: T) {
    if (this.data.indexOf(item) === -1) {
      return;
    }
    this.data.splice(this.data.indexOf(item), 1);
  }
  getItems() {
    return [...this.data];
  }
}

const stringData = new DataStore<string>();
const numberData = new DataStore<number>();
stringData.addItem("data1");
stringData.addItem("data2");
numberData.addItem(1);
numberData.addItem(2);
stringData.removeItem("data1");
numberData.removeItem(1);
console.log(stringData.getItems()); // [ "data2" ]
console.log(numberData.getItems()); // [ 2 ]
```

### Partial
一時的に別の型に切り替えることができる.
`Partial` で最終的にキャストされる型を指定する.
return するときは `as` でキャストする.
```ts
interface Todo {
  title: string;
  desc: string;
  date: Date;
}

function createTodo(
  title: string,
  desc: string,
  date: Date,
): Todo {
  let todo: Partial<Todo> = {};
  todo.title = title;
  todo.desc = desc;
  todo.date = date;
  return todo as Todo;
}
```
