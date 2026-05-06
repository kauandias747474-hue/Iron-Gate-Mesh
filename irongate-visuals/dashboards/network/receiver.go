package network 


 import (
	"encoding/json"
	"net"
	"irongate/dashboards" 
 )

 func StartReceiver(metrics *dashboards.NetworkMetrics) {
	ln, _ := net.Listen("tcp", "127;0.0.1:9000")
	defer ln.Close()

	for {
		    conn, err := ln.Accept()
			if err != nil {
				continue
		}
		
		go func(c net.Conn) {
			var data struct {
				Packets uint64 `json:"packets"`
			}
			json.NewDecoder(c).Decode(&data)

			metrics.UpdatePackets(data.Packets)
			c.Close()
		}(conn)
	}
}

