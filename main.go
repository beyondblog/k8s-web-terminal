package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	"log"
	"strconv"
)

func init() {
	iris.Static("/assets", "./public/assets", 1)
	iris.Static("/public", "./public", 1)
	iris.Config.Render.Template.Engine = iris.PongoEngine
	iris.Config.Render.Rest.IndentJSON = true
}

var (
	client  *DockerClient
	k8s     *K8sClient
	ws      websocket.Server
	k8sHost = flag.String("k8s_api", "http://127.0.0.1:8080", "Kubernetes api host")
	port    = flag.Int("port", 8088, "listen port")
)

func main() {
	flag.Parse()

	iris.Get("/", func(c *iris.Context) {
		c.Render("index.html", nil)
	})

	iris.Get("/api/nodes", listNodes)
	iris.Get("/api/nodes/containers", listContainers)
	iris.Get("/api/nodes/containers/shell/ws", shellContainer)
	iris.Get("/api/nodes/containers/shell/create", createContainer)
	iris.Get("/api/nodes/containers/shell/resize", resizeContainer)

	iris.Get("/container/terminal", containerTerminal)

	k8s = &K8sClient{*k8sHost, make(map[string]*DockerClient)}
	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	iris.Listen(addr)
}

func listContainers(ctx *iris.Context) {
	node := ctx.URLParam("node")
	if len(node) > 0 {
		containers := k8s.GetContainer(node)
		ctx.JSON(iris.StatusOK, containers)
		return
	}
	ctx.Write("node require!")
}

func listNodes(ctx *iris.Context) {
	nodes := k8s.Nodes()
	ctx.JSON(iris.StatusOK, nodes)
}

func shellContainer(ctx *iris.Context) {
	id := ctx.URLParam("id")
	node := ctx.URLParam("node")
	if len(id) > 0 && len(node) > 0 {
		client := k8s.GetDockerClient(node)
		input := make(chan []byte)

		ws := websocket.NewServer(iris.Config.Websocket)
		ws.OnConnection(func(c websocket.Connection) {

			log.Printf("\nConnection with ID: %s ", c.ID())
			log.Printf("\nCreateId: %s", id)
			output, err := client.ExecStart(id, input)
			if err != nil {
				log.Println("Error: %v", err)
			}

			go func() {
				for {
					if data, ok := <-output; ok {
						c.EmitMessage(data)
					} else {
						break
					}
				}
			}()

			c.OnMessage(func(data []byte) {
				input <- data
			})

			c.OnDisconnect(func() {
				log.Printf("\nConnection with ID: %s has been disconnected!", c.ID())

				//send EOF to close chan
				input <- []byte("EOF")
				close(input)
			})

		})

		if err := ws.Upgrade(ctx); err != nil {
			ctx.Write("Upgrade error!")
		}
		return
	}
	ctx.Write("param error!")
}

func containerTerminal(ctx *iris.Context) {
	ctx.Render("terminal.html", nil)
}

func createContainer(ctx *iris.Context) {
	containerId := ctx.URLParam("containerId")
	node := ctx.URLParam("node")
	command := ctx.URLParam("command")

	if len(containerId) > 0 && len(node) > 0 && len(command) > 0 {
		client := k8s.GetDockerClient(node)
		id, _ := client.CreateExec(containerId, command)
		ctx.JSON(iris.StatusOK, map[string]string{"id": id})
		return
	}
	ctx.Write("param error!")
}

func resizeContainer(ctx *iris.Context) {
	node := ctx.URLParam("node")
	id := ctx.URLParam("id")
	cols, _ := strconv.Atoi(ctx.URLParam("cols"))
	rows, _ := strconv.Atoi(ctx.URLParam("rows"))

	if len(node) > 0 && len(id) > 0 && cols > 0 && rows > 0 {
		log.Println("resize: ", id)
		client := k8s.GetDockerClient(node)
		client.ExecResize(id, cols, rows)
		ctx.JSON(iris.StatusOK, nil)
		return
	}
	ctx.Write("param error!")
}
