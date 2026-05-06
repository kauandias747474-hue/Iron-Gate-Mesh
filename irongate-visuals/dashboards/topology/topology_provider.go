package topology

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type NodeMetadata struct {
	ID       string    `json:"id"`
	Address  string    `json:"addr"`
	LastSeen time.Time `json:"last_seen"`
	IsLeader bool      `json:"is_leader"`
	// Novos campos para bater com seu network_map_cfg.json
	X        int       `json:"x"`
	Y        int       `json:"y"`
	Color    string    `json:"color"`
}

type TopologyProvider struct {
	mu    sync.RWMutex
	Nodes map[string]*NodeMetadata
}

func NewTopologyProvider() *TopologyProvider {
	tp := &TopologyProvider{
		Nodes: make(map[string]*NodeMetadata),
	}
	tp.LoadStaticConfig()
	return tp
}

func (tp *TopologyProvider) LoadStaticConfig() {
	file, err := os.ReadFile("dashboards/network_map_cfg.json")
	if err != nil {
		fmt.Println("[Topology] Aviso: Usando topologia dinâmica (config não encontrada).")
		return
	}

	var config struct {
		Nodes []struct {
			ID    string `json:"id"`
			X     int    `json:"x"`
			Y     int    `json:"y"`
			Color string `json:"color"`
		} `json:"nodes"`
	}

	if err := json.Unmarshal(file, &config); err == nil {
		tp.mu.Lock()
		defer tp.mu.Unlock()
		for _, n := range config.Nodes {
			tp.Nodes[n.ID] = &NodeMetadata{
				ID:    n.ID,
				X:     n.X,
				Y:     n.Y,
				Color: n.Color,
			}
		}
		fmt.Printf("[Topology] %d nós pré-configurados do mapa.\n", len(config.Nodes))
	}
}

func (tp *TopologyProvider) UpdateNode(id, addr string, leader bool) {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	if node, exists := tp.Nodes[id]; exists {
		node.Address = addr
		node.IsLeader = leader
		node.LastSeen = time.Now()
	} else {
		tp.Nodes[id] = &NodeMetadata{
			ID:       id,
			Address:  addr,
			LastSeen: time.Now(),
			IsLeader: leader,
		}
	}
}

func (tp *TopologyProvider) GetActiveNodes() []NodeMetadata {
	tp.mu.RLock()
	defer tp.mu.RUnlock()

	var active []NodeMetadata
	for _, node := range tp.Nodes {
		
		active = append(active, *node)
	}
	return active
}
