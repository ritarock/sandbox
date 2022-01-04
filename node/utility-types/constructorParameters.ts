type T0 = ConstructorParameters<ErrorConstructor>;
type T1 = ConstructorParameters<FunctionConstructor>;

class Person {
  constructor(public id: number, public name: string) {}
}

type P = ConstructorParameters<typeof Person>;
