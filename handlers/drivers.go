package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/agomezguru/cti-b2b/auth-service/middlewares"
	"github.com/agomezguru/cti-b2b/auth-service/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Start Web Server in selected PORT */ 
func Drivers() {

	fmt.Println("Entrando a drivers func")
	router := mux.NewRouter()
	
	// All functional endpoints go here

	// Only are two public endpoints because public register is not allowed.
  router.HandleFunc("/api/restful/user/login", middlewares.DBConnected(routers.Login)).Methods("POST")

	// Start private endpoints
	router.HandleFunc("/api/restful/user/register", middlewares.DBConnected(routers.Register)).Methods("POST")
	router.HandleFunc("/api/restful/user/profile", middlewares.DBConnected(middlewares.VerifyJWT(routers.UserProfile))).Methods("GET")
	//router.HandleFunc("/api/restful/dashboard", middlewares.DBConnected(middlewares.VerifyJWT(routers.Dashboard))).Methods("GET")
	
	// try to find env setting 
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // default OS port
	}
	
	// Set access permits to the world 
	handler := cors.AllowAll().Handler(router)

	// Print all available routes
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
    tpl, err1 := route.GetPathTemplate()
    met, err2 := route.GetMethods()
    fmt.Println(tpl, err1, met, err2)
    return nil
	})

	// Init web server
	log.Fatal(http.ListenAndServe(":" + PORT, handler))

}