/*
Copyright © 2026 jaronnie <jaron@jaronnie.com>

A simple example jzero plugin that demonstrates the plugin system.
*/

package main

import (
	"fmt"
	"os"

	"github.com/jzero-io/jzero/cmd/jzero/pkg/plugin"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jzero-hello",
	Short: "A simple hello world plugin for jzero",
	Long: `This is an example plugin for jzero that demonstrates
how to create and structure a jzero plugin.

Plugins must be named with the "jzero-" prefix and
are installed in ~/.jzero/plugins directory.`,
}

var descCmd = &cobra.Command{
	Use:   "desc",
	Short: "Show plugin descriptor information",
	Long:  `Display information about API, Proto, and Model specifications found in the plugin.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := plugin.New()
		if err != nil {
			return err
		}

		if p.Desc.Api.SpecMap != nil {
			fmt.Printf("get %d api file\n", len(p.Desc.Api.SpecMap))
			for file := range p.Desc.Api.SpecMap {
				fmt.Printf("api file: %s\n", file)
			}
		}

		if p.Desc.Proto.SpecMap != nil {
			fmt.Printf("get %d proto file\n", len(p.Desc.Proto.SpecMap))
			for file := range p.Desc.Proto.SpecMap {
				fmt.Printf("proto file: %s\n", file)
			}
		}

		if p.Desc.Model.SpecMap != nil {
			fmt.Printf("get %d model file\n", len(p.Desc.Model.SpecMap))
			for file := range p.Desc.Model.SpecMap {
				fmt.Printf("model file: %s\n", file)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(descCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
