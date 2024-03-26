package main

import (
	"fmt"
	"os"
	"template_cli/internal/args"

	"github.com/bradfordwagner/go-util/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "testcli",
}

var myVerb = &cobra.Command{
	Use:   "myVerb",
	Short: "myVerb does something",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	},
	Run: func(cmd *cobra.Command, args []string) {
		l := log.Log().With("cmd", "myVerb")
		l.Info("hi friends")
	},
}

var defaultArgs args.Args

func init() {
	rootCmd.AddCommand(myVerb)
}

func main() {
	// viper - environment variables
	viper.AutomaticEnv()
	viper.Unmarshal(&defaultArgs) // when verbs have divergent args this will need to be moved into cmd specific methods
	// cobra
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
