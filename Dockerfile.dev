FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o web-control

FROM alpine

WORKDIR /build

COPY --from=builder /build/web-control /build/web-control
COPY --from=builder /build/static /build/static

EXPOSE 3000

CMD ["./web-control"]