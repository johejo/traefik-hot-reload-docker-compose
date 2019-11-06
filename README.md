# traefik-hot-reload-docker-compose

Dynamic routing by [traefik](https://github.com/containous/traefik) using docker-compose and Golang hot reload example with [reflex](https://github.com/cespare/reflex) and [air](https://github.com/cosmtrek/air).

## What is traefik
https://github.com/containous/traefik/

## Requirements

- make
- docker
- docker-compose
- reflex https://github.com/cespare/reflex

## Install reflex

```
go get -u github.com/cespare/reflex
```

## Start

```
make compose-hot-reload
```

## Check

```
curl localhost/hello | jq .
```

## Hot reload

### Golang hot reload (powerd by [air](https://github.com/cosmtrek/air))

If you edit `hello/main.go` and save file, it will be automatically recompiled and the service will restart.

### docker-compose hot reload (powerd by [reflex](https://github.com/cespare/reflex))

If you edit `docker-compose.yaml` and save file, it will be automatically restart docker-compose.

Try changing traefik pathprefix label "/hello" to "/world".

The reflex detect change and restart traefik and traefik detect routing change.

```
curl localhost/world | jq .
```
