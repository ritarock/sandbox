import { useContext } from "react";
import { Config, MyAppContext } from "./HookContext";

export function HookContextChild() {
  return (
    <>
      <div id="c_child">
        <HookContextChildGrand />
      </div>
    </>
  );
}

export function HookContextChildGrand() {
  const { title, lang }: Config = useContext(MyAppContext);
  return (
    <>
      <div id="c_child_grand">
        {title} ({lang})
      </div>
    </>
  );
}
