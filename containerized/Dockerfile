FROM nginx:1.15.5-alpine

ARG REMOTE_USER
ARG REMOTE_DOMAIN
ARG API_KEY=

ENV REMOTE_USER=$REMOTE_USER
ENV REMOTE_DOMAIN=$REMOTE_DOMAIN
ENV API_KEY=$API_KEY

RUN apk add --no-cache curl openssh-client

# Install the Amplify plugin
RUN curl -L -O https://github.com/nginxinc/nginx-amplify-agent/raw/master/packages/install.sh
ENV API_KEY=$API_KEY
RUN if [ "$API_KEY" = "" ] ; then echo "API_KEY is undefined; skipping" ; else sh ./install.sh ; fi

RUN mkdir -p /etc/nginx/logs

COPY custom.mime.types /etc/nginx/custom.mime.types
COPY proxy.conf        /etc/nginx/proxy.conf
COPY nginx.conf        /etc/nginx/nginx.conf
COPY machine.pem       /etc/nginx/machine.pem
COPY setup-nginx.sh   /etc/nginx/setup-nginx.sh
COPY setup-tunnel.sh   /etc/nginx/setup-tunnel.sh

RUN chmod 600 /etc/nginx/machine.pem

CMD sh /etc/nginx/setup-tunnel.sh && sh /etc/nginx/setup-nginx.sh
