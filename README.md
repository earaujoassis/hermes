# Hermes

> Hermes (cryptonym to "web-bridge") is a containerized Nginx reverse-proxy + SSH tunnel installation

## Setup & Running

```sh
$ docker build -t earaujoassis/hermes --build-arg API_KEY=? --build-arg REMOTE_USER=? --build-arg REMOTE_DOMAIN=? .
$ docker run -d -p 8080:80 earaujoassis/hermes:latest
$ docker images --quiet --filter=dangling=true | xargs docker rmi
```

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
