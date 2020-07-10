FROM golang:alpine as build
WORKDIR /app
ADD *go* ./
ENV GOPATH /go
RUN go build -o sharkie 

FROM alpine:latest
COPY --from=build /app/sharkie /app/
WORKDIR /app
RUN chown 65534:65534 sharkie
USER 65534:65534
ADD static/ ./static
ENV SHARKIE_PORT ":5000"
ENTRYPOINT [ "./sharkie" ]
