import React, { useState } from "react";

let maxID = 0;

type Todo = {
  id: number,
  title: string,
  created: Date,
  isDone: boolean
}

export default function StateTodo() {
  const [title, setTitle] = useState<string>('');
  const [todo, setTodo] = useState<Todo[]>([]);
  const [desc, setDesc] = useState<boolean>(true);

  const handleChangeTitle = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };

  const handleClick = () => {
    setTodo([
      ...todo,
      {
        id: ++maxID,
        title,
        created: new Date(),
        isDone: false
      }
    ]);
  }

  const handleDone = (e: React.MouseEvent<HTMLButtonElement>) => {
    setTodo(todo.map(item => {
      if (item.id === Number((e.target as HTMLButtonElement).dataset.id)) {
        return {
          ...item,
          isDone: true
        };
      } else {
        return item;
      }
    }));
  }

  const handleRemove = (e: React.MouseEvent<HTMLButtonElement>) => {
    setTodo(todo.filter(item => item.id !== Number((e.target as HTMLButtonElement).dataset.id)))
  }

  const handleSort = () => {
    const sorted = [...todo];
    sorted.sort((m, n) => {
      if (desc) {
        return n.created.getTime() - m.created.getTime();
      } else {
        return m.created.getTime() - n.created.getTime();
      }
    })
    setDesc(d => !d);
    setTodo(sorted)
  }


  return (
    <>
      <div>
        <label>
          Todo:
          <input type="text" name="title"
            value={title} onChange={handleChangeTitle} />
        </label>
        <button type="button" onClick={handleClick}>
          add
        </button>
        <button type="button" onClick={handleSort}>
          sort ({desc ? 'up' : 'down'})
        </button>
        <hr />
        <ul>
          {todo.map(item => (
            <li key={item.id}
              className={item.isDone ? 'done' : ''}
            >
              {item.title}
              <button type="button"
                onClick={handleDone} data-id={item.id}
              >
                DONE
              </button>
              <button type="button"
                onClick={handleRemove} data-id={item.id}
              >
                REMOVE
              </button>
            </li>
          ))}
        </ul>
      </div>
    </>
  )
}
