package oauth

import "github.com/gorilla/mux"

type ServiceInterface interface {

    RegisterRouters(rout *mux.Router, prefix string)
    Register() error

}

