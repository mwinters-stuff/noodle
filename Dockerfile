FROM ubuntu:latest
RUN apt-get update && apt-get -y upgrade && apt-get -y install apt-utils ca-certificates
RUN DEBIAN_FRONTEND=noninteractive apt-get -y install tzdata
WORKDIR /app
RUN mkdir /config
COPY noodle /app/
ENTRYPOINT [ "/app/noodle", "--config", "/config/noodle.yaml" ]