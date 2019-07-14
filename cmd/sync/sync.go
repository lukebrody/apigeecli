package sync

import (
	"../shared"
	"./getsync"
	"./setsync"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "sync",
	Short: "Manage identities with grant access to control plane resources",
	Long:  "Manage identities with grant access to control plane resources",
}

var expand = false
var count string

func init() {

	Cmd.PersistentFlags().StringVarP(&shared.RootArgs.Org, "org", "o",
		"", "Apigee organization name")

	Cmd.MarkPersistentFlagRequired("org")
	Cmd.AddCommand(setsync.Cmd)
	Cmd.AddCommand(getsync.Cmd)
}