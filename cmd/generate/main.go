package main

import (
	"embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ChimeraCoder/gojson"
	"github.com/cborum/go-saxo"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

/*
{{- $DocUrl := printf "%s%s" $DocsBaseUrl .Refdocsurl }}
{{- getResponseBodyStruct $DocUrl .Key }}
*/

//go:embed saxo-api-service-groups-modified.json
var f embed.FS

const (
	forceStructUpdate = false
)

func main() {
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	dir, err := filepath.Abs(filepath.Dir(args[0]))
	panicIfErr(err)
	fmt.Println(dir)

	saxoOpenAPI := parseSaxoJson()

	err = os.Chdir(dir)
	panicIfErr(err)

	// Template
	funcs := template.FuncMap{
		"title":                 strings.Title,
		"lower":                 strings.ToLower,
		"getFuncParameters":     getFuncParameters,
		"hasNonRouteParameters": hasNonRouteParameters,
		"sanitize":              sanitize,
		"getParamType":          getParamType,
		"getResponseType":       getResponseType,
		"hasResponseType":       hasResponseType,
		"hasQueryParams":        hasQueryParams,
		"hasRouteParams":        hasRouteParams,
		"getNonRouteParams":     getNonRouteParams,
		"getRouteParams":        getRouteParams,
		"hasPostParams":         hasPostParams,
	}
	templateFile, err := os.ReadFile("cmd/generate/templates/service.tmpl")
	panicIfErr(err)
	t, err := template.New("service").Funcs(funcs).Parse(string(templateFile))
	// t, err := template.ParseFiles("templates/service.tmpl")
	panicIfErr(err)

	for _, serviceGroup := range saxoOpenAPI.Servicegroups {
		fmt.Println(serviceGroup.Key, serviceGroup.Title)

		err := os.Mkdir(strings.ToLower(serviceGroup.Key), 0777)
		if !os.IsExist(err) {
			panicIfErr(err)
		}

		for _, service := range serviceGroup.Services {
			createService(service, strings.ToLower(serviceGroup.Key), t)
		}
	}
}

func createService(service SaxoService, serviceGroupKey string, t *template.Template) {
	fmt.Println("-", service.Key, service.Title, len(service.Endpoints))
	for _, endpoint := range service.Endpoints {
		path := strings.ToLower(fmt.Sprintf("%s/%s_response.go", serviceGroupKey, endpoint.Key))
		if _, err := os.Stat(path); os.IsNotExist(err) || forceStructUpdate {
			saveResponseBodyStruct(saxo.DocBaseURL+endpoint.Refdocsurl, endpoint.Key, serviceGroupKey, path)
		} else {
			fmt.Println(" ", path, "exist")
		}
	}

	strippedServiceKey := strings.SplitAfterN(service.Key, "-", 2)[1]
	f, err := os.Create(strings.ToLower(fmt.Sprintf("%s/%s.go", serviceGroupKey, strippedServiceKey)))
	panicIfErr(err)
	defer f.Close()
	// TODO some post/patch may have query params and post body? -> two diff structs
	err = t.Execute(f, &TemplateData{
		service,
		serviceGroupKey,
		saxo.DocBaseURL,
	})
	panicIfErr(err)
}

func parseSaxoJson() *SaxoOpenAPI {
	jsonFile, err := f.Open("saxo-api-service-groups-modified.json")
	panicIfErr(err)

	saxoOpenAPI := &SaxoOpenAPI{}
	err = json.NewDecoder(jsonFile).Decode(saxoOpenAPI)
	panicIfErr(err)

	return saxoOpenAPI
}

func getFuncParameters(parameters []Parameter, key string, includeApiVersion bool, apiVersion string) string {
	pp := []string{}
	hasOptional := false

	// Route parameters
	for _, v := range parameters {
		if v.Origin == "Route" {
			pp = append(pp, strings.ToLower(fmt.Sprintf("%s string", v.Name)))
		} else {
			hasOptional = true
		}
	}

	// Query parameters
	if hasOptional {
		if includeApiVersion {
			key += apiVersion
		}
		pp = append(pp, fmt.Sprintf("params *%sParams", key))
	}

	return strings.Join(pp, ", ")
}

func hasNonRouteParameters(parameters []Parameter) bool {
	for _, v := range parameters {
		if v.Origin != "Route" {
			return true
		}
	}
	return false
}

func sanitize(str string) string {
	return strings.Trim(strings.Trim(str, "$"), "*")
}

func saveResponseBodyStruct(docUrl string, EndpointKey string, serviceGroupKey string, filePath string) bool {
	structBytes, err := getResponseBodyStruct(docUrl, EndpointKey, serviceGroupKey)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = os.WriteFile(filePath, structBytes, 0777)
	panicIfErr(err)

	return true
}

func getResponseBodyStruct(docUrl string, EndpointKey string, serviceGroupKey string) ([]byte, error) {
	resp, err := http.Get(docUrl)
	panicIfErr(err)

	root, err := html.Parse(resp.Body)
	panicIfErr(err)

	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.H5 {
			return scrape.Text(n) == "Response body"
		}
		return false
	}

	node, ok := scrape.Find(root, matcher)
	if !ok {
		return nil, errors.New("response body element not found")
	}

	node, ok = scrape.Find(node.Parent, scrape.ByTag(atom.Pre))
	if !ok {
		return nil, errors.New("response body json not found")
	}

	structBytes := jsonToStruct(scrape.Text(node), EndpointKey+"Response", serviceGroupKey)
	return structBytes, nil
}

func jsonToStruct(str string, structname string, serviceGroupKey string) []byte {
	res, err := gojson.Generate(
		strings.NewReader(str),
		gojson.ParseJson,
		structname,
		serviceGroupKey,
		[]string{"json"},
		false,
		false,
	)
	panicIfErr(err)
	return res
	// return strings.SplitAfterN(string(res), "\n", 3)[2]
}

func getParamType(str string) string {
	switch str {
	case "Bool":
		return "bool"
	case "Decimal":
		return "float64"
	case "Int":
		return "int64"
	default:
		return "string"
	}
}

func getResponseType(endpointKey string, serviceGroupKey string) string {
	path := strings.ToLower(fmt.Sprintf("%s/%s_response.go", serviceGroupKey, endpointKey))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "[]byte"
	} else {
		return fmt.Sprintf("*%sResponse", endpointKey)
	}
}

func hasResponseType(endpointKey string, serviceGroupKey string) bool {
	path := strings.ToLower(fmt.Sprintf("%s/%s_response.go", serviceGroupKey, endpointKey))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func hasQueryParams(params []Parameter) bool {
	for _, v := range params {
		if v.Origin == "Query-String" {
			return true
		}
	}
	return false
}

func hasRouteParams(params []Parameter) bool {
	for _, v := range params {
		if v.Origin == "Route" {
			return true
		}
	}
	return false
}

func getNonRouteParams(params []Parameter) []Parameter {
	pp := make([]Parameter, 0)
	for _, v := range params {
		if v.Origin != "Route" {
			pp = append(pp, v)
		}
	}
	return pp
}

func getRouteParams(params []Parameter) string {
	ss := []string{}
	for _, v := range params {
		if v.Origin == "Route" {
			ss = append(ss, fmt.Sprintf("saxo.RP(\"{%s}\", %s)", v.Name, strings.ToLower(v.Name)))
		}
	}
	return strings.Join(ss, ", ")
}

func hasPostParams(params []Parameter) bool {
	for _, v := range params {
		if v.Origin == "Body" {
			return true
		}
	}
	return false
}

type TemplateData struct {
	SaxoData         SaxoService
	ServiceGroupName string
	DocsBaseUrl      string
}

type SaxoOpenAPI struct {
	Servicegroups []struct {
		Title    string        `json:"Title"`
		Key      string        `json:"Key"`
		Services []SaxoService `json:"Services"`
	} `json:"serviceGroups"`
}

type SaxoService struct {
	Title     string `json:"Title"`
	Key       string `json:"Key"`
	Endpoints []struct {
		Httpmethod string      `json:"HttpMethod"`
		URL        string      `json:"Url"`
		Parameters []Parameter `json:"Parameters"`
		Key        string      `json:"Key"`
		Refdocsurl string      `json:"RefDocsUrl"`
	} `json:"Endpoints"`
	Includeapiversion bool   `json:"IncludeApiVersion"`
	Apiversion        string `json:"ApiVersion"`
}

type Parameter struct {
	Hasrequiredproperties bool        `json:"HasRequiredProperties"`
	Requiredproperties    interface{} `json:"RequiredProperties"`
	Name                  string      `json:"Name"`
	Isrequired            bool        `json:"IsRequired"`
	Origin                string      `json:"Origin"`
	Datatype              string      `json:"DataType"`
	Iscomplex             bool        `json:"IsComplex"`
	Isenum                bool        `json:"IsEnum"`
}

type RequiredProperty struct {
	Name      string `json:"Name"`
	DataType  string `json:"DataType"`
	IsComplex bool   `json:"IsComplex"`
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
