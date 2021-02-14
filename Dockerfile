FROM golang:1.15-alpine

WORKDIR /app

COPY . /app

RUN go build -o gostepper cmd/main.go

CMD [ "/app/gostepper" ]