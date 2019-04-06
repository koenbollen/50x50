import React from "react";

class Cell extends React.Component {
  constructor(props) {
    super(props);
    this.state = {isIncrement: false};

    this.onClick = this.onClick.bind(this)
  }

  onClick(e) {
    this.setState({
      ...this.state,
      isIncrement: true
    })
  }

  render() {
    var className = "cell";
    if(this.state.isIncrement) {
      className += " increment";
    } else if(this.state.isDecrement) {
      className += " decrement";
    }
    return (
      <div className={className} onClick={this.onClick}>0</div>
    );
  }
}

export default Cell
