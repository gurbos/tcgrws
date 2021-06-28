package endpoints

// Constants representing resource identifiers.
const (
	// Collection of all product line info
	ProductLinesURL = "http://127.0.0.1:8000/productLines"

	// List of all product line representations, and lists of all the sets,
	// card types,  and rarities for the product line specified in the
	// 'productLine' query parameter.
	MetaDataURL = "http://127.0.0.1:8000/metaData{?productLine}"

	// Collection of card representations
	CardsURL = "http://127.0.0.1:8000/cards?{productLineName,setName,from,size}"
)

type Endpoints struct {
	ProductLinesURL string `json:"productLinesUrl"`
	MetaDataURL     string `json:"productLineSetsUrl"`
	CardsURL        string `json:"productLineCardsUrl"`
}

func (ai *Endpoints) Init() {
	ai.ProductLinesURL = ProductLinesURL
	ai.MetaDataURL = MetaDataURL
	ai.CardsURL = CardsURL
}
