
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

1. na raiz do projeto execute `docker compose -f docker-compose.yaml up`
    - vai subir o banco de dados e fazer a migration
    - vai subir o RabbitMQ
    - vai subir a aplicação e disponibilizar as três APIs: rest, grpc, graphql

### PORTAS DAS APIS

- rest => 8000
- grpc => 50051
- graphql => 8080

