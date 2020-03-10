import React, { Component } from "react";
import "./CallJSON.scss";

class CallJSON extends Component {
  render() {
    const messages = this.props.callsJSON.map((msg, index) => (
      <p key={index}>{msg.data}</p>
    ));

    return (
      <div className="CallJSON">
        <h2>Llamada al JSON</h2>
        {messages}
      </div>
    );
  }
}

export default CallJSON;