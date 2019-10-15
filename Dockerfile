FROM alpine:3

COPY seed /seed

ENTRYPOINT ["/seed"]
