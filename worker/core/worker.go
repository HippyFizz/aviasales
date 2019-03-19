package core

import (
	"aviasales/modules/aviasales"
	"aviasales/worker/config"
	"encoding/json"
	"github.com/go-redis/redis"
	"io"
	"log"
	"net/url"
	"os"
)

var (
	worker *Worker
)

type UpdateItem struct {
	Key  *url.URL
	Data interface{}
}

type Worker struct {
	cache            *redis.Client
	Config           *config.Config
	AviasalesService *aviasales.Service
	logger           *log.Logger
}

func newWorker(config *config.Config, cache *redis.Client, aviasalesService *aviasales.Service) error {
	var workerHandle io.Writer
	logger := log.New(workerHandle, "Worker: ", log.Ldate|log.Ltime|log.Lshortfile)

	worker = &Worker{
		cache:            cache,
		Config:           config,
		AviasalesService: aviasalesService,
		logger:           logger,
	}
	worker.setUpLogger()
	return nil
}

func (w *Worker) setUpLogger() {
	if !w.Config.LogToFile {
		w.logger.SetOutput(os.Stdout)
		return
	}
	f, err := os.OpenFile(w.Config.LogFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		w.Config.LogToFile = false
		w.setUpLogger()
	}
	defer f.Close()
	w.logger.SetOutput(f)

}

func (w *Worker) UpdateCache(id int, jobs <-chan *UpdateItem) {
	for job := range jobs {
		response, err := w.AviasalesService.RequestPlaces(job.Key.Query(), nil)

		if err != nil {
			job.Data = err
			w.logger.Print(err)
			return
		}

		data, err := json.Marshal(response)

		if err != nil {
			job.Data = err
			w.logger.Print(err)
			return
		}

		_, err = w.cache.Set(job.Key.String(), data, w.Config.Expiration).Result()

		if err != nil {
			job.Data = err
			w.logger.Print(err)
			return
		}
		job.Data = response
	}
}

func (w *Worker) LogMessage(message interface{}) {
	w.logger.Print(message)
}

func (w *Worker) LoadCache() (interface{}, error) {
	data, err := w.cache.Do("KEYS", "*").Result()
	if err != nil {
		w.logger.Print(err)
		return nil, err
	}
	return data, nil
}

func Manager(config *config.Config, cache *redis.Client, aviasalesService *aviasales.Service) *Worker {
	config.ConvertNanosecondToSeconds()
	if err := newWorker(config, cache, aviasalesService); err != nil {
		panic(err)
	}
	return worker
}
