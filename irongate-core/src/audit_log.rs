pub struct AuditLogger;

impl AuditLogger {
    pub fn new() -> Self { Self }
    pub fn log(&self, msg: &str) {
        println!("[AUDIT - {}]: {}", chrono::Local::now(), msg);
    }
}
