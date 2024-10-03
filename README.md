# App para demonstrar o uso de testes unitários em Go

### Geração dos mocks
`
mockgen -source=internal/infra/repository.go -destination=internal/infra/repository_mock.go -package=infra
`

### Executar todos os testes
`
make test
`

### Gerar documentação Swagger
`
make swagger
`
