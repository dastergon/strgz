package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/blevesearch/bleve"
	"github.com/google/go-github/github"
)

func BleveIndex(indexName string) (bleve.Index, error) {
	index, err := bleve.Open(indexName)
	if err != nil {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexName, mapping)
	}
	return index, err
}

func Index(repo *github.Repository, index bleve.Index) error {
	document, err := json.Marshal(repo)
	if err != nil {
		log.Fatalln("Trouble JSON encoding")
	}
	err = index.Index(strconv.Itoa(repo.GetID()), string(document))
	return err
}

func SearchIndex(keyword string, index bleve.Index) (*bleve.SearchResult, error) {
	query := bleve.NewQueryStringQuery(keyword)
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		return nil, err
	}
	return searchResults, nil
}

func ShowResults(results *bleve.SearchResult, index bleve.Index) {
	if len(results.Hits) < 1 {
		fmt.Println(results)
	}
	for _, val := range results.Hits {
		id := val.ID
		doc, err := index.Document(id)
		if err != nil {
			fmt.Println(err)
		}
		for _, field := range doc.Fields {
			repo := github.Repository{}
			json.Unmarshal(field.Value(), &repo)
			fmt.Printf("%s - %s (%s)\n\t%s\n", *repo.Name, *repo.Description, *repo.Language, *repo.HTMLURL)
		}
	}
}
