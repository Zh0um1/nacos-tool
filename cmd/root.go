package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nacos/core/flag"
	"os"
)

func init() {

	rootCmd.PersistentFlags().StringVarP(&flag.Target, "target", "t", "", "nacos address, eg: http://127.0.0.1:8848")
	rootCmd.PersistentFlags().StringVar(&flag.Proxy, "proxy", "", "socks or http proxy, eg: socks5:127.0.0.1:8080")
	rootCmd.PersistentFlags().StringVarP(&flag.Key, "key", "k", "", "JWT sign key")
	configGroup := &cobra.Group{
		ID:    "config",
		Title: "Get or export config from nacos command",
	}
	userGroup := &cobra.Group{
		ID:    "user",
		Title: "Add user to nacos command",
	}
	rootCmd.AddGroup(configGroup, userGroup)
	rootCmd.AddCommand(userCmd, getCmd, exportCmd, listCmd)
}

var rootCmd = &cobra.Command{
	Use:   "nacos",
	Short: "nacos unauthorized exploit tools",
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
