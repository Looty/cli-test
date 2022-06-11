/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set your config file",
	Long: `Setting your configuration via a .yaml file
Example:
clit-test config -f config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		file, _ := cmd.Flags().GetString("file")

		if file == "" {
			log.Fatal("You must select a file")
		}

		fmt.Println("file:", file)
		viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
		viper.SetConfigName(file)   // name of config file (without extension)
		viper.AddConfigPath("config/")

		errFile := viper.ReadInConfig() // Find and read the config file
		if errFile != nil {             // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", errFile))
		}

		fmt.Println(viper.GetString("email"))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	configCmd.Flags().StringP("file", "f", "", "Config file")
}
