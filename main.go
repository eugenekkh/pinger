package main

import (
    "flag"
)

type ConfigHttp struct {
    Listen string
    Username string
    Password string
}

var configHttp ConfigHttp
var targets Targets

func main() {
    targets = make(Targets)

    flag.Var(&targets, "target", "List of hosts for ping")
    flag.StringVar(&(configHttp.Listen), "listen", "0.0.0.0:9123", "Bind internal http server to address and port. Default: 0.0.0.0:9123")
    flag.StringVar(&(configHttp.Username), "username", "", "Http basic auth username. Default empty")
    flag.StringVar(&(configHttp.Password), "password", "", "Http basic auth password. Default empty")
    flag.Parse()

    StartPinger()
    for host, _ := range targets {
        AddHostForPing(host)
    }

    StartHttpServer()
}
