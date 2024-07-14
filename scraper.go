package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var (
	keyPattern     = regexp.MustCompile(`"innertubeApiKey":"([^"]*)`)
	dataPattern    = regexp.MustCompile(`ytInitialData[^{]*(.*?);\s*<\/script>`)
	altDataPattern = regexp.MustCompile(`ytInitialData"[^{]*(.*);\s*window\["ytInitialPlayerResponse"\]`)
)

type videoResult struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	Duration     string `json:"duration"`
	ThumbnailUrl string `json:"thumbnail_src"`
	Views        string `json:"views"`
	Uploader     string `json:"username"`
}

type req struct {
	Version          string `json:"version"`
	Parser           string `json:"parser"`
	Key              string `json:"key"`
	EstimatedResults string `json:"estimatedResults"`
}

type videoRenderer struct {
	VideoId string `json:"videoId"`
	Title   struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"title"`
	LengthText struct {
		SimpleText string `json:"simpleText"`
	} `json:"lengthText"`
	Thumbnail struct {
		Thumbnails []struct {
			URL string `json:"url"`
		} `json:"thumbnails"`
	} `json:"thumbnail"`
	OwnerText struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"ownerText"`
	ViewCountText struct {
		SimpleText string   `json:"simpleText"`
		Runs       []string `json:"runs"`
	} `json:"viewCountText"`
	NavigationEndpoint struct {
		CommandMetadata struct {
			WebCommandMetadata struct {
				URL string `json:"url"`
			} `json:"webCommandMetadata"`
		} `json:"commandMetadata"`
	} `json:"navigationEndpoint"`
}

type ytSearchData struct {
	EstimatedResults string `json:"estimatedResults"`
	Contents         struct {
		TwoColumnSearchResultsRenderer struct {
			PrimaryContents struct {
				SectionListRenderer struct {
					Contents []struct { // sectionList
						ItemSectionRenderer struct {
							Contents []struct {
								ChannelRenderer  any           `json:"channelRenderer"`
								VideoRenderer    videoRenderer `json:"videoRenderer"`
								RadioRenderer    any           `json:"radioRenderer"`
								PlaylistRenderer any           `json:"playlistRenderer"`
							} `json:"contents"`
						} `json:"itemSectionRenderer"`
					} `json:"contents"`
				} `json:"sectionListRenderer"`
			} `json:"primaryContents"`
		} `json:"twoColumnSearchResultsRenderer"`
	} `json:"contents"`
}

func search(q string) ([]videoResult, error) {
	// get ze results
	url := "https://www.youtube.com/results?q=" + url.QueryEscape(q)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	jojo := req{
		Version: "0.1.5",
		Parser:  "json_format",
		Key:     "",
	}
	key := keyPattern.FindSubmatch(respBody)
	jojo.Key = string(key[1])

	matches := dataPattern.FindSubmatch(respBody)
	if len(matches) > 1 {
		jojo.Parser += ".object_var"
	} else {
		jojo.Parser += ".original"
		matches = altDataPattern.FindSubmatch(respBody)
	}
	data := ytSearchData{}
	err = json.Unmarshal(matches[1], &data)
	if err != nil {
		return nil, err
	}
	jojo.EstimatedResults = data.EstimatedResults

	// parse JSON data
	resSuka := make([]videoResult, 0)
	for _, sectionList := range data.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents {
		for _, content := range sectionList.ItemSectionRenderer.Contents {
			if content.VideoRenderer.VideoId == "" {
				continue
			}
			views := ""
			if content.VideoRenderer.ViewCountText.SimpleText != "" {
				views = content.VideoRenderer.ViewCountText.SimpleText
			} else if len(content.VideoRenderer.ViewCountText.Runs) > 0 {
				views = strings.Join(content.VideoRenderer.ViewCountText.Runs, "")
			}
			resSuka = append(resSuka, videoResult{
				Id:           content.VideoRenderer.VideoId,
				Title:        content.VideoRenderer.Title.Runs[0].Text,
				Url:          "https://youtube.com/watch?v=" + content.VideoRenderer.VideoId,
				Duration:     content.VideoRenderer.LengthText.SimpleText,
				ThumbnailUrl: content.VideoRenderer.Thumbnail.Thumbnails[len(content.VideoRenderer.Thumbnail.Thumbnails)-1].URL,
				Uploader:     content.VideoRenderer.OwnerText.Runs[0].Text,
				Views:        views,
			})
		}
	}

	return resSuka, nil
}
