package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Book is a list of book items to help store data associated with books.
type Book struct {
	Author string `xml:"author,attr"`
	Title  string `xml:"title,attr"`
	Year   string `xml:"hyr,attr"`
	ID     string `xml:"owi,attr"`
}

func main() {
	books, err := search("Harry Potter")
	checkError(err)

	fmt.Println(books)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {

	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

// ClassifySearchResponse is a struct for the results it is built from books.
type ClassifySearchResponse struct {
	Results []Book `xml:"works>work"`
}

func search(query string) ([]Book, error) {
	var resp *http.Response
	var err error

	if resp, err = http.Get("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query)); err != nil {
		return []Book{}, err
	}

	defer resp.Body.Close()
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return []Book{}, err
	}

	var c ClassifySearchResponse
	err = xml.Unmarshal(body, &c)

	return c.Results, err
}
