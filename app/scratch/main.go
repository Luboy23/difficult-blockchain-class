package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Tx struct {
	FromID	string	`json:"from"`
	ToID	string 	`json:"to"`
	Value	uint64	`json:"value"`
}


func main() {
	if err := run();
	err != nil {
		log.Fatalln(err)
	}
}


func run()  error {

	privateKey, err := crypto.LoadECDSA("zblock/accounts/kennedy.ecdsa")
	if err != nil {
		return fmt.Errorf("unable to load private key for node: %w", err)
		}

	tx := Tx{
		FromID: "Jack",
		ToID:"Jackie",
		Value: 1000,
	}
		
		data, err := json.Marshal(tx)
		if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
		}

		stamp := []byte(fmt.Sprintf("\x19Lu Signed Message:\n%d",len(data)))

		v := crypto.Keccak256(stamp,data)

		sig, err := crypto.Sign(v,privateKey)
		if err != nil {
		return fmt.Errorf("unable to sign: %w", err)
	}

	fmt.Println("Sig:",hexutil.Encode(sig))

	//	=======================================================
	publicKey, err := crypto.SigToPub(data, sig)
	if err != nil {
		return fmt.Errorf("unable to pub: %w", err)
	}

	fmt.Println("Pub:",crypto.PubkeyToAddress(*publicKey).String())

	//	=======================================================
	tx = Tx{
		FromID: "Jack",
		ToID:"Alice",
		Value: 1000,
	}
		
		data, err = json.Marshal(tx)
		if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
		}

		stamp = []byte(fmt.Sprintf("\x19Lu Signed Message:\n%d",len(data)))

		v2 := crypto.Keccak256(stamp,data)

		sig2, err := crypto.Sign(v2,privateKey)
		if err != nil {
		return fmt.Errorf("unable to sign: %w", err)
	}

	fmt.Println("Sig:",hexutil.Encode(sig2))

		//	=======================================================
		//	OVER THE WIRE
		tx2 := Tx{
			FromID: "Jack",
			ToID:"Alice",
			Value: 1000,
		}

		data, err = json.Marshal(tx2)
		if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
		}

		stamp = []byte(fmt.Sprintf("\x19Lu Signed Message:\n%d",len(data)))
		v2 = crypto.Keccak256(stamp,data)

		publicKey2, err := crypto.SigToPub(v2, sig2)
		if err != nil {
			return fmt.Errorf("unable to pub: %w", err)
		}
	
		fmt.Println("Pub:",crypto.PubkeyToAddress(*publicKey2).String())
	
	return nil
}

