package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// videoCmd represents the video command
var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "subcommand for video manipulation",
	Long:  `subcommand for video manipulation, which can be used to manipulate YouTube videos, such as uploading, updating, deleting, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("video called")
	},
}

func init() {
	rootCmd.AddCommand(videoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// videoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// videoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
