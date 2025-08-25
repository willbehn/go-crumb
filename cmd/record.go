package cmd

import (
	"encoding/json"
	"os"
	"willbehn/ht/internal"
	"willbehn/ht/models"

	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

func init() { rootCmd.AddCommand(recordCmd) }

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "fiks senere",
	RunE: func(cmd *cobra.Command, args []string) error {

		var ev models.CmdEvent

		if err := json.NewDecoder(os.Stdin).Decode(&ev); err != nil {
			return err
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
