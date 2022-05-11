package metrics

import (
	"context"
	"sort"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Top K IP.
var clientCounters *ClientCounters

var ClientCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
	// Namespace: plugin.Namespace,
	Name: "client_count_recent",
	Help: "Number of times client accessed recently",
}, []string{"client"})

// Report reports the metrics data associated with request.
func Report(ctx context.Context, zone, rcode string, size int, start time.Time) {
	if clientCounters != nil {
		clientCounters.put(rcode)
	}
}

// WithServer returns the current server handling the request.
func WithServer(ctx context.Context) string {
	// srv := ctx.Value(plugin.ServerCtx{})
	srv := ctx.Value("waf")
	if srv == nil {
		return ""
	}
	return srv.(string)
}

const other = "other"

const (
	slideInterval = time.Second * 15
	slideWindow   = time.Minute * 5
	slotNum       = slideWindow / slideInterval
	topK          = 10
)

type ClientCounters struct {
	counters [slotNum]map[string]uint
	ch       chan string
}

type ClientTuple struct {
	name string
	val  uint
}

func NewClientCounters() *ClientCounters {
	result := &ClientCounters{
		ch: make(chan string),
	}
	for i := range result.counters {
		result.counters[i] = make(map[string]uint)
	}
	return result
}

func (c *ClientCounters) put(client string) {
	c.ch <- client
}

func (c *ClientCounters) generateMetrics() {
	m := make(map[string]uint)
	for _, counter := range c.counters {
		for client, val := range counter {
			if _, ok := m[client]; !ok {
				m[client] = 0
			}
			m[client] += val
		}
	}
	result := make([]*ClientTuple, 0, len(m))
	for name, val := range m {
		tuple := &ClientTuple{
			name: name,
			val:  val,
		}
		result = append(result, tuple)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].val > result[j].val
	})
	ClientCount.Reset()
	end := topK
	if end > len(result) {
		end = len(result)
	}
	for _, t := range result[:end] {
		ClientCount.WithLabelValues(t.name).Set(float64(t.val))
	}
}

func (c *ClientCounters) run() {
	index := 0
	tick := time.Tick(slideInterval)
	for {
		select {
		case <-tick:
			c.generateMetrics()
			index = (index + 1) % int(slotNum)
			c.counters[index] = make(map[string]uint)
		case client := <-c.ch:
			if _, ok := c.counters[index][client]; !ok {
				c.counters[index][client] = 0
			}
			c.counters[index][client]++
		}
	}
}

func InitClientCounters() {
	clientCounters = NewClientCounters()
	go clientCounters.run()
}
