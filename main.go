package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	ExportVariable = "main"
)

func main() {
	logrus.Info("main")
	fmt.Println("vim-go")
}
