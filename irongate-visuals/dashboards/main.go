package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"irongate/dashboards"
	"irongate/topology"
)

type RemediationConfig struct {
	Rules []struct {
		Alert  string   `json:"alert"`
		Path   string   `json:"path"` // Ex: "scripts/build.ps1"
		Args   []string `json:"args"`
		Action string   `json:"action"`
	} `json:"rules"`
}

func main() {
	fmt.Println("---  IronGate Mesh: Iniciando Visuals (Windows Mode) ---")

	
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	
	if _, err := os.Stat("dashboard.json"); err == nil {
		fmt.Println("[Config] Arquivo dashboard.json detectado.")
	}

	
	topo := topology.NewTopologyProvider()
	topo.UpdateNode("self", "127.0.0.1:8080", true)
	fmt.Println("[Topology] Mapa configurado via dashboards/network_map_cfg.json")


	metrics := dashboards.NewMetricCollector()

	go metrics.Run(ctx)


	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				data := metrics.Export()
				
		
				totalPackets := 0
				if val, ok := data["total"].(int); ok {
					totalPackets = val
				}

				fmt.Printf("[%s] PPS: %d | Nós: %d | Status: ONLINE\n", 
					time.Now().Format("15:04:05"), 
					totalPackets, 
					len(topo.GetActiveNodes()))

	
				if totalPackets > 1000000 {
					fmt.Println("[ ALERT] Tráfego crítico detectado!")
					handleWindowsRemediation("HighTrafficAnomaly")
					saveIncidentSnapshot(data)
				}
			}
		}
	}()

	fmt.Println("[Main] Pronto. Pressione Ctrl+C para encerrar.")
	<-ctx.Done()
	
	fmt.Println("\n[Main] Encerrando processos...")
	time.Sleep(500 * time.Millisecond)
}


func handleWindowsRemediation(alertName string) {
	file, err := os.ReadFile("dashboards/alerts/remediation_scripts.json")
	if err != nil {
		log.Printf("[Erro] remediation_scripts.json não encontrado.")
		return
	}

	var config RemediationConfig
	json.Unmarshal(file, &config)

	for _, rule := range config.Rules {
		if rule.Alert == alertName {
			fmt.Printf("[Remediação] Acionando script: %s\n", rule.Path)

			var cmd *exec.Cmd
			
		
			if rule.Path[len(rule.Path)-4:] == ".ps1" {
				fullArgs := append([]string{"-ExecutionPolicy", "Bypass", "-File", rule.Path}, rule.Args...)
				cmd = exec.Command("powershell.exe", fullArgs...)
			} else {

				cmd = exec.Command("cmd.exe", append([]string{"/C", rule.Path}, rule.Args...)...)
			}

			if err := cmd.Start(); err != nil {
				log.Printf("[Erro] Falha ao executar %s: %v", rule.Path, err)
			}
		}
	}
}

func saveIncidentSnapshot(data map[string]interface{}) {
	snapshotPath := "dashboards/snapshots/last_incident.json"
	file, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile(snapshotPath, file, 0644)
	fmt.Printf("[Snapshot] Registro salvo em %s\n", snapshotPath)
}
