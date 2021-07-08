package endpoints

var endPoints = map[string]string{
	"productLines": "/productLines",
	"metaData":     "/metaData?{productLine}",
	"cards":        "/cards?{productLineName,setName,from,size}",
}

type ResourceURLs struct {
	host     string
	endpoint map[string]string
}

func (ru *ResourceURLs) URL(name string) string {
	return ru.host + ru.endpoint[name]
}

func (ru *ResourceURLs) ListEndpoints() []string {
	list := make([]string, 0, len(ru.endpoint))
	for key, _ := range ru.endpoint {
		list = append(list, ru.endpoint[key])
	}
	return list
}

var Urls *ResourceURLs

func Configure(publicHostName string) {
	Urls = new(ResourceURLs)
	Urls.host = publicHostName
	Urls.endpoint = endPoints
}
