FROM golang:1.14.5-alpine3.12

LABEL "com.quatrolabs.hermes"="quatroLABS Hermes"
LABEL "description"="Hermes (cryptonym to 'web-bridge') is an application for introspected tunnels to localhost"

RUN apk add --update --no-cache \
    binutils-gold \
    curl \
    g++ \
    gcc \
    gnupg \
    libgcc \
    linux-headers \
    make \
    python3 \
    postgresql \
    postgresql-contrib \
    postgresql-libs \
    postgresql-dev \
    git

ENV PATH=/usr/local/bin:$PATH

ENV PORT=80
ENV GIN_MODE=release
ENV GO111MODULE=on

RUN mkdir -p /app

WORKDIR /app

COPY . /app

EXPOSE 80

ENTRYPOINT [ "go", "run", "main.go" ]
CMD [ "web" ]
