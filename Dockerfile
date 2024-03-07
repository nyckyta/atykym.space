# package
FROM alpine:3.19
ENV PORT=$PORT
RUN mkdir /app && adduser -s /bin/sh -D -u 1001 app
COPY --from=build atykym.space /app/app
ADD static /app/static
RUN chown -R app:app /app
WORKDIR /app
USER app
ENTRYPOINT ["/app/app"]