FROM golang:1.22.10

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go 

CMD [ "./main" ]
