package cmd

import (
	"strings"
	"willbehn/ht/internal"
	"willbehn/ht/models"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long:  `Fiks senere`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := internal.OpenDB()

		if err != nil {
			return err
		}
		defer db.Close()

		var conditions []string
		var parameters []any

		query := `SELECT cmd, shell, dir, repo, branch, ts, exit_code, duration_ms  
		FROM commands `

		for _, arg := range args {
			conditions = append(conditions, " cmd LIKE ?")
			parameters = append(parameters, "%"+arg+"%")
		}

		if len(conditions) > 0 {
			query += "WHERE" + strings.Join(conditions, "AND")
		}

		rows, err := db.Query(query, parameters...)

		if err != nil {
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

		internal.ResultOutputShort(results)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
