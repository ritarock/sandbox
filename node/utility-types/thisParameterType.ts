function toHex(this: number) {
  return this.toString();
}

function numberToString(n: ThisParameterType<typeof toHex>) {
  return toHex.apply(n);
}

console.log(numberToString(123));
