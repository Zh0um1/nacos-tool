package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nacos/core"
	"nacos/core/flag"
)

func init() {
	getCmd.Flags().StringVarP(&flag.Namespace, "namespace", "n", "", "namespace to get or export, default: public, options")
	getCmd.Flags().StringVarP(&flag.Filename, "filename", "o", "", "export filename, options")
	exportCmd.Flags().StringVarP(&flag.Namespace, "namespace", "n", "", "namespace to get or export, default: public, options")
	exportCmd.Flags().StringVarP(&flag.Filename, "filename", "o", "", "export filename, options")
}

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all namespaces except public",
	GroupID: "config",
	Run: func(cmd *cobra.Command, args []string) {
		if flag.Target == "" {
			_ = cmd.Help()
			return
		}
		core.InitRequest()
		namespaces := core.GetNamespaces()
		fmt.Println(namespaces)
	},
}

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get config from nacos",
	Long:    "Get the configuration in the specified namespace, default: public",
	GroupID: "config",
	Run: func(cmd *cobra.Command, args []string) {
		if flag.Target == "" {
			_ = cmd.Help()
			return
		}
		core.InitRequest()
		configs, err := core.GetConfig(flag.Namespace)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, item := range configs {
			fmt.Printf("id: %v,\ndataId: %v,\ngroup: %v,\ncontent: %v,\ntenant:%v,\nappName: %v,\ntype: %v\n", item.Id, item.DataId, item.Group, item.Content, item.Tenant, item.AppName, item.AppName)
			fmt.Println()
		}
	},
}

var exportCmd = &cobra.Command{
	Use:     "export",
	Short:   "Export config from nacos",
	Long:    "Export the configuration in the specified namespace, if not provided will export all namespace",
	GroupID: "config",
	Run: func(cmd *cobra.Command, args []string) {
		if flag.Target == "" {
			_ = cmd.Help()
			return
		}
		core.InitRequest()
		err := core.Export(flag.Namespace, flag.Filename)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
