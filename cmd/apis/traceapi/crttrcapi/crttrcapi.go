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

package crttrcapi

import (
	"github.com/spf13/cobra"
	"github.com/srinandan/apigeecli/client/apis"
)

//Cmd to manage tracing of apis
var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new debug session for an API proxy",
	Long:  "Create a new debug session for Apigee API proxy revision deployed in an environment",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = apis.CreateTraceSession(name, revision, filter)
		return
	},
}

var name string
var filter map[string]string
var revision int

func init() {

	Cmd.Flags().StringVarP(&name, "name", "n",
		"", "API proxy name")
	Cmd.Flags().IntVarP(&revision, "rev", "v",
		-1, "API Proxy revision")
	Cmd.Flags().StringToStringVar(&filter, "filter",
		nil, "Filter Conditions; format is name1=value1,name2=value2...")

	_ = Cmd.MarkFlagRequired("name")
	_ = Cmd.MarkFlagRequired("rev")
}
