// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sharedflows

import (
	"github.com/spf13/cobra"
	"github.com/srinandan/apigeecli/apiclient"
	"github.com/srinandan/apigeecli/client/sharedflows"
)

//Cmd to delete shared flow
var DelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a shared flow",
	Long: "Deletes a shared flow and all associated policies, resources, and revisions." +
		"The flow must be undeployed first.",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		apiclient.SetApigeeOrg(org)
		apiclient.SetApigeeEnv(env)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = sharedflows.Delete(name, revision)
		return
	},
}

func init() {
	DelCmd.Flags().StringVarP(&name, "name", "n",
		"", "shared flow name")
	DelCmd.Flags().IntVarP(&revision, "rev", "v",
		-1, "shared flow revision")

	_ = DelCmd.MarkFlagRequired("name")
}