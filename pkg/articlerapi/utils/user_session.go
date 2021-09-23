package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	UserIDSessionAttribute string = "userID"
)

func GetUserIDFromContext(c *gin.Context) (uint64, error) {
	rawUserID, exist := c.Get(UserIDSessionAttribute)

	if !exist {
		return 0, fmt.Errorf("field %s is not exist in context", UserIDSessionAttribute)
	}

	stringifiedUserID, succeedConvertation := rawUserID.(string)

	if !succeedConvertation {
		return 0, fmt.Errorf("cannot convert session attribute %s to string type", UserIDSessionAttribute)
	}

	actualUserID, err := ParseUint64(stringifiedUserID)

	if err != nil {
		return 0, err
	}

	return actualUserID, nil
}
