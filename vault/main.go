package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"flag"
	"encoding/json"

	vault "github.com/hashicorp/vault/api"
)
var (
	vaultAddr   = os.Getenv("VAULT_ADDR")
	staticToken = os.Getenv("VAULT_TOKEN")
	nameSpace   = os.Getenv("VAULT_NAMESPACE")
)
// const (
// 	vaultAddr   = ""
// 	staticToken = ""
// )

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	flag.Parse()
	// first = flag.Args[0]
	client, err := vault.NewClient(&vault.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}
	client.SetToken(staticToken)
	client.SetNamespace(nameSpace)
	//client.SetOutputCurlString(true)
	// const secretPath = "efserviceaccounts/mailgun"
	var secretPath = flag.Args()
	// list, err := client.Logical().List(secretPath[0])
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("LIST: %+v\n", list)

	data, err := client.Logical().Read(secretPath[0]) ///apikey")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("AUTH: %+v\n", data.Auth)
	// fmt.Printf("DATA: %+v\n", data.Data)

	// b, _ := json.Marshal(data.Data)
	// fmt.Println(string(data.Data))
	// fmt.Printf("DATA: %+v\n", data)
	// fmt.Printf("DATA: %+v\n", data.Data.["data"][apikey])
	api := data.Data["data"].(map[string]interface{})
	// apiKey := api["apikey"].(string)
	// fmt.Printf("{ \"key\": %+v\n\"", apiKey, "\"}" )
	b, err := json.Marshal(api)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}