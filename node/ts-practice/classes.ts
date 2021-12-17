abstract class Department {
  static fiscalYear = 2021;
  // private id: string;
  // name: string;
  protected employees: string[] = [];

  static createEmpoyee(name: string) {
    return { name: name };
  }

  constructor(protected readonly id: string, public name: string) {
    console.log(Department.fiscalYear);
  }

  abstract describe(this: Department): void;

  addEmployee(employee: string) {
    this.employees.push(employee);
  }

  printEmployyInfomation() {
    console.log(this.employees.length);
    console.log(this.employees);
  }
}

class ItDepartment extends Department {
  admins: string[];
  constructor(id: string, admins: string[]) {
    super(id, "IT");
    this.admins = admins;
  }
  describe() {
    console.log(` IT Department (${this.id}: ${this.name})`);
  }
}

class AccountingDepartment extends Department {
  private lastReport: string;
  private static instance: AccountingDepartment;

  get mostRecentReport() {
    if (this.lastReport) {
      return this.lastReport;
    }
    throw new Error("not found report.");
  }

  set mostRecentReport(value: string) {
    if (!value) {
      throw new Error("not value.");
    }
    this.addReport(value);
  }

  private constructor(id: string, private reports: string[]) {
    super(id, "Accounting");
    this.lastReport = reports[0];
  }

  static getInstance() {
    if (AccountingDepartment.instance) {
      return this.instance;
    }
    this.instance = new AccountingDepartment("d2", []);
    return this.instance;
  }

  describe() {
    console.log(` Account Department (${this.id}: ${this.name})`);
  }

  addReport(text: string) {
    this.reports.push(text);
    this.lastReport = text;
  }

  printReports() {
    console.log(this.reports);
  }
  addEmployee(name: string) {
    if (name === "max") {
      return;
    }
    this.employees.push(name);
  }
}

const employee1 = Department.createEmpoyee("max");
console.log(employee1, Department.fiscalYear);

const it = new ItDepartment("d1", ["max"]);
console.log(it);

it.addEmployee("max");
it.addEmployee("manu");

it.describe();
it.printEmployyInfomation();

// const accounting = new AccountingDepartment('d2', [])
const accounting = AccountingDepartment.getInstance();
accounting.mostRecentReport = "something";
accounting.describe();
// accounting.printReports()
accounting.addEmployee("max");
accounting.addEmployee("manu");
// accounting.printEmployyInfomation()
console.log(accounting.mostRecentReport);
