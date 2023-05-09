package wlc

import "fmt"

type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (this Error) Error() string {
	return fmt.Sprintf("%d-%s", this.ErrCode, this.ErrMsg)
}
