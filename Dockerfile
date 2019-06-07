FROM golang:alpine
LABEL maintainer="Riaan Stegmann"

WORKDIR /app
COPY . .

RUN apk --update --no-cache add git && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /app/app .
CMD ["./app"]