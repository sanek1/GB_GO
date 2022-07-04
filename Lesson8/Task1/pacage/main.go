package pacage

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type MyUrl url.URL

func (valiUrl *MyUrl) Decode(Value string) error {
	res, err := url.Parse(Value)
	if err != nil {
		return err
	} else {
		fmt.Println("URL is valid")
	}
	*valiUrl = MyUrl(*res)
	return nil
}

type Specification struct {
	Port         int
	DB_url       MyUrl
	Jaeger_url   MyUrl
	Sentry_url   MyUrl
	Some_app_id  string
	Some_app_key string
	Kafka_broker map[string]int
}

func GetConfiguration() {

	var s Specification
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Port: %d\nSome_app_id: %s \nSome_app_key: %s \n"
	_, err = fmt.Printf(format, s.Port, s.Some_app_id, s.Some_app_key)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Kafka broker:")
	for k, v := range s.Kafka_broker {
		fmt.Printf("  %s: %d\n", k, v)
	}
}

func SetEnv() {
	os.Clearenv()
	os.Setenv("MYAPP_Port", "8080")
	os.Setenv("MYAPP_DB_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	os.Setenv("MYAPP_Jaeger_url", "http://jaeger:16686")
	os.Setenv("MYAPP_Sentry_url", "http://sentry:9000")
	os.Setenv("MYAPP_Kafka_broker", "kafka:9092")
	os.Setenv("MYAPP_Some_app_id", "testid")
	os.Setenv("MYAPP_Some_app_key", "testkey")
}
