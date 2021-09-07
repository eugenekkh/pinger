package monitor

// Metrics is a dumb data point computed from a history of Results.
type Metrics struct {
    Count       int
    Best        float64
    Last        float64
    Worst       float64
    Median      float64
    Mean        float64
    Stddev      float64
    Loss10      float64
    Loss30      float64
    Loss300     float64
}
