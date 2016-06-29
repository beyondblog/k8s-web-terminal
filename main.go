package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

func init() {
	iris.Static("/assets", "./public/assets", 1)
	iris.Static("/public", "./public", 1)
	iris.Static("/frontend", "./frontend", 1)
	iris.Config.Render.Template.Engine = iris.PongoEngine
	iris.Config.Render.Rest.IndentJSON = true
}

var (
	client *Client
)

func main() {

	//list all container
	iris.Get("/", func(c *iris.Context) {
		c.Render("index.html", nil)
	})

	iris.Get("/api/containers", func(ctx *iris.Context) {
		containers := client.ListContainers()
		ctx.JSON(iris.StatusOK, containers)
	})

	iris.Get("/docker", func(c *iris.Context) {
		fmt.Println("docker")
		c.Render("docker.html", map[string]interface{}{"Name": "iris"})
	})

	// the path which the websocket client should listen/registed to ->
	iris.Config.Websocket.Endpoint = "/bash"

	client = &Client{"http://127.0.0.1:2375"}

	input := make(chan []byte)

	ws := iris.Websocket // get the websocket server

	ws.OnConnection(func(c websocket.Connection) {

		id, _ := client.CreateExec("6544e9f086a6", "/bin/bash")
		fmt.Println("id:" + id)
		output, err := client.ExecStart(id, input)
		fmt.Println(err)
		fmt.Println(output)

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
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})

	})

	iris.Listen("0.0.0.0:8088")
}
