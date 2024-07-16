# Nome do arquivo binário
BINARY_NAME=lru_go

# Nome do arquivo principal (main package)
MAIN_PACKAGE=main.go

# Comando go
GO=go

.PHONY: all build run clean test

# Build o projeto
all: build

# Compila o projeto
build:
	$(GO) build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Executa o binário
run: build
	./$(BINARY_NAME)

# Limpa arquivos binários
clean:
	rm -f $(BINARY_NAME)

# Roda os testes
test:
	$(GO) test ./...