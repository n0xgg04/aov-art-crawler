package main

import (
	"github.com/fatih/color"
	"golang.org/x/exp/maps"
	"nox-art-crawler/constants"
	"nox-art-crawler/crawler"
)

func main() {
	_ = maps.Keys(constants.Heroes)
	color.Blue("Nox Art Crawler is running...")
	heroId := maps.Keys(constants.Heroes)
	ch := make(chan string, len(heroId))
	for _, h := range heroId {
		go crawler.CheckHero(h, ch)
	}

	for _, h := range heroId {
		color.Blue("Total %s skins found of hero %s", <-ch, h)
	}
}
