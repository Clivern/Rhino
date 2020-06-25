// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "rhino",
	Short: `Work seamlessly with Rhino from the command line.

We'd love to hear your feedback at <https://github.com/Clivern/Rhino>`,
}

// Execute runs cmd tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
