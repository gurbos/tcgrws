package endpoints

var endPoints = map[string]string{
	"productLines": "/productLines",
	"metaData":     "/metaData",
	"cards":        "/cards",
	"images":       "/images",
}

var apiRef = []string{
	"/productLines",
	"/metaData?{productLineId}",
	"/cards?{productLineName,setName,from,size}",
	"/images/{imageName}",
}

type ResourceURLs struct {
	host     string
	endpoint map[string]string
	apiRef   []string
}

func (ru *ResourceURLs) URL(name string) string {
	return ru.host + ru.endpoint[name]
}

func (ru *ResourceURLs) ListApiReference() []string {
	list := make([]string, 0, len(ru.apiRef))
	for _, val := range ru.apiRef {
		list = append(list, ru.host+val)
	}
	return list
}

var Urls *ResourceURLs

func Configure(publicHostName string) {
	Urls = new(ResourceURLs)
	Urls.host = publicHostName
	Urls.endpoint = endPoints
	Urls.apiRef = apiRef
}
