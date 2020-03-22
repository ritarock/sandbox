package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Bookmark struct {
	Checksum srting `json:"checksum"`
	Roots    struct {
		BookmarkBar struct {
			Children []struct {
				Children []struct {
					DataAdded string `json:"data_added"`
					Guid      string `json:"guid"`
					Id        string `json:"id"`
					MetaInfo  struct {
						LastVisitedDesktop string `json:"last_visited_desktop`
					}
					Name string `json:"name"`
					Type string `json:"type"`
					Url  string `json:"url"`
				} `json:"children"`
			} `json:"children"`
		} `json:"bookmark_bar"`
		Other struct {
			Children  []struct{} `json:"children"`
			DataAdded string     `json:"data_added"`
			Guid      string     `json:"guid"`
			Id        string     `json:"id"`
			MetaInfo  struct {
				LastVisitedDesktop string `json:"last_visited_desktop`
			}
			Name string `json:"name"`
			Type string `json:"type"`
		}
	} `json:"roots`
	Version int `json:"version"`
}

func main() {
	bytes, err := ioutil.ReadFile("json.json")
	if err != nil {
		log.Fatal(err)
	}

	var bookmark []Bookmark
	if err := json.Unmarshal(bytes, &bookmark); err != nil {
		log.Fatal(err)
	}

	for _, p := range bookmark {
		fmt.Printf("%d : %s\n", p.Roots.BookmarkBar.Children.Url)
	}
}
