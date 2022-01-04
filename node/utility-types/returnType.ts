declare function f1(): { a: number; b: string };

type T0 = ReturnType<() => string>;
type T1 = ReturnType<(s: string) => void>;
type T2 = ReturnType<<T>() => T>;
type T3 = ReturnType<<T extends U, U extends number[]>() => T>;
type T4 = ReturnType<typeof f1>;
