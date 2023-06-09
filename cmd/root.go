/*
Copyright © 2023 Martin Nirtl <martin.nirtl@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	etcHosts          bool
	dryRun            bool
	hostsFilePath     string
	sshConfigFilePath string
)

var rootCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Manage address mappings of SSH config and optionally entries of hosts file",
	Long:  `Manage address mappings of SSH config and optionally entries of hosts file.`,
	// DisableFlagsInUseLine: true,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		if len(args) >= 0 {
			comps = cobra.AppendActiveHelp(comps, "No args expected")
		}
		return comps, cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.ExactArgs(0),
	Run:  Print,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Only print updated /etc/hosts and ~/.ssh/config files")
	rootCmd.PersistentFlags().StringVar(&sshConfigFilePath, "ssh-config", "", "Set SSH Config file (e.g. /etc/ssh/config); default: ~/.ssh/config")
	rootCmd.PersistentFlags().StringVar(&hostsFilePath, "hosts-file", "", "Set host file (e.g. ~/hosts); default: /etc/hosts")
	rootCmd.PersistentFlags().BoolVar(&etcHosts, "etc-hosts", false, "Additionally add entry to /etc/hosts file (requires sudo)")
}

func getFilePaths() error {
	if hostsFilePath == "" {
		hostsFilePath = "/etc/hosts"
	}
	if sshConfigFilePath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		sshConfigFilePath = fmt.Sprintf("%s/.ssh/config", homeDir)
	}

	return nil
}
