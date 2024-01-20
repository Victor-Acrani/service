// Package web contains a small web framework extension.
package web

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dimfeld/httptreemux/v5"
)

// A Handler is a type that handles a http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
}
// MY NOTE: Its important to notice that the App struct embed *httptreemux.ContextMux so it can have all of its attributes.

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
	}
}

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) Handle(method string, path string, handler Handler) {

	h := func(w http.ResponseWriter, r *http.Request) {

		// ANY CODE I WANT

		if err := handler(r.Context(), w, r); err != nil {
			fmt.Println(err)
			return
		}

		// ANY CODE I WANT
	}

	a.ContextMux.Handle(method, path, h)
}


// MY NOTE: The handler type defined in this package overwrites the http.HandlerFunc in order to accept the usage of a context.Context.
// This is possible because we can leverage clousures in golang. The Handle function creates a annonymous http.HandlerFunc for adding
// to the httptreemux, but inside it we call our custom Handler so we can use a context.