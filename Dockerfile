FROM golang:1.20

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

COPY . .

RUN chmod -R 644 internal/database/migrations/*.sql

RUN chmod +x wait-for-it.sh

#CMD ["./wait-for-it.sh", "mysql:3306", "--", "migrate", "-path", "migrations", "-database", "mysql://root:root@tcp(mysql:3306)/wallet", "up", "&&", "go", "run", "cmd/walletcore/main.go"]
CMD ["go", "run", "cmd/walletcore/main.go"]