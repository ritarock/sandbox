import { invoke } from "@tauri-apps/api/core";
import "./App.css";
import { useState } from "react";

function App() {
  const [msg, setMsg] = useState("");

  function myCustomCommand() {
    invoke("my_custom_command");
  }

  function myCustomCommand2() {
    invoke("my_custom_command2", { invokeMessage: "hello"});
  }

  function myCustomCommand3() {
    // invoke("my_custom_command3").then((message) => console.log(message));
    invoke("my_custom_command3").then((message) => setMsg(message as string));
  }

  return (
    <main className="container">
      <button onClick={myCustomCommand}>myCustomCommand</button>
      <button onClick={myCustomCommand2}>myCustomCommand2</button>
      <button onClick={myCustomCommand3}>myCustomCommand3</button>
      <p>{msg}</p>
    </main>
  );
}

export default App;
