FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go mod download

ARG TOKEN

ENV PORT=8080
ENV TOKEN ${TOKEN}

CMD ["go", "run", "./cmd/main.go"]