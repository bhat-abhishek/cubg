package cmd

import (
	"github.com/bhat-abhishek/cubg/cmd/backup"
	"github.com/bhat-abhishek/cubg/cmd/network"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "cubg",
	Short: "CUBG - Centralized Universal Backup Grid",
	Long: "CUBG - Centralized Universal Backup Grid. Built to backup database and restore across all kind of infrastructure at ease",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubCommands() { 
	rootCmd.AddCommand(network.PingCmd)
	rootCmd.AddCommand(backup.BackupCmd)
}

func init() { 
	addSubCommands()
}

func Execute() error{ 
	return rootCmd.Execute()
}