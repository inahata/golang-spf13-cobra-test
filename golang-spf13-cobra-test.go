package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

func main() {
    rootCmd := &cobra.Command {
        Version: "0.0.1",
        Use: "golang-spf13-cobra-test [command] [options]",
    }

    initCmd := &cobra.Command{
        Use:     "init",
        Short:   "初期化します",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("初期化しました\n")
        },
    }
    rootCmd.AddCommand(initCmd)

    err := rootCmd.Execute()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
