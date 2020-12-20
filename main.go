package main

import (
	"bufio"
	"errors"
	"flag"
	"github.com/OlegVanyaGreatBand/architecture-lab-4/commands"
	"github.com/OlegVanyaGreatBand/architecture-lab-4/engine"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	inputFile = flag.String("f", "", "Input file name")
)

func scanFile(inputFile *string, loop *engine.EventLoop) error {
	var source io.Reader

	if *inputFile != "" {
		data, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			return err
		}
		source = strings.NewReader(string(data))
	} else {
		return errors.New("no file provided")
	}

	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		commandLine := scanner.Text()
		cmd := commands.Parse(commandLine)
		loop.Post(cmd)
	}

	return nil
}

func main()  {
	flag.Parse()

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if err := scanFile(inputFile, eventLoop); err != nil {
		_, _ = os.Stderr.WriteString(err.Error() + "\n")
	}

	eventLoop.AwaitFinish()
}
