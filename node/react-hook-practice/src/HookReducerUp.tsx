import { useReducer } from "react";

export default function HookReducerUp({ init }: { init: number }) {
  const [state, dispatch] = useReducer(
    (
      state: { count: number },
      action: { type: string; step?: number; init?: number },
    ): { count: number } => {
      switch (action.type) {
        case "update":
          return { count: state.count + action.step! };
        case "reset":
          return { count: action.init! };
        default:
          return state;
      }
    },
    {
      count: init,
    },
  );

  const handleUp = () => dispatch({ type: "update", step: 1 });
  const handleDown = () => dispatch({ type: "update", step: -1 });
  const handleReset = () => dispatch({ type: "reset", init: 0 });

  return (
    <>
      <button onClick={handleUp}>CountUp</button>
      <button onClick={handleDown}>CountDown</button>
      <button onClick={handleReset}>CountReset</button>
      <p>{state.count}</p>
    </>
  );
}
