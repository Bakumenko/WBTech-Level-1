package pkg

import "fmt"

type Client struct {
	Id   int
	Name string
}

func NewClient(id int, name string) *Client {
	return &Client{id, name}
}

func (c *Client) Doing(tar Target) {
	fmt.Printf("Client %v is doing\n", *c)
	tar.action()
}
