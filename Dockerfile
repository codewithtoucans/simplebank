FROM golang:alpine
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o main main.go
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN migrate -path /app/db/migration -database "postgresql://root:secret@postgres:5432/simple_bank?sslmode=disable" -verbose up

EXPOSE 3000
CMD [ "/app/main" ]
