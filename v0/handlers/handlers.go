package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gurbos/tcgrws/v0/api"
	"github.com/gurbos/tcgrws/v0/dbio"
	trans "github.com/gurbos/tcgrws/v0/repManipulators"
	rep "github.com/gurbos/tcgrws/v0/representations"
)

var defCardResultSize int64

type CardQueryParameters struct {
	ProductLineName []string
	SetName         []string
	From            int64
	Size            int64
}

func (cqp *CardQueryParameters) Set(query *url.Values) error {
	var cqpErr error = nil
	cqp.ProductLineName = (*query)["productLineName"]
	cqp.SetName = (*query)["setName"]

	from, err := strconv.Atoi(query.Get("from"))
	if err != nil {
		cqp.From = 0
		cqpErr = err
	}
	cqp.From = int64(from)

	size, err := strconv.Atoi(query.Get("size"))
	if err != nil {
		cqp.Size = defCardResultSize
		cqpErr = err
	}
	cqp.Size = int64(size)

	return cqpErr
}

// servHandler adds functionality to the handler reference by the 'mr' field.
// It implements the http.Handler interface. The 'ServeHTTP' method sets the
// HTTP header of the 'http.ResponseWriter' then calls the 'ServeHTTP' method
// that corresponds to the 'mr' field.
type servHandler struct {
	handler http.Handler
}

func (sh *servHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	{
		dateStr := time.Now().Format(time.RFC1123)
		w.Header().Set("Date", dateStr)
	}
	sh.handler.ServeHTTP(w, r)
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	var ae api.APIEndpoints
	ae.Init()
	jbuff, err := json.Marshal(ae)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jbuff)
}

// ProductLineHandler tranfers product line representations
func ProductLineHandler(w http.ResponseWriter, r *http.Request) {
	respErrs := make([]string, 0, 2) // Error list for the response payload

	productLines, dbErr := dbio.QueryAllProductLines() // Get a list of all product lines from database
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	productLineReps := trans.GetProductLineRepList(productLines)

	// Create and initialize response payload
	respPayload := new(rep.ResponsePayload)
	respPayload.Set(respErrs, productLineReps, []rep.SetInfoRep{}, nil, 0, 0)

	jbuff, _ := json.Marshal(respPayload) // Encode response payload to JSON

	_, err := w.Write(jbuff) // Send response to client
	if err != nil {
		log.Fatal("ProductFiltersCriteriaHandler\n http.Write error:", err)
	}
}

func MetaDataHandler(w http.ResponseWriter, r *http.Request) {
	respErrs := make([]string, 0, 2) // Error list for the response payload
	productLines, err := dbio.QueryAllProductLines()
	if err != nil {
		log.Fatal(err)
	}
	sets, err := dbio.QueryAllSets()
	if err != nil {
		log.Fatal(err)
	}
	productLineReps := trans.GetProductLineRepList(productLines)
	setReps := trans.GetSetRepList(sets)

	respPayload := new(rep.ResponsePayload)
	respPayload.Set(respErrs, productLineReps, setReps, []rep.Printer{}, 0, 0)
	jbuff, _ := json.Marshal(respPayload)
	_, err = w.Write(jbuff)
	if err != nil {
		log.Fatal("ProductFiltersCriteriaHandler\n http.Write error:", err)
	}
}

func CardsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var qParams CardQueryParameters
	qParams.Set(&query)
	productLines, err := dbio.QueryProductLines(qParams.ProductLineName)
	if err != nil {
		panic("CardHandler() --> queryProductLines(): " + err.Error())
	}
	plIDList := trans.GetProductLineIDs(productLines)

	sets, err := dbio.QuerySets(plIDList, qParams.SetName)
	if err != nil {
		panic("CardHandler() --> querySets(): " + err.Error())
	}
	setIDList := trans.GetSetIDs(sets)

	cards, err := dbio.QueryCards(plIDList, setIDList, qParams.From, qParams.Size)
	if err != nil {
		panic("CardHandler() --> queryCards(): " + err.Error())
	}

	plReps := trans.GetProductLineRepList(productLines)
	setReps := trans.GetSetRepList(sets)
	cardReps := trans.GetCardRepList(cards)

	var respPayload rep.ResponsePayload
	respPayload.Set(
		[]string{}, plReps, setReps, cardReps,
		uint(qParams.From), uint(qParams.Size),
	)
	jbuff, _ := json.Marshal(respPayload)
	w.Write(jbuff)
}
