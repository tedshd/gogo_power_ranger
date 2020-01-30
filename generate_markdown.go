package main

import (
    "fmt"
    "os"
    "time"
    "strings"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type YoutubeInfo struct {
    Category    interface{} `json:"category"`
    Description string      `json:"description"`
    Error       bool        `json:"error"`
    Errors      string      `json:"errors"`
    ID          string      `json:"id"`
    Tags        []string    `json:"tags"`
    Thumbnails  struct {
        Default struct {
            Height int    `json:"height"`
            URL    string `json:"url"`
            Width  int    `json:"width"`
        } `json:"default"`
        High struct {
            Height int    `json:"height"`
            URL    string `json:"url"`
            Width  int    `json:"width"`
        } `json:"high"`
        Medium struct {
            Height int    `json:"height"`
            URL    string `json:"url"`
            Width  int    `json:"width"`
        } `json:"medium"`
    } `json:"thumbnails"`
    Title string `json:"title"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    youtubeId := os.Args[1]

    fmt.Println("youtubeId: ", youtubeId)
    date := time.Now().Format("2006-01-02T15:04:05-0700")
    fmt.Println("date: ", date)

    resp, err := http.Get("https://tedshd.io/service/youtube_info/function.php?id=" + youtubeId + "&format=json")
	if err != nil {
		fmt.Println("http get Error: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

    var youtubeinfo = &YoutubeInfo{}

	errUnmarshal := json.Unmarshal(body, youtubeinfo)
	if errUnmarshal != nil {
		fmt.Println("errUnmarshal: ", errUnmarshal)
	}
    // fmt.Printf("%+v", youtubeinfo)
    fmt.Println(youtubeinfo.Error)
    if youtubeinfo.Error == true {
        fmt.Println("get youtubeinfo: ", youtubeinfo.Errors)
        return
    }

    titleOrigin := strings.Split(youtubeinfo.Title, "|")
    title := strings.TrimSpace(titleOrigin[0])

    fmt.Println(title)
    fmt.Println(youtubeinfo.Description)
    fmt.Println(youtubeinfo.ID)
    fmt.Println(youtubeinfo.Tags)
    fmt.Println(youtubeinfo.Thumbnails.High.URL)

    var tags string
    var tag string
    for i := 0; i < len(youtubeinfo.Tags); i++ {
        tag = "  - \"" + youtubeinfo.Tags[i] + "\"\n"
        tags += tag
    }

    tm := "---\n" +
        "date: " + date + "\n" +
        "title: \"" + title + "\"\n" +
        "description: \"" + youtubeinfo.Description + "\"\n" +
        "images: [\"" + youtubeinfo.Thumbnails.High.URL + "\"]\n" +
        "categories:\n" +
        "  - \"tutorial\"\n" +
        "  - \"feature\"\n" +
        "tags:\n" +
        tags +
        "# menu: main # Optional, add page to a menu. Options: main, side, footer\n" +
        "\n" +
        "# Theme-Defined params\n" +
        "comments: false # Enable Disqus comments for specific page\n" +
        "authorbox: true # Enable authorbox for specific page\n" +
        "toc: false # Enable Table of Contents for specific page\n" +
        "mathjax: true # Enable MathJax for specific page\n" +
        "\n" +
        "howto: false\n" +
        "\n" +
        "video:\n" +
        "  name: \"" + title + "\"\n" +
        "  description: \"" + youtubeinfo.Description + "\"\n" +
        "  thumbnailUrl: \"" + youtubeinfo.Thumbnails.High.URL + "\"\n" +
        "  image: \"" + youtubeinfo.Thumbnails.High.URL + "\"\n" +
        "  vid: \"" + youtubeinfo.ID + "\"\n" +
        "  duration: \"PT54S\"\n" +
        "  hasPart:\n" +
        "      -\n" +
        "        name: \"\"\n" +
        "        startoffset: 10\n" +
        "        endoffset: 14\n" +
        "  step:\n" +
        "      -\n" +
        "        name: \"\"\n" +
        "        text: \"\"\n" +
        "        image: \"\"\n" +
        "        url: \"\"\n" +
        "      -\n" +
        "        name: \"\"\n" +
        "        itemlistelement:\n" +
        "          -\n" +
        "            text: \"\"\n" +
        "          -\n" +
        "            text: \"\"\n" +
        "        image: \"\"\n" +
        "        url: \"\"\n" +
        "\n" +
        "---\n" +
        "\n" +
        youtubeinfo.Description + "\n" +
        "\n" +
        "{{< youtubelazy " + youtubeinfo.ID + " \"" + title + "\" >}}"

    content := []byte(tm)
    errFile := ioutil.WriteFile(strings.ReplaceAll(title, " ", "_") + ".md", content, 0644)
    check(errFile)

}
