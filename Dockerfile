FROM alpine:3.3

MAINTAINER beyondblog "beyondblog@outlook.com"

# add bash
RUN apk upgrade --update && \
    apk add --update  bash && \
    rm -rf /var/cache/apk/*

ENV K8S_API http://127.0.0.1:8080

COPY public /public

COPY templates /templates

COPY k8s-web-terminal /

COPY docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]

EXPOSE 8088
CMD ["/k8s-web-terminal"]
