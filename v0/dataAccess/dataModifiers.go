package dataAccess

import (
	res "github.com/gurbos/tcgrws/v0/resources"
	tcm "github.com/gurbos/tcmodels"
)

func toProductLineRepList(productLines []tcm.ProductLine) []res.ProductLineRep {
	plReps := make([]res.ProductLineRep, len(productLines))
	for i, elem := range productLines {
		setProductLine(&plReps[i], &elem)
	}
	return plReps
}

func toSetRepList(setInfos []tcm.SetInfo) []res.SetRep {
	cardSets := make([]res.SetRep, len(setInfos))
	for i, elem := range setInfos {
		setSet(&cardSets[i], &elem)
	}
	return cardSets
}

func toCardRepList(cards []tcm.YuGiOhCardInfo) []res.CardRep {
	cardReps := make([]res.CardRep, len(cards))
	for i, elem := range cards {
		setCard(&cardReps[i], &elem)
	}
	return cardReps
}

func setProductLine(plr *res.ProductLineRep, plm *tcm.ProductLine) {
	plr.ID = plm.ID
	plr.Name = plm.Name
	plr.URLName = plm.URLName
	plr.SetCount = plm.SetCount
	plr.CardCount = plm.CardCount
}

func setSet(sr *res.SetRep, sm *tcm.SetInfo) {
	sr.ID = sm.ID
	sr.Name = sm.Name
	sr.URLName = sm.URLName
	sr.ProductLineID = sm.ProductLineID
	sr.CardCount = sm.CardCount
}

func setCard(cr *res.CardRep, cm *tcm.YuGiOhCardInfo) {
	cr.ID = cm.ID
	cr.Name = cm.Name
	cr.Number = cm.Number
	cr.Attribute = cm.Attribute
	cr.CardType = cm.CardType
	cr.CardTypeB = cm.CardTypeB
	cr.Level = cm.Level
	cr.MonsterType = cm.MonsterType
	cr.LinkRating = cm.LinkRating
	cr.LinkArrows = cm.LinkArrows
	cr.Attack = cm.Attack
	cr.Defense = cm.Defense
	cr.Description = cm.Description
	cr.Rarity = cm.Rarity
	cr.ProductLineID = cm.ProductLineID
	cr.ProductLineName = cm.ProductLine.Name
	cr.ProductLineURLName = cm.ProductLine.URLName
	cr.SetID = cm.SetID
	cr.SetName = cm.SetInfo.Name
	cr.SetURLName = cm.SetInfo.URLName
	// cr.ImageURL = res.ImageLocator.ImgHost() + "/" + cr.ProductLineURLName + "/" + strconv.Itoa(int(cr.ID))
}
