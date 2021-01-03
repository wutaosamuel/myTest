import React from "react"

export interface Props {
  name: string;
}

class Welcome extends React.Component<Props, {}> {
  render() {
    return (
      <h1>
        Hello, {this.props.name}
      </h1>
    )
  };
}

export default Welcome