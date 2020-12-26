package liba

import "github.com/sirupsen/logrus"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Method1() {
	logrus.Infof("hello liba client 1.0.0")
}

func (c *Client) Method2() {
	logrus.Infof("hello liba client 2.0.0")
}
