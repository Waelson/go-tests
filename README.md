# App Demo com Stack de Observabilidade

### Stack
- Go 1.22
- Logrus
- Swagger
- Prometheus
- Grafana
- ElasticSearch
- Kibana

### Geração dos mocks
`
mockgen -source=internal/infra/repository.go -destination=internal/infra/repository_mock.go -package=internal/mock
`

### Executar todos os testes
`
make test
`

### Gerar documentação Swagger
`
make swagger
`
