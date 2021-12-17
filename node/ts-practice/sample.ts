interface Todo {
  title: string;
  desc: string;
  date: Date;
}

function createTodo(
  title: string,
  desc: string,
  date: Date,
): Todo {
  let todo: Partial<Todo> = {};
  todo.title = title;
  todo.desc = desc;
  todo.date = date;
  return todo as Todo;
}
