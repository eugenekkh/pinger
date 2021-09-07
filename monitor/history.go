package monitor

import "container/list"

import (
    "math"
    "sort"
    "sync"
    "time"
)

// Result stores the information about a single ping, in particular
// the round-trip time or whether the packet was lost.
type Result struct {
    RTT  time.Duration
    Lost bool
}

// History represents the ping history for a single node/device.
type History struct {
    results *list.List
    sync.RWMutex
}

// NewHistory creates a new History object with a specific capacity
func NewHistory(capacity int) History {
    history := History{}
    history.results = list.New()
    history.results.Init()
    return history
}

// AddResult saves a ping result into the internal history.
func (h *History) AddResult(rtt time.Duration, err error) {
    h.Lock()

    h.results.PushFront(Result{RTT: rtt, Lost: err != nil})
    for h.results.Len() > 301 {
        e := h.results.Back()
        h.results.Remove(e)
    }
    h.Unlock()
}

// Compute aggregates the result history into a single data point.
func (h *History) Compute() *Metrics {
    h.RLock()
    defer h.RUnlock()
    return h.compute()
}

func (h *History) compute() *Metrics {
    numFailure := 0
    µsPerMs := 1.0 / float64(time.Millisecond)

    data := make([]float64, 0, h.results.Len())
    var best, last, worst, mean, stddev, sumSquares, total float64
    var extremeFound bool

    for e := h.results.Front(); e != nil; e = e.Next() {
        curr := e.Value.(Result);
        if curr.Lost {
            numFailure++
        } else {
            rtt := float64(curr.RTT) * µsPerMs
            data = append(data, rtt)

            if !extremeFound || rtt < best {
                best = rtt
            }
            if !extremeFound || rtt > worst {
                worst = rtt
            }
            if (last == 0) {
                last = rtt
            }

            extremeFound = true
            total += rtt
        }
    }

    size := float64(h.results.Len() - numFailure)
    mean = total / size
    for _, rtt := range data {
        sumSquares += math.Pow(rtt - mean, 2)
    }

    stddev = math.Sqrt(sumSquares / size)

	median := math.NaN()
	if l := len(data); l > 0 {
		sort.Float64Slice(data).Sort()
		if l%2 == 0 {
			median = (data[l/2-1] + data[l/2]) / 2
		} else {
			median = data[l/2]
		}
	}

    return &Metrics{
        Count:       h.results.Len(),
        Best:        best,
        Last:        last,
        Worst:       worst,
        Median:      median,
        Mean:        mean,
        Stddev:      stddev,
        Loss10:      h.computeLoss(10),
        Loss30:      h.computeLoss(30),
        Loss300:     h.computeLoss(300),
    }
}

func (h *History) computeLoss(count int) float64 {
    if (h.results.Len() == 0) {
        return 0.0
    }

    i := 0
    var lost float64 = 0.0
    for e := h.results.Front(); e != nil; e = e.Next() {
        result := e.Value.(Result);
        if (result.Lost) {
            lost++;
        }
        if (i >= count) {
            break
        }
        i++
    }

    if (count > h.results.Len()) {
        return (lost / float64(h.results.Len())) * 100
    }

    return (lost / float64(count)) * 100
}
