import React from "react";
import fetch from "isomorphic-unfetch";

import Cell from "./cell";

class Grid extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};

    this.onCellClick = this.onCellClick.bind(this)
  }

  async componentDidMount() {
    const resp = await fetch("http://localhost:8080/state/create");
    const data = await resp.json();
    this.setState({...data})
  }

  async onCellClick(cell, i, e) {
    const x = i % this.state.size;
    const y = Math.floor(i / this.state.size);

    const resp = await fetch("http://localhost:8080/state/move", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        session: this.state.session,
        x, y
      })
    });
    const data = await resp.json();
    this.setState({
      ...this.state,
      ...data
    })
  }

  render() {
    const data = this.state.grid || [];
    return (
      <div className="grid">
        {data.map( (c, i) => {
          return (<Cell key={i} index={i} value={c} onClick={this.onCellClick} />);
        })}
      </div>
    );
  }
}

export default Grid
