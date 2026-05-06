mod config;
mod audit_log;
mod ebpf_manager;
mod policy_engine;
mod raft;
mod raft_node;
mod admission_control;
mod network_client;  

use tokio::signal;
use std::sync::Arc;
use std::time::Duration;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let logger = Arc::new(audit_log::AuditLogger::new());
    logger.log("Iniciando IronGate Mesh Engine...");

  
    let cfg = config::Config::load("config.json").unwrap_or_else(|err| {
        eprintln!("Erro crítico ao carregar config.json: {}", err);
        std::process::exit(1);
    });

    let ebpf_controller = Arc::new(ebpf_manager::EbpfController::new());
    let engine = policy_engine::PolicyEngine::new(Arc::clone(&ebpf_controller));
    let net_client = network_client::NetworkClient::new("127.0.0.1:9000");

 
    if !admission_control::AdmissionControl::is_authorized(cfg.node_id) {
        logger.log("Acesso negado: ID do nó não autorizado.");
        return Ok(());
    }

 
    let mut node = raft_node::RaftNode::new(cfg.clone(), engine);
    tokio::spawn(async move {
        node.run().await;
    });

    logger.log(&format!("Mesh operacional em {}. Enviando telemetria...", cfg.address));

   
    let mut fake_packet_count = 0;
    let telemetry_task = tokio::spawn(async move {
        loop {
            fake_packet_count += 10; // Simula pacotes capturados pelo eBPF
            net_client.send_metrics(fake_packet_count);
            tokio::time::sleep(Duration::from_millis(500)).await;
        }
    });

    match signal::ctrl_c().await {
        Ok(()) => {
            logger.log("Desligando IronGate Core...");
            telemetry_task.abort();
        }
        Err(err) => eprintln!("Erro no sinal do sistema: {}", err),
    }

    Ok(())
}
