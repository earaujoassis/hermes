#!/bin/sh

set -e

echo "Setting up SSH Tunnel at ${REMOTE_USER}@${REMOTE_DOMAIN}"
ssh -i /etc/nginx/machine.pem -o StrictHostKeyChecking=no -C -f -q -N -L 8000:localhost:80 $REMOTE_USER@$REMOTE_DOMAIN
