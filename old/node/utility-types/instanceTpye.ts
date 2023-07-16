class C {
  x = 0;
  y = 0;
}

class Foo {}

type T0 = InstanceType<typeof C>;
type T1 = InstanceType<typeof Foo>;
