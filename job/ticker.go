package job

import (
	"match_statistics_scrapper/notifier"
	"time"
)

func StartTicker() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			notifier.TgBot()
		}
	}
}
