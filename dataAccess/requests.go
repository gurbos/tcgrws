package dataAccess

import (
	"context"

	"github.com/gurbos/tcgrws/dbio"
	res "github.com/gurbos/tcgrws/resources"
)

func GetAllProductLines(ctx context.Context) ([]res.ProductLineRep, error) {
	var productReps []res.ProductLineRep
	dbProductLines, err := dbio.QueryProductLines([]string{})
	if err == nil {
		productReps = toProductLineRepList(dbProductLines)
	}
	return productReps, err
}

func GetProductLines(ctx context.Context, names []string) ([]res.ProductLineRep, error) {
	var productReps []res.ProductLineRep
	dbProductLines, err := dbio.QueryProductLines(names)
	if err == nil {
		productReps = toProductLineRepList(dbProductLines)
	}
	return productReps, err
}

func GetSets(ctx context.Context, productLineIds []int64, setNames []string) ([]res.SetRep, error) {
	var setReps []res.SetRep
	dbSets, err := dbio.QuerySets(productLineIds, setNames)
	if err == nil {
		setReps = toSetRepList(dbSets)
	}
	return setReps, err
}

func GetSetRepList(ctx context.Context, productLineIDs []int64, setNames []string) ([]res.SetRep, error) {
	var setReps []res.SetRep
	dbSets, err := dbio.QuerySets(productLineIDs, setNames)
	if err == nil {
		setReps = toSetRepList(dbSets)
	}
	return setReps, err
}

func GetCards(ctx context.Context, plIds []int64, setIds []int64, offset int64, length int64) ([]res.CardRep, error) {
	var cardReps []res.CardRep
	dbCards, err := dbio.QueryCards(plIds, setIds, offset, length)
	if err == nil {
		cardReps = toCardRepList(dbCards)
	}
	return cardReps, err
}
