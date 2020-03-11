package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/romenrf/websocket"
	_ "github.com/mattn/go-sqlite3"
)

//MODULOS CONEXION CON LA BASE DE DATOS
func llamarSqLite3() {
	log.Println("Creating sqlite-database.db...")
	sqlfile, err := os.Create("sqlite-database.db")
	if err != nil {
		os.Remove("sqlite-database.db")
		sqlfile2, err2 := os.Create("sqlite-database.db")
		if err2 != nil {
			log.Fatal(err2.Error())
		} else {
			sqlfile = sqlfile2
		}
	}
	sqlfile.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer sqliteDatabase.Close()

	createTable(sqliteDatabase)

	insertUser(sqliteDatabase, "2", "Paco", "francisco@gmail.com")

}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE users (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"mail" TEXT,		
	  );` // SQL Statement for Create Table

	log.Println("Creando tabla users...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Tabla usuarios creada")
}


func insertUser(db *sql.DB, id string, name string, mail string){
	log.Println("Insertando usuario...")
	insertUserSQL := `INSERT INTO users (id, name, mail) values (?,?,?)`

	statement, err := db.Prepare(insertUserSQL)

	injections 
		if err != nil{
			log.Fatalln(err.Error())
		}
		_, err = statement.Exec(id,name,mail)
		if err != nil{
			log.Fatalln(err.Error())
		}
}

func displayUsers(db *sql.DB){
	row, err := db.Query("SELECT * FROM users ORDER BY name")
	if err != nil{
		log.Fatalln(err)
	}
	defer row.Close()

	for row.Next(){
		var id integer
		var name string
		var mail string
		row.Scan(&id,&name,&mail)
		log.Println("Usuario: ",id," ",name," ",mail)
	}
}

//MODULO DE WEBSERVICE
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
