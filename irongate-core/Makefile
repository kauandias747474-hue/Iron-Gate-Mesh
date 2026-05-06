# Variáveis de comando
CARGO = cargo
BINARY_PATH = ./target/release/irongate-core

.PHONY: all build run clean test check

# Comando padrão: apenas digite 'make'
all: build

# Compila o projeto com todas as otimizações de performance
build:
	@echo " Compilando IronGate Core para Performance Máxima..."
	$(CARGO) build --release

# Compila (se necessário) e roda o projeto imediatamente
run:
	@echo " Iniciando IronGate Mesh..."
	$(CARGO) run

# Limpa todos os arquivos gerados pela compilação (libera espaço)
clean:
	@echo " Limpando arquivos de build..."
	$(CARGO) clean
	rm -f audit.log

# Verifica se há erros de digitação no código sem compilar tudo
check:
	@echo " Verificando integridade do código..."
	$(CARGO) check

# Roda testes unitários que você criar na pasta src
test:
	$(CARGO) test
