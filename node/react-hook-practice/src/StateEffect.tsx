import { useEffect, useState } from "react";

export default function StateEffect({ init }: { init: number }) {
  const [count, setCount] = useState(init);
  const [hoge, setHoge] = useState(0);

  useEffect(() => {
    console.log(`count is ${count}`);
  }, [count]);
  const handleClick = () => setCount(count + 1);

  return (
    <>
      <button onClick={() => setHoge(Date.now())}>Hoge ({hoge})</button>
      <button onClick={handleClick}>count</button>
      <p>{count} clicked</p>
    </>
  );
}
