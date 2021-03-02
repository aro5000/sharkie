FROM golang:1.15-alpine as build
WORKDIR /app
ADD cmd/ ./cmd
ENV GOPATH /go
ENV CGO_ENABLED=0

ARG DISABLE_TESTS
RUN if [[ "$DISABLE_TESTS" = "true" ]] ; then echo Skipping Tests ; else go test ./cmd/sharkie ; fi
RUN go build ./cmd/sharkie

FROM alpine:latest
COPY --from=build /app/sharkie /app/
WORKDIR /app
RUN chown 65534:65534 sharkie
USER 65534:65534
ADD static/ ./static
ENV SHARKIE_PORT ":5000"
ENTRYPOINT [ "./sharkie" ]
