package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
import (
	"./handlers"
	mws "./middlewares"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)


func main() {
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()
	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// We will setup our server so we can serve static assets like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.Handle("/status", handlers.StatusHandler).Methods("GET")
	r.Handle("/products", negroni.New(negroni.HandlerFunc(mws.JwtMiddleware.HandlerWithNext),negroni.Wrap(handlers.ProductsHandler))).Methods("GET")
	r.Handle("/products/{slug}/feedback", negroni.New(negroni.HandlerFunc(mws.JwtMiddleware.HandlerWithNext),negroni.Wrap(handlers.AddFeedbackHandler))).Methods("POST")

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3000", gorillaHandlers.LoggingHandler(os.Stdout, r))
}
