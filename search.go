package tm

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/txn2/ack"
	"github.com/txn2/es"
	"go.uber.org/zap"
)

// ModelSearchResults
type ModelSearchResults struct {
	es.SearchResults
	Hits struct {
		Total    int           `json:"total"`
		MaxScore float64       `json:"max_score"`
		Hits     []ModelResult `json:"hits"`
	} `json:"hits"`
}

// AccountSearchResultsAck
type ModelSearchResultsAck struct {
	ack.Ack
	Payload ModelSearchResults `json:"payload"`
}

// SearchModels
func (a *Api) SearchModels(account string, searchObj *es.Obj) (int, ModelSearchResults, error) {
	modelResults := &ModelSearchResults{}

	code, err := a.Elastic.PostObjUnmarshal(fmt.Sprintf("%s-%s/_search", account, IdxModel), searchObj, modelResults)
	if err != nil {
		a.Logger.Error("EsError", zap.Error(err))
		return code, *modelResults, err
	}

	return code, *modelResults, nil
}

// SearchAccountsHandler
func (a *Api) SearchModelsHandler(c *gin.Context) {
	ak := ack.Gin(c)

	obj := &es.Obj{}
	err := ak.UnmarshalPostAbort(obj)
	if err != nil {
		a.Logger.Error("Search failure.", zap.Error(err))
		return
	}

	// SearchModelsHandler must be security screened in
	// upstream middleware to protect account access.
	account := c.Param("account")

	code, esResult, err := a.SearchModels(account, obj)
	if err != nil {
		a.Logger.Error("EsError", zap.Error(err))
		ak.SetPayloadType("EsError")
		ak.SetPayload("Error communicating with database.")
		ak.GinErrorAbort(500, "EsError", err.Error())
		return
	}

	if code >= 400 && code < 500 {
		ak.SetPayload(esResult)
		ak.GinErrorAbort(500, "SearchError", "There was a problem searching")
		return
	}

	ak.SetPayloadType("ModelSearchResults")
	ak.GinSend(esResult)
}
