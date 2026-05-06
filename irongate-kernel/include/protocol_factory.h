#ifndef PROTOCOL_FACTORY_H
#define PROTOCOL_FACTORY_H

#define PROTO_ICMP 1
#define PROTO_TCP  6
#define PROTO_UDP  17


static inline int is_malformed_fragment(uint16_t frag_off) {

    return (frag_off & 0x2000) || (frag_off & 0x1FFF);
}


static inline int is_connection_init(uint8_t tcp_flags) {
  
    return (tcp_flags & 0x02) && !(tcp_flags & 0x10);
}


static inline uint16_t bpf_ntohs(uint16_t val) {
    return (val << 8) | (val >> 8);
}

#endif
