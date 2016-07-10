# k8s-web-terminal
Kubernetes web terminal
[![Build Status](https://travis-ci.org/beyondblog/k8s-web-terminal.svg?branch=master)](https://travis-ci.org/beyondblog/k8s-web-terminal)

## Screenshot

![dasbhoard](/docs/screenshot/dashboard.png?raw=true)

![bash](/docs/screenshot/bash.png?raw=true)

## Quick Start

You need to install the build environment

* Go 1.5+
* Node.js (npm、bower、jspm to build the Angular frontend)

```bash
$ godep restore

$ godep go build 

$ jspm install

$ ./k8s-web-terminal
```

## Run at Docker

```bash
$ docker run --name k8s-web-terminal -d -p 8088:8088 -e K8S_API=http://KUBERNET_API_HOST:8080 beyondblog/k8s-web-terminal
```

## Usage

```bash
Usage of ./k8s-web-terminal:
  -k8s_api string
        Kubernetes api host (default "http://127.0.0.1:8080")
  -port int
        listen port (default 8088)
```

## Reference

[0] [sourcelair/xterm.js](https://github.com/sourcelair/xterm.js)

[1] [beyondblog/go-docker-hijack](https://github.com/beyondblog/go-docker-hijack)

[2] [rancher/kubernetes-model](https://github.com/rancher/kubernetes-model)
