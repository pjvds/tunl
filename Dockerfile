FROM golang:1.15-alpine AS builder
ENV CGO_ENABLED=0

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o tunl .

FROM scratch 
COPY --from=builder /src/tunl /go/tunl

CMD ["/go/tunl"]
