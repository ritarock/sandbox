import { createContext } from "react";
import { HookContextChild } from "./HookContextChild";

export const MyAppContext = createContext<Config>({
  title: "",
  lang: "",
});

export type Config = {
  title: string;
  lang: string;
};

const config: Config = {
  title: "title",
  lang: "jp",
};

export default function HookContext() {
  return (
    <>
      <MyAppContext.Provider value={config}>
        <div id="c_main">
          <HookContextChild />
        </div>
      </MyAppContext.Provider>
    </>
  );
}
