// Copyright 2023 Google LLC
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

package env

import (
	"internal/apiclient"
	"internal/client/env"

	"github.com/spf13/cobra"
)

// GetSecReportCmd returns security incidents
var GetSecReportCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a security reports by name",
	Long:  "Returns a security reports by name",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		apiclient.SetApigeeEnv(environment)
		apiclient.SetRegion(region)
		return apiclient.SetApigeeOrg(org)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = env.GetSecurityReport(name)
		return
	},
}

func init() {
	GetSecReportCmd.Flags().StringVarP(&name, "name", "n",
		"", "Name of the security report")
	_ = GetSecReportCmd.MarkFlagRequired("name")
}
