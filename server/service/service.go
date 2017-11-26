// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package service

import (
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

type EtherService struct {
	client *rpc.Client
}

func NewEtherService() *EtherService {
	return &EtherService{
		client: nil,
	}
}

func (es *EtherService) call(result interface{}, method string, args ...interface{}) error{
	client, err := es.getClient()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := client.Call(result, method, args...); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (es *EtherService) getClient() (*rpc.Client, error) {
	if es.client != nil {
		log.Println("exist one client already")
		return es.client, nil
	}

	client, err := rpc.Dial("http://127.0.0.1:8201")
	if err != nil {
		log.Fatal(err)
		return client, err
	}

	es.client = client
	log.Println("connected new client")
	return es.client, nil
}
