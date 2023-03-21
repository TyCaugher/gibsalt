package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	fileName string
	subject  string
)

func getE6Posts(tags []string, pages int) {
	client := http.Client{}

	// Create the new url
	baseurl := "https://e621.net/posts.json?page="
	//url += strings.Join(tags[:], "+")

	for p := 1; p <= pages; p++ {

		url := baseurl + strconv.Itoa(p) + "&tags=" + strings.Join(tags[:], "+") // Incrament the page up and make the url

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		req.Header.Set("_client", "saucepls v0.1 (Kroren)")
		resp, err := client.Do(req)
		defer resp.Body.Close()

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var responseObject E6Response
		json.Unmarshal(responseData, &responseObject)

		fmt.Println("Found ", len(responseObject.Posts), " posts")

		downloadPostImages(responseObject)
		fmt.Println("Done downloading page " + strconv.Itoa(pages))
	}

}

func checkDir(path string) {
	// If directory exists, make it, return false
	// if true, continue
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func downloadPostImages(e6Data E6Response) {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	for index, post := range e6Data.Posts {

		fileName = post.Tags.Artist[0] + "_" + strconv.Itoa(post.ID) + "_." + post.File.Ext
		path := "e6downloads\\" + subject + "\\"
		// Check if the directory exists.

		file, err := os.Create(path + fileName)
		if err != nil {
			log.Println(err)
		}

		// create the request
		resp, err := client.Get(post.File.URL)
		if err != nil {
			log.Fatal(err)
		}
		println(string(index))
		defer resp.Body.Close()

		size, err := io.Copy(file, resp.Body)
		defer file.Close()

		fmt.Printf("Downloaded a file %s with size %d", fileName, size)
	}
	println("Finished...")
}

func main() {
	/*
		Command usage
		saucepls  pages=N tags
	*/
	tags := os.Args[1:] // arguments in an array
	fmt.Println("CMD Args: ", tags)
	subject = tags[1]

	pageFlag := flag.Int("pages", 1, "page count")

	flag.Parse()

	checkDir("e6downloads\\" + subject)
	getE6Posts(tags, *pageFlag)
}
