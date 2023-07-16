import { type } from "os"

function add(a: number, b: number): number {
  return a + b
}

function log(message: string, userId?: string) {
  let time = new Date().toLocaleDateString()
  console.log(time, message, userId || 'Not signed in')
}

function sum(number: number[]): number {
  return number.reduce((total, n) => total + n, 0)
}

sum([1, 2, 3]) //6
