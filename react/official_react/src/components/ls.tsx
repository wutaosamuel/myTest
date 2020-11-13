import React from 'react'

interface IProps {
  Numbers?: number[];
  IsDouble?: boolean;
}

class ls extends React.Component<IProps, {}> {
  static defaultProps = { 
    Numbers: [1, 2, 3, 4, 5],
    IsDouble: false
  }
  private doubleNumbers: number[] = [];
  private resultString: string = "";

  result() {
    if (this.props.IsDouble!) {
      for (var i:number = 0; i< this.props.Numbers!.length; i++) {
        this.resultString += `<li>${this.props.Numbers![i]}</li>`;
      }
    } 

    if (!this.props.IsDouble!) {
      for (var j:number = 0; j< this.props.Numbers!.length; j++) {
        this.resultString += `<li>${this.props.Numbers![j]*2}</li>`;
      }
    }
  }

  render() {
    this.result()

    return (
      <ul dangerouslySetInnerHTML={{__html: this.resultString}}>
      </ul>
    );
  }
}

export default ls