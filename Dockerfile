FROM golang:alpine AS builder

WORKDIR /build

RUN apk update && apk add git && git clone https://github.com/xaosBotTeam/web-control -b dev

RUN cd web-control && go get .  && go build -o web-control

FROM alpine

WORKDIR /build

COPY --from=builder /build/web-control/web-control /build/web-control
COPY --from=builder /build/web-control/static /build/static

EXPOSE 3000

CMD ["./web-control"]