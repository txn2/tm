/*
   Copyright 2019 txn2
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at
       http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"flag"
	"os"

	"github.com/txn2/micro"
	"github.com/txn2/provision"
	"github.com/txn2/tm"
)

var (
	elasticServerEnv = getEnv("ELASTIC_SERVER", "http://elasticsearch:9200")
)

func main() {
	esServer := flag.String("esServer", elasticServerEnv, "Elasticsearch Server")

	serverCfg, _ := micro.NewServerCfg("Type Model (tm)")
	server := micro.NewServer(serverCfg)

	tmApi, err := tm.NewApi(&tm.Config{
		Logger:        server.Logger,
		HttpClient:    server.Client,
		ElasticServer: *esServer,
	})
	if err != nil {
		server.Logger.Fatal("failure to instantiate the model API: " + err.Error())
		os.Exit(1)
	}

	// User token middleware
	server.Router.Use(provision.UserTokenHandler())

	// Get a model
	server.Router.GET("model/:account/:id",
		provision.AccountAccessCheckHandler(false),
		tmApi.GetModelHandler,
	)

	// Upsert a model
	server.Router.POST("model/:account",
		provision.AccountAccessCheckHandler(true),
		tmApi.UpsertModelHandler,
	)

	// run provisioning server
	server.Run()
}

// getEnv gets an environment variable or sets a default if
// one does not exist.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
