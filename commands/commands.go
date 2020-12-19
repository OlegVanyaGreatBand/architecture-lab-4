package commands

import (
	"fmt"
	"github.com/OlegVanyaGreatBand/architecture-lab-4/engine"
	"strings"
)

type PrintCmd struct {
	Msg string
}

func (pCmd *PrintCmd) Execute(h engine.Handler) {
	fmt.Println(pCmd.Msg)
}

type SplitCmd struct {
	Str string
	Sep string
}

func (sCmd *SplitCmd) Execute(h engine.Handler) {
	splitted := strings.Split(sCmd.Str, sCmd.Sep)
	joined := strings.Join(splitted, "\n")
	h.Post(&PrintCmd{ joined });
}

