FROM golang:1.19.3-alpine3.17 as build

WORKDIR /build

COPY go.mod .
COPY go.sum .
ADD gin_mongo ./gin_mongo
RUN go mod download
RUN go build -o server gin_mongo/main.go

FROM alpine as runtime
COPY --from=build /build/server /
EXPOSE 8000
CMD ["./server"]

