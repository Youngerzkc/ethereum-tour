package service

import (
	"log"

	"github.com/ethereum/go-ethereum/p2p"
)

// method: admin_nodeinfo
func (es *EtherService) GetNodeInfo() (p2p.NodeInfo, error) {
	var resp p2p.NodeInfo
	if err := es.call(&resp, "admin_nodeInfo"); err !=nil{
		log.Fatal(err)
		return resp, err
	}
	return resp, nil

}

func (es *EtherService) GetPeers() ([]p2p.PeerInfo, error){
	var resp []p2p.PeerInfo

	if err := es.call(&resp, "admin_peers"); err != nil {
		log.Fatal(err)
		return resp, err
	}

	return resp, nil
}

func (es *EtherService) GetRpcModules () (map[string]string, error) {
	var resp map[string]string

	if err := es.call(&resp, "rpc_modules"); err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}

func (es *EtherService) StartMiner() (string, error){
	var resp string

	if err := es.call(&resp, "miner_start", 4); err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}

func (es *EtherService) ListAccounts() ([]string, error){
	var resp []string
	if err := es.call(&resp, "personal_listAccounts"); err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}
