package routes

import (
    "github.com/urfave/negroni"
    "net/http"
)

// Route ...
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandleFunc  http.HandlerFunc
    Middlewares []negroni.Handler
}