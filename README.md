# k8s-web-terminal
Kubernetes web terminal

## Screenshot

![dasbhoard]((/docs/screenshot/dashboard.png?raw=true))

![bash]((/docs/screenshot/bash.png?raw=true))

## Quick Start

You need to install the build environment

* Go 1.4+
* Node.js (npm、bower、jspm to build the Angular frontend)

```bash
    go build 
    jspm install

    ./k8s-web-terminal
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
