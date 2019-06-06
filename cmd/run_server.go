package cmd

import (
    "github.com/gorilla/mux"
    "github.com/shangminlee/goauth/log"
)

func RunServer(configFile string) error{
    log.INFO.Print("Run server")

    // init configurations

    // init service

    // start a classic negroni app

    // create a router instance
    _ = mux.NewRouter()

    // add router

    return nil
}