package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"github.com/algorand/go-algorand-sdk/types"
)

// indexer host
const indexerAddress = "http://localhost:8980"
const indexerToken = ""

// query parameters
var txID = "QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA"

func main() {
	address, _ := types.DecodeAddress("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4")

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}

	// Query
	result, err := indexerClient.SearchForTransactions().Address(address).TXID(txID).Do(context.Background())

	// Print the results
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")
}
