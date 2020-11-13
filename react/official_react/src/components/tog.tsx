import React from "react"

interface IState {
  IsToggleOn: boolean;
}

class Tog extends React.Component<any, IState> {
  constructor(props: any) {
    super(props);
    this.state = {IsToggleOn: true};

    this.handleClick = this.handleClick.bind(this);
  }
  
  handleClick() {
    this.setState(state => ({
      IsToggleOn: !state.IsToggleOn
    }));
  }

  render() {
    return (
      < button onClick={this.handleClick}>
        {this.state.IsToggleOn ? 'ON' : 'OFF'}
      </button>
    );
  }
}

export default Tog