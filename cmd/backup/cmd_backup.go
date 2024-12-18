package backup

import (
	pgbackup "github.com/bhat-abhishek/cubg/cmd/backup/pg-backup"
	"github.com/spf13/cobra"
)

var BackupCmd = &cobra.Command{
	Use: "backup",
	Short: "Create a backup/snapshot of a Database",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() { 
	BackupCmd.AddCommand(pgbackup.PostgresCmd)
}


