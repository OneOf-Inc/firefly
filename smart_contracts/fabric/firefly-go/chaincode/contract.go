package chaincode

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/firefly/chaincode-go/batchpin"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) PinBatch(ctx contractapi.TransactionContextInterface, uuids, batchHash, payloadRef string, contexts []string) error {
	log.Println("Chaincode contract PinBatch: started")
	event, err := batchpin.BuildEvent(ctx, &batchpin.Args{
		UUIDs:      uuids,
		BatchHash:  batchHash,
		PayloadRef: payloadRef,
		Contexts:   contexts,
	})
	if err != nil {
		log.Println("failed to build event: %s", err)
		return err
	}
	log.Println("Chaincode contract PinBatch: BuildEvent done")
	bytes, err := json.Marshal(event)
	log.Println("Chaincode contract PinBatch: event marshalling done")
	if err != nil {
		log.Println("failed to marshal event: %s", err)
		return fmt.Errorf("failed to marshal event: %s", err)
	}
	batchPinEvent := ctx.GetStub().SetEvent("BatchPin", bytes)
	log.Println("Chaincode contract PinBatch: event is set")
	return batchPinEvent
}

func (s *SmartContract) NetworkAction(ctx contractapi.TransactionContextInterface, action, payload string) error {
	log.Println("Chaincode contract NetworkAction: started")
	event, err := batchpin.BuildEvent(ctx, &batchpin.Args{})
	log.Println("Chaincode contract NetworkAction: BuildEvent done")
	if err != nil {
		log.Println("failed to build event: %s", err)
		return err
	}
	event.Action = action
	event.PayloadRef = payload
	bytes, err := json.Marshal(event)
	log.Println("Chaincode contract NetworkAction: event marshalling done")
	if err != nil {
		log.Println("failed to marshal event: %s", err)
		return fmt.Errorf("failed to marshal event: %s", err)
	}
	batchPinEvent := ctx.GetStub().SetEvent("BatchPin", bytes)
	log.Println("Chaincode contract NetworkAction: event is set")
	return batchPinEvent
}

func (s *SmartContract) NetworkVersion() int {
	return 2
}
