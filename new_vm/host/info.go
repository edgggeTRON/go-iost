package host

import (
	"encoding/json"

	"github.com/iost-official/Go-IOS-Protocol/core/contract"
	"github.com/iost-official/Go-IOS-Protocol/new_vm/database"
)

type Info struct {
	ctx *Context
}

func NewInfo(ctx *Context) Info {
	return Info{ctx: ctx}
}

func (h *Info) BlockInfo() (info database.SerializedJSON, cost *contract.Cost) { // todo 清理并且保证正确

	blkInfo := make(map[string]interface{})

	blkInfo["parent_hash"] = h.ctx.Value("parent_hash")
	blkInfo["number"] = h.ctx.Value("number")
	blkInfo["witness"] = h.ctx.Value("witness")
	blkInfo["time"] = h.ctx.Value("time")

	bij, err := json.Marshal(blkInfo)
	if err != nil {
		panic(err)
	}

	return database.SerializedJSON(bij), contract.NewCost(1, 1, 1)
}
func (h *Info) TxInfo() (info database.SerializedJSON, cost *contract.Cost) {

	txInfo := make(map[string]interface{})
	txInfo["time"] = h.ctx.Value("time")
	txInfo["expiration"] = h.ctx.Value("expiration")
	txInfo["gas_limit"] = h.ctx.GValue("gas_limit")
	txInfo["gas_price"] = h.ctx.Value("gas_price")
	txInfo["auth_list"] = h.ctx.Value("auth_list")

	tij, err := json.Marshal(txInfo)
	if err != nil {
		panic(err)
	}

	return database.SerializedJSON(tij), contract.NewCost(1, 1, 1)
}
func (h *Info) ABIConfig(key, value string) {
	switch key {
	case "payment":
		if value == "contract_pay" {
			h.ctx.GSet("abi_payment", 1)
		}
	}
}

func (h *Info) GasLimit() int64 {
	return h.ctx.GValue("gas_limit").(int64)
}
