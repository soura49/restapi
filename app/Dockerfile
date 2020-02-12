FROM golang:alpine
RUN mkdir /go/src/app 
ADD . /go/src/app/
WORKDIR /go/src/app 
RUN apk update && \
    apk upgrade && \
    apk add git
RUN go get github.com/gorilla/mux && go get github.com/lib/pq
RUN go build -o main .
EXPOSE 8088
CMD ["./main"]
