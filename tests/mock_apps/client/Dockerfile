FROM golang:1.15.2

ENV CGO_ENABLED=0

WORKDIR /app

COPY . .

RUN go get github.com/dgrijalva/jwt-go && go build -o client .

CMD ["/app/client"]