package middleware

import (
	_ "context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	_ "math/big"
	"os"
	"strings"
)

func GetFileContent(address, dirname string) (string, error) {
	data, err := ioutil.ReadDir(dirname)//得到路径中所有文件名
	if err != nil {
		fmt.Println("read dir err", err)
		return "", err
	}
	//fmt.Println("得到路径全部文件名",data)
	for _, v := range data {
		if strings.Index(strings.ToLower(v.Name()), strings.ToLower(address)[2:]) > 0 {
			//代表找到文件
			file, err := os.Open(dirname + "\\" + v.Name())
			if err != nil {
				fmt.Printf("Failed to open file %v, err === %v\n", v.Name(), err)
				return "", err
			}
			data := make([]byte, 1024)
			count, err := file.Read(data)
			if err != nil || count < 1 {
				fmt.Printf("Failed to read file %v, err === %v\n", v.Name(), err)
				return "", err
			}
			fmt.Println("读取到的账号为",string(data[:count]))
			return string(data[:count]), nil
		}
	}

	return "", nil
}

func FileContent(address, dirname string) ([]byte, error) {
	data, err := ioutil.ReadDir(dirname)//得到路径中所有文件名
	if err != nil {
		fmt.Println("read dir err", err)
		return []byte(""), err
	}
	//fmt.Println("得到路径全部文件名",data)
	for _, v := range data {
		if strings.Index(strings.ToLower(v.Name()), strings.ToLower(address)[2:]) > 0 {
			//代表找到文件
			file, err := os.Open(dirname + "\\" + v.Name())
			if err != nil {
				fmt.Printf("Failed to open file %v, err === %v\n", v.Name(), err)
				return []byte(""), err
			}
			data := make([]byte, 1024)
			count, err := file.Read(data)
			if err != nil || count < 1 {
				fmt.Printf("Failed to read file %v, err === %v\n", v.Name(), err)
				return []byte(""), err
			}
			s:=string(data[:count])
			fmt.Println("读取到的账号为",string(data[:count]))
			return []byte(s), nil
		}
	}

	return []byte(""), nil
}

func GetOpts(addr,dirname,passwd string)(*bind.TransactOpts,error){
	ks,err := GetFileContent(addr,dirname)
	if err != nil{
		fmt.Println("Failed to util GetFileContent",err)

	}
	opts,err := bind.NewTransactor(strings.NewReader(ks),passwd)
	//key, err := keystore.DecryptKey(addr, passwd)
	//if err != nil {
	//	fmt.Println("获取账号的错误：",err)
	//}
	//fmt.Println("私钥为",key.PrivateKey, "公钥为",key.Address)
	if err != nil{
		fmt.Println("Failed to bind.NewTransctor",err)

	}
	fmt.Println(opts);
	return opts,nil
}
func Getaccout() (*ecdsa.PrivateKey, common.Address) {
	accout,err:=FileContent("0x35b4a6302d1bfc0c4acae31e59fe3d97308dfed6","D:/icegeth/icedata/keystore/")
	//const file = "D:/icegeth/icedata/keystore/UTC--2022-04-07T14-36-28.963758200Z--35b4a6302d1bfc0c4acae31e59fe3d97308dfed6"
	//account, err := ioutil.ReadFile(file)
	//if err != nil {
	//	panic(err)
	//}
	password := "123456"
	key, err := keystore.DecryptKey(accout, password)
	if err != nil {
		fmt.Println("获取秘钥的错误：",err)
	}
	fmt.Println("私钥为",key.PrivateKey, "公钥为",key.Address)
	return key.PrivateKey, key.Address

}
