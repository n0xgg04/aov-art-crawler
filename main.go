package main

import (
	"github.com/fatih/color"
	"golang.org/x/exp/maps"
	"nox-art-crawler/constants"
	"nox-art-crawler/crawler"
	"time"
)

func main() {
	const MaxQueueSize = 99999
	const DELAY = 5

	worker := make(chan constants.Works, MaxQueueSize)
	for {
		_ = maps.Keys(constants.Heroes)
		color.Blue("Nox Art Crawler is running...")
		heroId := maps.Keys(constants.Heroes)
		ch := make(chan string, len(heroId))
		for _, h := range heroId {
			go crawler.CheckHero(h, ch, worker)
		}

		for w := range worker {
			crawler.SendTelegram("[CRAWLER] Tìm thấy art skin mới ID "+w.SkinId+" tại server "+w.Server, w.PhotoUrl, "[CRAWLER] Tìm thấy art skin mới ID "+w.SkinId+" tại server server")
			time.Sleep(time.Second * 2)
		}
	}
}
