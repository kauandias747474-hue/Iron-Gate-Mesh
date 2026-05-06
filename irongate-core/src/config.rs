use serde::Deserialize;
use std::fs;

#[derive(Deserialize, Clone)]
pub struct Config {
    pub node_id: u64,
    pub address: String,
    pub _peers: Vec<String>,
}

impl Config {
  
    pub fn load(path: &str) -> Result<Self, Box<dyn std::error::Error>> {
        
        let content = fs::read_to_string(path)?;
        let config: Config = serde_json::from_str(&content)?;
        Ok(config)
    }
}
