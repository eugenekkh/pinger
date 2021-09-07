package main

import (
    "fmt"
    "net"
    "os"
    "time"

    "github.com/digineo/go-ping"
    "monitor"
)

var (
    pingInterval = 1 * time.Second
    pingTimeout = 1 * time.Second
    reportInterval = 1 * time.Second
    size uint = 56
    pinger *ping.Pinger
    mon *monitor.Monitor
    ticker *time.Ticker
)

func AddHostForPing(host string) {
    ipAddr, err := net.ResolveIPAddr("", host)
    if err != nil {
        fmt.Printf("invalid target '%s': %s", host, err)
        return
    }
    mon.AddTarget(host, *ipAddr)
}

func StartPinger() {
    if p, err := ping.New("0.0.0.0", "::"); err != nil {
        fmt.Printf("Unable to bind: %s\nRunning as root?\n", err)
        os.Exit(2)
    } else {
        pinger = p
    }
    pinger.SetPayloadSize(uint16(size))

    mon = monitor.New(pinger, pingInterval, pingTimeout)
    ticker = time.NewTicker(reportInterval)

    go func() {
        for range ticker.C {
            for host, metrics := range mon.Export() {
                targets[host].Count = metrics.Count
                targets[host].Best = float32(metrics.Best)
                targets[host].Last = float32(metrics.Last)
                targets[host].Worst = float32(metrics.Worst)
                targets[host].Median = float32(metrics.Median)
                targets[host].Mean = float32(metrics.Mean)
                targets[host].Stddev = float32(metrics.Stddev)
                targets[host].Loss10 = float32(metrics.Loss10)
                targets[host].Loss30 = float32(metrics.Loss30)
                targets[host].Loss300 = float32(metrics.Loss300)
            }
        }
    }()
}
