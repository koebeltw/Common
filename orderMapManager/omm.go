package orderMapManager

import (
	"diamond/tool/common"
	"errors"
	"math"
	"sync"
)

// OMM OrderMapManager blabla
type OMM struct {
	rw    sync.RWMutex
	datas map[uint64]interface{}
	index uint64
	max   uint64
	count uint64
}

// NewOMM blabla
func NewOMM(max uint64) *OMM {
	if max == 0 {
		return nil
	}

	return &OMM{max: max, datas: make(map[uint64]interface{}, max)}
}

// Add blabla
func (o *OMM) Add(data interface{}) (index uint64, err error) {
	defer o.rw.Unlock()
	var count uint64
	err = errors.New("255")

	if data != nil {
		o.rw.Lock()
		for {
			if o.index == math.MaxUint64 {
				o.index = 0
			} else {
				o.index = o.index + 1
			}

			if o.datas[o.index] == nil {
				o.datas[o.index] = data
				index = o.index
				o.count = o.count + 1
				err = errors.New("0")
				return
			}

			count = count + 1
			if count == o.max {
				err = errors.New("1")
				return
			}
		}
	}

	err = errors.New("1")
	return
}

// Del blabla
func (o *OMM) Del(index uint64) (err error) {
	defer o.rw.Unlock()
	err = errors.New("255")
	o.rw.Lock()
	delete(o.datas, index)
	o.count = o.count - 1
	err = errors.New("0")
	return
}

// Clear blabla
func (o *OMM) Clear() (err error) {
	defer o.rw.Unlock()
	err = errors.New("255")

	o.rw.Lock()
	for key := range o.datas {
		delete(o.datas, key)

	}

	o.datas = nil
	o.index = 0
	o.count = 0
	err = errors.New("0")
	return
}

//Put 存儲操作
func (o *OMM) Put(index uint64, datas interface{}) {
	defer o.rw.Unlock()
	o.rw.Lock()

	if o.datas[index] == nil {
		o.count = o.count + 1
	}
	o.datas[index] = datas
}

//Get 獲取操作
func (o *OMM) Get(index uint64) interface{} {
	defer o.rw.RUnlock()
	o.rw.RLock()

	return o.datas[index]
}

// ForEach blabla
func (o *OMM) ForEach(callBackfunc common.CallBackfunc) {
	o.rw.RLock()
	defer o.rw.RUnlock()

	for k, v := range o.datas {
		callBackfunc(k, v)
	}
}
