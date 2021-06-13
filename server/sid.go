package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
)

// SessIDManager数据结构及管理
type SessIDManager struct {
	m map[string]int
	l sync.RWMutex
}

// 添加一个用户
// 生成新sid, 添加并返回
func (s *SessIDManager) set(uid int) string {
	s.l.Lock()
	for sid, u := range s.m {
		if u == uid {
			delete(s.m, sid)
		}
	}

	sid := newSID()
	_, ok := s.m[sid]
	for ok {
		sid = newSID()
		_, ok = s.m[sid]
	}
	s.m[sid] = uid
	s.l.Unlock()
	return sid
}

// 用sid查询登录状态
// 已登录, 返回uid, nil
// 未登录, 返回0, err
func (s *SessIDManager) get(sid string) (int, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	uid, ok := s.m[sid]
	if ok {
		return uid, nil
	}
	return 0, errors.New(sid + " not found!\n")
}

// 随机生成一个sid
func newSID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
