# strgz
strgz is a CLI tool for Github that enables users to list, index and search starred repositories from their own Github account or others' account.

strgz comes in handy to users with thousands of starred repositories. strgz indexes the starred repositories and enables fast searching capabilities of the repositories. It uses the [Bleve](http://www.blevesearch.com/) indexing library to index all starred repositories and the [go-github](https://github.com/google/go-github) library to interact with the Github API. 

## Installation

    go get -u github.com/dastergon/strgz

## Examples

### Listing

List starred repositories of a user on the fly

    strgz ls username

List starred repositories of a user on the fly, but show only the URLs

    strgz ls username --url

List starred repositories of a user on the fly, but index repositories

    strgz ls username --index

### Searching

Search for a specific word in the index and show repositories

    strgz search vim

## Usage

        Usage:
            strgz [command]

        Available Commands:
            help        Help about any command
            ls          List starred repositories from a Github user
            search      Search index of starred repositories

        Flags:
                --config string   config file (default is $HOME/.strgz.yaml)
            -h, --help            help for strgz

        Use "strgz [command] --help" for more information about a command.
