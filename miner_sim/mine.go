package miner_sim

import (
    "log"
)

type Mine struct {
    miner_clients []*MinerClient
}

func (m * Mine) AddMiner(miner *MinerClient) {
    m.miner_clients = append(m.miner_clients, miner)
}

func (m * Mine) ConnectMiners(pool_server_address string) {
    log.Println("Connecting miners now...")

    for _, miner_client := range m.miner_clients {
        go miner_client.ConnectToPool(pool_server_address)
    }
}
