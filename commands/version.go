// Copyright (c) Alex Ellis 2017. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GitCommit injected at build-time
var GitCommit string
var (
	shortVersion bool
)

func init() {
	versionCmd.Flags().BoolVar(&shortVersion, "short-version", false, "Just print Git SHA")

	faasCmd.AddCommand(versionCmd)
}

// versionCmd displays version information
var versionCmd = &cobra.Command{
	Use:   "version [--short-version]",
	Short: "Display the clients version information",
	Long: fmt.Sprintf(`The version command returns the current clients version information.

This currently consists of the GitSHA from which the client was built.
- https://github.com/alexellis/faas-cli/tree/%s`, GitCommit),
	Example: `  faas-cli version
  faas-cli version --short-version`,
	Run: runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {

	if shortVersion {
		fmt.Println(GitCommit)
	} else {
		fmt.Printf(figletStr)
		fmt.Printf("Git Commit: %s\n", GitCommit)
	}
	return
}

const figletStr = `  ___                   _____           ____  
 / _ \ _ __   ___ _ __ |  ___|_ _  __ _/ ___| 
| | | | '_ \ / _ \ '_ \| |_ / _` + "`" + ` |/ _` + "`" + ` \___ \ 
| |_| | |_) |  __/ | | |  _| (_| | (_| |___) |
 \___/| .__/ \___|_| |_|_|  \__,_|\__,_|____/ 
      |_|                                     

`
