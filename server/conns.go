package server

import (
	"sync"

	"github.com/xiaolaji422/golink/lib/helper"
)

var Conns *connPool

func init() {
	Conns = &connPool{
		mutex: &sync.RWMutex{},
		sets:  make(map[string]*connSet, 0),
	}
}

// 管理所有链接
type connPool struct {
	mutex *sync.RWMutex
	sets  map[string]*connSet // 所有的链接
}

// 获取所有链接

// 获取set
func (c *connPool) getSet(AppId string) *connSet {
	var connLink *connSet
	if v, ok := c.sets[AppId]; ok {
		connLink = v
	} else {
		connLink = &connSet{
			mutex: &sync.RWMutex{},
			pool:  make(map[string]*Conn),
		}
		c.mutex.Lock()
		c.sets[AppId] = connLink
		c.mutex.Unlock()
	}
	return connLink
}

// func close
func (c *connPool) GetConn(AppId, UserId string) *Conn {
	var set = c.getSet(AppId)
	if helper.IsNil(set) {
		return nil
	}
	return set.getConn(UserId)
}

// add  conn to clients
func (c *connPool) AddConn(AppId, UserId string, conn *Conn) {
	var set = c.getSet(AppId)
	if helper.IsNil(set) {
		set = &connSet{}
	}
	set.addConn(UserId, conn)
}

// 链表
type connSet struct {
	mutex *sync.RWMutex
	pool  map[string]*Conn
}

// 获取链接
func (c *connSet) getConn(UserId string) *Conn {
	if v, ok := c.pool[UserId]; ok {
		return v
	} else {
		return nil
	}
}

// 获取链接
func (c *connSet) addConn(UserId string, conn *Conn) {
	if v, ok := c.pool[UserId]; ok { // 关闭持有的链接
		v.Close()
	}
	c.mutex.Lock()
	c.pool[UserId] = conn
	c.mutex.Unlock()
}
