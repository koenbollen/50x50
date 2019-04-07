import React from "react";
import ReactDOM from "react-dom";
import "./styles.scss";

import Grid from "./components/grid"

class App extends React.Component {
  render() {
    return (
      <div>
        <div>Try to create a fibonacci sequence in rows or columns!</div>
        <Grid/>
      </div>
    );
  }
}

var mountNode = document.getElementById("app");
ReactDOM.render(<App name="World" />, mountNode);
