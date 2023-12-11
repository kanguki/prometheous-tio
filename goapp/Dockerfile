FROM golang:1.20 as base
WORKDIR /
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app 

FROM alpine:3.19.0
WORKDIR /
COPY --from=base /app /
EXPOSE 9101
CMD ["/app"]

