import React from "react"

class Tick extends React.Component<any, any> {
  timerID: any;

  constructor(props:any) {
    super(props);
    this.state = {date: new Date()};
  }

  componentDidMount() {
    this.timerID = setInterval(
      () => this.tick(),
      1000
    );
  }

  componentWillUnmount() {
    clearInterval(this.timerID)
  }

  tick() {
    this.setState({
      date: new Date()
    })
  }

  render() {
    return (
      <h2>It is {this.state.date.toLocaleTimeString()}.</h2>
    )
  };
}

export default Tick