FROM golang:1.19.2-bullseye
WORKDIR /app
COPY . .
RUN go mod download
RUN go build cmd/api/main.go
CMD [ "./main" ]

