/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	api "github.com/celo-org/rosetta/api"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	log.Printf("Server started")

	rpcClient, err := rpc.Dial("https://alfajores-forno.celo-testnet.org/")
	if err != nil {
		log.Fatalf("Can't connect to node, %s", err)
	}

	AccountApiService := api.NewAccountApiService(rpcClient)
	AccountApiController := api.NewAccountApiController(AccountApiService)

	BlockApiService := api.NewBlockApiService()
	BlockApiController := api.NewBlockApiController(BlockApiService)

	ConstructionApiService := api.NewConstructionApiService()
	ConstructionApiController := api.NewConstructionApiController(ConstructionApiService)

	MempoolApiService := api.NewMempoolApiService()
	MempoolApiController := api.NewMempoolApiController(MempoolApiService)

	NetworkApiService := api.NewNetworkApiService(rpcClient)
	NetworkApiController := api.NewNetworkApiController(NetworkApiService)

	router := api.NewRouter(AccountApiController, BlockApiController, ConstructionApiController, MempoolApiController, NetworkApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
