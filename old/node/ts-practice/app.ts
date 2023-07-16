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


function Log(target: any, propatyName: string|Symbol) {
  console.log("propaty decorator")
  console.log(target, propatyName)
}

class Product {
  @Log
  title: string
  private _price: number

  set private(val: number) {
    if (val > 0) {
      this._price = val
    } else {
      throw new Error("errer")
    }
  }

  constructor(t: string, p: number) {
    this.title = t
    this._price = p
  }

  getPriceWithTax(tax: number) {
    return this._price * (1+tax)
  }
}
