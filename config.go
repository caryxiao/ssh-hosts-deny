package sshHostsDeny

import (
	"github.com/caryxiao/go-zlog"
	"os"
)

type CmdConfig struct {
	SecureFile      string // secure file path
	DenyFile        string // host deny file
	SshLoginFailCnt int    // default ssh login failed count
	PrintVer        bool   //print version
}

// check the config are correct
func (config CmdConfig) validate() (err error) {
	if _, err = os.Stat(config.SecureFile); err != nil {
		if os.IsNotExist(err) {
			zlog.Logger.Errorf("file is not exit: %s", config.SecureFile)
			return
		}
	}
	return
}
