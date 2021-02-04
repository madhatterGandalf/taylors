package crawler

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"taylors/global"
	"taylors/logger"
	"taylors/model"
	"time"
)

type dongFangCrawler struct {
	dongFangModel *dongFang
	url           string
	topCodes      []string
	sleepTime     int
	env           string // dev stage prod
	lock          sync.RWMutex
}

var (
	once               sync.Once
	dongFangCrawlerObj *dongFangCrawler
)

func NewDongFangCrawler() *dongFangCrawler {
	once.Do(func() {
		if !strings.Contains(global.GVA_CONFIG.Crawler.Url, "5000") {
			panic("Crawler.Url 配置错误")
		}
		dongFangCrawlerObj = &dongFangCrawler{
			dongFangModel: &dongFang{},
			url:           global.GVA_CONFIG.Crawler.Url,
			sleepTime:     global.GVA_CONFIG.Crawler.Sleep,
			env:           global.GVA_CONFIG.Crawler.Env,
		}
	})
	return dongFangCrawlerObj
}

func (crawler *dongFangCrawler) getBody(url string) (body []byte, err error) {
	rsp, err := http.Get(url)
	if err != nil {
		return
	}
	rspBody := rsp.Body
	defer rspBody.Close()

	bys, err := ioutil.ReadAll(rspBody)
	if err != nil {
		return
	}

	body = bys
	return
}

func (crawler *dongFangCrawler) checkTime() (flag bool) {
	if crawler.env == "dev" {
		return true
	}

	now := time.Now()
	weekday := now.Weekday()
	if weekday == time.Sunday || weekday == time.Saturday {
		return false
	}

	forenoonStart := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location())
	forenoonEnd := time.Date(now.Year(), now.Month(), now.Day(), 11, 40, 0, 0, now.Location())
	if now.After(forenoonStart) && now.Before(forenoonEnd) {
		return true
	}

	afternoonStart := time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, now.Location())
	afternoonEnd := time.Date(now.Year(), now.Month(), now.Day(), 15, 10, 0, 0, now.Location())
	if now.After(afternoonStart) && now.Before(afternoonEnd) {
		return true
	}

	return false
}

func (crawler *dongFangCrawler) Loop() {
	instence().store(crawler.process())
	for {
		time.Sleep(time.Second * time.Duration(crawler.sleepTime))
		if crawler.checkTime() {
			instence().store(crawler.process())
		}
	}
}

func (crawler *dongFangCrawler) process() (dongFang *dongFang) {
	defer func() {
		recover()
	}()

	body, err := crawler.getBody(crawler.url)
	if err != nil {
		return
	}
	dongFang, err = crawler.dongFangModel.JsonToModel(body)
	if err != nil {
		logger.Error("获取Json err:", err)
	}

	if dongFang.Data.Diff == nil || len(dongFang.Data.Diff) == 0 {
		return
	}

	codes := []string{}
	for i, diff := range dongFang.Data.Diff {
		if i <= 100 {
			codes = append(codes, diff.Code)
		}
	}

	crawler.lock.Lock()
	defer crawler.lock.Unlock()
	crawler.topCodes = codes

	return
}

func (crawler *dongFangCrawler) Top() (marketList []model.Stock) {
	marketList = instence().obtain(crawler.topCodes)
	return
}

func (crawler *dongFangCrawler) Monitor(codes []string) (marketList []model.Stock) {
	marketList = instence().obtain(codes)
	return
}

func (crawler *dongFangCrawler) All() (marketList []model.Stock) {
	marketList = instence().obtainAll()
	return
}
