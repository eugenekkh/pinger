package main

import (
    "crypto/subtle"
    "encoding/json"
    "log"
    "net/http"
)

type Output struct {
    Targets []*Target `json:"data"`
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
    var output Output
    output.Targets = make([]*Target, len(targets))

    i := 0
    for _, target := range targets {
        output.Targets[i] = target
        i++
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(output)
}

func BasicAuth(handler http.HandlerFunc, configHttp ConfigHttp) http.HandlerFunc {

    return func(w http.ResponseWriter, r *http.Request) {

        username, password, ok := r.BasicAuth()

        if (configHttp.Username == "") {
            handler(w, r)
            return
        }

        if !ok || subtle.ConstantTimeCompare([]byte(username), []byte(configHttp.Username)) != 1 || subtle.ConstantTimeCompare([]byte(password), []byte(configHttp.Password)) != 1 {
            w.WriteHeader(401)
            w.Write([]byte("Unauthorised.\n"))

            return
        }

        handler(w, r)
    }
}

func StartHttpServer() {
    http.HandleFunc("/", BasicAuth(DefaultHandler, configHttp))

    err := http.ListenAndServe(configHttp.Listen, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}