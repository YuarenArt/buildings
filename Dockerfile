FROM golang:1.21-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o awesome_project ./main.go

CMD ["./awesome_project"]