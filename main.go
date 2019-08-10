package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/dustin/go-wikiparse"
)

func fname(title string) string {
	return strings.Map(func(r rune) rune {
		if r == '/' || r == ' ' {
			return '-'
		}
		return r
	}, title)
}

func exclude(title string) bool {
	switch {
	case
		strings.HasPrefix(title, "Category comments:"),
		strings.HasPrefix(title, "Category:"),
		strings.HasPrefix(title, "Comments:"),
		strings.HasPrefix(title, "Image:"),
		strings.HasPrefix(title, "JAMWiki:"),
		strings.HasPrefix(title, "Template comments:"),
		strings.HasPrefix(title, "Template:"),
		strings.HasPrefix(title, "User comments:"),
		strings.HasPrefix(title, "User:"):
		return true
	}
	return false
}

type byTime []wikiparse.Revision

func (b byTime) Len() int          { return len(b) }
func (b byTime) Swap(i int, j int) { b[i], b[j] = b[j], b[i] }
func (b byTime) Less(i int, j int) bool {
	x, _ := time.Parse(time.RFC3339, b[i].Timestamp)
	y, _ := time.Parse(time.RFC3339, b[j].Timestamp)
	return x.After(y)
}

func main() {
	fd, err := os.Open("LOR-Wiki_All_Pages_2015-06-18.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	p, err := wikiparse.NewParser(fd)
	if err != nil {
		log.Fatal(err)
	}
	for err == nil {
		var page *wikiparse.Page
		page, err = p.Next()
		if err == nil {
			fmt.Println(page.Title, len(page.Revisions))
			sort.Sort(byTime(page.Revisions))
			rev := page.Revisions[0]
			if len(rev.Text) == 0 {
				continue
			}
			name := path.Join("dump", fname(page.Title)+".wiki")
			os.MkdirAll(path.Dir(name), 0755)
			if err := ioutil.WriteFile(name, []byte(rev.Text), 0644); err != nil {
				log.Fatal(err)
			}
		}
	}
}
