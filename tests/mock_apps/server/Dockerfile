FROM golang:1.15.2

ENV CGO_ENABLED=0

WORKDIR /app

COPY . .

RUN go get github.com/dgrijalva/jwt-go && go build -o server .

EXPOSE 8081

CMD ["/app/server"]