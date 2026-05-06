use std::io::Write;
use std::net::TcpStream;
use serde_json::json;

pub struct NetworkClient {
    address: String,
 }

 impl NetworkClient {
    pub fn new(addr: &str) -> Self {
        Self { address: addr.to_string() }
    }

    pub fn send_metrics(&self, packets: u64) {
       
        if let Ok(mut stream) = TcpStream::connect(&self.address) {
            let data = json!({
                "packets": packets,
                "status": "active"
            });
            let _ = stream.write_all(data.to_string().as_bytes());
        }
    }
}
