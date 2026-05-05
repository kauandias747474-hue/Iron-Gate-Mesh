#  IronGate Visuals & Control Plane

IronGate is a high-performance Service Mesh monitoring and security system. It integrates a Rust-based core (eBPF) with a Go-based visual dashboard for real-time traffic analysis and automated remediation.
IronGate é um sistema de monitoramento e segurança para Service Mesh de alta performance. Ele integra um núcleo em Rust (eBPF) com um painel visual em Go para análise de tráfego em tempo real e remediação automatizada.

---

## 🇺🇸 English Documentation

###  Project Structure & Refactoring
Today, the project underwent a major structural reorganization to comply with Go's workspace standards and Rust's strict safety patterns:
- **`/dashboards`**: (Refactored) Now a dedicated Go package. Contains `metrics.go`, UI templates, and theme configs.
- **`/topology`**: (New) Dedicated package for network graph logic and node state management.
- **`/dashboards/snapshots`**: Storage for critical incident logs and state captures.
- **`/scripts`**: Automation tools for benchmarking, building, and PowerShell-based remediation.

###  Advanced Concepts Applied Today
1.  **Concurrency & Parallelism (Go)**: Implemented `goroutines` for non-blocking metric collection. Used `sync.RWMutex` to prevent race conditions during dashboard exports and `sync/atomic` for high-speed packet counting.
2.  **Graceful Shutdown**: Implementation of `os/signal` and `context` to ensure the system closes all socket connections and saves logs before exiting (Ctrl+C).
3.  **Distributed Consensus (Rust/Raft)**: Initialization of the Raft protocol logic for state replication across the mesh.
4.  **Type Safety & Memory Management**: Strict enforcement of `uint64` for high-traffic telemetry to prevent integer overflow in critical systems.

###  Error Handling & Bug Fixes
- **Go Module Resolution**: Fixed the `"package ... is not in std"` error by correctly initializing `go mod init irongate` and organizing files into proper sub-packages.
- **Unused Dependency Cleanup**: Resolved compilation blockers by removing unused imports (`encoding/json`) and "declared but not used" variables in the metrics engine.
- **Rust Syntax & Pathing**: Fixed critical path separator errors (`::` vs `:`) and module visibility issues in the Core Engine.
- **Windows Pathing**: Optimized the executable execution logic for PowerShell using local scope referencing (`.\`).

###  Getting Started (Windows)
1.  **Prerequisites**: [Go 1.20+](https://go.dev/), [Rust Edition 2021](https://rustup.rs/), [PowerShell 7](https://github.com/PowerShell/PowerShell).
2.  **Module Initialization**: Inside the visuals folder, run `go mod tidy` to sync dependencies.
3.  **Build**: Run `go build -o irongate-visuals.exe .` to compile.
4.  **Execution**: Run `.\irongate-visuals.exe`.

---

## 🇧🇷 Documentação em Português

### Estrutura do Projeto e Refatoração
Hoje, o projeto passou por uma reorganização estrutural profunda para cumprir os padrões de workspace do Go e os padrões rígidos de segurança do Rust:
- **`/dashboards`**: (Refatorado) Agora um pacote Go dedicado. Contém o `metrics.go`, templates de interface e configurações de tema.
- **`/topology`**: (Novo) Pacote dedicado para lógica de grafos de rede e gerenciamento de estado de nós.
- **`/dashboards/snapshots`**: Armazenamento de logs de incidentes críticos e capturas de estado (snapshots).
- **`/scripts`**: Ferramentas de automação para benchmark, build e remediação baseada em PowerShell.

###  Conceitos Avançados Aplicados Hoje
1.  **Concorrência e Paralelismo (Go)**: Implementação de `goroutines` para coleta de métricas não bloqueante. Uso de `sync.RWMutex` para evitar condições de corrida (race conditions) e `sync/atomic` para contagem de pacotes em alta velocidade.
2.  **Encerramento Gracioso (Graceful Shutdown)**: Implementação de `os/signal` e `context` para garantir que o sistema feche conexões e salve logs corretamente ao receber um comando de parada (Ctrl+C).
3.  **Consenso Distribuído (Rust/Raft)**: Inicialização da lógica do protocolo Raft para replicação de estado em toda a mesh.
4.  **Segurança de Tipos**: Aplicação estrita de `uint64` para telemetria de alto tráfego, evitando estouro de inteiros (overflow) em sistemas críticos.

###  Tratamento de Erros e Correções
- **Resolução de Módulos Go**: Corrigido o erro `"package ... is not in std"` através da inicialização correta do `go mod init irongate` e organização de arquivos em sub-pacotes.
- **Limpeza de Dependências**: Resolvidos bloqueios de compilação removendo imports não utilizados (`encoding/json`) e variáveis declaradas mas não usadas no motor de métricas.
- **Sintaxe Rust**: Corrigidos erros críticos de separadores de caminho (`::` em vez de `:`) e visibilidade de módulos no Core Engine.
- **Execução no Windows**: Otimização da lógica de execução via PowerShell utilizando referenciamento de escopo local (`.\`).

###  Como Começar (Windows)
1.  **Pré-requisitos**: [Go 1.20+](https://go.dev/), [Rust Edition 2021](https://rustup.rs/), [PowerShell 7](https://github.com/PowerShell/PowerShell).
2.  **Sincronização**: Na pasta visuals, execute `go mod tidy` para organizar as dependências.
3.  **Compilação**: Execute `go build -o irongate-visuals.exe .` para gerar o binário.
4.  **Execução**: Use `.\irongate-visuals.exe` ou `go run .` para iniciar o monitoramento.

---

##  Configurações / Customization
- **Auto-Remediation**: Configure scripts em `dashboards/alerts/remediation_scripts.json` para respostas automáticas a ataques.
- **Topology**: Edite `dashboards/network_map_cfg.json` para definir o posicionamento dos nós no grafo de rede.
- **Logging**: O sistema gera automaticamente snapshots em `.json` na pasta `/snapshots` em caso de anomalias detectadas (ex: > 1M PPS).
