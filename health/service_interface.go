package health

import (
    "github.com/gorilla/mux"
    "github.com/shangminlee/goauth/util/routes"
)

// ServiceInterface defines exported methods
type ServiceInterface interface {
    // Export methods
    GetRoutes() []routes.Route
    RegisterRoutes(router *mux.Route, prefix string)
    Close()

}