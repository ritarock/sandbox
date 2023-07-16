declare function f1(arg: { a: number; b: string }): void;

type T0 = Parameters<() => string>;
type T1 = Parameters<(s: string) => void>;
type T2 = Parameters<<T>(arg: T) => T>;
type T3 = Parameters<typeof f1>;

type F = (arg1: string, arg2: number) => string;
type F1 = Parameters<F>;
const v: F1 = ["a", 123];
console.log(v);
