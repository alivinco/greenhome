FROM alpine
COPY ./bin/greenhome /opt/greenhome/
COPY ./static/ /opt/greenhome/static
COPY ./templates/ /opt/greenhome/templates
RUN apk --no-cache add ca-certificates && update-ca-certificates
VOLUME ["/etc/greenhome"]
EXPOSE 5010
WORKDIR /opt/greenhome/
ENTRYPOINT ["/opt/greenhome/greenhome","-c","/etc/blacktower/blacktower.toml"]