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

package listkvm

import (
	"github.com/spf13/cobra"
	"github.com/srinandan/apigeecli/apiclient"
	"github.com/srinandan/apigeecli/client/kvm"
)

//Cmd to list kvms
var Cmd = &cobra.Command{
	Use:   "list",
	Short: "Returns a list of KVMs",
	Long:  "Returns a list of KVMs",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = kvm.List()
		return
	},
}

func init() {

	Cmd.Flags().StringVarP(apiclient.GetApigeeEnvP(), "env", "e",
		"", "Environment name")
	_ = Cmd.MarkFlagRequired("env")
}
