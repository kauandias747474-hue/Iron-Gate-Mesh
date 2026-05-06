#ifndef SHARED_DEFS_H
#define SHARED_DEFS_H

#include <stdint.h>

typedef struct {

    uint32_t remote_ip;     
    uint16_t remote_port;    
    uint8_t  protocol;       
    uint8_t  action;         

   
    uint32_t connection_id; 
    uint8_t  tcp_flags;     
  
    uint64_t created_at;     
    uint64_t expiry_at;      
    
  
    uint32_t rule_source_id; 
    uint32_t priority;      
    
    uint64_t hit_count;      
    uint64_t total_bytes;    

   
    uint8_t  _reserved[4];   
} block_rule_t;

#endif
