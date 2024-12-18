package network

import (
	"fmt"
	"net"
	"time"
	"github.com/spf13/cobra"
)

var (
	host string
	port int
)

var PingCmd = &cobra.Command{
	Use: "ping-db",
	Short: "Check TCP service on a DB port",
	Long: "Just to check wether the TCP service is running on specified port and used for various reasons",
	Run: func(cmd *cobra.Command, args []string) {
		timeout := 2 * time.Second
		host := "localhost"
		port := 5432
		
		address := fmt.Sprintf("%s:%d", host, port)
		conn , err := net.DialTimeout("tcp", address, timeout)
		if err != nil { 
			fmt.Println("Database is NOT running on ", port)
			return 
		}
		fmt.Println("TCP service is running on",host,":",port)
		defer conn.Close()
	},
}

func init(){ 
	PingCmd.Flags().StringVarP(&host,"host","H","localhost","mentions the host")
	PingCmd.Flags().IntVarP(&port,"port","p",0,"PORT to connect")
	
	if err := PingCmd.MarkFlagRequired("port"); err != nil { 
		fmt.Print(err)
	}
}
