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
	file, _ := os.Create(path.Join("/mnt", vid.Title+".mp4"))
	defer file.Close()
	vid.Download(vid.Formats.Best(ytdl.FormatAudioEncodingKey)[1], file)
}
