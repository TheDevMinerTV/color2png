FROM golang:1.25-alpine AS builder
WORKDIR /build

COPY . /build

RUN CGO=0 GOOS=linux GOARCH=amd64 go build -o /colorgen

FROM alpine:3.22 AS runner
WORKDIR /data

COPY --from=builder /colorgen /bin/colorgen
CMD [ "/bin/colorgen" ]
