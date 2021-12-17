文字列を数値に変換するには `+` をつけるだけでできる.

```js
const result = +a + +b;
```

存在することが確定している場合 `!` をつければエラーを回避できる. `as` を付けて型キャストできる.

```ts
const input = document.getElementById("num")! as HTMLInputElement;
```

```ts
class Department {
  private id: string;
  name: string;
  private employees: string[] = [];

  constructor(id: string, n: string) {
    this.id = id;
    this.name = n;
  }

  describe(this: Department) {
    console.log("Department: " + this.name);
  }

  addEmployee(employee: string) {
    this.employees.push(employee);
  }

  printEmployyInfomation() {
    console.log(this.employees.length);
    console.log(this.employees);
  }
}
```

フィールドとコンストラクタはまとめることができる. やりかたは `constructor` にアクセス修飾子を書くだけ.

```ts
class Department {
  // private readonly id: string;
  // name: string;
  private employees: string[] = [];

  constructor(private readonly id: string, public name: string) {
  }

  describe(this: Department) {
    console.log("Department: " + this.name);
  }

  addEmployee(employee: string) {
    this.employees.push(employee);
  }

  printEmployyInfomation() {
    console.log(this.employees.length);
    console.log(this.employees);
  }
}

const accounting = new Department("1,", "Accounting");
console.log(accounting);

accounting.addEmployee("max");
accounting.addEmployee("manu");

accounting.describe();
accounting.printEmployyInfomation();

// const accountingCopy = { name: 'Dummy', describe: accounting.describe }
// accountingCopy.describe()
```

readonly : 初期化したあと変更ができない protected : スブクラスからアクセスを許可する

getter は()不要. プロパティのように実行する.

```ts
  get mostRecentReport() {
    if (this.lastReport) {
      return this.lastReport
    }
    throw new Error("not found report.")
  }


console.log(accounting.mostRecentReport)
```

setter も同様

```ts
  set mostRecentReport(value: string) {
    if (!value) {
    throw new Error("not value.")
    }
    this.addReport(value)
  }

accounting.mostRecentReport = 'something'
```

static メソッドや static プロパティには this でアクセスできない。(インスタンスからアクセスできない)

抽象メソッドは抽象クラス内でのみ使える. 抽象メソッドは関数の構造だけを定義しておく
抽象クラスからはインスタンスをつくれない。継承したサブクラスからはインスタンスを使える

```ts
abstract describe(this: Department): void
```

シングルトンパターン オブジェクトを1つしか作らせたくない場合に使う.

```ts
  private static instance: AccountingDepartment

  static getInstance() {
    if (AccountingDepartment.instance) {
      return this.instance
    }
    this.instance = new AccountingDepartment("d2", [])
    return this.instance
  }



// const accounting = new AccountingDepartment('d2', [])
const accounting = AccountingDepartment.getInstance()
```

interface とは オブジェクトがどんな形であるか定義するもの.
オブジェクトの設計図ではなくオリジナルの型を定義している.

interface と カスタムタイプの違い.
インターフェースはオブジェクトの構造を記述するためだけに使う
カスタムタイプはunion型などさまざまな型を定義できる
インターフェースをつかえばオブジェクトの構造を定義したいという意図を明確にできる.

interfaceを実装する
```ts
class Person implements Greetable {
}
```

class がどんな構造を持つのかを定義するときに使う.複数のクラスで同じ機能を持たせたいときに使える.
interfaceは実際の機能や値をもつことはない.
抽象クラスに似ている.違いはinterfaceは値や実装を持たない.抽象クラスは実際の値や実装を混在させてもったり持たなかったりできる.

interfaceはreadonlyできる

interfaceでも継承できる
```ts
interface Named {
  readonly name: string
}

interface Greetable extends Named {
  greet(phrase: string): void
}
```
