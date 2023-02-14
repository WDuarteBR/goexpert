/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// testeCmd represents the teste command
var testeCmd = &cobra.Command{
	Use:   "teste",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		comando, _ := cmd.Flags().GetString("comando")
		if comando == "ping" {
			cmd.Print("PING")
		} else {
			cmd.Print("PONG")
		}
	},
}

func init() {
	rootCmd.AddCommand(testeCmd)
	testeCmd.Flags().StringP("comando", "c", "", "Escolha ping ou pong")
	testeCmd.MarkFlagRequired("comando")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
