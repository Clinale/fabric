/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main
/*
import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/examples/chaincode/go/example02"
)
*/
import (
	"fmt"
	"github.com/Clinale/fabric/tree/release-1.3/core/chaincode/shim"
	"github.com/Clinale/fabric/tree/release-1.3/examples/chaincode/go/example02"
)

func main() {
	err := shim.Start(new(example02.SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
