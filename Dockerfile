# Etapa 1: Construção da imagem Go
FROM golang:1.25 AS builder

WORKDIR /app

# Copiar o código para o container
COPY . .

# Instalar dependências
RUN go mod tidy

# Compilar a aplicação
RUN go build -o app .

# Etapa 2: Execução do container
FROM gcr.io/distroless/base

WORKDIR /root/

# Copiar o binário compilado para a imagem final
COPY --from=builder /app/app .

# Expor a porta da aplicação
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./app"]
