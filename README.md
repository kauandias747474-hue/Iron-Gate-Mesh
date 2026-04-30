# Iron Gate Mesh

[![Rust](https://img.shields.io/badge/Rust-000000?style=for-the-badge&logo=rust&logoColor=white)](https://www.rust-lang.org/)
[![C](https://img.shields.io/badge/C-A8B9CC?style=for-the-badge&logo=c&logoColor=white)](https://learn.microsoft.com/cpp/c-language/)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![eBPF](https://img.shields.io/badge/eBPF-185697?style=for-the-badge&logo=linux&logoColor=white)](https://ebpf.io/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)

## 🇧🇷 Português: Documentação Técnica de Engenharia

O **IronGate Mesh** é uma malha de serviço distribuída de próxima geração, projetada para fornecer segurança e observabilidade em ambientes de missão crítica. Diferente de proxies convencionais que operam inteiramente no espaço de usuário (User Space), o IronGate utiliza um modelo híbrido para maximizar a taxa de transferência de dados e minimizar a superfície de ataque.

### 🔬 O Problema da Latência de Rede Tradicional
Em arquiteturas padrão, cada pacote atravessa múltiplas camadas de abstração do kernel antes de chegar a uma aplicação de segurança. Isso envolve:
1. Interrupções de hardware.
2. Alocação de memória no Kernel (`sk_buff`).
3. Troca de contexto (Context Switch) para o User Space.
4. Processamento da lógica de firewall.

O IronGate resolve isso utilizando **eBPF/XDP**, interceptando o pacote no estágio 1, permitindo que a decisão de segurança seja tomada em nanossegundos, antes mesmo do Kernel alocar memória pesada para o pacote.

### 🏗️ Arquitetura do Ecossistema

#### 1. Plano de Dados (Data Plane - C/eBPF)
Escrito em C restrito e compilado para bytecode eBPF.
- **XDP Hook:** O programa é anexado ao ponto mais baixo da pilha de rede.
- **Stateless Filtering:** Avaliação de cabeçalhos IP/TCP/UDP em tempo real.
- **Shared Maps:** Utiliza tabelas de hash atômicas para receber listas de IPs bloqueados vindas do plano de controle sem interrupção de serviço.

#### 2. Plano de Controle (Control Plane - Rust)
O núcleo inteligente que orquestra a malha.
- **Consenso com Raft:** Implementa uma máquina de estados distribuída. Se um administrador altera uma política em um nó, o algoritmo Raft garante que todos os outros nós confirmem e apliquem a mudança, mantendo a consistência eventual e rigorosa do cluster.
- **Segurança Pragmática:** A escolha do Rust elimina erros de *Segmentation Fault* e *Memory Leaks*, críticos em sistemas que rodam 24/7.

#### 3. Camada de Observabilidade (Visuals - Go)
Focada em alta concorrência e telemetria.
- **Scraping de Mapas:** O agente Go monitora os contadores de performance dentro do Kernel.
- **Exposição de Dados:** Transforma eventos binários complexos em métricas estruturadas, facilitando a integração com pipelines de dados e visualização executiva.

###  Conceitos Avançados Aplicados
- **Eleição de Líder:** Mecanismo automático para resiliência do cluster.
- **Bytecode Verification:** Segurança garantida pelo verificador JIT do Kernel.
- **Lock-Free Data Structures:** Uso de mapas eBPF para evitar travas (locks) que reduziriam a performance.

---

## 🇺🇸 English: Engineering Technical Documentation

**IronGate Mesh** is a next-generation distributed service mesh designed to provide security and observability in mission-critical environments. Unlike conventional proxies that operate entirely in User Space, IronGate utilizes a hybrid model to maximize throughput and minimize the attack surface.

### 🔬 The Traditional Network Latency Problem
In standard architectures, every packet traverses multiple kernel abstraction layers before reaching a security application. This involves:
1. Hardware interrupts.
2. Kernel memory allocation (`sk_buff`).
3. Context switching to User Space.
4. Firewall logic processing.

IronGate solves this by using **eBPF/XDP**, intercepting the packet at stage 1, allowing the security decision to be made in nanoseconds—before the Kernel even allocates heavy memory for the packet.

### 🏗️ Ecosystem Architecture

#### 1. Data Plane (Kernel Space - C/eBPF)
Written in restricted C and compiled to eBPF bytecode.
- **XDP Hook:** The program is attached to the lowest point of the network stack.
- **Stateless Filtering:** Real-time IP/TCP/UDP header evaluation.
- **Shared Maps:** Uses atomic hash tables to receive blocklists from the control plane without service interruption.

#### 2. Control Plane (User Space - Rust)
The intelligent core orchestrating the mesh.
- **Raft Consensus:** Implements a distributed state machine. If an admin changes a policy on one node, the Raft algorithm ensures all other nodes confirm and apply the change, maintaining strict and eventual cluster consistency.
- **Pragmatic Safety:** Rust was chosen to eliminate *Segmentation Faults* and *Memory Leaks*, which are critical for systems running 24/7.

#### 3. Observability Layer (Visuals - Go)
Focused on high concurrency and telemetry.
- **Map Scraping:** The Go agent monitors performance counters inside the Kernel.
- **Data Exposure:** Converts complex binary events into structured metrics, facilitating integration with data pipelines and executive dashboards.

###  Advanced Engineering Concepts
- **Leader Election:** Automatic mechanism for cluster resilience.
- **Bytecode Verification:** Security guaranteed by the Kernel's JIT verifier.
- **Lock-Free Data Structures:** Leveraging eBPF maps to avoid performance-degrading locks during high-traffic scenarios.
- **Zero-Copy Architecture:** Maximizing packet processing efficiency by avoiding memory duplication.

## 📁 Repository Structure

| Module | Language | Responsibility |
| :--- | :--- | :--- |
| **irongate-kernel** | C / eBPF | Packet interception and filtering (XDP) |
| **irongate-core** | Rust | Cluster consensus and state management |
| **irongate-bridge** | Rust (Aya) | Kernel/Userspace communication bridge |
| **irongate-visuals** | Go | Telemetry, API, and Metrics export |
| **deploy** | Docker | Environment orchestration and scaling |
