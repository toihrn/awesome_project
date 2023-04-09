package pinecone

const (
	pineCone = "https://index_name-project_id.svc.environment.pinecone.io/query"
)

//func Query(ctx context.Context, pineConeQueryParam *model.PineConeQueryParam) ([]byte, error) {
//payLoad := strings.NewReader(xjson.ToString(pineConeQueryParam))
//req ,_ =http.
//}

//func main() {
//
//	url := "https://index_name-project_id.svc.environment.pinecone.io/query"
//
//	payload := strings.NewReader("{\"includeValues\":\"false\",\"includeMetadata\":\"false\"}")
//
//	req, _ := http.NewRequest("POST", url, payload)
//
//	req.Header.Add("accept", "application/json")
//	req.Header.Add("content-type", "application/json")
//
//	res, _ := http.DefaultClient.Do(req)
//
//	defer res.Body.Close()
//	body, _ := ioutil.ReadAll(res.Body)
//
//	fmt.Println(res)
//	fmt.Println(string(body))
//
//}
