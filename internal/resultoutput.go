package internal

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"willbehn/ht/models"
)

func ResultOutputShort(results []models.CmdEvent) {

	fmt.Printf("%-4s  %-8s  %-20s  %s\n", "ID", "SINCE", "DIR", "COMMAND")
	fmt.Println(strings.Repeat("-", 72))

	for _, ev := range results {
		t := time.Unix(ev.TS, 0).Local()
		since := TimeSince(t)

		fmt.Printf("%-4d  %-8s  %-20s  %s\n",
			ev.Id,
			since,
			prettyDir(ev.Dir),
			ev.Cmd,
		)
	}
}

func ResultOutputLong(results []models.CmdEvent) {
	fmt.Printf("%-4s  %-19s  %-7s  %-20s  %s\n", "ID", "TIME", "SHELL", "DIR", "COMMAND")
	fmt.Println(strings.Repeat("-", 80))

	for _, ev := range results {
		t := time.Unix(ev.TS, 0).Local()
		absT := t.Format("2006-01-02 15:04:05")

		fmt.Printf("%-4d  %-19s  %-7s  %-20s  %s\n",
			ev.Id,
			absT,
			ev.Shell,
			prettyDir(ev.Dir),
			ev.Cmd,
		)
	}
}

func prettyDir(dir string) string {
	if dir == "" {
		return "-"
	}

	clean := filepath.Clean(dir)
	split := strings.Split(clean, string(filepath.Separator))
	if len(split) > 3 {
		keep := split[len(split)-2:]
		return "…" + string(filepath.Separator) + filepath.Join(keep...)
	}
	return dir
}
