package cmd

import (
	"log"

	"github.com/dastergon/strgz/lib"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search index of starred repositories",
	Long: `strgz is a CLI tool for Github that enables users to list, index and search
repositories that have starred themselves or from other Github users.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("Please specify a word to search for.")
		}

		word := args[0]

		index, err := lib.BleveIndex("starred.bleve")
		if err != nil || index == nil {
			log.Fatalln("Bleve index failure\n", err)
		}

		results, err := lib.SearchIndex(word, index)
		if err != nil {
			log.Fatalln("Index search failed\n", err)
		}

		lib.ShowResults(results, index)
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
}
