FROM       alpine:latest
MAINTAINER demo <demo@demo.com>

#RUN apk add --update curl bash && \
#    rm -rf /var/cache/apk/*


WORKDIR /data

COPY Main .
COPY webroot   ./webroot

EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["./Main"]