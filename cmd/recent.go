package cmd

import (
	"fmt"
	"time"
	"willbehn/ht/internal"
	"willbehn/ht/models"

	"github.com/spf13/cobra"
)

var recentCmd = &cobra.Command{
	Use:   "recent",
	Short: "A brief description of your command",
	Long:  `Fiks senere`,
	RunE: func(cmd *cobra.Command, args []string) error {

		db, err := internal.OpenDB()

		if err != nil {
			return err
		}

		defer db.Close()

		tx, err := db.BeginTx(cmd.Context(), nil)

		if err != nil {
			return err
		}

		rows, err := tx.Query(`
			SELECT cmd, shell, dir, repo, branch, ts, exit_code, duration_ms  
			FROM commands
			ORDER BY ts DESC
			LIMIT 20`)

		if err != nil {
			tx.Rollback()
			return err
		}

		var results []models.CmdEvent

		for rows.Next() {
			var ev models.CmdEvent

			if err := rows.Scan(
				&ev.Cmd,
				&ev.Shell,
				&ev.Dir,
				&ev.Repo,
				&ev.Branch,
				&ev.TS,
				&ev.Exit,
				&ev.Dur,
			); err != nil {
				return err
			}
			results = append(results, ev)
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		for _, ev := range results {

			t := time.Unix(ev.TS, 0).Local()
			tRel := internal.TimeSince(t)

			fmt.Printf("\033[32m (%s) \033[0m %s\n", tRel, ev.Cmd)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(recentCmd)

}
