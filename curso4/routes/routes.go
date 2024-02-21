package routes

import (
	"log"
	"net/http"

	"curso4/controllers"
	"curso4/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandlerRequest() {
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)

	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")

	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.AtualizaPersonalidade).Methods("Put")
	r.HandleFunc("/api/personalidades/{id}", controllers.RemovePersonalidade).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
