FROM golang:1.24-alpine

# Install dependencies
RUN apk add --no-cache \
    bash curl nano wget rsync git ca-certificates postgresql-client

# Install nodejs
RUN wget http://nodejs.org/dist/v20.13.1/node-v20.13.1-linux-x64.tar.gz \
    && tar -C /usr/local --strip-components 1 -xzf node-v20.13.1-linux-x64.tar.gz \
    && rm node-v20.13.1-linux-x64.tar.gz

WORKDIR /var/www/html
