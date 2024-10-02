# go-tests

### Geração dos mocks
`
mockgen -source=internal/infra/repository.go -destination=internal/infra/repository_mock.go -package=infra
`

### Executar todos os testes
`
make test
`