FROM golang:1.22 as build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/atykym.space

FROM alpine:3.19
ENV PORT=$PORT
RUN mkdir /app && adduser -s /bin/sh -D -u 1001 app
COPY --from=build /app/atykym.space /app/app
ADD static /app/static
RUN chown -R app:app /app
WORKDIR /app
USER app
ENTRYPOINT ["/app/app"]