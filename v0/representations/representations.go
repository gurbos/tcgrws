package representations

import (
	tcm "github.com/gurbos/tcmodels"
)

type Printer interface {
	Print() string
}

type ProductLineInfoRep struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	URLName   string `json:"urlName"`
	SetCount  uint   `json:"setCount"`
	CardCount uint   `json:"cardCount"`
}

func (pli *ProductLineInfoRep) Set(pl *tcm.ProductLine) {
	pli.ID = pl.ID
	pli.Name = pl.Name
	pli.URLName = pl.URLName
	pli.SetCount = pl.SetCount
	pli.CardCount = pl.CardCount
}

type SetInfoRep struct {
	ID            uint   `json:"setId"`
	Name          string `json:"setName"`
	URLName       string `json:"setUrlName"`
	ProductLineID uint   `json:"productLineId"`
	CardCount     uint   `json:"cardCount"`
}

func (csi *SetInfoRep) Set(setInfo *tcm.SetInfo) {
	csi.ID = setInfo.ID
	csi.Name = setInfo.Name
	csi.URLName = setInfo.URLName
	csi.ProductLineID = setInfo.ProductLineID
	csi.CardCount = setInfo.CardCount
}

type SetID struct {
	ID uint `json:"id"`
}

type CardRep struct {
	ID                 uint   `json:"cardId"`
	Name               string `json:"cardName"`
	URLName            string `json:"cardUrlName"`
	Number             string `json:"number"`
	Attribute          string `json:"attribute"`
	CardType           string `json:"cardType"`
	CardTypeB          string `json:"cardTypeB"`
	Level              string `json:"level"`
	MonsterType        string `json:"monsterType"`
	LinkRating         string `json:"linkRating"`
	LinkArrows         string `json:"linkArrows"`
	Attack             string `json:"attack"`
	Defense            string `json:"defense"`
	Description        string `json:"description"`
	Rarity             string `json:"rarity"`
	ProductLineID      uint   `json:"productLineId"`
	ProductLineName    string `json:"productLineName"`
	ProductLineURLName string `json:"productLineUrlName"`
	SetID              uint   `json:"setId"`
	SetName            string `json:"setName"`
	SetURLName         string `json:"setUrlName"`
	ImageURL           string `json:"imageURL"`
}

func (card *CardRep) Set(ci *tcm.YuGiOhCardInfo) {
	card.ID = ci.ID
	card.Name = ci.Name
	card.Number = ci.Number
	card.Attribute = ci.Attribute
	card.CardType = ci.CardType
	card.CardTypeB = ci.CardTypeB
	card.Level = ci.Level
	card.MonsterType = ci.MonsterType
	card.LinkRating = ci.LinkRating
	card.LinkArrows = ci.LinkArrows
	card.Attack = ci.Attack
	card.Defense = ci.Defense
	card.Description = ci.Description
	card.Rarity = ci.Rarity
	card.ProductLineID = ci.ProductLineID
	card.ProductLineName = ci.ProductLine.Name
	card.ProductLineURLName = ci.ProductLine.URLName
	card.SetID = ci.SetID
	card.SetName = ci.SetInfo.Name
	card.SetURLName = ci.SetInfo.URLName
	// card.ImageURL = ImagesDir + "/" + card.ProductLineURLName + "/" + strconv.Itoa(int(card.ID)) + "_200w.jpg"
}

func (card *CardRep) Print() string {
	return card.Name
}

type ResponsePayload struct {
	Errors      []string             `json:"errors"`
	ProductLine []ProductLineInfoRep `json:"productLine"`
	CardSet     []SetInfoRep         `json:"cardSet"`
	Cards       []Printer            `json:"cards"`
	From        uint                 `json:"from"`
	Size        uint                 `json:"size"`
	TotalCards  uint                 `json:"totalCards"`
}

func (rp *ResponsePayload) Set(errs []string, plInfos []ProductLineInfoRep, csInfos []SetInfoRep, cards []Printer, from uint, size uint) {
	rp.Errors = errs
	rp.ProductLine = plInfos
	rp.CardSet = csInfos
	rp.Cards = cards
	rp.From = from
	rp.Size = size
	for _, elem := range rp.CardSet {
		rp.TotalCards += elem.CardCount
	}
}
