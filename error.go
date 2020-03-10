// Copyright 2018 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package wechat

import "fmt"

type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (we *Error) Error() string {
	return fmt.Sprintf("code:[%d] error:[%s]", we.ErrCode, we.ErrMsg)
}
