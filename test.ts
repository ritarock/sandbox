type hogeType = "h" | "ho" | "hog" | "hoge";
type fugaType = "f" | "fu" | "fug" | "fuga";

const mapping: { [k1: string]: { [k2: string]: string } } = {
  "hoge": {
    "h": "A",
    "ho": "B",
    "hog": "C",
    "hoge": "D",
  },
  "fuga": {
    "f": "E",
    "fu": "F",
    "fug": "G",
    "fuga": "H"
  }
};

function yobikata(hoge: hogeType, fuga: fugaType) {
  const x = mapping[a];
  if(!x) return undefined;
  return x[b];
}

console.info(yobikata("hoge", "hog")); // C
console.info(yobikata("fuga", "f")); // E
console.info(yobikata("X", "1")); // ERROR Argument of type '"X"' is not assignable to parameter of type 'hogeType'.
