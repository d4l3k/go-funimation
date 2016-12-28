package main

import (
	"log"

	funimation "github.com/d4l3k/go-funimation"
	dlfunimation "golang.ssttevee.com/funimation/lib"
)

func main() {
	c := funimation.NewClient()
	_, err := c.Login("rice@fn.lc", "2oR1bkyKPTRv9ic3")
	if err != nil {
		log.Fatal(err)
	}
	queue, err := c.Queue(100, 0)
	if err != nil {
		log.Fatal(err)
	}
	dl := c.DownloadClient()
	for _, v := range queue.NextVideo {
		log.Printf("%+v", v)
		url := v.VideoURL
		s, err := dl.GetSeries(v.ShowURL)
		if err != nil {
			log.Fatal(err)
		}
		e, err := s.GetEpisodeBySlug(v.VideoURL)
		if err != nil {
			log.Fatal(err)
		}
		for _, lang := range e.Languages() {
			for _, quality := range e.Qualities(lang) {
				log.Printf("lang, quality: %q, %q", lang, quality)
			}
		}
		url, err = e.GuessVideoUrl(dlfunimation.Dubbed, dlfunimation.FullHighDefinition)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(url)
		resp, err := c.Client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.StatusCode)
		break
	}
}
