package logs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
)

// Api request log model for fiber web
type LogApiModel struct {
	Ip            string `json:"ip"`
	Authorization string `json:"authorization"`
	Method        string `json:"method"`
	Path          string `json:"path"`
	Url           string `json:"url"`
}

// Customize api request logger for filelog
func CustomApiLogger() logger.Config {
	file, err := os.OpenFile("./apiLog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return logger.Config{
		Output:     file,
		Format:     getCustomApiLogModel(),
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Istanbul",
	}

}

// fill api log model an return using tagbuilder func.
func getCustomApiLogModel() string {
	var model = LogApiModel{
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

// Basic tag builder for api request log.
func TagBuilder(param string) string {
	return fmt.Sprintf("${%s}", param)
}

// file log with zero log
func Filelogger() *zerolog.Logger {
	file, err := os.OpenFile("App.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	logger := zerolog.New(file).With().Timestamp().Logger()

	return &logger
}
