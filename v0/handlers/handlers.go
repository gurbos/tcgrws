package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	data "github.com/gurbos/tcgrws/v0/dataAccess"
	ep "github.com/gurbos/tcgrws/v0/endpoints"
	res "github.com/gurbos/tcgrws/v0/resources"
)

func Configure(staticContDir string) {
	staticContentDir = staticContDir
}

var (
	staticContentDir    string
	defCardResultOffset int64 = 0
	defCardResultLength int64 = 10
)

type CardQueryParameters struct {
	ProductLineName []string
	SetName         []string
	Offset          int64
	Length          int64
}

func (cqp *CardQueryParameters) Set(query *url.Values) error {
	var cqpErr error = nil
	cqp.ProductLineName = (*query)["productLineName"]
	cqp.SetName = (*query)["setName"]

	ostr := query.Get("offset")
	offset, err := strconv.Atoi(ostr)
	if err == nil {
		cqp.Offset = int64(offset)
	} else {
		cqp.Offset = defCardResultOffset
	}

	lstr := query.Get("length")
	length, err := strconv.Atoi(lstr)
	if err == nil {
		cqp.Length = int64(length)
	} else {
		cqp.Length = defCardResultLength
	}

	return cqpErr
}

func EndPointsHandler(w http.ResponseWriter, r *http.Request) {
	endPoints := ep.Urls.ListEndpoints()
	jbuff, err := json.Marshal(&endPoints)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jbuff)
}

// ProductLineHandler tranfers product line representations
func ProductLineHandler(w http.ResponseWriter, r *http.Request) {
	respErrs := make([]string, 0, 2) // Error list for the response payload

	productLineReps, dbErr := data.GetAllProductLines() // Get a list of all product lines from database
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	// Create and initialize response payload
	respPayload := new(res.ResponsePayload)
	respPayload.Set(respErrs, productLineReps, []res.SetRep{}, nil, 0, 0)

	jbuff, _ := json.Marshal(respPayload) // Encode response payload to JSON

	_, err := w.Write(jbuff) // Send response to client
	if err != nil {
		log.Fatal("ProductFiltersCriteriaHandler\n http.Write error:", err)
	}
}

func MetaDataHandler(w http.ResponseWriter, r *http.Request) {
	respErrs := make([]string, 0, 2) // Error list for the response payload
	productLineReps, err := data.GetAllProductLines()
	if err != nil {
		log.Fatal(err)
	}
	setReps, err := data.GetSets([]int64{}, []string{})
	if err != nil {
		log.Fatal(err)
	}

	respPayload := new(res.ResponsePayload)
	respPayload.Set(respErrs, productLineReps, setReps, []res.CardRep{}, 0, 0)
	jbuff, _ := json.Marshal(respPayload)
	_, err = w.Write(jbuff)
	if err != nil {
		log.Fatal("ProductFiltersCriteriaHandler\n http.Write error:", err)
	}
}

func CardsHandler(w http.ResponseWriter, r *http.Request) {
	urlStr := r.URL.RawQuery
	_, err := r.URL.Parse(urlStr)
	if err != nil {
		// Handle errors
	}

	query := r.URL.Query()
	var qParams CardQueryParameters
	qParams.Set(&query)
	productLineReps, err := data.GetProductLines(qParams.ProductLineName)
	if err != nil {
		panic("CardHandler() --> queryProductLines(): " + err.Error())
	}
	plIds := GetProductLineIDs(productLineReps)

	setReps, err := data.GetSets(plIds, qParams.SetName)
	if err != nil {
		panic("CardHandler() --> querySets(): " + err.Error())
	}
	setIds := GetSetIDs(setReps)

	cardReps, err := data.GetCards(plIds, setIds, qParams.Offset, qParams.Length)
	if err != nil {
		panic("CardHandler() --> queryCards(): " + err.Error())
	}

	var respPayload res.ResponsePayload
	respPayload.Set(
		[]string{}, productLineReps, setReps, cardReps,
		qParams.Offset, qParams.Length)
	jbuff, _ := json.Marshal(respPayload)
	w.Write(jbuff)
}

// getProductLineIDs creates a list of product line IDs
func GetProductLineIDs(productLine []res.ProductLineRep) []int64 {
	idList := make([]int64, len(productLine))
	for i, elem := range productLine {
		idList[i] = int64(elem.ID)
	}
	return idList
}

// getSetIDs creates a list of set info IDs
func GetSetIDs(sets []res.SetRep) []int64 {
	setIDs := make([]int64, len(sets))
	for i, elem := range sets {
		setIDs[i] = int64(elem.ID)
	}
	return setIDs
}
