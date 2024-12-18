package pgbackup

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	backupType string
	cfgFile string
)

var PostgresCmd = &cobra.Command{
	Use: "pg",
	Short: "Takes a snap shot of Postgres Database",
	Run: func(cmd *cobra.Command, args []string) {
		dbName := viper.GetString("dbName")
		fmt.Print("Database Name ", dbName)
	},
}

func init() { 
	PostgresCmd.Flags().StringVarP(&backupType,"type","t","dump","Backup types like SQL dump, File System level and Continous Archiving")
	PostgresCmd.Flags().StringVar(&cfgFile,"config","","config path of Configuration file")
	PostgresCmd.MarkFlagRequired("config")
	cobra.OnInitialize(initConfig)
}

func initConfig() { 
	if cfgFile != "" { 
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err != nil { 
		log.Fatalf("Error reading config file: %v", err)
	}

	fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())

}