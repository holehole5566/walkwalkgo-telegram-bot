FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go mod download

ARG TOKEN
ARG MONGO_URI

ENV PORT=8080
ENV TOKEN ${TOKEN}
ENV MONGO_URI ${MONGO_URI}

CMD ["go", "run", "./cmd/main.go"]