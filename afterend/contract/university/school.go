// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UniversityABI is the input ABI used to generate the binding from.
const UniversityABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_collegeName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_collegeAccounts\",\"type\":\"address\"}],\"name\":\"addCollege\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"schoolName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"addschoolEvnt\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"subjectName\",\"type\":\"string\"}],\"name\":\"addSticker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"studentAccounts\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"studentName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"CET4\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Credit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"GraduationDissertation\",\"type\":\"uint256\"}],\"name\":\"addstudentcertsEvnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"Subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Score\",\"type\":\"uint256\"}],\"name\":\"addstudentsubjectEvnt\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"studentAccounts\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"studentName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CET4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Credit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"GraduationDissertation\",\"type\":\"uint256\"}],\"name\":\"createoneCerts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_studentAccounts\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"subjectName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_Score\",\"type\":\"uint256\"}],\"name\":\"createoneEducational\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_schoolName\",\"type\":\"string\"}],\"name\":\"createSchool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_studentSchool\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"collegeName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_studentAccounts\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_studentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_studentSex\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_studentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_studentSpecialized\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_studentLevel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_studentSystem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_studentstartTime\",\"type\":\"uint256\"}],\"name\":\"createStudent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"colleges\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"collegeName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collegeAccounts\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"educationalCerts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"studentAccounts\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"studentName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CET4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Credit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"GraduationDissertation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"educationalSubjects\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"studentAccounts\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"Subject\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"Score\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockCertsTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockCollegetime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockSchooltime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockStudenttime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockSubjectTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collegetime\",\"type\":\"uint256\"}],\"name\":\"getcollegeNews\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"collegeName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collegeAccounts\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"educationalCertsTime\",\"type\":\"uint256\"}],\"name\":\"geteducationalCerts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"studentAccounts\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"studentName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CET4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Credit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"GraduationDissertation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"schooltime\",\"type\":\"uint256\"}],\"name\":\"getschoolNews\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"schoolName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"schoolAccounts\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"studenttime\",\"type\":\"uint256\"}],\"name\":\"getstudentNews\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"studentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"studentSchool\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"collegeName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"studentAccounts\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"studentSex\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"studentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"studentSpecialized\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"studentLevel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"studentSystem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"studentstartTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"educationalSubjecttime\",\"type\":\"uint256\"}],\"name\":\"getsubjectNews\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"studentAccounts\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"Subject\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"Score\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_collegeName\",\"type\":\"string\"}],\"name\":\"iscollegekExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"isEqual\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_schoolName\",\"type\":\"string\"}],\"name\":\"isschoolExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_studentName\",\"type\":\"string\"}],\"name\":\"isstudentExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"subjectName\",\"type\":\"string\"}],\"name\":\"isSubjectExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"schools\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"schoolName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"schoolAccounts\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"schoolIsExist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"students\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"studentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"studentSchool\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"collegeName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"studentAccounts\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"studentSex\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"studentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"studentSpecialized\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"studentLevel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"studentSystem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"studentstartTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// University is an auto generated Go binding around an Ethereum contract.
type University struct {
	UniversityCaller     // Read-only binding to the contract
	UniversityTransactor // Write-only binding to the contract
	UniversityFilterer   // Log filterer for contract events
}

// UniversityCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversitySession struct {
	Contract     *University       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniversityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversityCallerSession struct {
	Contract *UniversityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UniversityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversityTransactorSession struct {
	Contract     *UniversityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UniversityRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversityRaw struct {
	Contract *University // Generic contract binding to access the raw methods on
}

// UniversityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversityCallerRaw struct {
	Contract *UniversityCaller // Generic read-only contract binding to access the raw methods on
}

// UniversityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversityTransactorRaw struct {
	Contract *UniversityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversity creates a new instance of University, bound to a specific deployed contract.
func NewUniversity(address common.Address, backend bind.ContractBackend) (*University, error) {
	contract, err := bindUniversity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &University{UniversityCaller: UniversityCaller{contract: contract}, UniversityTransactor: UniversityTransactor{contract: contract}, UniversityFilterer: UniversityFilterer{contract: contract}}, nil
}

// NewUniversityCaller creates a new read-only instance of University, bound to a specific deployed contract.
func NewUniversityCaller(address common.Address, caller bind.ContractCaller) (*UniversityCaller, error) {
	contract, err := bindUniversity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversityCaller{contract: contract}, nil
}

// NewUniversityTransactor creates a new write-only instance of University, bound to a specific deployed contract.
func NewUniversityTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversityTransactor, error) {
	contract, err := bindUniversity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversityTransactor{contract: contract}, nil
}

// NewUniversityFilterer creates a new log filterer instance of University, bound to a specific deployed contract.
func NewUniversityFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversityFilterer, error) {
	contract, err := bindUniversity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversityFilterer{contract: contract}, nil
}

// bindUniversity binds a generic wrapper to an already deployed contract.
func bindUniversity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniversityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_University *UniversityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _University.Contract.UniversityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_University *UniversityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _University.Contract.UniversityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_University *UniversityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _University.Contract.UniversityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_University *UniversityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _University.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_University *UniversityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _University.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_University *UniversityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _University.Contract.contract.Transact(opts, method, params...)
}

// Colleges is a free data retrieval call binding the contract method 0x91bbf4f3.
//
// Solidity: function colleges(uint256 ) view returns(string collegeName, address collegeAccounts)
func (_University *UniversityCaller) Colleges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CollegeName     string
	CollegeAccounts common.Address
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "colleges", arg0)

	outstruct := new(struct {
		CollegeName     string
		CollegeAccounts common.Address
	})

	outstruct.CollegeName = out[0].(string)
	outstruct.CollegeAccounts = out[1].(common.Address)

	return *outstruct, err

}

// Colleges is a free data retrieval call binding the contract method 0x91bbf4f3.
//
// Solidity: function colleges(uint256 ) view returns(string collegeName, address collegeAccounts)
func (_University *UniversitySession) Colleges(arg0 *big.Int) (struct {
	CollegeName     string
	CollegeAccounts common.Address
}, error) {
	return _University.Contract.Colleges(&_University.CallOpts, arg0)
}

// Colleges is a free data retrieval call binding the contract method 0x91bbf4f3.
//
// Solidity: function colleges(uint256 ) view returns(string collegeName, address collegeAccounts)
func (_University *UniversityCallerSession) Colleges(arg0 *big.Int) (struct {
	CollegeName     string
	CollegeAccounts common.Address
}, error) {
	return _University.Contract.Colleges(&_University.CallOpts, arg0)
}

// EducationalCerts is a free data retrieval call binding the contract method 0xb38b836f.
//
// Solidity: function educationalCerts(uint256 ) view returns(address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation)
func (_University *UniversityCaller) EducationalCerts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StudentAccounts        common.Address
	StudentName            string
	CET4                   *big.Int
	Credit                 *big.Int
	GraduationDissertation *big.Int
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "educationalCerts", arg0)

	outstruct := new(struct {
		StudentAccounts        common.Address
		StudentName            string
		CET4                   *big.Int
		Credit                 *big.Int
		GraduationDissertation *big.Int
	})

	outstruct.StudentAccounts = out[0].(common.Address)
	outstruct.StudentName = out[1].(string)
	outstruct.CET4 = out[2].(*big.Int)
	outstruct.Credit = out[3].(*big.Int)
	outstruct.GraduationDissertation = out[4].(*big.Int)

	return *outstruct, err

}

// EducationalCerts is a free data retrieval call binding the contract method 0xb38b836f.
//
// Solidity: function educationalCerts(uint256 ) view returns(address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation)
func (_University *UniversitySession) EducationalCerts(arg0 *big.Int) (struct {
	StudentAccounts        common.Address
	StudentName            string
	CET4                   *big.Int
	Credit                 *big.Int
	GraduationDissertation *big.Int
}, error) {
	return _University.Contract.EducationalCerts(&_University.CallOpts, arg0)
}

// EducationalCerts is a free data retrieval call binding the contract method 0xb38b836f.
//
// Solidity: function educationalCerts(uint256 ) view returns(address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation)
func (_University *UniversityCallerSession) EducationalCerts(arg0 *big.Int) (struct {
	StudentAccounts        common.Address
	StudentName            string
	CET4                   *big.Int
	Credit                 *big.Int
	GraduationDissertation *big.Int
}, error) {
	return _University.Contract.EducationalCerts(&_University.CallOpts, arg0)
}

// EducationalSubjects is a free data retrieval call binding the contract method 0x66c23dfd.
//
// Solidity: function educationalSubjects(uint256 ) view returns(address studentAccounts, string Subject, uint256 Score)
func (_University *UniversityCaller) EducationalSubjects(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StudentAccounts common.Address
	Subject         string
	Score           *big.Int
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "educationalSubjects", arg0)

	outstruct := new(struct {
		StudentAccounts common.Address
		Subject         string
		Score           *big.Int
	})

	outstruct.StudentAccounts = out[0].(common.Address)
	outstruct.Subject = out[1].(string)
	outstruct.Score = out[2].(*big.Int)

	return *outstruct, err

}

// EducationalSubjects is a free data retrieval call binding the contract method 0x66c23dfd.
//
// Solidity: function educationalSubjects(uint256 ) view returns(address studentAccounts, string Subject, uint256 Score)
func (_University *UniversitySession) EducationalSubjects(arg0 *big.Int) (struct {
	StudentAccounts common.Address
	Subject         string
	Score           *big.Int
}, error) {
	return _University.Contract.EducationalSubjects(&_University.CallOpts, arg0)
}

// EducationalSubjects is a free data retrieval call binding the contract method 0x66c23dfd.
//
// Solidity: function educationalSubjects(uint256 ) view returns(address studentAccounts, string Subject, uint256 Score)
func (_University *UniversityCallerSession) EducationalSubjects(arg0 *big.Int) (struct {
	StudentAccounts common.Address
	Subject         string
	Score           *big.Int
}, error) {
	return _University.Contract.EducationalSubjects(&_University.CallOpts, arg0)
}

// GetBlockCertsTime is a free data retrieval call binding the contract method 0xe9824993.
//
// Solidity: function getBlockCertsTime() view returns(uint256)
func (_University *UniversityCaller) GetBlockCertsTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "getBlockCertsTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockCertsTime is a free data retrieval call binding the contract method 0xe9824993.
//
// Solidity: function getBlockCertsTime() view returns(uint256)
func (_University *UniversitySession) GetBlockCertsTime() (*big.Int, error) {
	return _University.Contract.GetBlockCertsTime(&_University.CallOpts)
}

// GetBlockCertsTime is a free data retrieval call binding the contract method 0xe9824993.
//
// Solidity: function getBlockCertsTime() view returns(uint256)
func (_University *UniversityCallerSession) GetBlockCertsTime() (*big.Int, error) {
	return _University.Contract.GetBlockCertsTime(&_University.CallOpts)
}

// GetBlockCollegetime is a free data retrieval call binding the contract method 0xaeed31bb.
//
// Solidity: function getBlockCollegetime() view returns(uint256)
func (_University *UniversityCaller) GetBlockCollegetime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "getBlockCollegetime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockCollegetime is a free data retrieval call binding the contract method 0xaeed31bb.
//
// Solidity: function getBlockCollegetime() view returns(uint256)
func (_University *UniversitySession) GetBlockCollegetime() (*big.Int, error) {
	return _University.Contract.GetBlockCollegetime(&_University.CallOpts)
}

// GetBlockCollegetime is a free data retrieval call binding the contract method 0xaeed31bb.
//
// Solidity: function getBlockCollegetime() view returns(uint256)
func (_University *UniversityCallerSession) GetBlockCollegetime() (*big.Int, error) {
	return _University.Contract.GetBlockCollegetime(&_University.CallOpts)
}

// GetBlockSchooltime is a free data retrieval call binding the contract method 0x8359d8bd.
//
// Solidity: function getBlockSchooltime() view returns(uint256)
func (_University *UniversityCaller) GetBlockSchooltime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "getBlockSchooltime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockSchooltime is a free data retrieval call binding the contract method 0x8359d8bd.
//
// Solidity: function getBlockSchooltime() view returns(uint256)
func (_University *UniversitySession) GetBlockSchooltime() (*big.Int, error) {
	return _University.Contract.GetBlockSchooltime(&_University.CallOpts)
}

// GetBlockSchooltime is a free data retrieval call binding the contract method 0x8359d8bd.
//
// Solidity: function getBlockSchooltime() view returns(uint256)
func (_University *UniversityCallerSession) GetBlockSchooltime() (*big.Int, error) {
	return _University.Contract.GetBlockSchooltime(&_University.CallOpts)
}

// GetBlockStudenttime is a free data retrieval call binding the contract method 0xffe7bc7e.
//
// Solidity: function getBlockStudenttime() view returns(uint256)
func (_University *UniversityCaller) GetBlockStudenttime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "getBlockStudenttime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockStudenttime is a free data retrieval call binding the contract method 0xffe7bc7e.
//
// Solidity: function getBlockStudenttime() view returns(uint256)
func (_University *UniversitySession) GetBlockStudenttime() (*big.Int, error) {
	return _University.Contract.GetBlockStudenttime(&_University.CallOpts)
}

// GetBlockStudenttime is a free data retrieval call binding the contract method 0xffe7bc7e.
//
// Solidity: function getBlockStudenttime() view returns(uint256)
func (_University *UniversityCallerSession) GetBlockStudenttime() (*big.Int, error) {
	return _University.Contract.GetBlockStudenttime(&_University.CallOpts)
}

// GetBlockSubjectTime is a free data retrieval call binding the contract method 0xe34fb5c2.
//
// Solidity: function getBlockSubjectTime() view returns(uint256)
func (_University *UniversityCaller) GetBlockSubjectTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "getBlockSubjectTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockSubjectTime is a free data retrieval call binding the contract method 0xe34fb5c2.
//
// Solidity: function getBlockSubjectTime() view returns(uint256)
func (_University *UniversitySession) GetBlockSubjectTime() (*big.Int, error) {
	return _University.Contract.GetBlockSubjectTime(&_University.CallOpts)
}

// GetBlockSubjectTime is a free data retrieval call binding the contract method 0xe34fb5c2.
//
// Solidity: function getBlockSubjectTime() view returns(uint256)
func (_University *UniversityCallerSession) GetBlockSubjectTime() (*big.Int, error) {
	return _University.Contract.GetBlockSubjectTime(&_University.CallOpts)
}

// GetcollegeNews is a free data retrieval call binding the contract method 0xe7ae0afb.
//
// Solidity: function getcollegeNews(uint256 collegetime) view returns(string collegeName, address collegeAccounts)
func (_University *UniversityCaller) GetcollegeNews(opts *bind.CallOpts, collegetime *big.Int) (struct {
	CollegeName     string
	CollegeAccounts common.Address
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "getcollegeNews", collegetime)

	outstruct := new(struct {
		CollegeName     string
		CollegeAccounts common.Address
	})

	outstruct.CollegeName = out[0].(string)
	outstruct.CollegeAccounts = out[1].(common.Address)

	return *outstruct, err

}

// GetcollegeNews is a free data retrieval call binding the contract method 0xe7ae0afb.
//
// Solidity: function getcollegeNews(uint256 collegetime) view returns(string collegeName, address collegeAccounts)
func (_University *UniversitySession) GetcollegeNews(collegetime *big.Int) (struct {
	CollegeName     string
	CollegeAccounts common.Address
}, error) {
	return _University.Contract.GetcollegeNews(&_University.CallOpts, collegetime)
}

// GetcollegeNews is a free data retrieval call binding the contract method 0xe7ae0afb.
//
// Solidity: function getcollegeNews(uint256 collegetime) view returns(string collegeName, address collegeAccounts)
func (_University *UniversityCallerSession) GetcollegeNews(collegetime *big.Int) (struct {
	CollegeName     string
	CollegeAccounts common.Address
}, error) {
	return _University.Contract.GetcollegeNews(&_University.CallOpts, collegetime)
}

// GeteducationalCerts is a free data retrieval call binding the contract method 0xaa656578.
//
// Solidity: function geteducationalCerts(uint256 educationalCertsTime) view returns(address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation)
func (_University *UniversityCaller) GeteducationalCerts(opts *bind.CallOpts, educationalCertsTime *big.Int) (struct {
	StudentAccounts        common.Address
	StudentName            string
	CET4                   *big.Int
	Credit                 *big.Int
	GraduationDissertation *big.Int
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "geteducationalCerts", educationalCertsTime)

	outstruct := new(struct {
		StudentAccounts        common.Address
		StudentName            string
		CET4                   *big.Int
		Credit                 *big.Int
		GraduationDissertation *big.Int
	})

	outstruct.StudentAccounts = out[0].(common.Address)
	outstruct.StudentName = out[1].(string)
	outstruct.CET4 = out[2].(*big.Int)
	outstruct.Credit = out[3].(*big.Int)
	outstruct.GraduationDissertation = out[4].(*big.Int)

	return *outstruct, err

}

// GeteducationalCerts is a free data retrieval call binding the contract method 0xaa656578.
//
// Solidity: function geteducationalCerts(uint256 educationalCertsTime) view returns(address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation)
func (_University *UniversitySession) GeteducationalCerts(educationalCertsTime *big.Int) (struct {
	StudentAccounts        common.Address
	StudentName            string
	CET4                   *big.Int
	Credit                 *big.Int
	GraduationDissertation *big.Int
}, error) {
	return _University.Contract.GeteducationalCerts(&_University.CallOpts, educationalCertsTime)
}

// GeteducationalCerts is a free data retrieval call binding the contract method 0xaa656578.
//
// Solidity: function geteducationalCerts(uint256 educationalCertsTime) view returns(address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation)
func (_University *UniversityCallerSession) GeteducationalCerts(educationalCertsTime *big.Int) (struct {
	StudentAccounts        common.Address
	StudentName            string
	CET4                   *big.Int
	Credit                 *big.Int
	GraduationDissertation *big.Int
}, error) {
	return _University.Contract.GeteducationalCerts(&_University.CallOpts, educationalCertsTime)
}

// GetschoolNews is a free data retrieval call binding the contract method 0x7ae414cd.
//
// Solidity: function getschoolNews(uint256 schooltime) view returns(string schoolName, address schoolAccounts)
func (_University *UniversityCaller) GetschoolNews(opts *bind.CallOpts, schooltime *big.Int) (struct {
	SchoolName     string
	SchoolAccounts common.Address
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "getschoolNews", schooltime)

	outstruct := new(struct {
		SchoolName     string
		SchoolAccounts common.Address
	})

	outstruct.SchoolName = out[0].(string)
	outstruct.SchoolAccounts = out[1].(common.Address)

	return *outstruct, err

}

// GetschoolNews is a free data retrieval call binding the contract method 0x7ae414cd.
//
// Solidity: function getschoolNews(uint256 schooltime) view returns(string schoolName, address schoolAccounts)
func (_University *UniversitySession) GetschoolNews(schooltime *big.Int) (struct {
	SchoolName     string
	SchoolAccounts common.Address
}, error) {
	return _University.Contract.GetschoolNews(&_University.CallOpts, schooltime)
}

// GetschoolNews is a free data retrieval call binding the contract method 0x7ae414cd.
//
// Solidity: function getschoolNews(uint256 schooltime) view returns(string schoolName, address schoolAccounts)
func (_University *UniversityCallerSession) GetschoolNews(schooltime *big.Int) (struct {
	SchoolName     string
	SchoolAccounts common.Address
}, error) {
	return _University.Contract.GetschoolNews(&_University.CallOpts, schooltime)
}

// GetstudentNews is a free data retrieval call binding the contract method 0xc7b41977.
//
// Solidity: function getstudentNews(uint256 studenttime) view returns(string studentName, string studentSchool, string collegeName, address studentAccounts, string studentSex, uint256 studentId, string studentSpecialized, string studentLevel, uint256 studentSystem, uint256 studentstartTime)
func (_University *UniversityCaller) GetstudentNews(opts *bind.CallOpts, studenttime *big.Int) (struct {
	StudentName        string
	StudentSchool      string
	CollegeName        string
	StudentAccounts    common.Address
	StudentSex         string
	StudentId          *big.Int
	StudentSpecialized string
	StudentLevel       string
	StudentSystem      *big.Int
	StudentstartTime   *big.Int
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "getstudentNews", studenttime)

	outstruct := new(struct {
		StudentName        string
		StudentSchool      string
		CollegeName        string
		StudentAccounts    common.Address
		StudentSex         string
		StudentId          *big.Int
		StudentSpecialized string
		StudentLevel       string
		StudentSystem      *big.Int
		StudentstartTime   *big.Int
	})

	outstruct.StudentName = out[0].(string)
	outstruct.StudentSchool = out[1].(string)
	outstruct.CollegeName = out[2].(string)
	outstruct.StudentAccounts = out[3].(common.Address)
	outstruct.StudentSex = out[4].(string)
	outstruct.StudentId = out[5].(*big.Int)
	outstruct.StudentSpecialized = out[6].(string)
	outstruct.StudentLevel = out[7].(string)
	outstruct.StudentSystem = out[8].(*big.Int)
	outstruct.StudentstartTime = out[9].(*big.Int)

	return *outstruct, err

}

// GetstudentNews is a free data retrieval call binding the contract method 0xc7b41977.
//
// Solidity: function getstudentNews(uint256 studenttime) view returns(string studentName, string studentSchool, string collegeName, address studentAccounts, string studentSex, uint256 studentId, string studentSpecialized, string studentLevel, uint256 studentSystem, uint256 studentstartTime)
func (_University *UniversitySession) GetstudentNews(studenttime *big.Int) (struct {
	StudentName        string
	StudentSchool      string
	CollegeName        string
	StudentAccounts    common.Address
	StudentSex         string
	StudentId          *big.Int
	StudentSpecialized string
	StudentLevel       string
	StudentSystem      *big.Int
	StudentstartTime   *big.Int
}, error) {
	return _University.Contract.GetstudentNews(&_University.CallOpts, studenttime)
}

// GetstudentNews is a free data retrieval call binding the contract method 0xc7b41977.
//
// Solidity: function getstudentNews(uint256 studenttime) view returns(string studentName, string studentSchool, string collegeName, address studentAccounts, string studentSex, uint256 studentId, string studentSpecialized, string studentLevel, uint256 studentSystem, uint256 studentstartTime)
func (_University *UniversityCallerSession) GetstudentNews(studenttime *big.Int) (struct {
	StudentName        string
	StudentSchool      string
	CollegeName        string
	StudentAccounts    common.Address
	StudentSex         string
	StudentId          *big.Int
	StudentSpecialized string
	StudentLevel       string
	StudentSystem      *big.Int
	StudentstartTime   *big.Int
}, error) {
	return _University.Contract.GetstudentNews(&_University.CallOpts, studenttime)
}

// GetsubjectNews is a free data retrieval call binding the contract method 0x128f3022.
//
// Solidity: function getsubjectNews(uint256 educationalSubjecttime) view returns(address studentAccounts, string Subject, uint256 Score)
func (_University *UniversityCaller) GetsubjectNews(opts *bind.CallOpts, educationalSubjecttime *big.Int) (struct {
	StudentAccounts common.Address
	Subject         string
	Score           *big.Int
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "getsubjectNews", educationalSubjecttime)

	outstruct := new(struct {
		StudentAccounts common.Address
		Subject         string
		Score           *big.Int
	})

	outstruct.StudentAccounts = out[0].(common.Address)
	outstruct.Subject = out[1].(string)
	outstruct.Score = out[2].(*big.Int)

	return *outstruct, err

}

// GetsubjectNews is a free data retrieval call binding the contract method 0x128f3022.
//
// Solidity: function getsubjectNews(uint256 educationalSubjecttime) view returns(address studentAccounts, string Subject, uint256 Score)
func (_University *UniversitySession) GetsubjectNews(educationalSubjecttime *big.Int) (struct {
	StudentAccounts common.Address
	Subject         string
	Score           *big.Int
}, error) {
	return _University.Contract.GetsubjectNews(&_University.CallOpts, educationalSubjecttime)
}

// GetsubjectNews is a free data retrieval call binding the contract method 0x128f3022.
//
// Solidity: function getsubjectNews(uint256 educationalSubjecttime) view returns(address studentAccounts, string Subject, uint256 Score)
func (_University *UniversityCallerSession) GetsubjectNews(educationalSubjecttime *big.Int) (struct {
	StudentAccounts common.Address
	Subject         string
	Score           *big.Int
}, error) {
	return _University.Contract.GetsubjectNews(&_University.CallOpts, educationalSubjecttime)
}

// IsEqual is a free data retrieval call binding the contract method 0x465c4105.
//
// Solidity: function isEqual(string a, string b) pure returns(bool)
func (_University *UniversityCaller) IsEqual(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "isEqual", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEqual is a free data retrieval call binding the contract method 0x465c4105.
//
// Solidity: function isEqual(string a, string b) pure returns(bool)
func (_University *UniversitySession) IsEqual(a string, b string) (bool, error) {
	return _University.Contract.IsEqual(&_University.CallOpts, a, b)
}

// IsEqual is a free data retrieval call binding the contract method 0x465c4105.
//
// Solidity: function isEqual(string a, string b) pure returns(bool)
func (_University *UniversityCallerSession) IsEqual(a string, b string) (bool, error) {
	return _University.Contract.IsEqual(&_University.CallOpts, a, b)
}

// IsSubjectExist is a free data retrieval call binding the contract method 0xe68d4a80.
//
// Solidity: function isSubjectExist(string subjectName) view returns(bool)
func (_University *UniversityCaller) IsSubjectExist(opts *bind.CallOpts, subjectName string) (bool, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "isSubjectExist", subjectName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubjectExist is a free data retrieval call binding the contract method 0xe68d4a80.
//
// Solidity: function isSubjectExist(string subjectName) view returns(bool)
func (_University *UniversitySession) IsSubjectExist(subjectName string) (bool, error) {
	return _University.Contract.IsSubjectExist(&_University.CallOpts, subjectName)
}

// IsSubjectExist is a free data retrieval call binding the contract method 0xe68d4a80.
//
// Solidity: function isSubjectExist(string subjectName) view returns(bool)
func (_University *UniversityCallerSession) IsSubjectExist(subjectName string) (bool, error) {
	return _University.Contract.IsSubjectExist(&_University.CallOpts, subjectName)
}

// IscollegekExist is a free data retrieval call binding the contract method 0x8fe17b11.
//
// Solidity: function iscollegekExist(string _collegeName) view returns(bool)
func (_University *UniversityCaller) IscollegekExist(opts *bind.CallOpts, _collegeName string) (bool, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "iscollegekExist", _collegeName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IscollegekExist is a free data retrieval call binding the contract method 0x8fe17b11.
//
// Solidity: function iscollegekExist(string _collegeName) view returns(bool)
func (_University *UniversitySession) IscollegekExist(_collegeName string) (bool, error) {
	return _University.Contract.IscollegekExist(&_University.CallOpts, _collegeName)
}

// IscollegekExist is a free data retrieval call binding the contract method 0x8fe17b11.
//
// Solidity: function iscollegekExist(string _collegeName) view returns(bool)
func (_University *UniversityCallerSession) IscollegekExist(_collegeName string) (bool, error) {
	return _University.Contract.IscollegekExist(&_University.CallOpts, _collegeName)
}

// IsschoolExist is a free data retrieval call binding the contract method 0x3031471f.
//
// Solidity: function isschoolExist(string _schoolName) view returns(bool)
func (_University *UniversityCaller) IsschoolExist(opts *bind.CallOpts, _schoolName string) (bool, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "isschoolExist", _schoolName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsschoolExist is a free data retrieval call binding the contract method 0x3031471f.
//
// Solidity: function isschoolExist(string _schoolName) view returns(bool)
func (_University *UniversitySession) IsschoolExist(_schoolName string) (bool, error) {
	return _University.Contract.IsschoolExist(&_University.CallOpts, _schoolName)
}

// IsschoolExist is a free data retrieval call binding the contract method 0x3031471f.
//
// Solidity: function isschoolExist(string _schoolName) view returns(bool)
func (_University *UniversityCallerSession) IsschoolExist(_schoolName string) (bool, error) {
	return _University.Contract.IsschoolExist(&_University.CallOpts, _schoolName)
}

// IsstudentExist is a free data retrieval call binding the contract method 0x0b2882e0.
//
// Solidity: function isstudentExist(string _studentName) view returns(bool)
func (_University *UniversityCaller) IsstudentExist(opts *bind.CallOpts, _studentName string) (bool, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "isstudentExist", _studentName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsstudentExist is a free data retrieval call binding the contract method 0x0b2882e0.
//
// Solidity: function isstudentExist(string _studentName) view returns(bool)
func (_University *UniversitySession) IsstudentExist(_studentName string) (bool, error) {
	return _University.Contract.IsstudentExist(&_University.CallOpts, _studentName)
}

// IsstudentExist is a free data retrieval call binding the contract method 0x0b2882e0.
//
// Solidity: function isstudentExist(string _studentName) view returns(bool)
func (_University *UniversityCallerSession) IsstudentExist(_studentName string) (bool, error) {
	return _University.Contract.IsstudentExist(&_University.CallOpts, _studentName)
}

// Schools is a free data retrieval call binding the contract method 0x719c0d04.
//
// Solidity: function schools(uint256 ) view returns(string schoolName, address schoolAccounts, bool schoolIsExist)
func (_University *UniversityCaller) Schools(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SchoolName     string
	SchoolAccounts common.Address
	SchoolIsExist  bool
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "schools", arg0)

	outstruct := new(struct {
		SchoolName     string
		SchoolAccounts common.Address
		SchoolIsExist  bool
	})

	outstruct.SchoolName = out[0].(string)
	outstruct.SchoolAccounts = out[1].(common.Address)
	outstruct.SchoolIsExist = out[2].(bool)

	return *outstruct, err

}

// Schools is a free data retrieval call binding the contract method 0x719c0d04.
//
// Solidity: function schools(uint256 ) view returns(string schoolName, address schoolAccounts, bool schoolIsExist)
func (_University *UniversitySession) Schools(arg0 *big.Int) (struct {
	SchoolName     string
	SchoolAccounts common.Address
	SchoolIsExist  bool
}, error) {
	return _University.Contract.Schools(&_University.CallOpts, arg0)
}

// Schools is a free data retrieval call binding the contract method 0x719c0d04.
//
// Solidity: function schools(uint256 ) view returns(string schoolName, address schoolAccounts, bool schoolIsExist)
func (_University *UniversityCallerSession) Schools(arg0 *big.Int) (struct {
	SchoolName     string
	SchoolAccounts common.Address
	SchoolIsExist  bool
}, error) {
	return _University.Contract.Schools(&_University.CallOpts, arg0)
}

// Students is a free data retrieval call binding the contract method 0x06ead22e.
//
// Solidity: function students(uint256 ) view returns(string studentName, string studentSchool, string collegeName, address studentAccounts, string studentSex, uint256 studentId, string studentSpecialized, string studentLevel, uint256 studentSystem, uint256 studentstartTime)
func (_University *UniversityCaller) Students(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StudentName        string
	StudentSchool      string
	CollegeName        string
	StudentAccounts    common.Address
	StudentSex         string
	StudentId          *big.Int
	StudentSpecialized string
	StudentLevel       string
	StudentSystem      *big.Int
	StudentstartTime   *big.Int
}, error) {
	var out []interface{}
	err := _University.contract.Call(opts, &out, "students", arg0)

	outstruct := new(struct {
		StudentName        string
		StudentSchool      string
		CollegeName        string
		StudentAccounts    common.Address
		StudentSex         string
		StudentId          *big.Int
		StudentSpecialized string
		StudentLevel       string
		StudentSystem      *big.Int
		StudentstartTime   *big.Int
	})

	outstruct.StudentName = out[0].(string)
	outstruct.StudentSchool = out[1].(string)
	outstruct.CollegeName = out[2].(string)
	outstruct.StudentAccounts = out[3].(common.Address)
	outstruct.StudentSex = out[4].(string)
	outstruct.StudentId = out[5].(*big.Int)
	outstruct.StudentSpecialized = out[6].(string)
	outstruct.StudentLevel = out[7].(string)
	outstruct.StudentSystem = out[8].(*big.Int)
	outstruct.StudentstartTime = out[9].(*big.Int)

	return *outstruct, err

}

// Students is a free data retrieval call binding the contract method 0x06ead22e.
//
// Solidity: function students(uint256 ) view returns(string studentName, string studentSchool, string collegeName, address studentAccounts, string studentSex, uint256 studentId, string studentSpecialized, string studentLevel, uint256 studentSystem, uint256 studentstartTime)
func (_University *UniversitySession) Students(arg0 *big.Int) (struct {
	StudentName        string
	StudentSchool      string
	CollegeName        string
	StudentAccounts    common.Address
	StudentSex         string
	StudentId          *big.Int
	StudentSpecialized string
	StudentLevel       string
	StudentSystem      *big.Int
	StudentstartTime   *big.Int
}, error) {
	return _University.Contract.Students(&_University.CallOpts, arg0)
}

// Students is a free data retrieval call binding the contract method 0x06ead22e.
//
// Solidity: function students(uint256 ) view returns(string studentName, string studentSchool, string collegeName, address studentAccounts, string studentSex, uint256 studentId, string studentSpecialized, string studentLevel, uint256 studentSystem, uint256 studentstartTime)
func (_University *UniversityCallerSession) Students(arg0 *big.Int) (struct {
	StudentName        string
	StudentSchool      string
	CollegeName        string
	StudentAccounts    common.Address
	StudentSex         string
	StudentId          *big.Int
	StudentSpecialized string
	StudentLevel       string
	StudentSystem      *big.Int
	StudentstartTime   *big.Int
}, error) {
	return _University.Contract.Students(&_University.CallOpts, arg0)
}

// AddCollege is a paid mutator transaction binding the contract method 0xbf0dc139.
//
// Solidity: function addCollege(string _collegeName, address _collegeAccounts) returns(uint256)
func (_University *UniversityTransactor) AddCollege(opts *bind.TransactOpts, _collegeName string, _collegeAccounts common.Address) (*types.Transaction, error) {
	return _University.contract.Transact(opts, "addCollege", _collegeName, _collegeAccounts)
}

// AddCollege is a paid mutator transaction binding the contract method 0xbf0dc139.
//
// Solidity: function addCollege(string _collegeName, address _collegeAccounts) returns(uint256)
func (_University *UniversitySession) AddCollege(_collegeName string, _collegeAccounts common.Address) (*types.Transaction, error) {
	return _University.Contract.AddCollege(&_University.TransactOpts, _collegeName, _collegeAccounts)
}

// AddCollege is a paid mutator transaction binding the contract method 0xbf0dc139.
//
// Solidity: function addCollege(string _collegeName, address _collegeAccounts) returns(uint256)
func (_University *UniversityTransactorSession) AddCollege(_collegeName string, _collegeAccounts common.Address) (*types.Transaction, error) {
	return _University.Contract.AddCollege(&_University.TransactOpts, _collegeName, _collegeAccounts)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string subjectName) returns()
func (_University *UniversityTransactor) AddSticker(opts *bind.TransactOpts, subjectName string) (*types.Transaction, error) {
	return _University.contract.Transact(opts, "addSticker", subjectName)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string subjectName) returns()
func (_University *UniversitySession) AddSticker(subjectName string) (*types.Transaction, error) {
	return _University.Contract.AddSticker(&_University.TransactOpts, subjectName)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string subjectName) returns()
func (_University *UniversityTransactorSession) AddSticker(subjectName string) (*types.Transaction, error) {
	return _University.Contract.AddSticker(&_University.TransactOpts, subjectName)
}

// CreateSchool is a paid mutator transaction binding the contract method 0xf0dcb15c.
//
// Solidity: function createSchool(string _schoolName) returns(uint256)
func (_University *UniversityTransactor) CreateSchool(opts *bind.TransactOpts, _schoolName string) (*types.Transaction, error) {
	return _University.contract.Transact(opts, "createSchool", _schoolName)
}

// CreateSchool is a paid mutator transaction binding the contract method 0xf0dcb15c.
//
// Solidity: function createSchool(string _schoolName) returns(uint256)
func (_University *UniversitySession) CreateSchool(_schoolName string) (*types.Transaction, error) {
	return _University.Contract.CreateSchool(&_University.TransactOpts, _schoolName)
}

// CreateSchool is a paid mutator transaction binding the contract method 0xf0dcb15c.
//
// Solidity: function createSchool(string _schoolName) returns(uint256)
func (_University *UniversityTransactorSession) CreateSchool(_schoolName string) (*types.Transaction, error) {
	return _University.Contract.CreateSchool(&_University.TransactOpts, _schoolName)
}

// CreateStudent is a paid mutator transaction binding the contract method 0x3c36a1db.
//
// Solidity: function createStudent(string _studentSchool, string collegeName, address _studentAccounts, string _studentName, string _studentSex, uint256 _studentId, string _studentSpecialized, string _studentLevel, uint256 _studentSystem, uint256 _studentstartTime) returns(uint256)
func (_University *UniversityTransactor) CreateStudent(opts *bind.TransactOpts, _studentSchool string, collegeName string, _studentAccounts common.Address, _studentName string, _studentSex string, _studentId *big.Int, _studentSpecialized string, _studentLevel string, _studentSystem *big.Int, _studentstartTime *big.Int) (*types.Transaction, error) {
	return _University.contract.Transact(opts, "createStudent", _studentSchool, collegeName, _studentAccounts, _studentName, _studentSex, _studentId, _studentSpecialized, _studentLevel, _studentSystem, _studentstartTime)
}

// CreateStudent is a paid mutator transaction binding the contract method 0x3c36a1db.
//
// Solidity: function createStudent(string _studentSchool, string collegeName, address _studentAccounts, string _studentName, string _studentSex, uint256 _studentId, string _studentSpecialized, string _studentLevel, uint256 _studentSystem, uint256 _studentstartTime) returns(uint256)
func (_University *UniversitySession) CreateStudent(_studentSchool string, collegeName string, _studentAccounts common.Address, _studentName string, _studentSex string, _studentId *big.Int, _studentSpecialized string, _studentLevel string, _studentSystem *big.Int, _studentstartTime *big.Int) (*types.Transaction, error) {
	return _University.Contract.CreateStudent(&_University.TransactOpts, _studentSchool, collegeName, _studentAccounts, _studentName, _studentSex, _studentId, _studentSpecialized, _studentLevel, _studentSystem, _studentstartTime)
}

// CreateStudent is a paid mutator transaction binding the contract method 0x3c36a1db.
//
// Solidity: function createStudent(string _studentSchool, string collegeName, address _studentAccounts, string _studentName, string _studentSex, uint256 _studentId, string _studentSpecialized, string _studentLevel, uint256 _studentSystem, uint256 _studentstartTime) returns(uint256)
func (_University *UniversityTransactorSession) CreateStudent(_studentSchool string, collegeName string, _studentAccounts common.Address, _studentName string, _studentSex string, _studentId *big.Int, _studentSpecialized string, _studentLevel string, _studentSystem *big.Int, _studentstartTime *big.Int) (*types.Transaction, error) {
	return _University.Contract.CreateStudent(&_University.TransactOpts, _studentSchool, collegeName, _studentAccounts, _studentName, _studentSex, _studentId, _studentSpecialized, _studentLevel, _studentSystem, _studentstartTime)
}

// CreateoneCerts is a paid mutator transaction binding the contract method 0x4c39dfbe.
//
// Solidity: function createoneCerts(address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation) returns(uint256)
func (_University *UniversityTransactor) CreateoneCerts(opts *bind.TransactOpts, studentAccounts common.Address, studentName string, CET4 *big.Int, Credit *big.Int, GraduationDissertation *big.Int) (*types.Transaction, error) {
	return _University.contract.Transact(opts, "createoneCerts", studentAccounts, studentName, CET4, Credit, GraduationDissertation)
}

// CreateoneCerts is a paid mutator transaction binding the contract method 0x4c39dfbe.
//
// Solidity: function createoneCerts(address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation) returns(uint256)
func (_University *UniversitySession) CreateoneCerts(studentAccounts common.Address, studentName string, CET4 *big.Int, Credit *big.Int, GraduationDissertation *big.Int) (*types.Transaction, error) {
	return _University.Contract.CreateoneCerts(&_University.TransactOpts, studentAccounts, studentName, CET4, Credit, GraduationDissertation)
}

// CreateoneCerts is a paid mutator transaction binding the contract method 0x4c39dfbe.
//
// Solidity: function createoneCerts(address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation) returns(uint256)
func (_University *UniversityTransactorSession) CreateoneCerts(studentAccounts common.Address, studentName string, CET4 *big.Int, Credit *big.Int, GraduationDissertation *big.Int) (*types.Transaction, error) {
	return _University.Contract.CreateoneCerts(&_University.TransactOpts, studentAccounts, studentName, CET4, Credit, GraduationDissertation)
}

// CreateoneEducational is a paid mutator transaction binding the contract method 0x4134d0e4.
//
// Solidity: function createoneEducational(address _studentAccounts, string subjectName, uint256 _Score) returns(uint256)
func (_University *UniversityTransactor) CreateoneEducational(opts *bind.TransactOpts, _studentAccounts common.Address, subjectName string, _Score *big.Int) (*types.Transaction, error) {
	return _University.contract.Transact(opts, "createoneEducational", _studentAccounts, subjectName, _Score)
}

// CreateoneEducational is a paid mutator transaction binding the contract method 0x4134d0e4.
//
// Solidity: function createoneEducational(address _studentAccounts, string subjectName, uint256 _Score) returns(uint256)
func (_University *UniversitySession) CreateoneEducational(_studentAccounts common.Address, subjectName string, _Score *big.Int) (*types.Transaction, error) {
	return _University.Contract.CreateoneEducational(&_University.TransactOpts, _studentAccounts, subjectName, _Score)
}

// CreateoneEducational is a paid mutator transaction binding the contract method 0x4134d0e4.
//
// Solidity: function createoneEducational(address _studentAccounts, string subjectName, uint256 _Score) returns(uint256)
func (_University *UniversityTransactorSession) CreateoneEducational(_studentAccounts common.Address, subjectName string, _Score *big.Int) (*types.Transaction, error) {
	return _University.Contract.CreateoneEducational(&_University.TransactOpts, _studentAccounts, subjectName, _Score)
}

// UniversityAddschoolEvntIterator is returned from FilterAddschoolEvnt and is used to iterate over the raw logs and unpacked data for AddschoolEvnt events raised by the University contract.
type UniversityAddschoolEvntIterator struct {
	Event *UniversityAddschoolEvnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniversityAddschoolEvntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversityAddschoolEvnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniversityAddschoolEvnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniversityAddschoolEvntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversityAddschoolEvntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversityAddschoolEvnt represents a AddschoolEvnt event raised by the University contract.
type UniversityAddschoolEvnt struct {
	EventType  common.Hash
	SchoolName string
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddschoolEvnt is a free log retrieval operation binding the contract event 0x1c97dfb8018aa9e77de1468ed826db3490275e7a59497fb793f8ccff1e97e28a.
//
// Solidity: event addschoolEvnt(string indexed eventType, string schoolName, uint256 timestamp)
func (_University *UniversityFilterer) FilterAddschoolEvnt(opts *bind.FilterOpts, eventType []string) (*UniversityAddschoolEvntIterator, error) {

	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _University.contract.FilterLogs(opts, "addschoolEvnt", eventTypeRule)
	if err != nil {
		return nil, err
	}
	return &UniversityAddschoolEvntIterator{contract: _University.contract, event: "addschoolEvnt", logs: logs, sub: sub}, nil
}

// WatchAddschoolEvnt is a free log subscription operation binding the contract event 0x1c97dfb8018aa9e77de1468ed826db3490275e7a59497fb793f8ccff1e97e28a.
//
// Solidity: event addschoolEvnt(string indexed eventType, string schoolName, uint256 timestamp)
func (_University *UniversityFilterer) WatchAddschoolEvnt(opts *bind.WatchOpts, sink chan<- *UniversityAddschoolEvnt, eventType []string) (event.Subscription, error) {

	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _University.contract.WatchLogs(opts, "addschoolEvnt", eventTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversityAddschoolEvnt)
				if err := _University.contract.UnpackLog(event, "addschoolEvnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddschoolEvnt is a log parse operation binding the contract event 0x1c97dfb8018aa9e77de1468ed826db3490275e7a59497fb793f8ccff1e97e28a.
//
// Solidity: event addschoolEvnt(string indexed eventType, string schoolName, uint256 timestamp)
func (_University *UniversityFilterer) ParseAddschoolEvnt(log types.Log) (*UniversityAddschoolEvnt, error) {
	event := new(UniversityAddschoolEvnt)
	if err := _University.contract.UnpackLog(event, "addschoolEvnt", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniversityAddstudentcertsEvntIterator is returned from FilterAddstudentcertsEvnt and is used to iterate over the raw logs and unpacked data for AddstudentcertsEvnt events raised by the University contract.
type UniversityAddstudentcertsEvntIterator struct {
	Event *UniversityAddstudentcertsEvnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniversityAddstudentcertsEvntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversityAddstudentcertsEvnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniversityAddstudentcertsEvnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniversityAddstudentcertsEvntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversityAddstudentcertsEvntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversityAddstudentcertsEvnt represents a AddstudentcertsEvnt event raised by the University contract.
type UniversityAddstudentcertsEvnt struct {
	EventType              common.Hash
	Timestamp              *big.Int
	StudentAccounts        common.Address
	StudentName            string
	CET4                   *big.Int
	Credit                 *big.Int
	GraduationDissertation *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAddstudentcertsEvnt is a free log retrieval operation binding the contract event 0xeed2d09852776448538f60d68a9c324628faf5fcf8f8e851971543a1a09c21d7.
//
// Solidity: event addstudentcertsEvnt(string indexed eventType, uint256 timestamp, address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation)
func (_University *UniversityFilterer) FilterAddstudentcertsEvnt(opts *bind.FilterOpts, eventType []string) (*UniversityAddstudentcertsEvntIterator, error) {

	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _University.contract.FilterLogs(opts, "addstudentcertsEvnt", eventTypeRule)
	if err != nil {
		return nil, err
	}
	return &UniversityAddstudentcertsEvntIterator{contract: _University.contract, event: "addstudentcertsEvnt", logs: logs, sub: sub}, nil
}

// WatchAddstudentcertsEvnt is a free log subscription operation binding the contract event 0xeed2d09852776448538f60d68a9c324628faf5fcf8f8e851971543a1a09c21d7.
//
// Solidity: event addstudentcertsEvnt(string indexed eventType, uint256 timestamp, address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation)
func (_University *UniversityFilterer) WatchAddstudentcertsEvnt(opts *bind.WatchOpts, sink chan<- *UniversityAddstudentcertsEvnt, eventType []string) (event.Subscription, error) {

	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _University.contract.WatchLogs(opts, "addstudentcertsEvnt", eventTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversityAddstudentcertsEvnt)
				if err := _University.contract.UnpackLog(event, "addstudentcertsEvnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddstudentcertsEvnt is a log parse operation binding the contract event 0xeed2d09852776448538f60d68a9c324628faf5fcf8f8e851971543a1a09c21d7.
//
// Solidity: event addstudentcertsEvnt(string indexed eventType, uint256 timestamp, address studentAccounts, string studentName, uint256 CET4, uint256 Credit, uint256 GraduationDissertation)
func (_University *UniversityFilterer) ParseAddstudentcertsEvnt(log types.Log) (*UniversityAddstudentcertsEvnt, error) {
	event := new(UniversityAddstudentcertsEvnt)
	if err := _University.contract.UnpackLog(event, "addstudentcertsEvnt", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniversityAddstudentsubjectEvntIterator is returned from FilterAddstudentsubjectEvnt and is used to iterate over the raw logs and unpacked data for AddstudentsubjectEvnt events raised by the University contract.
type UniversityAddstudentsubjectEvntIterator struct {
	Event *UniversityAddstudentsubjectEvnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniversityAddstudentsubjectEvntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversityAddstudentsubjectEvnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniversityAddstudentsubjectEvnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniversityAddstudentsubjectEvntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversityAddstudentsubjectEvntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversityAddstudentsubjectEvnt represents a AddstudentsubjectEvnt event raised by the University contract.
type UniversityAddstudentsubjectEvnt struct {
	EventType common.Hash
	Timestamp *big.Int
	Subject   string
	Score     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddstudentsubjectEvnt is a free log retrieval operation binding the contract event 0x322b0b25cc9d84a982f6aad4f0c7f62bac7ec85731a5116d44499ef4e053833c.
//
// Solidity: event addstudentsubjectEvnt(string indexed eventType, uint256 timestamp, string Subject, uint256 Score)
func (_University *UniversityFilterer) FilterAddstudentsubjectEvnt(opts *bind.FilterOpts, eventType []string) (*UniversityAddstudentsubjectEvntIterator, error) {

	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _University.contract.FilterLogs(opts, "addstudentsubjectEvnt", eventTypeRule)
	if err != nil {
		return nil, err
	}
	return &UniversityAddstudentsubjectEvntIterator{contract: _University.contract, event: "addstudentsubjectEvnt", logs: logs, sub: sub}, nil
}

// WatchAddstudentsubjectEvnt is a free log subscription operation binding the contract event 0x322b0b25cc9d84a982f6aad4f0c7f62bac7ec85731a5116d44499ef4e053833c.
//
// Solidity: event addstudentsubjectEvnt(string indexed eventType, uint256 timestamp, string Subject, uint256 Score)
func (_University *UniversityFilterer) WatchAddstudentsubjectEvnt(opts *bind.WatchOpts, sink chan<- *UniversityAddstudentsubjectEvnt, eventType []string) (event.Subscription, error) {

	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _University.contract.WatchLogs(opts, "addstudentsubjectEvnt", eventTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversityAddstudentsubjectEvnt)
				if err := _University.contract.UnpackLog(event, "addstudentsubjectEvnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddstudentsubjectEvnt is a log parse operation binding the contract event 0x322b0b25cc9d84a982f6aad4f0c7f62bac7ec85731a5116d44499ef4e053833c.
//
// Solidity: event addstudentsubjectEvnt(string indexed eventType, uint256 timestamp, string Subject, uint256 Score)
func (_University *UniversityFilterer) ParseAddstudentsubjectEvnt(log types.Log) (*UniversityAddstudentsubjectEvnt, error) {
	event := new(UniversityAddstudentsubjectEvnt)
	if err := _University.contract.UnpackLog(event, "addstudentsubjectEvnt", log); err != nil {
		return nil, err
	}
	return event, nil
}
