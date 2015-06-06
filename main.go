package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"poc/cli"
	"poc/server/internal"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	{ // Test a silo generated on server side and viewed in Client
		fmt.Println("========= Silo Generated on server side ========")
		siloInServer := serverTypes.NewSilo()
		bserver, _ := json.Marshal(siloInServer)
		fmt.Printf("Server JSON Representation: %s\n", string(bserver))
		fmt.Printf("Server OBJ Representation: %v\n", siloInServer)

		var siloInCli cliTypes.Silo
		json.Unmarshal(bserver, &siloInCli)
		fmt.Printf("CLI Obj Representation: %v\n", siloInCli)
		bcli, _ := json.Marshal(siloInCli)
		fmt.Printf("CLI JSON Representation: %s\n", string(bcli))
	}

	{ // Test a silo that was defined from blue print
		fmt.Println("========= Silo Generated from blueprint ========")
		bluePrint := "{\"SiloName\":\"nico\",\"SiloVersion\":4,\"Spec\":{\"Fieldstr\":\"shoppingPod\",\"Fieldint\":7,\"Fieldruntime\":{\"R1\":\"orange\",\"R2\":67}}}"
		rawSiloInCli, siloInCli, _ := cliTypes.NewSiloFromBluePrint(bluePrint)
		fmt.Printf("CLI Obj Representation: %v\n", *siloInCli)
		bcli, _ := json.Marshal(siloInCli)
		fmt.Printf("CLI JSON Representation: %s\n", string(bcli))
		brawcli, _ := json.Marshal(rawSiloInCli)
		fmt.Printf("CLI RawJSON Representation: %s\n", string(brawcli))

		// Read the raw json output in the server types
		var siloInServer serverTypes.Silo
		json.Unmarshal(brawcli, &siloInServer)
		bserver, _ := json.Marshal(siloInServer)
		fmt.Printf("Server JSON Representation: %s\n", string(bserver))
		fmt.Printf("Server OBJ Representation: %v\n", siloInServer)
	}

}
