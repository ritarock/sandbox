import React from "react";
import ReactDOM from "react-dom/client";
import "./App.css";
import "./index.css";
import StateEffect from "./StateEffect.tsx";
import HookTimer from "./HookTimer.tsx";
import App from "./App.tsx";
import HookRef from "./HookRef.tsx";
import HookReducer from "./useReducer.tsx";
import HookReducerUp from "./HookReducerUp.tsx";
import HookContext from "./HookContext.tsx";
import HookMemo from "./HookMemo.tsx";
import HookCustom from "./HookCustom.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App />
    <StateEffect init={0} />
    <HookTimer init={10} />
    <HookRef />
    <HookReducer init={0} />
    <HookReducerUp init={0} />
    <HookContext />
    <HookMemo />
    <HookCustom />
  </React.StrictMode>,
);
