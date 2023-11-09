# builder image...
FROM golang:1.17-alpine3.16 AS builder
LABEL stage=builder
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN go build -o main .

## generate clean, final image...
FROM alpine:3.16 AS runner
LABEL stage=runner
LABEL maintainer="Made by ayessenb && roshakba"
LABEL org.label-schema.description="ascii-art-web"
WORKDIR /app
COPY --from=builder /app/main ./
COPY /assets /app/assets
COPY /templates /app/templates
COPY /data /app/data
EXPOSE 8080

# Run the executable
CMD ["/app/main"]