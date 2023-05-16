package server

import "github.com/americanas-go/config"

const server = "app.server"

func init() {
	config.Add(server, "rest", "decides what server to use. Grpc or Rest")
}

func Value() string {
	return config.String(server)
}
