package src

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func getAddressStats(addr string) (*SyncResponse, error) {
	endpoint := fmt.Sprintf("https://blockchain.info/rawaddr/%s", addr)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	syncRes := &SyncResponse{}
	err = json.Unmarshal(body, syncRes)

	return syncRes, nil
}

func getLatestBlock() (*LatestBlock, error) {
	endpoint := "https://blockchain.info/latestblock"
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	latestBlock := &LatestBlock{}
	err = json.Unmarshal(body, latestBlock)

	return latestBlock, nil
}

func getBlock(hash string) (*BlockResponse, error) {
	endpoint := fmt.Sprintf("https://blockchain.info/rawblock/%s", hash)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	block := &BlockResponse{}
	err = json.Unmarshal(body, block)

	return block, nil
}
