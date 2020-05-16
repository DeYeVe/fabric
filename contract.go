package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type chaincode struct {
}

//key-value 입력 메소드
func putArgs(key, val, keyArg, valArg String){
    key = keyArg
    val = valArg
}

//Init
func (t *chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var Country, Number, Inventor, Title, StartDate, EndDate, Duration string    //계약서 정보 key. 국가, 특허번호, 출원인, 특허명칭, 등록일, 만료일, 계약기간
	var CountryVal, NumberVal, InventorVal, TitleVal, StartDateVal, EndDateVal, DurationVal string   // value
	var err error

	if len(args) != 14 {
		return nil, errors.New("error. number of arguments is incorrect")
	} // arg counts*2

	// Initialize the chaincode
	putArgs(Country, CountryVal, args[0], args[1])
	putArgs(Number, NumberVal, args[2], args[3])
	putArgs(Inventor, InventorVal, args[4], args[5])
	putArgs(Title, TitleVal, args[6], args[7])
	putArgs(StartDate, StartDateVal, args[8], args[9])
	putArgs(EndDate, EndDateVal, args[10], args[11])
	putArgs(Duration, DurationVal, args[12], args[13])

	// Write the state to the ledger
	stub.PutState(Country, []byte(CountryVal))
	stub.PutState(Number, []byte(NumberVal))
	stub.PutState(Inventor, []byte(InventorVal))
	stub.PutState(Title, []byte(TitleVal))
	stub.PutState(StartDate, []byte(StartDateVal))
	stub.PutState(EndDate, []byte(EndDateVal))
	stub.PutState(Duration, []byte(DurationVal))

	return nil, nil
}

//Invoke
func (t *chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}

	var Country, Number, Inventor, Title, StartDate, EndDate, Duration string    //계약서 정보 key. 국가, 특허번호, 출원인, 특허명칭, 등록일, 만료일, 계약기간
	var CountryVal, NumberVal, InventorVal, TitleVal, StartDateVal, EndDateVal, DurationVal string   // value
	var X int          // Transaction value
	var err error

	if len(args) != 7 {
		return nil, errors.New("error. number of arguments is incorrect")
	}// args counts

	Country = args[0]
	Number = args[1]
	Inventor = args[2]
	Title = args[3]
	StartDate = args[4]
	EndDate = args[5]
	Duration = args[6]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	CountryValbytes := stub.GetState(Country)
	CountryVal, _ = CountryValbytes

	NumberValbytes := stub.GetState(Number)
	NumberVal, _ = NumberValbytes

	InventorValbytes := stub.GetState(Inventor)
	InventorVal, _ = InventorValbytes

	TitleValbytes := stub.GetState(Title)
	TitleVal, _ = TitleValbytes

	StartDateValbytes := stub.GetState(StartDate)
	StartDateVal, _ = StartDateValbytes

	EndDateValbytes := stub.GetState(EndDate)
	EndDateVal, _ = EndDateValbytes

	DurationValbytes := stub.GetState(Duration)
	DurationVal, _ = DurationValbytes

	// Write the state back to the ledger
	stub.PutState(Country, []byte(CountryVal))
	stub.PutState(Number, []byte(NumberVal))
	stub.PutState(Inventor, []byte(InventorVal))
	stub.PutState(Title, []byte(TitleVal))
	stub.PutState(StartDate, []byte(StartDateVal))
	stub.PutState(EndDate, []byte(EndDateVal))
	stub.PutState(Duration, []byte(DurationVal))

	return nil, nil
}
func main() {
	err := shim.Start(new(chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
