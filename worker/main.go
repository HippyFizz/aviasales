package worker

import (
	"aviasales/cache"
	"aviasales/config"
	"aviasales/modules/aviasales"
	"aviasales/worker/core"
	"fmt"
	"net/url"
	"time"
)

var (
	conf *config.Config
)

func StartUp(filepath string) {
	conf = config.Manager(filepath)
	client := cache.Manager(&conf.Redis)
	aviasalesService := aviasales.Manager(&conf.Avisales)
	worker := core.Manager(&conf.Worker, client, aviasalesService)

	for {
		worker.LogMessage("Start updating cache")
		data, err := worker.LoadCache()
		if err == nil {
			jobs := make(chan *core.UpdateItem, len(data.([]interface{})))
			for _, value := range data.([]interface{}) {
				requestUrl, err := url.ParseRequestURI(fmt.Sprintf("%s", value))
				if err != nil {
					worker.LogMessage(err)
				} else {
					jobs <- &core.UpdateItem{Key: requestUrl}
				}
			}
			for w := 1; w <= conf.Worker.Count; w++ {
				go worker.UpdateCache(w, jobs)
			}
		}
		time.Sleep(conf.Worker.UpdateInterval)
	}
}
