package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sa-webapp/common"

	"github.com/gin-gonic/gin"
)

type SentimentRequest struct {
	Sentence string `json:"sentence"`
}

type SentimentResponse struct {
	Sentence string  `json:"sentence"`
	Polarity float64 `json:"polarity"`
}

func SentimentHandler(ctx *gin.Context) {
	req := &SentimentRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.Error(err)
		return
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		ctx.Error(err)
		return
	}
	httpReq, err := http.NewRequest("POST", common.URL+"/analyse/sentiment", bytes.NewBuffer(jsonBytes))
	if err != nil {
		ctx.Error(err)
		return
	}
	c := http.DefaultClient
	httpResp, err := c.Do(httpReq)
	if err != nil {
		ctx.Error(err)
		return
	}
	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := &SentimentResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, resp)
}
