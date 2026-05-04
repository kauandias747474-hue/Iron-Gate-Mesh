# Iron Gate Mesh 

[![Rust](https://img.shields.io/badge/Rust-000000?style=for-the-badge&logo=rust&logoColor=white)](https://www.rust-lang.org/)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![C](https://img.shields.io/badge/C-A8B9CC?style=for-the-badge&logo=c&logoColor=white)](https://learn.microsoft.com/cpp/c-language/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)

---

## 🇧🇷 Português: Documentação Técnica de Engenharia

O **IronGate Mesh** é uma malha de serviço distribuída projetada para segurança e observabilidade de alta performance. Esta versão opera integralmente em **User Space**, utilizando captura de pacotes otimizada e um motor de consenso distribuído para garantir proteção resiliente sem a complexidade de dependências diretas do Kernel.

### 🏗️ Arquitetura do Ecossistema

#### 1. Plano de Dados (Data Plane - Go / C)
Responsável pela ingestão, análise e filtragem de pacotes em tempo real.
*   **Packet Sniffing**: Utiliza bibliotecas de captura de baixo nível para interceptação direta de interfaces de rede.
*   **Software Filtering**: Avaliação de regras de firewall processadas por goroutines paralelas, garantindo o processamento de alto tráfego com baixa latência.
*   **Atomic Policies**: Aplica listas de bloqueio (blacklists) vindas do plano de controle instantaneamente através de operações atômicas em memória.

#### 2. Plano de Controle (Control Plane - Rust)
O cérebro da malha distribuída, focado em consistência rigorosa e segurança de memória.
*   **Consenso com Raft**: Implementa o algoritmo Raft para garantir que todos os nós da malha concordem sobre as políticas de segurança de forma síncrona.
*   **Segurança Pragmática**: A escolha do Rust elimina *data races* e vazamentos de memória, garantindo estabilidade para sistemas que operam 24/7.
*   **Distribuição de Estado**: Propaga atualizações de segurança para todo o cluster em milissegundos após a detecção de uma ameaça.

#### 3. Camada de Observabilidade (Visuals - Go)
*   **Telemetry Stream**: Transforma o fluxo binário de pacotes em métricas estruturadas e telemetria em tempo real.
*   **Dashboard API**: Expõe dados agregados de performance e saúde dos nós para visualização executiva e técnica.

---

## 🇺🇸 English: Engineering Technical Documentation

**IronGate Mesh** is a distributed service mesh designed for high-performance security and observability. This version operates entirely in **User Space**, leveraging optimized packet capture and a distributed consensus engine to provide resilient protection without the complexity of direct Kernel dependencies.

### 🏗️ Ecosystem Architecture

#### 1. Data Plane (User Space - Go / C)
Responsible for real-time packet ingestion, analysis, and filtering.
*   **Packet Sniffing**: Utilizes low-level capture libraries for direct network interface interception.
*   **Software Filtering**: Firewall rule evaluation processed by parallel goroutines, ensuring high-throughput processing with low latency.
*   **Atomic Policies**: Instantly applies blocklists from the control plane using memory-safe atomic operations.

#### 2. Control Plane (User Space - Rust)
The distributed mesh intelligence, focused on strict consistency and memory safety.
*   **Raft Consensus**: Implements the Raft algorithm to ensure all mesh nodes agree on security policies synchronously.
*   **Pragmatic Safety**: Rust eliminates data races and memory leaks, ensuring stability for 24/7 security systems.
*   **State Distribution**: Propagates security updates to the entire cluster within milliseconds upon threat detection.

#### 3. Observability Layer (Visuals - Go)
*   **Telemetry Stream**: Converts raw binary packet flow into structured metrics and real-time telemetry.
*   **Dashboard API**: Exposes aggregated performance data and node health metrics for executive and technical visualization.

---

## 📁 Estrutura do Repositório / Repository Structure

| Módulo / Module | Linguagem / Language | Responsabilidade / Responsibility |
| :--- | :--- | :--- |
| **irongate-core** | Rust | Consenso de cluster (Raft) e gestão de estado / Cluster consensus and state |
| **irongate-network** | Go / C | Captura e filtragem de pacotes / Packet capture and filtering |
| **irongate-visuals** | Go | Telemetria e Dashboard / Telemetry and Dashboard |
| **deploy** | Docker | Orquestração e escalonamento / Orchestration and scaling |
