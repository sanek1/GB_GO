package pacage

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type MyUrl url.URL

func (m *MyUrl) UnmarshalJSON(data []byte) error {

	//return []byte(fmt.Sprintf("\"%s\"", strings.Join(m, ","))), nil
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	res, err := url.Parse(v)
	if err != nil {
		return err
	} else {
		fmt.Println("URL is valid")
	}
	*m = MyUrl(*res)
	return nil
}

type SpecificationJSON struct {
	Port         int    `json:"port"`
	DB_url       MyUrl  `json:"dB_url"`
	Jaeger_url   MyUrl  `json:"jaeger_url"`
	Sentry_url   MyUrl  `json:"sentry_url"`
	Some_app_id  string `json:"some_app_id"`
	Some_app_key int    `json:"some_app_key"`
	Kafka_broker string `json:"kafka_broker"`
}

func GetConfigFromJSON() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var specification SpecificationJSON
	if err = json.Unmarshal(data, &specification); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", specification.Some_app_key)
}
