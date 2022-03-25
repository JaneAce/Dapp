package eth

import (
	_ "fmt"
	_ "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	_ "log"
)
func Eth()  {
	conn,err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("连接失败",err)
	}
	//延迟关闭连接
	defer conn.Close()
	//2.生成合约实例
	//demoIns,err := NewCalldemo(common.HexToAddress("合约地址"),conn)
	if err != nil {
		log.Fatalf("连接合约失败",err)
	}
	//3.调用合约函数

	//val,err := demoIns.GetCount(nil)

}
