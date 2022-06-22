package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/MatthewKirik/architecture-lab-4/command"
	"github.com/MatthewKirik/architecture-lab-4/engine"
)

var inputFile = flag.String("f", "./assets/commands-examples.txt",
	"Input file with loop commands")

func main() {
	flag.Parse()

	input, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		commandLine := scanner.Text()
		cmd := command.Parse(commandLine)
		eventLoop.Post(cmd)
	}
	eventLoop.AwaitFinish()
}
