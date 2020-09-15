FROM alpine:latest
WORKDIR /bin
COPY kubectl-yaml_writer ./

ENTRYPOINT ["sh"]

