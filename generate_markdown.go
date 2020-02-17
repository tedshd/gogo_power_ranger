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

type YoutubeResponse struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
				Standard struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"standard"`
			} `json:"thumbnails"`
			ChannelTitle         string   `json:"channelTitle"`
			Tags                 []string `json:"tags"`
			CategoryID           string   `json:"categoryId"`
			LiveBroadcastContent string   `json:"liveBroadcastContent"`
			Localized            struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"localized"`
		} `json:"snippet"`
	} `json:"items"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    youtubeId := os.Args[1]
    // youtubeId := "A8iFrfCxbYs"
    youtubeAPIKey := ""

    fmt.Println("youtubeId: ", youtubeId)
    date := time.Now().Format("2006-01-02T15:04:05-0700")
    fmt.Println("date: ", date)

    resp, err := http.Get("https://www.googleapis.com/youtube/v3/videos?part=snippet&id=" + youtubeId + "&key=" + youtubeAPIKey)
	if err != nil {
		fmt.Println("http get Error: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

    var YoutubeResponse = &YoutubeResponse{}

	errUnmarshal := json.Unmarshal(body, YoutubeResponse)
	if errUnmarshal != nil {
		fmt.Println("errUnmarshal: ", errUnmarshal)
    }

    // fmt.Printf("%+v", YoutubeResponse)

    youtubeinfo := YoutubeResponse.Items[0].Snippet

    titleOrigin := strings.Split(youtubeinfo.Title, "|")
    title := strings.TrimSpace(titleOrigin[0])

    fmt.Println(title)
    fmt.Println(youtubeinfo.Description)
    fmt.Println(YoutubeResponse.Items[0].ID)
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
        "  vid: \"" + YoutubeResponse.Items[0].ID + "\"\n" +
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
        "{{< youtubelazy " + YoutubeResponse.Items[0].ID + " \"" + title + "\" >}}"

    content := []byte(tm)
    errFile := ioutil.WriteFile(strings.ReplaceAll(title, " ", "_") + ".md", content, 0644)
    check(errFile)

}
