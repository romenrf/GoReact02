package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/romenrf/websocket"
)

/* Creamos un UPGRADER para leer y escribit en el buffer
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	//Chekeamos el origen de nuestra conexión
	//Esto nos permite hacer peticiones desde nuestro React
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Creamos un READER el cual escucha los nuevos mensajes enviados por nuestro WEBSOCKET
func reader(conn *websocket.Conn) {
	for {
		// Leemos el mensaje
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// Mostramos el mensaje
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}*/

// Creamos nuestro websocket
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server Go listo en: ", r.Host)

	//Hacemos UPGRADE de nuestra conexión
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	//Escuchamos los nuevos mensajes entrantes por el WEBSOCKET
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Servidor Go con WEBSOCKET")
	})*/
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8085", nil)
}
