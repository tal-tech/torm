/*===============================================================
*   Copyright (C) 2020 All rights reserved.
*
*   FileName：hook.go
*   Author：WuGuoFu
*   Date： 2020-06-17
*   Description：
*
================================================================*/
package torm

var (
	TormBefore func() error
	TormAfter  func() error
)

func init() {
	TormBefore = func() error {
		return nil
	}
	TormAfter = func() error {
		return nil
	}
}
