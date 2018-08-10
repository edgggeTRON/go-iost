package new_vm

import (
	"github.com/iost-official/Go-IOS-Protocol/core/contract"
	"github.com/iost-official/Go-IOS-Protocol/new_vm/host"
)

//go:generate mockgen -destination vm_mock.go -package new_vm github.com/iost-official/Go-IOS-Protocol/new_vm VM

type VM interface {
	Init() error
	LoadAndCall(host host.IHost, contract *contract.Contract, api string, args ...interface{}) (rtn []interface{}, cost *contract.Cost, err error)
	Release()
}