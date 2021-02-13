function squareOf(n: number) {
  return n * n
}

squareOf(2)
squareOf('z')

let a: {b: number} = {
  b: 12
}

let c : {
  firstName: string
  lastName: string
} = {
  firstName: 'john',
  lastName: 'barrowman'
}

class Person {
  constructor(
    public firstName: string,
    public lastName: string
  ) {}
}
c = new Person('matt', 'smith')

let aa: {
  b: number
  c?: string
  [key: number]: boolean
}

aa = {b:1}
aa = {b: 1, c: undefined }
aa = {b:1, 10: true, 20: false}

let user: {
  readonly firstName: string
} = {
  firstName: 'abby'
}

user.firstName
user.firstName = 'abby with an'
