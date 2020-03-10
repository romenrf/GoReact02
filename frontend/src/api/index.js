// api/index.js
var socket = new WebSocket("ws://localhost:8085/ws");
//TODO: Esto es una prueba  
let connect = cb => {
  console.log("Conectando...");

  socket.onopen = () => {
    console.log("Conección completada");
  };

  socket.onmessage = msg => {
    console.log("Mensaje: ",msg);
    cb(msg)
  };

  socket.onclose = event => {
    console.log("Socket Conexión Cerrada: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = msg => {
  console.log("Enviando msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };