FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN export GO11MODULE=on
RUN go get "go-shortner/shortner"
RUN go get "go-shortner/store"
RUN go get "github.com/gin-gonic/gin"
RUN go get "github.com/itchyny/base58-go"
RUN go get "github.com/stretchr/testify/assert"
RUN go get "github.com/go-redis/redis/v8"
RUN cd /app && git clone https://github.com/Mitulsharma123/go-shortner 

RUN cd /app/go-shortner && go build 

EXPOSE 9808

CMD ["/app/go-shortner/main"]
