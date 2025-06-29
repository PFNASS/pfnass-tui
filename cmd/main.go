package cmd

import (
    "context"
    // "errors"

    "github.com/spf13/cobra"
    // "github.com/spf13/viper"
    // "golang.org/x/sync/errgroup"
)

// Execute is the command line applications entry function
func Execute() error {
    rootCmd := &cobra.Command{
        Version: "v0.0.1",
        Use:     "pfnassctl",
        Long:    "pfnassctl is a command line tool for managing my homelab server - PFNASS",
        Example: "pfnassctl --help",
        RunE: func(cmd *cobra.Command, args []string) error {
            return nil
        },
    }

    return rootCmd.ExecuteContext(context.Background())
}
