# Nome do módulo
MODULE_NAME := github.com/Waelson/go-tests

# Arquivo que armazena o resultado da cobertura de código
COVERAGE_FILE := coverage.out

# Diretório do pacote
PKG := ./...

# Comando para rodar os testes
test:
	@echo "Running unit tests..."
	go test -v $(PKG)

# Rodar os testes com geração de relatório de cobertura de código
coverage:
	@echo "Running unit tests with coverage..."
	go test -v -coverprofile=$(COVERAGE_FILE) $(PKG)
	@echo "Generating coverage report..."
	go tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@echo "Coverage report generated at coverage.html"

# Limpar arquivos de cobertura gerados
clean:
	@echo "Cleaning up..."
	@rm -f $(COVERAGE_FILE) coverage.html
	@echo "Cleanup complete"

# Rodar verificação de formatação
fmt:
	@echo "Running gofmt..."
	gofmt -s -w $(PKG)

# Rodar linter para verificar o estilo de código
lint:
	@echo "Running golint..."
	golint $(PKG)

# Rodar todos os comandos acima
all: fmt lint test coverage

swagger:
	@echo "Generating swagger documentation..."
	swag init -d cmd/,internal/controller --parseDependency -o internal/docs
	mv internal/docs/swagger.json internal/docs/swagger.yaml docs/specs
	@echo "Done!"
.PHONY: test coverage clean fmt lint all
