type Admin = {
  name: string
  privileges: string[]
}

type Employee = {
  name: string
  startDate: Date
}

// interface Admin {
//   name: string
//   privileges: string[]
// }
// interface Employee {
//   name: string
//   startDate: Date
// }

// interface ElevatedEmploee extends Admin, Employee {}
type ElevatedEmploee = Admin & Employee

const e1: ElevatedEmploee = {
  name: 'max',
  privileges: ['create-server'],
  startDate: new Date(),
}

type Combinable = string | number
type Numeric = number | boolean
type Universal = Combinable & Numeric

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

type UnknownEmployee = Employee | Admin

function printEmployeeInformation(emp: UnknownEmployee) {
  console.log(emp.name)
  if ('privileges' in emp) {
    console.log(emp.privileges)
  }
  if ('startDate' in emp) {
    console.log(emp.startDate)
  }
}

class Car {
  drive() {
    console.log('driving')
  }
}

class Truck {
  drive() {
    console.log('driving truck')
  }
  loadCargo(amount: number) {
    console.log('truck')
  }
}

type Vehicle = Car | Truck
const v1 = new Car()
const v2 = new Truck()

function useVehicle(vehicle: Vehicle) {
  if (vehicle instanceof Truck) {
    vehicle.loadCargo(100)
  }
}

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

// const userInputElement = <HTMLInputElement>document.getElementById("user-input")!
const userInputElement = document.getElementById("user-input")! as HTMLInputElement
userInputElement.value = 'hello'

interface ErrorContainer { // {email: '正しいメールアドレスではありません', username: '名前に記号を含めることはできません'}
  [prop: string]: string
}

const errorBag: ErrorContainer = {
  email: '正しいメールアドレスではありません',
  username: '名前に記号を含めることはできません'
}


const fetchedUserData = {
  id: 'u1',
  name: 'user1',
  job: {
    title: 'Developer',
    description: 'TypeScript',
  },
}

console.log(fetchedUserData && fetchedUserData.job && fetchedUserData.job.title)
console.log(fetchedUserData?.job?.title)

const userInput = ''
const storedData = userInput ?? 'DEFAULT'
