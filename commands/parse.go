package commands

import (
	"fmt"
	"github.com/OlegVanyaGreatBand/architecture-lab-4/engine"
	"strings"
)

func Parse(command string) engine.Command {
	splitted := strings.FieldsFunc(command, func(r rune) bool {
		return r == ' '
	})


	var parsed engine.Command = &PrintCmd{
		fmt.Sprintf("ParsingError: Invalid command: %s", command),
	}

	l := len(splitted)
	switch command := splitted[0]; command {
	case "print":
		if l != 2 {
			parsed = &PrintCmd{
				fmt.Sprintf("SyntaxError: Invalid count of arguments for print: %d", l),
			}
		} else {
			parsed = &PrintCmd{splitted[1]}
		}
	case "split":
		if l != 3 {
			parsed = &PrintCmd{
				fmt.Sprintf("SyntaxError: Invalid count of arguments for split: %d", l),
			}
		} else {
			parsed = &SplitCmd{
				Str: splitted[1],
				Sep: splitted[2],
			}
		}
	}

	return parsed
}


