package main

import (
	"bufio"
	"github.com/OlegVanyaGreatBand/architecture-lab-4/commands"
	"github.com/OlegVanyaGreatBand/architecture-lab-4/engine"
	"os"
)

func main()  { /*
	l := new(engine.Loop)
	l.Start()
	l.Post(&commands.PrintCmd{"ki"})
	l.AwaitFinish()*/

	eventLoop := new(engine.Loop)
	eventLoop.Start()
	inputFile := "test.txt"
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := commands.Parse(commandLine) // parse the line to get an instance of Command
			eventLoop.Post(cmd)
		}
	}

	eventLoop.AwaitFinish()
}