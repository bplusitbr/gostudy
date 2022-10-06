FROM golang:1.18.4
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -trimpath -o ./bplusauth main.go
CMD ["sh", "-c", "./bplusauth"]
EXPOSE 8080

