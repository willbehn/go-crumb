package internal

import (
	"fmt"
	"time"
	"willbehn/ht/models"
)

func ResultOutputShort(results []models.CmdEvent) {
	for _, ev := range results {

		t := time.Unix(ev.TS, 0).Local()
		tRel := TimeSince(t)

		//fmt.Printf("(%s) %s\n", tRel, ev.Cmd)
		fmt.Printf("\033[32m %-8s \033[0m %s %d %s\n", tRel, ev.Cmd, ev.Exit, ev.Dir)
	}
}
