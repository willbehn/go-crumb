package cmd

import (
	"willbehn/ht/internal"
	"willbehn/ht/models"

	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

var (
	recCmdStr string
	recShell  string
	recDir    string
	recRepo   string
	recBranch string
	recTS     int64
	recExit   int
	recDur    int64
)

func init() {
	f := recordCmd.Flags()
	f.StringVar(&recCmdStr, "cmd", "", "full command line")
	f.StringVar(&recShell, "shell", "zsh", "shell name (zsh/bash/fish)")
	f.StringVar(&recDir, "dir", "", "working directory")
	f.StringVar(&recRepo, "repo", "", "git repo (optional)")
	f.StringVar(&recBranch, "branch", "", "git branch (optional)")
	f.Int64Var(&recTS, "ts", 0, "unix timestamp (seconds)")
	f.IntVar(&recExit, "exit", 0, "exit code")
	f.Int64Var(&recDur, "dur", 0, "duration ms (optional)")
	rootCmd.AddCommand(recordCmd)
}

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "fiks senere",
	RunE: func(cmd *cobra.Command, args []string) error {

		ev := models.CmdEvent{
			Cmd:    recCmdStr,
			Shell:  recShell,
			Dir:    recDir,
			Repo:   strptr(recRepo),
			Branch: strptr(recBranch),
			TS:     recTS,
			Exit:   recExit,
			Dur:    recDur,
		}

		db, err := internal.OpenDB()

		if err != nil {
			return err
		}

		defer db.Close()

		tx, err := db.BeginTx(cmd.Context(), nil)
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(cmd.Context(),
			`INSERT INTO commands (ts,shell,dir,repo,branch,cmd,exit_code,duration_ms)
			 VALUES (?,?,?,?,?,?,?,?)`,
			ev.TS, ev.Shell, ev.Dir, ev.Repo, ev.Branch, ev.Cmd, ev.Exit, ev.Dur)

		if err != nil {
			_ = tx.Rollback()
			return err
		}

		return tx.Commit()
	},
}

func strptr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
