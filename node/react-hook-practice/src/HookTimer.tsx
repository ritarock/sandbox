import { useEffect, useState } from "react";

export default function HookTimer({ init }: { init: number }) {
  const [count, setCount] = useState(init);

  useEffect(() => {
    const t = setInterval(() => {
      setCount((c) => c - 1);
    }, 1000);

    return () => {
      clearInterval(t);
    };
  }, []);

  return (
    <>
      <div>Count: {count}</div>
    </>
  );
}
