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


func NewNetworkMetrics() *NetworkMetrics {
	return &NetworkMetrics{
		History: make([]MetricPoint, 0, MaxHistoryPoints),
	}
}

func (m *NetworkMetrics) UpdatePackets(n uint64) {
	atomic.StoreUint64(&m.TotalPackets, n)
	
	m.mu.Lock()
	m.lastUpdate = time.Now()
	m.mu.Unlock()
}


func (m *NetworkMetrics) Run(ctx context.Context) {
	ticker := time.NewTicker(SamplingRate)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n[Metrics] Encerrando coleta graciosamente...")
			return
		case <-ticker.C:
			m.pollKernelMaps()
		}
	}
}

func (m *NetworkMetrics) pollKernelMaps() {
	currentTotal := atomic.LoadUint64(&m.TotalPackets)

	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.History) >= MaxHistoryPoints {
		m.History = m.History[1:]
	}

	m.History = append(m.History, MetricPoint{
		Timestamp: time.Now(),
		Value:     currentTotal,
	})
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

// RunDashboard exibe os dados processados no terminal
func RunDashboard(m *NetworkMetrics) {
	fmt.Println("\n====================================")
	fmt.Println("    IRONGATE - MESH TELEMETRY       ")
	fmt.Println("====================================")
	
	for {
		data := m.Export()
		status := "OFFLINE"
		if data["up"].(bool) {
			status = "ONLINE"
		}

		
		fmt.Printf("\r[NODE-01] Status: %s | Pacotes: %d | Histórico: %d pts", 
			status, data["total"], len(data["history"].([]MetricPoint)))
		
		time.Sleep(200 * time.Millisecond)
	}
}
