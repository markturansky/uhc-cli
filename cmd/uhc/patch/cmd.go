/*
Copyright (c) 2018 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package patch

import (
	"fmt"
	"os"

	"github.com/openshift-online/uhc-cli/cmd/uhc/common"
	"github.com/openshift-online/uhc-cli/cmd/uhc/urls"
	"github.com/spf13/cobra"

	"github.com/openshift-online/uhc-cli/pkg/util"
)

var args common.Args

var Cmd = &cobra.Command{
	Use:   "patch PATH",
	Short: "Send a PATCH request",
	Long:  "Send a PATCH request to the given path.",
	Run:   run,
}

func init() {
	flags := Cmd.Flags()
	common.AddCommonFlags(flags, &args)
	flags.StringVar(
		&args.Body,
		"body",
		"",
		"Name fo the file containing the request body. If this isn't given then "+
			"the body will be taken from the standard input.",
	)
}

func run(cmd *cobra.Command, argv []string) {
	path, err := urls.Expand(argv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create URI: %v\n", err)
		os.Exit(1)
	}
	connection, err := util.NewConnection(args.Debug)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create connection: %v\n", err)
		os.Exit(1)
	}

	// Create and populate the request:
	request := connection.Patch().Path(path)
	util.AddParamsAndHeaders(request, args.Parameter, args.Header)
	util.AddBody(request, args.Body)
	util.DoHTTP(request)
}
