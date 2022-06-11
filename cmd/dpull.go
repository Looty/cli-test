/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// dpullCmd represents the dpull command
var dpullCmd = &cobra.Command{
	Use:   "dpull",
	Short: "Runs docker pull with params",
	Long: `Runs docker pull with with options:
A well defined image: [image:tag]
A combination of image and tag

Examples:
cli-test dpull -f image:tag
cli-test dpull -i image -t tag (default is latest for tag)
	`,
	Run: func(cmd *cobra.Command, args []string) {
		pullFull, _ := cmd.Flags().GetString("full")
		pullImage, _ := cmd.Flags().GetString("image")
		pullTag, _ := cmd.Flags().GetString("tag")

		fmt.Println(pullFull, pullImage, pullTag)

		dockerPull(pullFull, pullImage, pullTag)
	},
}

func dockerPull(full string, image string, tag string) {
	input := ""
	if full != "" {
		input = full
	} else if image != "" {
		if tag == "" {
			tag = "latest"
		}
		input = fmt.Sprintf("%s:%s", image, tag)
	}

	if input == "" {
		log.Fatal("no image selected")
	}

	fmt.Println(input)
	var out, err = exec.Command("docker", "pull", input).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func init() {
	rootCmd.AddCommand(dpullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dpullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	dpullCmd.Flags().BoolP("full", "f", false, "Help message for full")
	dpullCmd.Flags().BoolP("image", "i", false, "Help message for image")
	dpullCmd.Flags().BoolP("tag", "t", false, "Help message for tag")
}
