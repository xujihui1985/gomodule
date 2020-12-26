package liba

import "github.com/sirupsen/logrus"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Method1() {
	logrus.Infof("hello liba client")
}
