package core

import (
	"AQChain/models"
	"errors"
	"sync"
)

// 清除交易缓存
func clearTxCache(cache *TxSet, txs []*models.Transaction) {
	for _, tx := range txs {
		cache.Remove(tx)
	}
}

type TxSet struct {
	mtx  sync.Mutex
	txs  map[string]*txSetItem
	list []*models.Transaction
}

type txSetItem struct {
	tx    *models.Transaction
	index int
}

func NewTxSet() *TxSet {
	return &TxSet{
		txs:  make(map[string]*txSetItem, 0),
		list: make([]*models.Transaction, 0),
	}
}

func (ps *TxSet) Add(tx *models.Transaction) error {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.txs[tx.Hash] != nil {
		return errors.New("重复的交易")
	}
	ps.txs[tx.Hash] = &txSetItem{tx, len(ps.list)}
	ps.list = append(ps.list, tx)

	return nil
}

func (ps *TxSet) Get(peerID string) *models.Transaction {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	item, ok := ps.txs[peerID]
	if ok {
		return item.tx
	}
	return nil
}

func (ps *TxSet) Has(peerID string) bool {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	_, ok := ps.txs[peerID]
	return ok
}

func (ps *TxSet) List() []*models.Transaction {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.list
}

func (ps *TxSet) Remove(tx *models.Transaction) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	item := ps.txs[tx.Hash]
	if item == nil {
		return
	}

	index := item.index
	// Copy the list but without the last element.
	// (we must copy because we're mutating the list)
	newList := make([]*models.Transaction, len(ps.list)-1)
	copy(newList, ps.list)
	// If it's the last con, that's an easy special case.
	if index == len(ps.list)-1 {
		ps.list = newList
		delete(ps.txs, tx.Hash)
		return
	}

	// Move the last item from ps.list to "index" in list.
	lastTx := ps.list[len(ps.list)-1]
	lastTxHash := lastTx.Hash
	lastTxItem := ps.txs[lastTxHash]
	newList[index] = lastTx
	lastTxItem.index = index
	ps.list = newList
	delete(ps.txs, tx.Hash)
}

func (ps *TxSet) Size() int {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return len(ps.list)
}
