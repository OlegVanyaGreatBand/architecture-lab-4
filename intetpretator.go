package main

import (
	"bufio"
	"flag"
	"github.com/OlegVanyaGreatBand/architecture-lab-4/commands"
	"github.com/OlegVanyaGreatBand/architecture-lab-4/engine"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	inputFile = flag.String("f", "", "Expression to compute")
)
func main()  {
	flag.Parse()

	var source io.Reader

    if *inputFile != "" {
		data, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error())
			return
		}
		source = strings.NewReader(string(data))
	} else {
		_, _ = os.Stderr.WriteString("No expression provided")
		return
	}

	eventLoop := new(engine.Loop)
	eventLoop.Start()
		scanner := bufio.NewScanner(source)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := commands.Parse(commandLine) // parse the line to get an instance of Command
			eventLoop.Post(cmd)
		}
	eventLoop.AwaitFinish()
}
