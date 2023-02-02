/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"os"

	"github.com/spf13/cobra"

	dbInfra "github.com/infezek/product_ports_adapters_golang/adapters/db"
	"github.com/infezek/product_ports_adapters_golang/application"
)

var cfgFile string

var db, _ = sql.Open("sqlite3", "sqlite.db")
var productDb = dbInfra.NewProductDb(db)
var productService = application.ProductService{Persistence: productDb}

var rootCmd = &cobra.Command{
	Use:   "product_ports_adapters_golang",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
