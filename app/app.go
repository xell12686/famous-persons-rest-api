package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/famous-persons-rest-api/app/handler"
	"github.com/famous-persons-rest-api/app/model"
	"github.com/famous-persons-rest-api/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/persons", a.GetAllPersons)
	a.Post("/persons", a.CreatePerson)
	a.Get("/persons/{name}", a.GetPerson)
	a.Put("/persons/{name}", a.UpdatePerson)
	a.Delete("/persons/{name}", a.DeletePerson)
	a.Put("/persons/{name}/disable", a.DisablePerson)
	a.Put("/persons/{name}/enable", a.EnablePerson)
}

// Get : Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post : Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put : Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete : Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// GetAllPersons : Handlers to manage Person Data
func (a *App) GetAllPersons(w http.ResponseWriter, r *http.Request) {
	handler.GetAllPersons(a.DB, w, r)
}

// CreatePerson ...
func (a *App) CreatePerson(w http.ResponseWriter, r *http.Request) {
	handler.CreatePerson(a.DB, w, r)
}

// GetPerson ...
func (a *App) GetPerson(w http.ResponseWriter, r *http.Request) {
	handler.GetPerson(a.DB, w, r)
}

// UpdatePerson ...
func (a *App) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	handler.UpdatePerson(a.DB, w, r)
}

// DeletePerson ...
func (a *App) DeletePerson(w http.ResponseWriter, r *http.Request) {
	handler.DeletePerson(a.DB, w, r)
}

// DisablePerson ...
func (a *App) DisablePerson(w http.ResponseWriter, r *http.Request) {
	handler.DisablePerson(a.DB, w, r)
}

// EnablePerson ...
func (a *App) EnablePerson(w http.ResponseWriter, r *http.Request) {
	handler.EnablePerson(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(host, handlers.CORS(headers, methods, origins)(a.Router)))
}
