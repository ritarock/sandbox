function toHex(this: number) {
  return this.toString();
}

const fiveToHex: OmitThisParameter<typeof toHex> = toHex.bind(123);

console.log(fiveToHex());
