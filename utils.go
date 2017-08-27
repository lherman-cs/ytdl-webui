package main

import (
	"log"
	"os"
	"path"

	"github.com/otium/ytdl"
)

func Download(url string) {
	vid, err := ytdl.GetVideoInfo(url)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(path.Join("data", vid.Title+".mp4"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	vid.Download(vid.Formats.Best(ytdl.FormatAudioEncodingKey)[1], file)
}
