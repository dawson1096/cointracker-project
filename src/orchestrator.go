package src

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Orchestrator struct {
	mu          sync.Mutex
	AddressMap  map[string]*Ledger
	LatestBlock *BlockResponse
}

func (orc *Orchestrator) Init() {
	orc.AddressMap = make(map[string]*Ledger)

	go func() {
		for true {
			block, err := getLatestBlock()
			if err != nil {
				fmt.Println(err)
				continue
			}

			if orc.LatestBlock == nil || orc.LatestBlock.Height+1 == block.Height {
				latestBlock, err := getBlock(block.Hash)
				if err != nil {
					fmt.Println(err)
					continue
				}

				orc.LatestBlock = latestBlock
				orc.updateAddresses()
			}
		}
	}()
}

func (orc *Orchestrator) AddAddress(c *gin.Context) {
	addr := c.Param("address")

	//Verify address is valid btc addr
	if !validateAddress(addr) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid address format"})
		return
	}

	orc.mu.Lock()
	defer orc.mu.Unlock()
	if _, ok := orc.AddressMap[addr]; ok {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "address already exists"})
		return
	}
	syncRes, err := getAddressStats(addr)
	if err != nil {
		delete(orc.AddressMap, addr)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "internal error"})
		return
	}

	ledger := &Ledger{
		Address: addr,
		Balance: syncRes.Final_balance,
		Txns:    syncRes.Txs,
	}
	orc.AddressMap[addr] = ledger
	c.JSON(http.StatusCreated, gin.H{"success": true, "result": addr})
}

func (orc *Orchestrator) RemoveAddress(c *gin.Context) {
	addr := c.Param("address")

	if !validateAddress(addr) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid address format"})
		return
	}

	orc.mu.Lock()
	defer orc.mu.Unlock()
	if _, ok := orc.AddressMap[addr]; ok {
		delete(orc.AddressMap, addr)
		c.JSON(http.StatusOK, gin.H{"success": true, "address": addr})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "address not found"})
}

func (orc *Orchestrator) SyncAddress(c *gin.Context) {
	addr := c.Param("address")

	if !validateAddress(addr) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid address format"})
		return
	}

	orc.mu.Lock()
	defer orc.mu.Unlock()
	if _, ok := orc.AddressMap[addr]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "address not found"})
		return
	}

	syncRes, err := getAddressStats(addr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "internal error"})
		return
	}

	ledger := &Ledger{
		Address: addr,
		Balance: syncRes.Final_balance,
		Txns:    syncRes.Txs,
	}
	orc.AddressMap[addr] = ledger

	c.JSON(http.StatusCreated, gin.H{"success": true, "result": ledger})
}

func (orc *Orchestrator) GetAddress(c *gin.Context) {
	addr := c.Param("address")

	if !validateAddress(addr) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid address format"})
		return
	}

	orc.mu.Lock()
	defer orc.mu.Unlock()
	if val, ok := orc.AddressMap[addr]; ok {
		c.JSON(http.StatusOK, gin.H{"success": true, "result": val})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "address not found"})
}

func (orc *Orchestrator) ListAddresses(c *gin.Context) {
	addresses := []string{}
	orc.mu.Lock()
	defer orc.mu.Unlock()
	for addr := range orc.AddressMap {
		addresses = append(addresses, addr)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "result": addresses})
}

func (orc *Orchestrator) updateAddresses() {
	orc.mu.Lock()
	defer orc.mu.Unlock()

	if len(orc.AddressMap) == 0 {
		return
	}

	checkAddrs := []string{}
	checkAddrsMap := make(map[string]bool)

	for _, txn := range orc.LatestBlock.Tx {
		for _, input := range txn.Inputs {
			addr := input.Prev_out.Addr
			if _, ok := orc.AddressMap[addr]; ok {
				if _, ok := checkAddrsMap[addr]; !ok {
					checkAddrs = append(checkAddrs, addr)
					checkAddrsMap[addr] = true
				}
			}
		}

		for _, out := range txn.Out {
			addr := out.Addr
			if _, ok := orc.AddressMap[addr]; ok {
				if _, ok := checkAddrsMap[addr]; !ok {
					checkAddrs = append(checkAddrs, addr)
					checkAddrsMap[addr] = true
				}
			}
		}
	}

	for _, addr := range checkAddrs {
		syncRes, err := getAddressStats(addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ledger := &Ledger{
			Address: addr,
			Balance: syncRes.Final_balance,
			Txns:    syncRes.Txs,
		}
		orc.AddressMap[addr] = ledger
	}

}
