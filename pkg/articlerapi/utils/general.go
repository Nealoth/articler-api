package utils

import (
	"github.com/Nealoth/articler-api/pkg/articlerapi/logger"
	"strconv"
	"time"
)

const (
	UintStringifyBase = 36
)

func StringifyUint64(n uint64) string {
	return strconv.FormatUint(n, UintStringifyBase)
}

func ParseUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, UintStringifyBase, 64)
}

func SafeParseDuration(duration string) time.Duration {
	parsedDuration, err := time.ParseDuration(duration)

	if err != nil {
		logger.Panic(err)
	}

	return parsedDuration
}