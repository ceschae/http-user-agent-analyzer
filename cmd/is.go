/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"http-user-agent-analyzer/client"
	"log"

	"github.com/spf13/cobra"
)

// isCmd represents the is command
var isCmd = &cobra.Command{
	Use:   "is <string in quotes> <adjective>",
	Short: "Use this question word to expect a true/false response",
	Long: `An example usage:
'http-user-agent-analyzer is "./filepath/or/rawstring.txt" allowed'
The string in quotes is allowed to be a filepath or a verbatim string, either way it will work.
Currently, the only allowed adjective is 'allowed' (ironic) and it returns 'true' if the 
User-Agent header string indicates the HTTP request is from Firefox, false otherwise (espcially Safari).`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO(cait) the 'allowed' requirement isn't enforced, nor is the 'file path/raw string' detection,
		// it's all raw string at the moment
		log.Println("'is' called")
		err := client.Allow(args[0])
		if err != nil {
			log.Fatalf("error encountered when calling 'allowed': %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(isCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
