import useCounter from "./UseCounter";

export default function HookCustom() {
  const [state, handle] = useCounter(0, 1);

  return (
    <>
      <button
        onClick={handle.handleUp}
      >
        CountUp
      </button>
      <button
        onClick={handle.handleDown}
      >
        CountDown
      </button>
      <button
        onClick={handle.handleReset}
      >
        CountReset
      </button>
      <p>{state.count}</p>
    </>
  );
}
