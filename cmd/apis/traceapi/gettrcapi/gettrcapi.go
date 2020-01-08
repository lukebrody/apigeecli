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

package gettrcapi

import (
	"github.com/spf13/cobra"
	"github.com/srinandan/apigeecli/client/apis"
)

//Cmd to manage tracing of apis
var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get a debug session for an API proxy revision",
	Long:  "Get a debug session for an API proxy revision deployed in an environment",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = apis.GetTraceSession(name, revision, sessionID, messageID)
		return
	},
}

var name, sessionID, messageID string
var revision int

func init() {

	Cmd.Flags().StringVarP(&name, "name", "n",
		"", "API proxy name")
	Cmd.Flags().IntVarP(&revision, "rev", "v",
		-1, "API Proxy revision")
	Cmd.Flags().StringVarP(&sessionID, "ses", "s",
		"", "Debug session Id")
	Cmd.Flags().StringVarP(&messageID, "msg", "m",
		"", "Debug session Id")

	_ = Cmd.MarkFlagRequired("name")
	_ = Cmd.MarkFlagRequired("rev")
	_ = Cmd.MarkFlagRequired("ses")
}
