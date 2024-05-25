FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go api ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./qbit

RUN chmod +x qbit

CMD [ "./qbit" ]