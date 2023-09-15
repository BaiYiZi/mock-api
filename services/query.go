package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/BaiYiZi/mock-api/util"
	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) (map[string]any, int, error) {
	p, _ := os.Getwd()
	filename := c.GetString("req_path") + ".json"
	filePath := path.Join(p, "mock-data", filename)

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("method not allowed")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := gin.H{}
	json.Unmarshal(byteValue, &result)

	canMethod := result["info"].(map[string]any)["method"].(string)
	if canMethod != c.GetString("req_method") {
		return nil, http.StatusNotFound, fmt.Errorf("method not allowed")
	}

	apiParams := result["info"].(map[string]any)["params"].(map[string]any)
	if len(apiParams) > 0 {
		reqParams := map[string]any{}
		err := c.ShouldBind(&reqParams)

		if err != nil || !util.CompareTwoMapInterface(apiParams, reqParams) {
			return nil, http.StatusBadRequest, fmt.Errorf("parameter validation error")
		}

	}

	delete(result, "info")
	return result["data"].(map[string]any), http.StatusOK, nil
}
