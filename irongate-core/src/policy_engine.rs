use std::sync::Arc;
use crate::ebpf_manager::EbpfController;

pub struct PolicyEngine {
    controller: Arc<EbpfController>,
}

impl PolicyEngine {

    pub fn new(controller: Arc<EbpfController>) -> Self {
        Self { controller }
    }

    pub fn process_decision(&self, decision: &str, target: &str) {
        println!("[PolicyEngine] Comando recebido: {} para o alvo: {}", decision, target);
     
        // if decision == "BLOCK" { self.controller.block_ip(target); }
    }
}
