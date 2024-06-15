package utils

import (
	"encoding/json"
	"io"
)

func UnmarshalBody(body io.Reader, obj interface{}) {
	bodyStr, err := io.ReadAll(body)
	LogError(err)

	err = json.Unmarshal(bodyStr, obj)
	LogError(err)
}
