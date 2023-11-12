import { RouterProvider } from "react-router-dom";
import routeBasic from "./routeBasic";
import "./App.css";

function App() {
  return (
    <>
      <RouterProvider router={routeBasic} />
    </>
  );
}

export default App;
