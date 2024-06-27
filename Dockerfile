# package
FROM alpine:3.19
ENV CONFIG_PATH=/app/config.json
ENV CONTENT_DIR=/app/anon/home
RUN mkdir -p /app/anon && adduser -s /bin/sh -D -u 1001 anon
ADD atykym.space /app/app
ADD prod.config.json /app/config.json
ADD static /app/static
ADD templates /app/templates
ADD anon/home /app/anon/home
RUN chown -R anon:anon /app
WORKDIR /app
RUN ["chmod", "+x", "./app"]
USER anon
ENTRYPOINT ["/app/app"]