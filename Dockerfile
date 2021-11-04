FROM golang:1.17.3

ARG RHINO_VERSION=1.6.1

ENV GO111MODULE=on

RUN mkdir -p /app/configs
RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/Clivern/Rhino/releases/download/${RHINO_VERSION}/Rhino_${RHINO_VERSION}_Linux_x86_64.tar.gz | tar xz

RUN rm LICENSE
RUN rm README.md
RUN mv Rhino rhino

COPY ./config.dist.json /app/configs/

RUN ./rhino version

EXPOSE 8080

VOLUME /app/configs

CMD ["./rhino", "serve", "-c", "/app/configs/config.dist.json"]