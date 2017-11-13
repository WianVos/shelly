// MIT License

// Copyright (c) 2017 Wian Vos

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
// SOFTWARE.

package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/wianvos/shelly/pkg/templates"
)

var subDirs = []string{"tasks", "vars", "templates", "test", "files", "handlers", "meta"}
var outputDir string
var roleName string
var license string

// CreateCommand represents the base command when called without any subcommands
var CreateCommand = &cobra.Command{
	Use:   "create",
	Short: "create a new boilerplate role project",
	Run:   createCommand,
}

func init() {
	CreateCommand.PersistentFlags().StringVarP(&roleName, "name", "n", "", "the name of the role to create")
	CreateCommand.PersistentFlags().StringVarP(&outputDir, "out", "o", getCurrentDir()+roleName, "specify an output Directory for the code")
	CreateCommand.PersistentFlags().StringVarP(&license, "license", "l", "MIT", "specify a license to use in regard to this role")

	RootCmd.AddCommand(CreateCommand)

}

func createCommand(cmd *cobra.Command, args []string) {

	basePath := filepath.Join(outputDir, roleName)
	fts := templates.NewFileTemplates()

	// get the variables we need to do our thing
	data := splitArgsIntoKVPair(args[0])
	fts.AddData(data)

	// create directory

	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		jww.FATAL.Printf("%s: unable to create directory: %s -> %s", cmd.CommandPath(), basePath, err)
	}
	// create sub directories
	for _, s := range subDirs {
		sp := filepath.Join(basePath, s)
		err := os.MkdirAll(sp, os.ModePerm)
		if err != nil {
			jww.FATAL.Printf("%s: unable to create directory: %s -> %s", cmd.CommandPath(), sp, err)
		}
	}

	// create files

	for _, t := range fts {
		t.Write(basePath)

	}

}

func splitArgsIntoKVPair(s string) map[string]string {

	m := make(map[string]string)
	ss := strings.Split(s, ",")

	for _, pair := range ss {
		z := strings.Split(pair, "=")
		m[z[0]] = z[1]
	}

	return m
}
