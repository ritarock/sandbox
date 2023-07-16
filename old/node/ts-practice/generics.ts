// const names: Array<string> = []

// const promise = new Promise<string>((resolve, reject) => {
//   setTimeout(() => {
//     resolve('finish')
//   }, 2000)
// })

// promise.then(data => {
//   data.split(' ')
// })

function merge<T extends object, U extends object>(objA: T, objB: U) {
  return Object.assign(objA, objB)
}

const mergedObj = merge({name: 'max', hobbies: ['Sports']}, {age: 10})

interface Lengthy {
  length: number
}

function countAndDescribe<T extends Lengthy>(element: T): [T, string] {
  let descriptionText = 'not value'
  if (element.length > 0) {
    descriptionText = "count " + element.length
  }
  return [element, descriptionText]
}

console.log(countAndDescribe("hello"))
console.log(countAndDescribe(['cooking', 'sports']))
console.log(countAndDescribe([]))

function extractAndConvert<T extends object, U extends keyof T>(obj: T, key: U) {
  return 'value: ' + obj[key]
}

extractAndConvert({name: "max"}, "name")

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

// const names: Readonly<string[]> = ['max', 'anna']
// names.push('manu')
