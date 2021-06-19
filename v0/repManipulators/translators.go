package repManipulators

import (
	rep "github.com/gurbos/tcgrws/v0/representations"
	tcm "github.com/gurbos/tcmodels"
)

// func ExtractSetIDList(setIDList []SetID) ([]uint, error) {
// 	var setIDs []uint
// 	if setIDList != nil {
// 		idList := make([]uint, len(setIDList))
// 		for i, elem := range setIDList {
// 			idList[i] = elem.ID
// 		}
// 	}
// 	return setIDs, nil
// }

func GetProductLineRepList(productLines []tcm.ProductLine) []rep.ProductLineInfoRep {
	plReps := make([]rep.ProductLineInfoRep, len(productLines))
	for i, elem := range productLines {
		plReps[i].Set(&elem)
	}
	return plReps
}

func GetSetRepList(setInfos []tcm.SetInfo) []rep.SetInfoRep {
	cardSets := make([]rep.SetInfoRep, len(setInfos))
	for i, elem := range setInfos {
		cardSets[i].Set(&elem)
	}
	return cardSets
}

func GetCardRepList(cards []tcm.YuGiOhCardInfo) []rep.Printer {
	cardReps := make([]rep.CardRep, len(cards))
	printers := make([]rep.Printer, len(cardReps))
	for i, elem := range cards {
		cardReps[i].Set(&elem)
		printers[i] = rep.Printer(&cardReps[i])
	}
	return printers
}

// // Return a list of cards based on the argument type.
// func makeCardList(data interface{}) interface{} {
// 	var cards interface{}
// 	dataKind := reflect.ValueOf(data).Kind()
// 	if dataKind == reflect.Slice {
// 		dataVal := reflect.ValueOf(data)
// 		switch data.(type) {
// 		case []tcm.YuGiOhCardInfo:
// 			list := make([]tcm.YuGiOhCardInfo, dataVal.Len())
// 			tmp := dataVal.Interface().([]tcm.YuGiOhCardInfo)
// 			for i, elem := range tmp {
// 				list[i].Set(&elem)
// 			}
// 			cards = list
// 		}
// 	}
// 	return cards
// }

// getProductLineIDs creates a list of product line IDs
func GetProductLineIDs(productLine []tcm.ProductLine) []int64 {
	idList := make([]int64, len(productLine))
	for i, elem := range productLine {
		idList[i] = int64(elem.ID)
	}
	return idList
}

// getSetIDs creates a list of set info IDs
func GetSetIDs(sets []tcm.SetInfo) []int64 {
	setIDs := make([]int64, len(sets))
	for i, elem := range sets {
		setIDs[i] = int64(elem.ID)
	}
	return setIDs
}
