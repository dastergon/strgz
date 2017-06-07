package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
	"github.com/dastergon/strgz/lib"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

var ghToken string
var index bleve.Index
var err error

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List starred repositories from a Github user",
	Long: `strgz is a CLI tool for Github that enables users to list, index and search
repositories that have starred themselves or from other Github users.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("Please specify a user to search for.")
		}

		user := args[0]
		isIndex, _ := cmd.Flags().GetBool("index")
		isUrl, _ := cmd.Flags().GetBool("url")
		isName, _ := cmd.Flags().GetBool("name")

		client := lib.ConnectGithub(ghToken)

		var starredRepos []*github.StarredRepository

		for {
			starred_repos, resp, err := client.Activity.ListStarred(context.Background(), user, nil)
			if err != nil {
				log.Fatalln(args[0], "is not a valid Github user\n", err)
			}
			starredRepos = append(starredRepos, starred_repos...)
			if resp.NextPage == 0 {
				break
			}
		}

		if isIndex {
			index, err = lib.BleveIndex("starred.bleve")
			if err != nil || index == nil {
				log.Fatalln("Bleve index failure\n", err)
			}
		}

		for _, repo := range starredRepos {
			if isIndex {
				err = lib.Index(repo.Repository, index)
				if err != nil {
					log.Fatalln("Failure to index data\n", err)
				}
			}
			if isUrl {
				fmt.Println(*repo.Repository.HTMLURL)
			} else if isName {
				fmt.Println(*repo.Repository.FullName)
			} else {
				fmt.Printf("%s ", *repo.Repository.Name)
				// some repos do not include a description
				if repo.Repository.Description != nil {
					fmt.Printf("- %s ", *repo.Repository.Description)
				}
				// some repos are not assigned to a specific language
				if repo.Repository.Language != nil {
					fmt.Printf("(%s)", *repo.Repository.Language)
				}
				fmt.Printf("\n\t%s\n", *repo.Repository.HTMLURL)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
	lsCmd.PersistentFlags().StringVarP(&ghToken, "token", "t", "", "Github authentication token")
	lsCmd.Flags().BoolP("index", "i", false, "Enable indexing of the starred repositories")
	lsCmd.Flags().BoolP("url", "u", false, "Show only URLs of the repository")
	lsCmd.Flags().BoolP("name", "n", false, "Show only the full name of the repository")
}
