FROM golang:latest

RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN export GO111MODULE=on

RUN git clone https://github.com/sachins602/stockApi.git

RUN go build

EXPOSE 8080

ENTRYPOINT [ "/build/goapi" ]