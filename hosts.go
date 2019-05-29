package sshHostsDeny

import (
	"github.com/caryxiao/go-zlog"
	"strings"
)

type host struct {
	HType string
	Ip    string
}

// 用来记录ssh登录出错的IP
type recordHost struct {
	host
	Cnt int // 记录出现的错误次数
}

type hosts struct {
	Data       map[string]host
	RecordHost map[string]*recordHost
}

func NewHosts() *hosts {
	return &hosts{
		Data:       make(map[string]host),
		RecordHost: make(map[string]*recordHost),
	}
}

func (h *hosts) add(item host) {
	if !h.FindKey(item.Ip) {
		h.Data[item.Ip] = item
	}
}

func (h *hosts) FindKey(key string) bool {
	if _, ok := h.Data[key]; ok {
		return true
	}
	return false
}

func (h *hosts) addFromStrByte(b []byte) {
	hostStr := strings.TrimSpace(string(b))
	if hostStr != "" {
		hostSlice := strings.Split(hostStr, ":")
		if len(hostSlice) == 2 {
			h.formatHost(&hostSlice)
			if hostSlice[0] != "" && hostSlice[1] != "" {
				h.add(host{HType: hostSlice[0], Ip: hostSlice[1]})
			}
		}
	}
}

func (h hosts) formatHost(host *[]string) {
	for i, v := range *host {
		// 处理值的空格问题
		(*host)[i] = strings.TrimSpace(v)
	}
}

func (h *hosts) AddRecordHost(item host) {
	if _, ok := h.RecordHost[item.Ip]; ok {
		h.RecordHost[item.Ip].Cnt += 1
	} else {
		h.RecordHost[item.Ip] = &recordHost{
			host: item,
			Cnt:  1,
		}
	}
}

func (h *hosts) DelRecordHost(item host) {
	delete(h.RecordHost, item.Ip)
}

func (h *hosts) GetRecordHostCnt(key string) int {
	if _, ok := h.RecordHost[key]; ok {
		return h.RecordHost[key].Cnt
	}
	return 0
}

func (h *hosts) GetRecordHost(key string) recordHost {
	if _, ok := h.RecordHost[key]; ok {
		return *h.RecordHost[key]
	}
	return recordHost{}
}

func getSystemHostsDeny(denyFile string) (hs hosts, err error) {
	hs = *NewHosts()
	err = ReadFile(denyFile, hs.addFromStrByte)
	zlog.Logger.Debugf("get system hosts deny: %+v", hs.Data)
	if err != nil {
		return
	}

	return
}
