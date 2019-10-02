FROM alpine:3

COPY dev-sec-ops-seed /seed

ENTRYPOINT ["/seed"]
