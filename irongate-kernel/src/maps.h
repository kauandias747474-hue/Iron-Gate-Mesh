#ifndef SRC_MAPS_H
#define SRC_MAPS_H

#include "shared_defs.h"


#define SEC(NAME) __attribute__((section(NAME), used))


struct {
    uint32_t type;       
    uint32_t key_size;     
    uint32_t value_size;   
    uint32_t max_entries;  
} BLOCKLIST_MAP SEC(".maps") = {
    .type = 1, 
    .key_size = sizeof(uint32_t),
    .value_size = sizeof(block_rule_t),
    .max_entries = 10000,
};


/ 3. Ring Buffer de Eventos (Alertas em Tempo Real)

struct {
    uint32_t type;
    uint32_t max_entries;
} EVENTS_RINGBUF SEC(".maps") = {
    .type = 27, 
    .max_entries = 64 * 1024,
};

#endif
