import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import "./App.css";

function CalculateArea() {
  const [length, setLength] = useState(0);
  const [width, setWidth] = useState(0);

  const calculate = (length: number, width: number) => {
    return length * width;
  };

  const incrementLength = () => {
    setLength((l) => l + 1);
  };
  const incrementWidth = () => {
    setWidth((w) => w + 1);
  };

  const calculateArea = useMemo(() => calculate(length, width), [
    length,
    width,
  ]);

  return (
    <>
      <div>
        <p>length: {length}, width: {width}</p>
        <p>Area: {calculateArea}</p>
        <button onClick={incrementLength}>length + 1</button>
        <button onClick={incrementWidth}>width + 1</button>
      </div>
    </>
  );
}

function SampleCallBack() {
  const [message, setMessage] = useState("");
  const [count, setCount] = useState(0);

  // useCallback を使わないと console.log がクリックの度に出力される
  // const outputLog = (value: string) => {
  //   console.log(value)
  // }
  const outputLog = useCallback(
    (value: string) => {
      console.log(value);
    },
    [],
  );

  useEffect(() => {
    outputLog(message);
  }, [message, outputLog]);

  return (
    <>
      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
      />
      <button onClick={() => setCount(count + 1)}>click me</button>
    </>
  );
}

function App() {
  const [count, setCount] = useState(0);
  useEffect(() => {
    console.log(`${count} 回クリックされました`);
  }, [count]);

  function handleClick() {
    setCount(count + 1);
  }

  const ref = useRef(0);

  const handleRefClick = () => {
    ref.current = ref.current + 1;
    console.log(ref.current);
  };

  return (
    <>
      <div>
        <button onClick={handleClick}>
          You pressed me {count} times
        </button>
        <CalculateArea />
        <div>
          <SampleCallBack />
        </div>
        <button onClick={handleRefClick}>click(ref)</button>
      </div>
    </>
  );
}

export default App;
