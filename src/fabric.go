package main

import{
        "fmt"
        "strconv"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        "github.com/hyperledger/fabric/protos/peer"
}

type BasicChain struct{ //블록체인 객체 생성

}
//체인코드는 다음과 같은 구성으로 이뤄져 있다.
//1. Init 
//2. Invoke
//1번 함수로 트랜잭션 발생에 대해 적절한지 판단하고
//2번 함수로 원하는 기능을 수행한다.

//나가면 투표를 기권하는 것으로 알겠다.

//DID 를 만들 때 투표 여부를 해놔야 겠다.

- 블록 몇 개 생성되는지
- 중복 투표에 대해 막을 것
//투표의 경우, Invoke 함수 내에서 필요한 기능은 무엇인가
// 집계 == total count
// 각 후보에 대한 집계 == n's count




func (t *BasicChain) Init(stub shim.CahincodeStubInterface) peer.Response{
        args := stub.GetStringArgs()

        if len(args) != 2 {
                return shim.Error("Error Incorrect arguments")
        }
        err := stub.PutState(args[0],[]byte(args[1]))
        if err != nil{
                return shim.Error(fmt.Sprintf("Failed to create asset : %s", args[0]))
        }
        return shim.Success(nil)
}
func (t *BasicCahin) Invoke(stub shim.CahincodeStubInterface) peer.Response {
        fn, args := stub.GetFunctionAndParameters()
        var result string
        var err error
        if fn == "set"{
                result, err = set(stub, args)
        } else if fn == "count"{
                result, err = count(stub, args)
        } else {
                result, err = get(stub, args)
}
        if err != nil {
                return shim.Error(err.Error())
        }

        return shim.Success([]byte(result))
}

func set(stub shim.CahincodeStubInterface, args []string) (string, error){
        if len(args) != 2 {
                return "", fmt.Error("Error Incorrect args")
        }
        err := stub.PutState(args[0], []byte(args[1]))
        if err != nil {
                return "", fmt.Error("Failed to set asset : %s", args[0])
        }
        return args[1], nil
}
func get(stub shim.ChaincodeStubInterface, args []string) (string, error){
        if len(args) != 1{
                return "", fmt.Errorf("Incorrect arguments. Expecting a key, %s", args[0])
        }
        value, err := stub.GetState(args[0])
        if err != nil {
                return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
        }
        if value == nil{
                return "", fmt.Errorf("Asset not found: %s", args[0])
        }
        return string(value), nil
}
func count(stub shim.ChaincodeStubInterface, args []string) (string, error){
        var A, B string
        var Aval, Bval, int
        var X int
        var err error
        if len(args) != 3{
                return "", fmt.Errorf("Incorrect number of args, expecting 3")
        }
        A = args[0]
        B = args[1]
        Avalbytes, err := stub.GetState(A)
        Aval, _ = strconv.Atoi(string(Avalbytes))
        Bvalbytes, err := stub.GetState(B)
        Bval, _ = strconv.Atoi(string(Bvalbytes))

        X, err = strconv.Atoi(args[2])

        Aval = Aval - X
        Bval = Bval + X
        fmt.Printf("Aval = %d, Bvale = %d\n", Aval, Bval)
        err = stub.PutState(A, []byte(strconv.Itoa(Aval))
        if err != nil{
                return "", fmt.Errorf(err.Error())
        }
		err = stub.PutState(B, []byte(strconv.Itoa(Bval))
        if err != nil{
                return "", fmt.Errorf(err.Error())
        }
        return args[2], nil
}
func main(){
        if err := shim.start(new(BasicChain)); err != nil {
                fmt.Printf("Error BasicChain : %s ", err)
        }
}
