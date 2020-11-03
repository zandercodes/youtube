package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print metadata of the desired video",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		video, err := getDownloader().GetVideo(args[0])
		exitOnError(err)

		fmt.Printf("Title:    %s\n", video.Title)
		fmt.Printf("Author:   %s\n", video.Author)
		fmt.Printf("Duration: %v\n", video.Duration)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"itag", "quality", "MimeType"})

		for _, itag := range video.Formats {
			table.Append([]string{strconv.Itoa(itag.ItagNo), itag.Quality, itag.MimeType})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
