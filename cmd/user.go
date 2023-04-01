package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nacos/core"
	"nacos/core/flag"
)

func init() {
	userCmd.Flags().StringVarP(&flag.Username, "username", "u", "", "username to add nacos")
	userCmd.Flags().StringVarP(&flag.Password, "password", "p", "", "password for user")
}

var userCmd = &cobra.Command{
	Use:     "adduser",
	Short:   "Add user",
	Long:    "Add the specified user name and password to the system.\nIf no user name and password are provided, they are random values",
	GroupID: "user",
	Run: func(cmd *cobra.Command, args []string) {
		if flag.Target == "" {
			_ = cmd.Help()
			return
		}
		core.InitRequest()
		err := core.AddUser(flag.Username, flag.Password)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
