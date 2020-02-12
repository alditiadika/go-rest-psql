package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	//lib/pq
	_ "github.com/lib/pq"

	"github.com/alditiadika/go-rest-psql/app/handle"
	"github.com/alditiadika/go-rest-psql/config"
)

// App struct
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize App
func (a *App) Initialize() {
	a.connectDB()
	a.Router = mux.NewRouter()
	a.setRouters()
}

//Run app
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) connectDB() {
	conf := config.GetConf()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error when calling database!")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error when ping database!")
		panic(err)
	}
	a.DB = db
	fmt.Println("Connected to database!")
}

func (a *App) request(method string, path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(method)
}

func (a *App) handleRequest(handlerF func(res http.ResponseWriter, req *http.Request, db *sql.DB)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		handlerF(res, req, a.DB)
	}
}

func (a *App) setRouters() {
	a.request("GET", "/users", a.handleRequest(handle.GetUserByParam))
}
