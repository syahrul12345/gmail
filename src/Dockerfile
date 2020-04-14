FROM golang:alpine as stage1
WORKDIR /app
COPY ./go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
COPY --from=stage1 /app/main .

CMD ["./main"]