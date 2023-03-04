/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wduartebr/goexpert/cli/internal/database"
)

func newCreateCommand(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create new category",
		Long:  "Create new category",
		RunE:  runCreate(categoryDb),
	}
}

func runCreate(category database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("description")

		_, err := category.Create(name, desc)
		if err != nil {
			return err
		}

		return nil
	}
}

func init() {
	createCmd := newCreateCommand(GetCategory(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of category")
	createCmd.Flags().StringP("description", "d", "", "Description of category")
	createCmd.MarkFlagsRequiredTogether("name", "description")

}
