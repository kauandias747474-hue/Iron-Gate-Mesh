pub struct EbpfController;

impl EbpfController {
    pub fn new() -> Self { Self }

    pub fn send_block_command(&self, target_ip: &str) {
        println!(">>> COMANDO ENVIADO AO DATA PLANE: Bloquear {}", target_ip);
    }
}
