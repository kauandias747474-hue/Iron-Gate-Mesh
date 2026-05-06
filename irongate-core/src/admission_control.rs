pub struct AdmissionControl;

impl AdmissionControl {
    pub fn is_authorized(node_id: u64) -> bool {

        node_id > 0
    }
}
