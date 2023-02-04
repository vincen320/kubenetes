FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o go-env

RUN chmod +x /app/go-env

#build a tiny docker image
FROM alpine:3

RUN mkdir /app

#SANGAT PENGARUH!
#AGAR PATH SEKARANG MENJADI SAMA SEPERTI DI BINARY, KARENA GOLANG MULAI MEMBACA PATH SESUAI DENGAN PATH SEKARANG (JIKA ADA PROGRAM GOLANG YANG BUTUH BACA FILE INI SANGAT PERLU, CONTOH: godotenv())
#CONTOH COBA LIHAT DI TERMINAL PASTI POSISI NYA DI /d/test Microservices/Application 1/user-service
#COBA di cd ..
#LALU run main.go nya seperti ini -> product-service/main.go
#PASTI ERROR karena golangnya tidak ketemu file go.mod di path yang sekarang (/d/test Microservices/Application 1) (kalo go run sepertinya masih butuh go mod)
WORKDIR /app

COPY --from=builder /app/go-env /app

ENTRYPOINT [ "./go-env" ]