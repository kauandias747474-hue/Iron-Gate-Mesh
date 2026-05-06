#include "maps.h"
#include "protocol_factory.h"
#include "shared_defs.h"


SEC("xdp")
int irongate_filter_main(void* ctx) {
   
    uint32_t src_ip = 0x0100007F; 


    block_rule_t *rule = bpf_map_lookup_elem(&BLOCKLIST_MAP, &src_ip);

    if (rule) {
       
        uint64_t now = bpf_ktime_get_ns();
        if (rule->expiry_at != 0 && now > rule->expiry_at) {
            return 1; // Passa o pacote (Regra expirada)
        }

     
        __sync_fetch_and_add(&rule->hit_count, 1);
        
     
        alert_event_t event = { .src_ip = src_ip, .reason = 0 };
        bpf_ringbuf_output(&EVENTS_RINGBUF, &event, sizeof(event), 0);

        return 2; 
    }

    return 1; // Permite o pacote (XDP_PASS)
}
