package logs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

type LogModel struct {
	Ip            string `json:"ip"`
	Authorization string `json:"authorization"`
	Method        string `json:"method"`
	Path          string `json:"path"`
	Url           string `json:"url"`
}

func CustomLogger() logger.Config {
	file, err := os.OpenFile("./appLog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return logger.Config{
		Output:     file,
		Format:     getCustomLogModel(),
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Istanbul",
	}

}

func getCustomLogModel() string {
	var model = LogModel{
		Ip:            TagBuilder(logger.TagIP),
		Authorization: TagBuilder(logger.TagHeader + "Authorization"),
		Method:        TagBuilder(logger.TagMethod),
		Path:          TagBuilder(logger.TagPath),
		Url:           TagBuilder(logger.TagURL),
	}

	stringLog, err := json.Marshal(model)
	if err != nil {
		return fmt.Sprintf("error serialize log model: %v", err)
	}

	return string(stringLog) + "\n"
}

func TagBuilder(param string) string {
	return fmt.Sprintf("${%s}", param)
}
