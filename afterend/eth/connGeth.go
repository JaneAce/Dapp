package eth

import (
	"context"
	"crypto/ecdsa"
	university "dapp/afterend/contract/university"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math"
	"strconv"
	"time"

	"dapp/afterend/middleware"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

var (
	client *ethclient.Client
)

//连接geth
func init() {
	rpcDel, err := rpc.Dial("http://localhost:8577")
	if err != nil {
		panic(err)
	}
	client = ethclient.NewClient(rpcDel)
}

//获取合约
func GetContract(contractAddr string) (contract *university.University,err error){
	ins, err := university.NewUniversity(common.HexToAddress(contractAddr), client)
	if err != nil {
		panic(err)
	}
	return ins,nil
}

//获取账号和私钥
func Getaccout(addr,dirname,password string) (*ecdsa.PrivateKey, common.Address) {
	accout,err:=middleware.FileContent(addr,dirname)

	key, err := keystore.DecryptKey(accout, password)
	if err != nil {
		fmt.Println("获取秘钥的错误：",err)
	}
	fmt.Println("私钥为",key.PrivateKey, "公钥为",key.Address)
	return key.PrivateKey, key.Address

}

//获取gasprice
func GetgasPrice() (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return big.NewInt(0), err
	} else {
		return gasPrice, nil
	}

}

//获取nonce
func Getnonce(address common.Address) (uint64, error) {
	nonce, err := client.NonceAt(context.Background(), address,nil)
	if err != nil {
		return 0, err
	} else {
		return nonce, nil
	}

}

//获取区块数
func GetBlockNumber() (*types.Header,error) {
	header,err :=client.HeaderByNumber(context.Background(),nil)
	if err!=nil {
		panic(err)
	}
	//fmt.Println(header)
	return header, err
}

//设置TransactOpts
func setopts(privateKey *ecdsa.PrivateKey, address common.Address) *bind.TransactOpts {
	//chainID
	chainld, _ := new(big.Int).SetString("12345654321", 10)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainld)
	if err != nil {
		panic(err)
	}
	nonce, err := Getnonce(address)
	if err != nil {
		panic(err)
	}
	gasPrice, err := GetgasPrice()
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	return auth
}

func GetTxopts(addr,dirname,passwprd string) *bind.TransactOpts{
	privateKey, publicKey := Getaccout(addr,dirname,passwprd)
	opts := setopts(privateKey, publicKey)
	return opts
}

//创建新账号
func NewAcc(passwd string)(account string,err error){
	cli,err := rpc.Dial("http://localhost:8577")
	if err != nil {
		fmt.Println("connect to geth error:",err)
	}
	defer cli.Close()
	err =cli.Call(&account,"personal_newAccount",passwd)
	if err != nil {
		fmt.Println("Failed to new Account personal",err)
	}
	err1:= UnlockAc(account,passwd)
	if err1 != nil{
		fmt.Println("解锁账户错误",err1)
	}
	err2,result:=BatchTranserEth(account,"10")
	if err2 != nil{
		fmt.Println("转账错误",err2)
	}
	fmt.Println("转账成功结果",result)
	fmt.Println("创建的新账号account=",account)
	return account,nil
}

//wei转化成ether
func weiToEther(wei *big.Int) string {
	fbal := new(big.Float)
	fbal.SetString(wei.String())
	ether := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(18)))
	sprintf := fmt.Sprintf("%.18f", ether )
	return sprintf
}

//给账户解锁
func UnlockAc(account,passwd string)(err error){
	cli,err := rpc.Dial("http://localhost:8577")
	if err != nil {
		fmt.Println("connect to geth error:",err)
	}
	var result bool
	defer cli.Close()
	err =cli.Call(&result,"personal_unlockAccount",account,passwd)
	if err != nil {
		fmt.Println("Failed to unlock Account personal",err)
	}
	fmt.Println("账号解锁结果",result)
	return nil
}

//创建学校
func CreateSchool(schoolName string,ops *bind.TransactOpts) (*types.Transaction,error){
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err !=nil{
		fmt.Println(err)
	}
	school,err := university.CreateSchool(ops,schoolName)
	if err!=nil {
		fmt.Println("创建学校的错误",err)
	}
	return school,err
}

//创建学院
func CreateCollege(passwd string,collegeName string,collegeAddr common.Address,ops *bind.TransactOpts)(err error,transaction *types.Transaction){
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err != nil{
		fmt.Println(err)
	}
	college,err := university.AddCollege(ops,collegeName,collegeAddr)
	if err != nil{
		fmt.Println("创建学院的错误",err)
	}
	return nil,college
}

////创建证书
//func CreateCert(studentAccount common.Address,studentName string,CET4,credit,GraduationDissertation int64,opts *bind.TransactOpts)(*types.Transaction,error){
//	university, err := GetContract("0x23dC3128b46aA220b17afc981D412D0A4FF5A68f")
//	if err != nil {
//		fmt.Println(err)
//	}
//	Bcet4 := big.NewInt(CET4)
//	Bcredit := big.NewInt(credit)
//	BGraduationDissertation := big.NewInt(GraduationDissertation)
//	student,err := university.CreateoneCerts(opts,studentAccount,studentName,Bcet4,Bcredit,BGraduationDissertation)
//	if err != nil{
//		fmt.Println("创建证书错误",err)
//	}
//	return student,err
//}

//添加学生
func CreateStudent(sutdnetschool,studentName,studentSex,studentSpecialized,studentLevel,collegeName string,studnetAc common.Address,studentId,studentSystem,studentStartTime *big.Int,ops *bind.TransactOpts)(*types.Transaction,error)  {
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err !=nil{
		fmt.Println(err)
	}
	student,err := university.CreateStudent(ops,sutdnetschool,collegeName,studnetAc,studentName,studentSex,studentId,studentSpecialized,studentLevel,studentSystem,studentStartTime)
	if err != nil {
		fmt.Println("创建学生失败", err)
	}
	return student,nil
}

//添加学生的pid
func GetStudentPid()(result *big.Int){
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err != nil {
		fmt.Println(err)
	}
	result, err = university.GetBlockStudenttime(nil)
	if err != nil {
		fmt.Println("获取创建学生区块时间失败",err)
	}
	fmt.Println("查询学生pid结果" ,result)
	return result
}

//获取学院的pid
func GetCollagePid()(collagePid *big.Int)  {
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err != nil {
		fmt.Println(err)
	}
	collagePid,err = university.GetBlockCollegetime(nil)
	if err !=nil {
		log.Println("获取创建学院的pid失败",err)
	}
	fmt.Println("查询学院pid的结果为",collagePid)
	return collagePid
}
//添加成绩
func AddScore(studnetAc common.Address,subjectName string,score *big.Int,ops *bind.TransactOpts)(*types.Transaction,error){
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err !=nil{
		fmt.Println(err)
	}
	educational,err := university.CreateoneEducational(ops,studnetAc,subjectName,score)
	if err != nil{
		fmt.Println("添加成绩失败",err)
	}
	return educational,nil
}

//添加科目
func AddSubject(subjectName string, ops *bind.TransactOpts) (*types.Transaction, error) {
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err != nil {
		fmt.Println(err)
	}
	subject,err := university.AddSticker(ops,subjectName)
	if err != nil {
		fmt.Println("创建科目",err)
	}
	return subject, nil

}
//获取科目成绩的pid
func GetSubjectPid()(subjectPid *big.Int){
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err != nil {
		fmt.Println(err)
	}
	subjectPid,err = university.GetBlockSubjectTime(nil)
	if err !=nil {
		log.Println("获取添加科目的pid失败",err)
	}
	fmt.Println("查询科目pid的结果为",subjectPid)
	return subjectPid
}

//生成证书
func CreateCerts(studentName string,studentAccount common.Address,CET4,cerdit,graduate int64,ops *bind.TransactOpts)(*types.Transaction,error){
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err != nil {
		fmt.Println(err)
	}
	CET4_ := big.NewInt(CET4)
	cerdit_ := big.NewInt(cerdit)
	graduate_ := big.NewInt(graduate)
	certs,err := university.CreateoneCerts(ops,studentAccount,studentName,CET4_,cerdit_,graduate_)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("生成证书的返回值",certs)
	fmt.Println("生成证书的哈希值",certs.Hash().Hex())
	return certs,nil
}

//获取证书的pid
func GetCertPid()(certPid *big.Int){
	university,err := GetContract("0x348Cb500b2A9849C6715e01aE7da05BB41b42f39")
	if err != nil {
		fmt.Println(err)
	}
	certPid,err = university.GetBlockCollegetime(nil)
	if err !=nil {
		log.Println("获取添加科目的pid失败",err)
	}
	fmt.Println("查询科目pid的结果为",certPid)
	return certPid
}
//给账户转账
func BatchTranserEth(to string,amount string)(err error,result common.Hash){
	//client, err := ethclient.Dial("HTTP://127.0.0.1:8577")
	//if err != nil {
	//	fmt.Println(err)
	//}
	privateKey,_:=Getaccout("0x4261a1cDBC33461C27718a54c61A5ACCab9321f1","F:\\Dappgeth\\dappdata\\keystore","123456")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// chainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	var count uint8
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			fmt.Println(err)
		}
	valuef, err := strconv.ParseFloat(amount,64)   //先转换为 float64

	if err != nil {

		log.Println("is not a number")

	}

	// 再通过sprintf格式化为*Int

	valueWei, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", valuef*1000000000000000000), 10)

	if !isOk {

		log.Println("float to bigInt failed!")

	}
		//value := big.NewInt(amount * 1000000000) // in wei
		gasLimit := uint64(21000)              // in units
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		toAddress := common.HexToAddress(to)
		var data []byte
		tx := types.NewTransaction(nonce, toAddress, valueWei, gasLimit, gasPrice, data)

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			fmt.Println(err)
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			fmt.Println(err)
		}
		count++
		fmt.Printf("id: %d , -->> tx sent: %s \n", count, signedTx.Hash().Hex())
		time.Sleep(time.Second * 1)
		return nil,signedTx.Hash()
}




