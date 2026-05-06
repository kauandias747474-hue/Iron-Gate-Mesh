use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
pub enum Command {
    BlockIp(String),
    AllowIp(String),
}

pub struct RaftState {
    pub current_term: u64,
    pub voted_for: Option<u64>,
    pub log: Vec<Command>,     
    pub commit_index: usize, 
}

impl RaftState { 
    pub fn new() -> Self {
        Self {
            current_term: 0,
            voted_for: None,
            log: Vec::new(),
            commit_index: 0,
        }
    }

    pub fn append_command(&mut self, cmd: Command) {
        self.log.push(cmd);
        
       
        if !self.log.is_empty() {
            self.commit_index = self.log.len() - 1;
        }
    }
}
