## ジェネリクス関数
関数の後ろに <> を付けて表現

## 制約の付け方
extends キーを使って表現する.
```ts
function merge<T extends object, U extends object>(objA: T, objB: U) {
  return Object.assign(objA, objB)
}
```

## key of
`key of` を使うことでオブジェクトのキーの制約をもたせる
```ts
function extractAndConvert<T extends object, U extends keyof T>(obj: T, key: U) {
  return 'value: ' + obj[key]
}

extractAndConvert({name: "max"}, "name")
```

## ジェネリクスクラス
class の後ろに <> を付けて表現
```ts
class DataStorage<T extends string|number|boolean> {
  private data: T[] = []

  addItem(item: T) {
    this.data.push(item)
  }
  removeItem(item: T) {
    if (this.data.indexOf(item) === -1) {
      return
    }
    this.data.splice(this.data.indexOf(item), 1) // -1
  }
  getItems() {
    return [...this.data]
  }
}

const textStorage = new DataStorage<string>()
textStorage.addItem('data1')
textStorage.addItem('data2')
textStorage.removeItem('data1')
console.log(textStorage.getItems())
```

## Partial
一時的に別の型に切り替えたりできる.
```ts
interface CourseGoal {
  title: string,
  description: string,
  date: Date
}

function createCourseGoal(
  title: string,
  description: string,
  date: Date
): CourseGoal {
  let courseGoal: Partial<CourseGoal> = {}
  courseGoal.title = title
  courseGoal.description = description
  courseGoal.date = date
  return courseGoal as CourseGoal
}
```

## readonly
オブジェクトや配列に対して値を追加したり削除したりできな
```ts
const names: Readonly<string[]> = ['max', 'anna']
// names.push('manu')
```
