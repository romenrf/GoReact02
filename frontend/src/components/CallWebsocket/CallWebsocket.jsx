import React, { Component } from "react";
import "./CallWebsocket.scss";

class CallWebsocket extends Component {
  render() {
    const messages = this.props.callsWebsocket.map((msg, index) => (
      <p key={index}>{msg.data}</p>
    ));

    return (
      <div className="CallWebsocket">
        <h2>Llamada al Websocket</h2>
        {messages}
      </div>
    );
  }
}

export default CallWebsocket;