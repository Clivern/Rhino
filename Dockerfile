FROM golang:1.14.4

ARG RHINO_VERSION=0.1.1

ENV GO111MODULE=on

RUN mkdir -p /app/configs
RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/Clivern/Rhino/releases/download/${RHINO_VERSION}/Rhino_${RHINO_VERSION}_Linux_x86_64.tar.gz | tar xz

RUN rm LICENSE
RUN rm README.md
RUN mv Rhino rhino

COPY ./config.dist.json /app/configs/

RUN ./rhino --get=release

EXPOSE 8080

VOLUME /app/configs

CMD ["./rhino", "--config", "/app/configs/config.dist.json"]