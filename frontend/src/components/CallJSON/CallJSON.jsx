import React, { Component } from "react";
import "./CallJSON.scss";


/*const ListadoUsuarios = ({usuarios}) => {
  return (
    <div>
      <h1>listado de sqlite3</h1>
      {usuarios.map((usuario) => (
        <div className="card">
          <div className="card-body">
            <h5 className="card-title">{usuario.url}</h5>            
          </div>
        </div>
      ))}
    </div>
  )
}*/


class CallJSON extends Component {
  constructor(){
    super()
    this.state = {usuarios: []}
  }
  componentDidMount () {
    fetch("http://localhost:8085/users")
      .then( data => data.json())
      .then((response) => {
        this.setState({usuarios: response})
        console.log(JSON.stringify(response,null,3))
      })
      .catch(console.log());
  }

  render() {    
    if (this.state.usuarios.length > 0){
      return (
        <div className="container-fluid">
          <p>Llamada realizada con éxito: {this.state.usuarios}</p>
        </div>
      )
    }
    else{
      return (
        <div className="CallJSON">
          <p>No hay valores.</p>
        </div>
      )
    }
  }
}

export default CallJSON;