import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import CallJSON from "./components/CallJSON/CallJSON";


class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      callsJSON: []
    }
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        callsJSON: [...this.state.callsJSON, msg]
      }))
      console.log(this.state);
    });
  }

  send() {
    console.log("Enviando mensaje: ");
    sendMsg("Petición");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <CallJSON callsJSON={this.state.callsJSON} />
        <button onClick={this.send}>Hacer llamada al JSON</button>
      </div>
    );
  }
}

export default App;
