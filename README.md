# Iron Gate Mesh 

[![Rust](https://img.shields.io/badge/Rust-000000?style=for-the-badge&logo=rust&logoColor=white)](https://www.rust-lang.org/)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![C](https://img.shields.io/badge/C-A8B9CC?style=for-the-badge&logo=c&logoColor=white)](https://learn.microsoft.com/cpp/c-language/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)

---

## 🇧🇷 Português: Documentação Técnica de Engenharia

O **IronGate Mesh** é uma malha de serviço distribuída projetada para segurança e observabilidade de alta performance. Esta versão opera através de um modelo híbrido: processamento de rede otimizado em **User Space** para facilidade de deploy, e um plano de dados em **C** para futuras extensões em nível de Kernel.

### 🏗️ Arquitetura do Ecossistema

#### 1. Plano de Controle (Control Plane - `irongate-core`) [Rust]
O cérebro distribuído da malha, focado em consistência rigorosa.
* **Consenso Raft**: Implementação de máquina de estados distribuída para garantir que políticas de segurança sejam idênticas em todo o cluster.
* **Segurança Pragmática**: Uso intensivo de Rust para eliminar vazamentos de memória e *race conditions* em operações 24/7.
* **Gestão de Auditoria**: Registro centralizado de eventos e controle de admissão de novos nós.

#### 2. Plano de Dados e Kernel (`irongate-kernel` & `irongate-network`) [C / Go]
Responsável pela ingestão e filtragem de pacotes.
* **Filtragem de Software**: Avaliação de pacotes em tempo real utilizando goroutines paralelas e operações atômicas.
* **Otimização C**: Estrutura preparada para filtragem *stateless* e balanceamento de carga de baixo nível.
* **Políticas Atômicas**: Atualização imediata de blacklists sem necessidade de reiniciar o serviço.

#### 3. Camada de Observabilidade (`irongate-visuals`) [Go]
* **Telemetria**: Extração de métricas de performance diretamente do processamento de rede.
* **Dashboard API**: Exposição de dados estruturados sobre tráfego, pacotes descartados e integridade dos nós.

---

## 🇺🇸 English: Engineering Technical Documentation

**IronGate Mesh** is a distributed service mesh designed for high-performance security and observability. This version operates through a hybrid model: optimized **User Space** network processing for deployment flexibility, and a **C** data plane for future Kernel-level extensions.

### 🏗️ Ecosystem Architecture

#### 1. Control Plane (`irongate-core`) [Rust]
The distributed intelligence of the mesh, focused on strict consistency.
* **Raft Consensus**: Distributed state machine implementation to ensure security policies remain identical across the cluster.
* **Pragmatic Safety**: Leverages Rust to eliminate memory leaks and race conditions in 24/7 operations.
* **Audit Management**: Centralized event logging and node admission control.

#### 2. Data Plane & Kernel (`irongate-kernel` & `irongate-network`) [C / Go]
Responsible for packet ingestion and filtering.
* **Software Filtering**: Real-time packet evaluation using parallel goroutines and atomic operations.
* **C Optimization**: Architecture prepared for stateless filtering and low-level load balancing.
* **Atomic Policies**: Immediate blacklist updates without service interruption.

#### 3. Observability Layer (`irongate-visuals`) [Go]
* **Telemetry**: Performance metrics extraction directly from network processing.
* **Dashboard API**: Exposure of structured data regarding traffic, dropped packets, and node health.th.

---

**🇧🇷 Versão em Português (Original)**

# Deep Engineering Analysis: A "Death Star" de Rede (Iron Gate Mesh)

## 1. Camada de Hardware e Baixo Nível (Performance Pura)
*   **XDP (eXpress Data Path):** Ponto de entrada ultra-rápido que processa pacotes diretamente no driver da placa de rede.
*   **Per-CPU Maps:** Mapas de memória exclusivos para cada núcleo da CPU, eliminando a disputa por travas (lock contention).
*   **Data Structure Alignment (Padding):** Inserção de bytes vazios em structs para que tenham múltiplos de 8 bytes, otimizando o acesso ao cache L1/L2.
*   **Network Byte Order (Endianness):** Tradução de dados entre o padrão da rede (Big-Endian) e o da CPU (Little-Endian) via `bpf_ntohs`.
*   **Cache Locality:** Organização de dados para que a CPU encontre as informações na memória ultra-rápida interna, evitando idas lentas à RAM.

## 2. Camada de Networking e Segurança (Filtros e Rotas)
*   **Early Drop:** Descarte imediato de pacotes malformados ou fragmentados para proteger o Kernel de ataques DoS.
*   **Stateful Inspection:** Capacidade de entender o estado da conexão (flags TCP) em vez de apenas olhar pacotes isolados.
*   **Connection Init:** Monitoramento de handshakes para detectar padrões de ataque como Port Scanning ou SYN Flood.
*   **Flow ID (Flow Hashing):** Geração de um identificador único para cada conversa (5-tuple), permitindo o uso de um "Caminho Rápido" (Fast Path).
*   **DNAT (Destination NAT):** Redirecionamento de pacotes para IPs internos no nível do Kernel, transformando o firewall em um Load Balancer.
*   **Checksum Recalculation:** Ajuste matemático obrigatório no cabeçalho do pacote após qualquer alteração de IP ou porta para manter a integridade.

## 3. Camada de Arquitetura e Memória (Estabilidade)
*   **TTL (Time-To-Live / expiry_at):** Selo de validade em cada regra, permitindo que o Kernel ignore registros expirados sem ajuda externa.
*   **Kernel-side Garbage Collection:** Estratégia passiva de limpeza onde o Kernel invalida regras antigas, poupando CPU do plano de controle.
*   **Ring Buffer (EVENTS_RINGBUF):** Fila circular para envio de eventos do Kernel para o User-space.
*   **Zero-Copy Telemetry:** Leitura direta da memória compartilhada entre Kernel e Go, sem cópia de buffers, garantindo monitoramento em tempo real.
*   **Tail Calls:** Técnica para encadear programas eBPF modulares, superando o limite de complexidade de um único programa sem perder performance.

## 4. Camada de Sistemas Distribuídos (O Mesh)
*   **Separação de Planos (Control Plane vs. Data Plane):** O Rust/Go decide as políticas (Cérebro), o C/Kernel executa as ordens (Músculo).
*   **Consenso Raft:** Algoritmo que garante que todos os nós concordem sobre quem é o líder e quais são as regras de bloqueio atuais.
*   **Rule Source ID:** Identificador de autoridade que evita o conflito de regras e o estado de Split-Brain (onde nós diferentes tomam decisões contraditórias).
*   **Distributed Policy Enforcement:** A aplicação sincronizada de uma regra em todos os nós do cluster ao mesmo tempo.

## 📈 Resumo de Impacto
| Categoria | Principal Benefício | Conceito "Chave" |
| :--- | :--- | :--- |
| **Hardware** | Latência Mínima | Padding / Per-CPU |
| **Networking** | Segurança Proativa | Early Drop / Flow ID |
| **Memória** | Estabilidade do SO | TTL / Ring Buffer |
| **Distribuído** | Consistência Global | Raft / Source ID |

## 🛠️ Log de "Cicatrizes": Erros Identificados e Tratados

### Erros de Estrutura e Configuração
*   **Dependências no Workspace:** O erro "this virtual manifest specifies a dependencies section" ocorreu porque o Cargo.toml da raiz não pode baixar bibliotecas; ele é apenas um "gerente" de pastas.
*   **Conflito de SO:** A tentativa de rodar o código da Aya (Linux) no Windows sem as flags `#[cfg(target_os = "linux")]` causou falhas ao buscar cabeçalhos inexistentes.
*   **Cargo.toml Duplicado:** Ajustamos a estrutura para que o arquivo da raiz contenha apenas a lista de membros do workspace, delegando dependências aos subprojetos.
*   **irongate-kernel (C vs Rust):** Identificamos que o Rust não compila arquivos .c diretamente, exigindo Clang para gerar o objeto usado pelo loader.rs.
*   **Arquivos Soltos:** Corrigimos a localização do loader.rs para o irongate-core, centralizando a execução e deixando o bridge apenas para contratos.

### Erros no Assembly e eBPF
*   **O Erro de "Target" (unknown target bpfel):** O Clang para Windows muitas vezes não inclui suporte nativo ao eBPF. A solução foi migrar para ambientes com suporte LLVM completo.
*   **O Erro de Instrução (no such instruction):** Ocorreu ao tentar ler instruções exclusivas eBPF (como `ldxdw`) como se fossem instruções Intel/AMD (x86).
*   **O Erro de Permissão (sudo: not found):** No ambiente WSL/Docker logado como root, o comando sudo é redundante e muitas vezes inexistente.
  
    **Por que removemos o Assembly direto?**

*A remoção do Assembly manual foi estratégica: o eBPF utiliza um conjunto de instruções (ISA) próprio e extremamente rigoroso que é verificado pelo Kernel antes da execução. Ao delegar a geração desse código ao Clang/LLVM a partir de C ou Rust, eliminamos o risco de instruções inválidas para a arquitetura alvo e garantimos que o Verificador do Kernel aceite o programa sem rejeições por violações de memória ou segurança.*

### Erros de Sintaxe e Runtime (Rust)
*   **Double Colon (::):** Correção do uso de `:` (atribuição de tipo) por `::` (separador de módulos) em `fs::read_to_string`.
*   **Mismatched Types (E0308):** Ajuste em funções que retornavam Result quando era esperado `()`, harmonizando o `tokio::spawn`.
*   **Campos de Struct:** Adição de vírgulas faltantes em structs no `raft.rs`.
*   **Case Sensitivity:** Correção de chamadas como `ebpf_Controller` (maiúsculo) quando definido como `ebpf_controller`.
*   **Os Error 2:** Erro de runtime onde o `config.json` não era encontrado no disco apesar da compilação perfeita.

### Erros no Go (Visuals)
*   **Pacotes na Raiz:** Erro "package ... is not in std" resolvido ao definir corretamente o módulo Go, evitando conflitos com a biblioteca padrão.
*   **Unused Imports:** Limpeza de imports como "encoding/json" exigida pelo rigor do compilador Go.
*   **Redeclared in this block:** Correção de structs e funções repetidas em arquivos diferentes dentro da mesma pasta (Go trata como escopo único).

## 🎓 Conceitos de Engenharia Aprendidos na Prática
*   **Ownership & Arcs:** Uso de "Ponteiros Inteligentes" (Arc) para compartilhar o EbpfController sem que o Rust destrua a variável prematuramente.
*   **Async Runtime (Tokio):** Uso de `spawn` para rodar tarefas em segundo plano (background), permitindo processamento de rede e escuta de comandos simultâneos.
*   **Modularização:** Divisão de lógica em `mod` e `pub struct` para garantir escalabilidade.
*   **Iron Gate Mesh:** Cada erro resolvido tornou-se uma camada de proteção.
  

**🇺🇸 English Version (Translated)**
# Deep Engineering Analysis: The Network "Death Star" (Iron Gate Mesh)

## 1. Hardware & Low-Level Layer (Pure Performance)
*   **XDP (eXpress Data Path):** Ultra-fast entry point processing packets directly at the NIC driver level.
*   **Per-CPU Maps:** Exclusive memory maps for each CPU core, eliminating lock contention.
*   **Data Structure Alignment (Padding):** Inserting empty bytes into structs to align them to 8-byte multiples, optimizing L1/L2 cache access.
*   **Network Byte Order (Endianness):** Data translation between Network standard (Big-Endian) and CPU standard (Little-Endian) via `bpf_ntohs`.
*   **Cache Locality:** Organizing data so the CPU finds information in internal ultra-fast memory, avoiding slow RAM fetches.

## 2. Networking & Security Layer (Filters & Routes)
*   **Early Drop:** Immediate discarding of malformed or fragmented packets to protect the Kernel from DoS attacks.
*   **Stateful Inspection:** Ability to understand connection states (TCP flags) rather than just looking at isolated packets.
*   **Connection Init:** Monitoring handshakes to detect attack patterns like Port Scanning or SYN Flooding.
*   **Flow ID (Flow Hashing):** Generating a unique identifier for each conversation (5-tuple), enabling a "Fast Path."
*   **DNAT (Destination NAT):** Kernel-level packet redirection to internal IPs, turning the firewall into a Load Balancer.
*   **Checksum Recalculation:** Mandatory mathematical adjustment of the packet header after any IP or Port modification to maintain integrity.

## 3. Architecture & Memory Layer (Stability)
*   **TTL (Time-To-Live / expiry_at):** Validity stamp on each rule, allowing the Kernel to ignore expired entries without external help.
*   **Kernel-side Garbage Collection:** Passive cleanup strategy where the Kernel invalidates old rules, saving Control Plane CPU.
*   **Ring Buffer (EVENTS_RINGBUF):** Circular queue for sending events from Kernel-space to User-space.
*   **Zero-Copy Telemetry:** Direct reading of shared memory between Kernel and Go, without buffer copying, ensuring real-time monitoring.
*   **Tail Calls:** Technique for chaining modular eBPF programs, overcoming single-program complexity limits without losing performance.

## 4. Distributed Systems Layer (The Mesh)
*   **Plane Separation (Control Plane vs. Data Plane):** Rust/Go decides the policies (The Brain), C/Kernel executes the orders (The Muscle).
*   **Raft Consensus:** Algorithm ensuring all nodes agree on who the leader is and what the current blocking rules are.
*   **Rule Source ID:** Authority identifier that prevents rule conflicts and Split-Brain states (where different nodes make contradictory decisions).
*   **Distributed Policy Enforcement:** Synchronized application of a rule across all cluster nodes simultaneously.

## 📈 Impact Summary
| Category | Main Benefit | Key Concept |
| :--- | :--- | :--- |
| **Hardware** | Minimum Latency | Padding / Per-CPU |
| **Networking** | Proactive Security | Early Drop / Flow ID |
| **Memory** | OS Stability | TTL / Ring Buffer |
| **Distributed** | Global Consistency | Raft / Source ID |

## 🛠️ Log of "Scars": Identified and Resolved Errors

### Structural and Configuration Errors
*   **Workspace Dependencies:** The "this virtual manifest specifies a dependencies section" error occurred because the root Cargo.toml cannot download libraries; it is only a folder manager.
*   **OS Conflict:** Attempting to run Aya (Linux) code on Windows without `#[cfg(target_os = "linux")]` flags caused failures when searching for non-existent headers.
*   **Duplicate Cargo.toml:** Adjusted the structure so the root file only contains the workspace member list, delegating dependencies to subprojects.
*   **irongate-kernel (C vs Rust):** Identified that Rust does not compile .c files directly, requiring Clang to generate the object used by `loader.rs`.
*   **Loose Files:** Fixed `loader.rs` location to `irongate-core`, centralizing execution and leaving `bridge` strictly for contracts.

### Assembly and eBPF Errors
*   **Target Error (unknown target bpfel):** Clang on Windows often lacks native eBPF support. Solution: Migrated to environments with full LLVM support.
*   **Instruction Error (no such instruction):** Occurred when trying to read eBPF-exclusive instructions (like `ldxdw`) as if they were Intel/AMD (x86) instructions.
*   **Permission Error (sudo: not found):** In WSL/Docker environments logged as root, the `sudo` command is redundant and often missing.
  
 **Why did we remove direct Assembly?**

*The removal of manual Assembly was strategic: eBPF uses its own strictly regulated Instruction Set Architecture (ISA) that must pass a Kernel Verifier before execution. By delegating code generation to Clang/LLVM from C or Rust, we eliminated the risk of invalid instructions for the target architecture and ensured the Kernel Verifier accepts the program without rejection due to memory or safety violations.*

### Syntax and Runtime Errors (Rust)
*   **Double Colon (::):** Corrected use of `:` (type assignment) to `::` (module separator) in `fs::read_to_string`.
*   **Mismatched Types (E0308):** Adjusted functions returning `Result` when `()` was expected, harmonizing `tokio::spawn`.
*   **Struct Fields:** Added missing commas in structs within `raft.rs`.
*   **Case Sensitivity:** Corrected calls like `ebpf_Controller` (uppercase) when defined as `ebpf_controller`.
*   **Os Error 2:** Runtime error where `config.json` was not found on disk despite perfect compilation.

### Go Errors (Visuals)
*   **Root Packages:** "package ... is not in std" error resolved by correctly defining the Go module, avoiding standard library conflicts.
*   **Unused Imports:** Cleaned up imports like "encoding/json" as required by the strict Go compiler.
*   **Redeclared in this block:** Corrected repeated structs and functions across different files in the same folder (Go treats the folder as a single scope).

## 🎓 Engineering Concepts Learned in Practice
*   **Ownership & Arcs:** Using "Smart Pointers" (Arc) to share the `EbpfController` without Rust destroying the variable prematurely.
*   **Async Runtime (Tokio):** Using `spawn` to run background tasks, allowing simultaneous network processing and command listening.
*   **Modularization:** Dividing logic into `mod` and `pub struct` to ensure scalability.
*   **Iron Gate Mesh:** Every resolved error became a layer of protection.


## 🇧🇷 Português: Descrição dos Módulos

* **`irongate-core`**: Responsável pela inteligência. Garante que todos os nós sigam a mesma regra através do protocolo **Raft**.
* **`irongate-kernel`**: Onde o "trabalho pesado" acontece. Escrito em **C** para rodar o mais próximo possível do hardware.
* **`irongate-network`**: Gerencia o fluxo de pacotes que não foram descartados pelo kernel, usando a concorrência do **Go**.
* **`irongate-visuals`**: Transforma os dados brutos de rede em gráficos e métricas compreensíveis.
* **Nota de Engenharia**: Esta estrutura segue o padrão de **Monorepo com Workspace**, permitindo que os componentes em Rust, Go e C coexistam enquanto mantêm seus próprios ciclos de compilação e dependências isoladas.
  
## 🇺🇸 English: Module Description
* **`irongate-core`**: Responsible for intelligence. Ensures all nodes follow the same rule through the **Raft** protocol.
* **`irongate-kernel`**: Where the "heavy lifting" happens. Written in **C** to run as close as possible to the hardware.
* **`irongate-network`**: Manages the flow of packets that were not discarded by the kernel, using **Go**'s concurrency.
* **`irongate-visuals`**: Transforms raw network data into understandable charts and metrics.
* **Engineering Note**: This structure follows the **Monorepo with Workspace** pattern, allowing Rust, Go, and C components to coexist while maintaining their own build cycles and isolated dependencies.

  ## 📂 Estrutura de Pastas / Directory Structure
```text
irongate/
├── .github/                   # CI/CD Automation (GitHub Actions)
├── irongate-core/             # Control Plane
│   ├── src/
│   │   ├── raft/              # Consensus Logic
│   │   ├── ebpf/              # eBPF Lifecycle Management
│   │   └── main.rs            # Rust Entry Point
│   └── Cargo.toml             
├── irongate-kernel/           # Data Plane (C / eBPF)
│   ├── include/               
│   ├── filter.c               # Packet Filtering Programs
│   └── Makefile               
├── irongate-network/          # Network Processing (User Space)
│   ├── engine/                # Parallel Processing Goroutines
│   └── main.go                # Kernel-User Space Bridge
├── irongate-visuals/          # Observability Layer
│   ├── api/                   # Telemetry REST Endpoints
│   ├── dashboard/             # Monitoring Dashboard
│   └── go.mod                 
├── scripts/                   # Support Tools
│   ├── setup.sh               
│   └── chaos_test.py          
├── config/                    # Configuration Files
└── README.md                  # Definitive Documentation
  
