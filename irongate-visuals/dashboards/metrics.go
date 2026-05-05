package dashboards

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (

	MaxHistoryPoints = 100

	SamplingRate = 500 * time.Millisecond
)


type MetricPoint struct {
	Timestamp time.Time `json:"ts"`
	Value     uint64    `json:"val"`
}


type NetworkMetrics struct {
	TotalPackets   uint64
	DroppedPackets uint64
	mu             sync.RWMutex
	History        []MetricPoint
	lastUpdate     time.Time
}

func NewMetricCollector() *NetworkMetrics {
	return &NetworkMetrics{
		History: make([]MetricPoint, 0, MaxHistoryPoints),
	}
}


func (m *NetworkMetrics) Run(ctx context.Context) {
	ticker := time.NewTicker(SamplingRate)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[Metrics] Encerrando coleta graciosamente...")
			return
		case <-ticker.C:
			m.pollKernelMaps()
		}
	}
}

func (m *NetworkMetrics) pollKernelMaps() {

	newTotal := atomic.AddUint64(&m.TotalPackets, 250)

	m.mu.Lock()
	defer m.mu.Unlock()


	if len(m.History) >= MaxHistoryPoints {
		m.History = m.History[1:]
	}


	m.History = append(m.History, MetricPoint{
		Timestamp: time.Now(),
		Value:     newTotal,
	})

	m.lastUpdate = time.Now()
}


func (m *NetworkMetrics) Export() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return map[string]interface{}{
		"total":   atomic.LoadUint64(&m.TotalPackets),
		"dropped": atomic.LoadUint64(&m.DroppedPackets),
		"history": m.History,
		"up":      !m.lastUpdate.IsZero() && time.Since(m.lastUpdate).Seconds() < 2,
	}
}
