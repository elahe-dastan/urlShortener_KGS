package store

import (
	"time"

	"github.com/elahe-dastan/urlShortener/generator"
	"github.com/elahe-dastan/urlShortener/metric"
	"github.com/elahe-dastan/urlShortener/model"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus"
)

type TinyURL interface {
}

type ShortURL struct {
	DB        *gorm.DB
	Length    int
	Histogram prometheus.Histogram
}

func NewShortURL(d *gorm.DB) ShortURL {
	return ShortURL{DB: d,
		Histogram: metric.NewHistogram("choosing_short_url_histogram")}
}

// Connects to the database and saves all the random short urls generated by the key generator service in it
func (url ShortURL) Save() {
	if url.DB.HasTable(&model.ShortURL{}) {
		return
	}

	url.DB.Debug().AutoMigrate(&model.ShortURL{})

	generator.Generate(url.DB, url.Length)
}

func (url ShortURL) Choose() string {
	var selectedURL model.ShortURL

	start := time.Now()

	defer func() {
		duration := time.Since(start)
		url.Histogram.Observe(duration.Seconds())
	}()

	url.DB.Raw("UPDATE short_urls SET is_used = ? WHERE url = "+
		"(SELECT url FROM short_urls WHERE is_used = ? LIMIT 1 FOR UPDATE) "+
		"RETURNING *;", true, false).Scan(&selectedURL) //O(lgn)

	return selectedURL.URL
}
