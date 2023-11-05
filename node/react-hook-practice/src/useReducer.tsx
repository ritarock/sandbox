import { useReducer } from "react";

export default function HookReducer({ init }: { init: number }) {
  const [state, dispatch] = useReducer(
    (state: { count: number }, action: { type: string }): {
      count: number;
    } => {
      switch (action.type) {
        case "update":
          return { count: state.count + 1 };
        default:
          return state;
      }
    },
    {
      count: init,
    },
  );

  const handleClick = () => {
    dispatch({ type: "update" });
  };

  return (
    <>
      <button onClick={handleClick}>Count</button>
      <p>{state.count}</p>
    </>
  );
}
