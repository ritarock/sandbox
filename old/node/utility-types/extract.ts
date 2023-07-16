type T0 = Extract<"a" | "b" | "c", "a" | "f">;

type T1 = Extract<string | number | (() => void), Function>;
