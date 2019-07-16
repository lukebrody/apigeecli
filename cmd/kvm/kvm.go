package kvm

import (
	"../shared"
	"./crtkvm"
	"./delkvm"
	"./listkvm"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "kvms",
	Short: "Manage Key Value Maps",
	Long:  "Manage Key Value Maps",
}

var expand = false
var count string

func init() {

	Cmd.PersistentFlags().StringVarP(&shared.RootArgs.Org, "org", "o",
		"", "Apigee organization name")

	Cmd.MarkPersistentFlagRequired("org")
	Cmd.AddCommand(listkvm.Cmd)
	Cmd.AddCommand(delkvm.Cmd)
	Cmd.AddCommand(crtkvm.Cmd)
}
