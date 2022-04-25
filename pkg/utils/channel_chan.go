package utils

import (
	"awesomeProject/app/models"
	"awesomeProject/platform/database"
	"github.com/google/uuid"
	"sync"
)

type WaitMapObject struct {
	wg   map[uuid.UUID]int
	mu   sync.Mutex
	cond sync.Cond
}

func WaitMap() *WaitMapObject {
	m := &WaitMapObject{}
	m.wg = make(map[uuid.UUID]int)
	m.cond.L = &m.mu
	return m
}

func (m *WaitMapObject) Wait(name uuid.UUID) {
	m.mu.Lock()
	for m.wg[name] != 0 {
		m.cond.Wait()
	}
	m.mu.Unlock()
}

func (m *WaitMapObject) Done(name uuid.UUID) {
	m.mu.Lock()
	no := m.wg[name] - 1
	if no < 0 {
		panic("")
	}
	m.wg[name] = no
	m.mu.Unlock()
	m.cond.Broadcast()
}

func (m *WaitMapObject) Add(name uuid.UUID, no int) {
	m.mu.Lock()
	m.wg[name] = m.wg[name] + no
	m.mu.Unlock()
}

var clients []models.Client
var wgm *WaitMapObject
var Cm map[uuid.UUID]chan uuid.UUID

func worker(transactionChan <-chan uuid.UUID, id uuid.UUID) {
	defer wgm.Done(id)

	for transaction := range transactionChan {
		ApproveTransaction(transaction)
	}
}

func init() {
	wgm := WaitMap()
	Cm = make(map[uuid.UUID]chan uuid.UUID)
	db, err := database.OpenDBConnection()
	if err != nil {
		return
	}
	clients, err := db.GetClients()
	for _, client := range clients {
		wgm.Add(client.Id, 1)
		Cm[client.Id] = make(chan uuid.UUID, 100000)
		go worker(Cm[client.Id], client.Id)
	}
	// wait for all workers to exit
	/*
		for _, client := range clients {
			wgm.Wait(client.Id)
		}*/

}
