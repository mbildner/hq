package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	htmlBytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println("error reading from stdin!")
		os.Exit(1)
	}

	query := os.Args[1]

	if strings.HasPrefix(query, "attr") {
		re := regexp.MustCompile(`"(.*)"`)
		str := re.FindStringSubmatch(query)[1]

		doc, _ := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlBytes))
		q := doc.Find("body > *")

		q.Each(func(_ int, s *goquery.Selection) {
			attributeValue, exists := s.Attr(str)
			if exists {
				fmt.Println(attributeValue)
			}
		})

	} else if strings.HasPrefix(query, "text") {
		doc, _ := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlBytes))
		q := doc.Find("body > *")

		q.Each(func(_ int, s *goquery.Selection) {
			fmt.Println(s.Text())
		})

	} else {
		doc, _ := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlBytes))
		doc.Find(query).Each(func(_ int, s *goquery.Selection) {
			html, _ := goquery.OuterHtml(s)
			fmt.Println(html)
		})
	}

}
