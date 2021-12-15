package main

import (
	"fmt"
	"go-grpc-blog/cms/handler"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)




func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("cms/env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Printf("error loading configuration: %v", err)
	}

	

	var decoder = schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	store := sessions.NewCookieStore([]byte(config.GetString("session.secret")))
	r := handler.New(decoder, store)

	host,port:=config.GetString("server.host"),config.GetString("server.port")

	log.Printf("Server Starting no : http://%s:%s",host,port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s",host,port), r); err != nil {
		log.Fatal(err)
	}
}
