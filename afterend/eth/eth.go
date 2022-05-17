package eth

import (
	_ "context"
	_ "fmt"
	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "log"
)

//var(
//	//连接geth
//	conn, err1  = ethclient.Dial("HTTP://127.0.0.1:8545")
//	//实例化合约
//	demoIns,err2 = contract.NewUniversity(common.HexToAddress("0x067203d11DD99d40F33De299c3f95bbCa071DCd0"),conn)
//	FromAddress = common.HexToAddress("0xd073a33f878Ed1d6f52DEEE58183720271EA5d4c")
//)
//
//func GetSchoolTime(c *gin.Context){
//	row ,_:= demoIns.GetBlockSchooltime(nil)
//	fmt.Println("eth查询结果",row)
//}
//func GetFileContent(address, dirname string) (string, error) {
//	data, err := ioutil.ReadDir(dirname)//得到路径中所有文件名
//	if err != nil {
//		fmt.Println("read dir err", err)
//		return "", err
//	}
//	for _, v := range data {
//		if strings.Index(strings.ToLower(v.Name()), strings.ToLower(address)[2:]) > 0 {
//			//代表找到文件
//			file, err := os.Open(dirname + "\\" + v.Name())
//			if err != nil {
//				fmt.Printf("Failed to open file %v, err === %v\n", v.Name(), err)
//				return "", err
//			}
//			data := make([]byte, 1024)
//			count, err := file.Read(data)
//			if err != nil || count < 1 {
//				fmt.Printf("Failed to read file %v, err === %v\n", v.Name(), err)
//				return "", err
//			}
//			return string(data[:count]), nil
//		}
//	}
//	return "", nil
//}
//func GetOpts(addr ,dirname,passwd string)(*bind.TransactOpts,error){
//	ks,err := GetFileContent(addr,dirname)
//	if err != nil{
//		fmt.Println("Failed to util GetFileContent",err)
//		return nil,err
//	}
//
//	opts,err := bind.NewTransactor(strings.NewReader(ks),passwd)
//
//	if err != nil{
//		fmt.Println("Failed to bind.NewTransctor",err)
//		return nil,err
//	}
//	return opts,nil
//}

//func Eth(schoolName,schoolAddress string)(err error){
//	//genache
//	conn,err := ethclient.Dial("HTTP://127.0.0.1:8545")
//	if err != nil {
//		log.Fatalf("连接失败",err)
//	}
	//延迟关闭连接
	//defer conn.Close()
	//2.生成合约实例
	//demoIns,err :=NewUniversity(common.HexToAddress("0x067203d11DD99d40F33De299c3f95bbCa071DCd0"),conn)
	//if err != nil {
	//	log.Fatalf("连接合约失败",err)
	//}
	//auth,err := GetOpts(devAddr,dirName,"123456")
	//auth:=GetTxopts()
	//_,err = demoIns.CreateSchool(auth,schoolName,schoolAddress)
	//gasPrice,err := conn.SuggestGasPrice(context.Background())
	//if err != nil{
	//	fmt.Printf("gas消耗有误")
	//}
	//if err != nil{
	//	fmt.Printf("私钥有误")
	//}
	//reslut, err := demoIns.CreateSchool(&bind.TransactOpts{
	//	From:   common.HexToAddress(to),
	//	GasPrice: gasPrice,
	//}, schoolName,schoolAddress)
	//fmt.Println("创建学校返回的值：",)
	//3.调用合约函数
	//rows,err := demoIns.CreateSchool(opts,schoolName,schoolAddress)
	//if err !=nil {
	//	return nil, err
	//}
	//	var school SCHOOL
	//	err = rows.Scan(&school.SchoolName)
	//	if err != nil {
	//		return nil, err
	//	}
	//	schools  = append(schools, school)
//	return err
//}
//var client *ethclient.Client

//连接geth
//func init() {
//
//	rpcDel, err := rpc.Dial("http://localhost:8577")
//	if err != nil {
//		panic(err)
//	}
//	client = ethclient.NewClient(rpcDel)
//	//fmt.Println(client)
//}

//func Getaccout() (*ecdsa.PrivateKey, common.Address) {
//	const file = "D:/icegeth/icedata/keystore/UTC--2021-09-12T11-45-17.413734000Z--a775270c810db24029044c81a6567d55126de5ae"
//	account, err := ioutil.ReadFile(file)
//	if err != nil {
//		fmt.Println("获取文件有误",err)
//	}
//	password := "123456"
//	key, err := keystore.DecryptKey(account, password)
//	if err != nil {
//		panic(err)
//	}
//	//fmt.Println(key.PrivateKey, key.Address)
//	return key.PrivateKey, key.Address
//
//}
////获取gasprice
//func GetgasPrice() (*big.Int, error) {
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		return big.NewInt(0), err
//	} else {
//		return gasPrice, nil
//	}
//
//}
////获取nonce
//func Getnonce(address common.Address) (uint64, error) {
//	nonce, err := client.PendingNonceAt(context.Background(), address)
//	if err != nil {
//		return 0, err
//	} else {
//		return nonce, nil
//	}
//
//}
//func setopts(privateKey *ecdsa.PrivateKey, address common.Address) *bind.TransactOpts {
//	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(10001))
//	if err != nil {
//		panic(err)
//	}
//	nonce, err := Getnonce(address)
//	if err != nil {
//		panic(err)
//	}
//	gasPrice, err := GetgasPrice()
//	if err != nil {
//		panic(err)
//	}
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)
//	auth.GasLimit = uint64(300000)
//	auth.GasPrice = gasPrice
//	return auth
//}
//func GetTxopts() *bind.TransactOpts{
//	privateKey, publicKey := Getaccout()
//	opts := setopts(privateKey, publicKey)
//	return opts
//}



