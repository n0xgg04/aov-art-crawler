package crawler

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/imroc/req/v3"
	"io"
	"io/fs"
	"nox-art-crawler/constants"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	client = req.C()
	f, _   = os.Getwd()
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Crawl(heroId string) int {
	heroPath := filepath.Join(filepath.Dir(f), filepath.Base(f), "data", "art", heroId)
	failed := 0
	total := 0
skinScan:
	for i := 0; i <= 25; i++ {
		skinId := heroId
		if i < 10 {
			skinId = skinId + "0"
		}
		skinId += strconv.Itoa(i)
		if fileExists(path.Join(heroPath, skinId+".jpg")) {
			total++
			color.White("[Skip] Found before %s", skinId)
			continue skinScan
		}

	serverScan:
		for server, api := range constants.ArtAPI {
			resp, err := client.R().
				Get(strings.ReplaceAll(api, "##ID##", skinId))
			if err == nil {
				if resp.StatusCode == 200 {
					buf := new(bytes.Buffer)
					_, err = io.Copy(buf, resp.Body)
					if buf.Len() < 10*1024 {
						color.Yellow("[Not Found] Skin %s in server %s.", skinId, server)
						failed++
					} else {
						f, e := os.Create(path.Join(heroPath, skinId+".jpg"))
						if e != nil {
							color.Red("[Save failed] Failed to save image")
						} else {
							_, err = f.ReadFrom(buf)
							if err == nil {
								total++
								color.Green("[Found art] Found new art %s in server %s.", skinId, server)
								_ = f.Close()
								break serverScan
							}
						}
					}
				} else {
					color.Yellow("[Not Found] Skin %s in server %s.", skinId, server)
					failed++
				}
			}
		}
	}
	return total
}

func CheckHero(heroId string, ch chan<- string) {
	path2 := filepath.Join(filepath.Dir(f), filepath.Base(f), "data", "art", heroId)
	isExisted := fs.ValidPath(path2)
	if !isExisted {
		err := os.MkdirAll(path2, fs.ModePerm)
		if err != nil {
			panic("Can't create folder for " + heroId)
		}
	}
	ch <- strconv.Itoa(Crawl(heroId))
}
