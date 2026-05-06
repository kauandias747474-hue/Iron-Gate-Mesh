use crate::raft::{RaftState, Command};
use crate::policy_engine::PolicyEngine;
use crate::config::Config;
use tokio::time::{sleep, Duration};

pub struct RaftNode {
    pub config: Config,
    pub state: RaftState,
    pub engine: PolicyEngine,

}

impl RaftNode {
    pub fn new(config: Config, engine: PolicyEngine) -> Self {
        Self {
            config, 
            state: RaftState::new(),
            engine,
        }
    }

    pub async fn run(&mut self) {
        println!("Nó {} iniciado em {}", self.config.node_id, self.config.address);

        loop {


            let fake_detection = "192.168.1.50".to_string();
            
            let cmd = Command::BlockIp(fake_detection.clone());
            self.state.append_command(cmd);

            println!("[RAFT] Consenso atingido para bloquear: {}", fake_detection);
            self.engine.process_decision("BLOCK", &fake_detection);

            sleep(Duration::from_secs(15)).await;
        }
    }
}
