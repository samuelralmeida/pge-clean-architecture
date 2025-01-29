
# pge-clean-architecture

**DESAFIO DA PÓS GO EXPERT: Clean Architecture**

## ORIENTAÇÕES DO DESAFIO:

Partindo do código da aula: https://github.com/devfullcycle/goexpert/tree/main/20-CleanArch

Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.

Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## ORIENTAÇÕES PARA EXECUÇÃO

1. na raiz do projeto execute `docker compose -f docker-compose.yaml up` para criação do banco de dados
2. navegue até o diretório com o main do projeto: `cd cmd/ordersystem`
3. execute o projeto: `go run main.go wire_gen.go`

### PORTAS DAS APIS

- webserver => 8000
- grpc => 50051
- graphql => 8080

