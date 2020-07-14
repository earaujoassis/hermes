# Hermes

> Hermes (cryptonym to "web-bridge") is a containerized NGINX reverse-proxy + SSH tunnel installation +
> an application for introspected tunnels to localhost

This project was once used to setup a proxy-server using NGINX + SSH Tunnel. The problem and solution are: (1) there
is a *hidden* server in the Web, which is not directly open to the open Web; (2) a Docker container connects to
the *hidden* server through SSH, creating a SSH Tunnel inside the container and making it available to the
NGINX setup through a localhost port; (3) the open Web can access the *hidden* HTTP server through this setup.
Basically, it is intended to be used in a home-private context with TCP-port restrictions imposed by an ISP or network setup.

Currently, there's an application for introspected tunnels to localhost, just like [ngrok v1](inconshreveable/ngrok)
used to work. You may define several applications running in a *hidden* server and the application will make them
available to the open Web.

## Setup & Running for the containerized (initial) solution

The following image installs [NGINX Amplify](https://amplify.nginx.com/), if the `API_KEY` build-argument is available.
`REMOTE_USER` and `REMOTE_DOMAIN` is mandatory, as well as the `machine.pem` file in the project root; that file is
copied into the Docker image to access the remote/*hidden* server. The following commands create a new image, tagged
`earaujoassis/hermes`, and create a Docker container mapping the port `8080 -> 80` (port 80 is exposed by default through
the official `nginx` Docker image):

```sh
$ docker build -t earaujoassis/hermes --build-arg API_KEY=? --build-arg REMOTE_USER=? --build-arg REMOTE_DOMAIN=? .
$ docker run -d -p 8080:80 earaujoassis/hermes:latest
```

The following command can be used to remove dangling images created through the command above:

```sh
$ docker images --quiet --filter=dangling=true | xargs docker rmi
```

## Setup & Running introspected tunnels to localhost

A `.env` file is used to load environment information for the application – `.env.sample` is available with important
environment variables. Since the application uses two processes, we use `goreman` (a Golang fork of Foreman) to initiate
all necessary processes.

Once the configuration is complete, the following commands will run the application locally:

```sh
$ go get github.com/mattn/goreman
$ goreman start
$ open http://localhost:5000
```

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
