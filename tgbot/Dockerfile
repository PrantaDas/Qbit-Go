FROM golang:1.18 AS build

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./qbit

# RUN chmod +x qbit

# CMD [ "./qbit" ]


# Run stage
FROM golang:1.18

WORKDIR /app

COPY --from=build /app/qbit ./
RUN chmod +x qbit

CMD ["./qbit"]