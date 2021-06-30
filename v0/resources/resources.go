package resources

func Config(imgHost string, imgDir string) {
	ImageLocator = new(imageLocator)
	ImageLocator.init(imgHost, imgDir)
}

type imageLocator struct {
	imgHost string
	imgDir  string
}

func (i *imageLocator) init(host string, dir string) {
	i.imgHost = host
	i.imgDir = dir
}

var ImageLocator *imageLocator

func (i *imageLocator) ImgHost() string {
	return i.imgHost
}

func (i *imageLocator) ImgDir() string {
	return i.imgDir
}

type ProductLineRep struct {
	ID        uint   `json:"Id"`
	Name      string `json:"name"`
	URLName   string `json:"urlName"`
	SetCount  uint   `json:"setCount"`
	CardCount uint   `json:"cardCount"`
}

type SetRep struct {
	ID            uint   `json:"setId"`
	Name          string `json:"setName"`
	URLName       string `json:"setUrlName"`
	ProductLineID uint   `json:"productLineId"`
	CardCount     uint   `json:"cardCount"`
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

type ResponsePayload struct {
	Errors      []string         `json:"errors"`
	ProductLine []ProductLineRep `json:"productLine"`
	CardSet     []SetRep         `json:"cardSet"`
	Cards       []CardRep        `json:"cards"`
	Offset      int64            `json:"offset"`
	Length      int64            `json:"length"`
	TotalCards  int64            `json:"totalCards"`
}

func (rp *ResponsePayload) Set(errs []string, plInfos []ProductLineRep, csInfos []SetRep, cards []CardRep, offset int64, length int64) {
	rp.Errors = errs
	rp.ProductLine = plInfos
	rp.CardSet = csInfos
	rp.Cards = cards
	rp.Offset = offset
	rp.Length = length
	if len(rp.CardSet) != 0 {
		for _, elem := range rp.CardSet {
			rp.TotalCards += int64(elem.CardCount)
		}
	} else {
		rp.TotalCards = int64(rp.ProductLine[0].CardCount)
	}
}
