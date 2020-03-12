import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import CallJSON from "./components/CallJSON/CallJSON";
import CallsWebsocket from "./components/CallWebsocket/CallWebsocket";


class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      callsWebsocket: [],
    }
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        callsWebsocket: [...this.state.callsWebsocket, msg]
      }))
      console.log(this.state);
    });
  }

  send() {
    console.log("Enviando mensaje: ");
    sendMsg("Petición Websocket");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <div className="container">
          <div className="row">
            <div className="col-sm">
              <CallsWebsocket callsWebsocket={this.state.callsWebsocket} />
              <button onClick={this.send}>Hacer llamada al Websocket</button>
            </div>
            <div className="col-sm">
              <CallJSON />
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default App;
