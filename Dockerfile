FROM arm32v7/golang:1.15 AS builder
LABEL maintainer="Daniel Lynch <danplynch@gmail.com>"
RUN mkdir -p /go/src/github.com/randomtask1155/hqweb
RUN mkdir /app
WORKDIR $GOPATH/src/github.com/randomtask1155/hqweb

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=7
ADD . .
RUN ls /go/src/github.com/randomtask1155/hqweb
RUN go get github.com/revel/cmd/revel 
RUN cd /go/src/github.com/randomtask1155/hqweb
RUN /go/bin/revel build . /app

FROM arm32v7/ubuntu:bionic
COPY --from=builder /app/* /app/
RUN mkdir /app/src
RUN mv /app/github.com /app/src
RUN ls -l /app/
RUN ls /app/src/github.com/randomtask1155/hqweb
RUN cat /app/run.sh
EXPOSE 9000
ENTRYPOINT ["/app/run.sh"]
