## デコレータ

## クラスデコレータ
```ts
function Logger(constructor: Function) {
  console.log('ログ出力中...')
  console.log(constructor)
}

@Logger
class Person {
  name = 'max'

  constructor() {
    console.log('creating person object...')
  }
}

const pers = new Person()

console.log(pers)
```

## デコレータファクトリ
```ts
function Logger(logString: string) {
  return function(constructor: Function) {
    console.log(logString)
    console.log(constructor)
  }
}

@Logger('ログ出力中...')
class Person {
  name = 'max'

  constructor() {
    console.log('creating person object...')
  }
}

const pers = new Person()

console.log(pers)
```

複数のデコレータがある場合下から順に実行される.(呼び出し順は上から)
```ts
@Logger
@WithTemplate
```

デコレータは class が定義されたときに実行する
