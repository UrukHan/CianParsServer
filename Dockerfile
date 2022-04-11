FROM alpine
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app
COPY . /app
EXPOSE 8000
CMD ["./main"]