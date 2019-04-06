import React from "react";
import ReactDOM from "react-dom";
import "./styles.scss";

import Grid from "./components/grid"

class App extends React.Component {
  render() {
    return (
      <div>
        <div>Hello {this.props.name}</div>
        <Grid/>
      </div>
    );
  }
}

var mountNode = document.getElementById("app");
ReactDOM.render(<App name="World" />, mountNode);
