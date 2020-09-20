FROM alpine
COPY app /
COPY init.sql /
RUN chmod +x /app
RUN apk add --no-cache tzdata
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
EXPOSE 8080
ENTRYPOINT ["/app"]
