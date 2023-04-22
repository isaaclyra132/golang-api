# Define a imagem base do Docker a ser utilizada
FROM golang:1.20.3-alpine

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod ./
COPY go.sum ./

# Faz o download das dependências do projeto
RUN go mod download

# Copia todos os arquivos do projeto para o diretório de trabalho
COPY . .

# Compila o projeto
RUN go build -o main .

# Define a porta em que a aplicação irá escutar
EXPOSE 8043

# Inicia a aplicação quando o contêiner for iniciado
CMD ["./main"]
