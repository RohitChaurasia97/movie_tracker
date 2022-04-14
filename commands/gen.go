// MIT License

// Copyright (c) The RAI Authors

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.)
package commands

import (
	"fmt"
	"os"

	str "github.com/retail-ai-inc/bean/string"
	"github.com/spf13/cobra"
)

var (
	// gen represents the `gen` command.
	genCmd = &cobra.Command{
		Use:   "gen [command]",
		Short: "Generate a resource or key.",
		Long:  `This command requires a sub command parameter to generate a resource like a file or server or a key.`,
	}
)

var (
	// genSecretCmd represents the `gen secret` command.
	genSecretCmd = &cobra.Command{
		Use:   "secret",
		Short: "Generate 32 character long random alphanumeric string to use as secret.",
		Long:  `This alphanumeric random string is useful to use as secret in env.json file.`,
		Args:  cobra.ExactArgs(0),
		Run:   genSecret,
	}
)

func init() {

	genCmd.AddCommand(genSecretCmd)
	rootCmd.AddCommand(genCmd)
}

func genSecret(cmd *cobra.Command, args []string) {

	// Generate a 32 character long random secret string.
	secret, err := str.GenerateRandomString(32, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(secret)
}
