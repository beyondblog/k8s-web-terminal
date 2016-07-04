package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	"log"
)

func init() {
	iris.Static("/assets", "./public/assets", 1)
	iris.Static("/public", "./public", 1)
	iris.Static("/frontend", "./frontend", 1)
	iris.Config.Render.Template.Engine = iris.PongoEngine
	iris.Config.Render.Rest.IndentJSON = true
}

var (
	client *DockerClient
	k8s    *K8sClient
	ws     websocket.Server
)

func main() {

	//list all container
	iris.Get("/", func(c *iris.Context) {
		c.Render("index.html", nil)
	})

	iris.Get("/api/nodes", listNodes)
	iris.Get("/api/nodes/containers", listContainers)
	iris.Get("/api/nodes/containers/shell/ws", shellContainer)

	iris.Get("/container/terminal", containerTerminal)

	k8s = &K8sClient{"http://127.0.0.1:8080", make(map[string]*DockerClient)}

	iris.Listen("0.0.0.0:8088")
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
	containerId := ctx.URLParam("containerId")
	node := ctx.URLParam("node")
	if len(containerId) > 0 && len(node) > 0 {
		client := k8s.GetDockerClient(node)
		input := make(chan []byte)

		ws := websocket.NewServer(iris.Config.Websocket)
		ws.OnConnection(func(c websocket.Connection) {

			id, _ := client.CreateExec(containerId, "/bin/bash")
			output, err := client.ExecStart(id, input)
			if err != nil {
				log.Println("Error: %v", err)
			}

			go func() {
				for {
					data := <-output
					c.EmitMessage(data)
				}
			}()

			c.OnMessage(func(data []byte) {
				input <- data
			})

			c.OnDisconnect(func() {
				log.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
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
