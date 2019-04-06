import React from "react";

class Cell extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isIncrement: false,
      isDecrement: false
    };

    this.onClick = this.onClick.bind(this)
  }

  componentWillReceiveProps(nextProps) {
    if(nextProps.value > this.props.value) {
      this.setState({isIncrement: true});
      clearTimeout(this.timeout);
      this.timeout = setTimeout(() => {
        this.setState({isIncrement: false});
      }, 510);
    } else if(nextProps.value < this.props.value) {
      this.setState({isDecrement: true});
    }
  }

  componentWillUnmount() {
    clearTimeout(this.timeout)
  }

  onClick(e) {
    this.props.onClick(this, this.props.index, e);
  }

  render() {
    var className = "cell";
    if(this.state.isIncrement) {
      className += " increment";
    } else if(this.state.isDecrement) {
      className += " decrement";
    }
    return (
      <div className={className} onClick={this.onClick}>{this.props.value}</div>
    );
  }
}

export default Cell
