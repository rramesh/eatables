package server

import (
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/rramesh/eatables/data"
	"github.com/rramesh/eatables/handlers"
)

// RESTServer contains logger
type RESTServer struct {
	l      hclog.Logger
	itemDB *data.ItemDB
}

// NewRESTServer creates a new REST server
func NewRESTServer(l hclog.Logger, idb *data.ItemDB) *RESTServer {
	return &RESTServer{l, idb}
}

// Server creates a REST Server instance
func (r *RESTServer) Server(v *data.Validation) *http.Server {
	ih := handlers.NewItems(r.l, v, r.itemDB)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/items", ih.ListAll)
	getRouter.HandleFunc("/items/{id:[0-9]+}", ih.ListSingle)
	getRouter.HandleFunc("/items/sku/{uuid}", ih.ListItemBySKU)
	getRouter.HandleFunc("/items/vendor/{uuid}", ih.ListItemsByVendor)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/items", ih.Create)
	postRouter.Use(ih.MiddlewareValidateItem)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/items", ih.Update)
	putRouter.Use(ih.MiddlewareValidateItem)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/items/{id:[0-9]+}", ih.Delete)

	// API Documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.HandleFunc("/docs", sh.ServeHTTP)
	getRouter.HandleFunc("/swagger.yaml", http.FileServer(http.Dir("./")).ServeHTTP)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := &http.Server{
		// Addr:         *bindAddress,      //configure the bind address
		Handler:      ch(sm),                                             // set the default handler
		ErrorLog:     r.l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		IdleTimeout:  120 * time.Second,                                  // max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,                                    // max time to read request from the client
		WriteTimeout: 1 * time.Second,                                    // max time to write response to the client
	}
	return s
}
