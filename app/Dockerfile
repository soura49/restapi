FROM golang:alpine
ADD src /go/src/
WORKDIR /go/src 
RUN apk update && \
    apk upgrade && \
    apk add git && \
    apk add util-linux
RUN go get github.com/gorilla/mux && go get github.com/lib/pq
RUN go build -o main .
EXPOSE 8088
CMD ["./main"]
