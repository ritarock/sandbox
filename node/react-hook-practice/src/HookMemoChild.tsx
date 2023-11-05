import { memo } from "react";

export const MyButton = memo(
  (
    { id, handleClick, children }: {
      id: string;
      handleClick: () => void;
      children: React.ReactNode;
    },
  ) => {
    console.log(`MyButton is called: ${id}`);

    return <button onClick={handleClick}>{children}</button>;
  },
);

export const MyCounter = memo(
  ({ id, value }: { id: string; value: number }) => {
    console.log(`MyCounter is called: ${id}`);

    return <p>NOW: {value}</p>;
  },
);
