package protocol

import (
	"errors"
	"strings"
)

type ProtocolType string

const (
	None        = ""
	Vmess       = "vmess"
	Shadowsocks = "ss"
)

func GetProtocol(uri string) (ProtocolType, error) {
	split := strings.Split(uri, "://")
	if len(split) == 0 {
		return noneProtocolType(uri)
	}
	switch split[0] {
	case "ss":
		return Shadowsocks, nil
	case "vmess":
		return Vmess, nil
	}
	return noneProtocolType(uri)
}

func noneProtocolType(uri string) (ProtocolType, error) {
	return None, errors.New("no supported protocol found: " + uri)
}
