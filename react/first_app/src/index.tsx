import React from "react"
import ReactDOM from "react-dom"

import Hello from "./components/hello";

ReactDOM.render(
  <Hello name="TypeScript" enthusiasmLevel={10} />,
  document.getElementById("root") as HTMLElement
);