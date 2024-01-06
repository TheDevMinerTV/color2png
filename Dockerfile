FROM golang:1.19-alpine3.16 as builder
WORKDIR /build

RUN apk add gcc g++

COPY . /build

RUN GOOS=linux GOARCH=amd64 go build -o /colorgen

FROM alpine:3.19 AS runner
WORKDIR /data

COPY --from=builder /colorgen /bin/colorgen
CMD [ "/bin/colorgen" ]
