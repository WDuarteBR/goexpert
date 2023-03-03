/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// name, _ := cmd.Flags().GetString("name")
		fmt.Println("Category name called with name " + category)
		enable, _ := cmd.Flags().GetBool("enable")
		fmt.Println("Category enable called with " + fmt.Sprint(enable))
		id, _ := cmd.Flags().GetInt16("id")
		fmt.Println("Category id called with " + fmt.Sprint(id))
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamdo antes do RUN !")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado após o RUN !")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("Ocorreu um erro!")
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	// categoryCmd.PersistentFlags().StringP("name", "n", "w", "nome da categoria") // com StringP posso usar nomes abreviados
	categoryCmd.PersistentFlags().BoolP("enable", "e", false, "Categoria é válida")
	categoryCmd.PersistentFlags().Int16P("id", "i", 0, "Id da categoria")
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "Nome da categoria")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
