import { useReducer } from "react";

type CounterState = {
  count: number;
};

type CounterActions = {
  handleUp: () => void;
  handleDown: () => void;
  handleReset: () => void;
};

export default function useCounter(
  init: number,
  step: number,
): [CounterState, CounterActions] {
  const [state, dispatch] = useReducer(
    (
      state: { count: number },
      action: { type: string; step?: number; init?: number },
    ): {
      count: number;
    } => {
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

  const handleUp = () => dispatch({ type: "update", step });
  const handleDown = () => dispatch({ type: "update", step: -step });
  const handleReset = () => dispatch({ type: "reset", init });

  return [state, { handleUp, handleDown, handleReset }];
}
