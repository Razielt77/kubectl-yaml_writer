FROM alpine:latest
WORKDIR /bin
COPY kyml ./

ENTRYPOINT ["sh"]

