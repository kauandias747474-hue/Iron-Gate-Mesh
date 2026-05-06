package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"irongate/dashboards" 
)

func main() {

	metrics := dashboards.NewNetworkMetrics()


	go func() {
		listener, err := net.Listen("tcp", "127.0.0.1:9000")
		if err != nil {
			fmt.Println("Erro ao iniciar receptor:", err)
			return
		}
		defer listener.Close()

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			
			go func(c net.Conn) {
				defer c.Close()
				var data struct {
					Packets uint64 `json:"packets"`
				}
				if err := json.NewDecoder(c).Decode(&data); err == nil {
					metrics.UpdatePackets(data.Packets)
				}
			}(conn)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\n[INFO] Encerrando Dashboard...")
		os.Exit(0)
	}()

	dashboards.RunDashboard(metrics)
}
