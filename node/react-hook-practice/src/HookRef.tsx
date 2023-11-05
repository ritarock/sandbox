import { useRef, useState } from "react";

export default function HookRef() {
  const id = useRef<number>(0);
  const [count, setCount] = useState(0);

  const handleStart = () => {
    if (id.current === 0) {
      id.current = setInterval(() => setCount((c) => c + 1), 1000);
    }
  };
  const handleEnd = () => {
    clearInterval(id.current);
    id.current = 0;
  };

  return (
    <>
      <button onClick={handleStart}>start</button>
      <button onClick={handleEnd}>end</button>
      <p>{count}</p>
    </>
  );
}
