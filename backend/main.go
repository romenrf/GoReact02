package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"github.com/romenrf/websocket"
	_ "github.com/mattn/go-sqlite3"
	"errors"
	"encoding/json"
)

//MODULOS CONEXION CON LA BASE DE DATOS
type User struct{
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Mail string `json:mail`
}

var db *sql.DB

func GetConnection() *sql.DB {
    // Para evitar realizar una nueva conexión en cada llamada a
    // la función GetConnection.
    if db != nil {
        return db
    }
    // Declaramos la variable err para poder usar el operador
    // de asignación “=” en lugar que el de asignación corta,
    // para evitar que cree una nueva variable db en este scope y
    // en su lugar que inicialice la variable db que declaramos a
    // nivel de paquete.
    var err error
    // Conexión a la base de datos
    db, err = sql.Open("sqlite3", "./users.db")
    if err != nil {
        panic(err)
    }
    return db
}


func (itemUser User)crearUser() error {
	db := GetConnection()
	insertUserSQL := `INSERT INTO users (id, name, mail) values (?,?,?)`
	statement, err := db.Prepare(insertUserSQL)
	if err != nil{
		return err
	}

	defer statement.Close()

	resultRow, err := statement.Exec(itemUser.ID,itemUser.Name,itemUser.Mail)
	if err != nil{
		return err
	}

	if itemRow,err := resultRow.RowsAffected(); err != nil || itemRow != 1{
		return errors.New("Error: Se esperaba una fila afectada")
	}

	return nil
	
}

func (itemUser *User)displayUsers() ([]User, error){
	db := GetConnection()
	rows, err := db.Query("SELECT * FROM users ORDER BY name")
	if err != nil{
		return []User{}, err
	}
	defer rows.Close()

	resultUsers := []User{}

	for rows.Next(){
		rows.Scan(
			&itemUser.ID,
			&itemUser.Name,
			&itemUser.Mail)
		resultUsers = append(resultUsers, *itemUser)
	}
	return resultUsers,nil
}

// GetNotesHandler nos permite manejar las peticiones a la ruta
// ‘/notes’ con el método GET vía JSON
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    // Puntero a una estructura de tipo Note.
    newUser := new(User)
    // Solicitando todas las usuarios en la base de datos.
    users, err := newUser.displayUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    // Convirtiendo el slice de usuarios a formato JSON,
    // retorna un []byte y un error.
    resultJSON, err := json.Marshal(users)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Escribiendo el código de respuesta.
    w.WriteHeader(http.StatusOK)
    // Estableciendo el tipo de contenido del cuerpo de la
    // respuesta.
    w.Header().Set("Content-Type", "application/json")
    // Escribiendo la respuesta, es decir nuestro slice de notas
    // en formato JSON.
    w.Write(resultJSON)
}


func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var newUser User
// Tomando el cuerpo de la petición, en formato JSON, y
    // decodificándola e la variable note que acabamos de
    // declarar.
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
// Creamos la nueva nota gracias al método Create.
    err = newUser.crearUser()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}




func handlerUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            GetUsersHandler(w, r)
        case http.MethodPost:
            CreateUserHandler(w, r)
        /*case http.MethodPut:
            UpdateNotesHandler(w, r)
        case http.MethodDelete:
            DeleteNotesHandler(w, r)*/
        default:
            // Caso por defecto en caso de que se realice una
            // petición con un método diferente a los esperados.
            http.Error(w, "Metodo no permitido",
                http.StatusBadRequest)
            return
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


// setupRoutes nos permite manejar la petición a la ruta ‘/users // y pasa el control a la función correspondiente según el método
// de la petición.
//modifico y por defecto intento activar el websocket
func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/users", handlerUsers)    
}


func main() {
	setupRoutes()
	http.ListenAndServe(":8085", nil)
}
