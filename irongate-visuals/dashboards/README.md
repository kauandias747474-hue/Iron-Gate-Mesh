#  IronGate Visuals & Control Plane

IronGate is a high-performance Service Mesh monitoring and security system. It integrates a Rust-based core (eBPF) with a Go-based visual dashboard.
IronGate é um sistema de monitoramento e segurança para Service Mesh de alta performance. Ele integra um núcleo em Rust (eBPF) com um painel visual em Go.

---

## 🇺🇸 English Documentation

###  Project Structure
- **`/dashboards`**: Contains UI templates, alert thresholds, and theme configs.
- **`/dashboards/snapshots`**: Historical data, incident reports, and chaos test results.
- **`/scripts`**: Automation tools for benchmarking, building, and environment setup.

###  Getting Started (Windows)
1. **Prerequisites**: [Go](https://go.dev/), [Python 3](https://www.python.org/), [PowerShell 7](https://github.com/PowerShell/PowerShell).
2. **Build the project**: Run `.\scripts\build.ps1` to compile the Go binary.
3. **Execution**: Run `.\irongate-visuals.exe`. The system will automatically load `dashboard.json`.

###  Automation & Testing
- **Benchmarking**: Use `bash scripts/bench_pps.sh` to simulate high traffic.
- **Chaos Engineering**: Run `python scripts/chaos_test.py` to test Raft consensus by killing nodes.
- **Remediation**: The system automatically triggers scripts in `/scripts` when thresholds in `dashboards/alerts/thresholds.yaml` are exceeded.

---

## 🇧🇷 Documentação em Português

###  Estrutura do Projeto
- **`/dashboards`**: Contém templates de interface, limites de alerta e configurações de tema.
- **`/dashboards/snapshots`**: Dados históricos, relatórios de incidentes e resultados de testes de caos.
- **`/scripts`**: Ferramentas de automação para benchmark, build e configuração de ambiente.

###  Como Começar (Windows)
1. **Pré-requisitos**: [Go](https://go.dev/), [Python 3](https://www.python.org/), [PowerShell 7](https://github.com/PowerShell/PowerShell).
2. **Compilar**: Execute `.\scripts\build.ps1` para gerar o executável Go.
3. **Execução**: Execute `.\irongate-visuals.exe`. O sistema carregará o `dashboard.json` automaticamente.

###  Automação e Testes
- **Benchmark**: Use `bash scripts/bench_pps.sh` para simular tráfego intenso.
- **Testes de Caos**: Execute `python scripts/chaos_test.py` para testar o consenso Raft derrubando nós.
- **Remediação**: O sistema dispara scripts em `/scripts` automaticamente quando os limites em `dashboards/alerts/thresholds.yaml` são ultrapassados.

---

##  Config / Configurações
- **Theme**: Edit `dashboards/theme_config.json` to change UI colors.
- **Topology**: Edit `dashboards/network_map_cfg.json` to arrange node positions.
