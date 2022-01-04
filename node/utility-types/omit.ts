interface Todo {
  title: string;
  desc: string;
  completed: boolean;
  createdAt: number;
}

type TodoPreview = Omit<Todo, "desc">;

const todo: TodoPreview = {
  title: "Clean room",
  completed: false,
  createdAt: 1000,
};

console.log(todo);

type TodoInfo = Omit<Todo, "completed" | "createdAt">;

const todoInfo: TodoInfo = {
  title: "Pick up kids",
  desc: "aaa",
};

console.log(todoInfo);
