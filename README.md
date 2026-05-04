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
*   **Consenso Raft**: Implementação de máquina de estados distribuída para garantir que políticas de segurança sejam idênticas em todo o cluster.
*   **Segurança Pragmática**: Uso intensivo de Rust para eliminar vazamentos de memória e *race conditions* em operações 24/7.
*   **Gestão de Auditoria**: Registro centralizado de eventos e controle de admissão de novos nós.

#### 2. Plano de Dados e Kernel (`irongate-kernel` & `irongate-network`) [C / Go]
Responsável pela ingestão e filtragem de pacotes.
*   **Filtragem de Software**: Avaliação de pacotes em tempo real utilizando goroutines paralelas e operações atômicas.
*   **Otimização C**: Estrutura preparada para filtragem *stateless* e balanceamento de carga de baixo nível.
*   **Políticas Atômicas**: Atualização imediata de blacklists sem necessidade de reiniciar o serviço.

#### 3. Camada de Observabilidade (`irongate-visuals`) [Go]
*   **Telemetria**: Extração de métricas de performance diretamente do processamento de rede.
*   **Dashboard API**: Exposição de dados estruturados sobre tráfego, pacotes descartados e integridade dos nós.

---

## 🇺🇸 English: Engineering Technical Documentation

**IronGate Mesh** is a distributed service mesh designed for high-performance security and observability. This version operates through a hybrid model: optimized **User Space** network processing for deployment flexibility, and a **C** data plane for future Kernel-level extensions.

### 🏗️ Ecosystem Architecture

#### 1. Control Plane (`irongate-core`) [Rust]
The distributed intelligence of the mesh, focused on strict consistency.
*   **Raft Consensus**: Distributed state machine implementation to ensure security policies remain identical across the cluster.
*   **Pragmatic Safety**: Leverages Rust to eliminate memory leaks and race conditions in 24/7 operations.
*   **Audit Management**: Centralized event logging and node admission control.

#### 2. Data Plane & Kernel (`irongate-kernel` & `irongate-network`) [C / Go]
Responsible for packet ingestion and filtering.
*   **Software Filtering**: Real-time packet evaluation using parallel goroutines and atomic operations.
*   **C Optimization**: Architecture prepared for stateless filtering and low-level load balancing.
*   **Atomic Policies**: Immediate blacklist updates without service interruption.

#### 3. Observability Layer (`irongate-visuals`) [Go]
*   **Telemetry**: Performance metrics extraction directly from network processing.
*   **Dashboard API**: Exposure of structured data regarding traffic, dropped packets, and node health.

---

## 📁 Estrutura do Projeto / Project Structure
```text
irongate/
├── .github/workflows/       # CI/CD Pipelines
├── irongate-core/           # Control Plane (Rust: Raft, Audit, Admission)
├── irongate-kernel/         # Data Plane (C: Filter, Load Balancer)
├── irongate-visuals/        # Observability (Go: Metrics, Dashboard)
├── scripts/                 # Automation (Build, Chaos Testing)
└── README.md                # Project Documentation
