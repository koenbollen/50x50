import React from "react";

import Cell from "./Cell";

class Grid extends React.Component {
  render() {
    const data = new Array(50*50).fill(0);
    return (
      <div className="grid">
        {data.map( (c, i) => {
          return (<Cell key={i} value={c} />);
        })}
      </div>
    );
  }
}

export default Grid
