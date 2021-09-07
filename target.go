package main

import (
    "fmt"
    "strings"
)

type Target struct {
    Count       int     `json:"count"`
    Best        float32 `json:"best"`
    Last        float32 `json:"last"`
    Worst       float32 `json:"wrost"`
    Median      float32 `json:"median"`
    Mean        float32 `json:"mean"`
    Stddev      float32 `json:"stddev"`
    Loss10      float32 `json:"loss10"`
    Loss30      float32 `json:"loss30"`
    Loss300     float32 `json:"loss300"`
}

type Targets map[string]*Target

func (targets *Targets) String() string {
    data := make([]string, len(*targets))
    i := 0
    for host, _ := range *targets {
        data[i] = fmt.Sprint(host)
        i++
    }

    return strings.Join(data, ", ")
}

func (targets *Targets) Set(host string) error {
    (*targets)[host] = new(Target)

    return nil
}
