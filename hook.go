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

import "database/sql"

var (
	HookExec   func(func() (sql.Result, error)) (sql.Result, error)
	HookGet    func(func() (bool, error)) (bool, error)
	HookFind   func(func() error) error
	HookInsert func(func() (int64, error)) (int64, error)
	HookDelete func(func() (int64, error)) (int64, error)
	HookQuery  func(func() ([]map[string][]byte, error)) ([]map[string][]byte, error)
	HookUpdate func(func() (int64, error)) (int64, error)
)

func init() {
	HookExec = func(fn func() (sql.Result, error)) (sql.Result, error) {
		return fn()
	}
	HookGet = func(fn func() (bool, error)) (bool, error) {
		return fn()
	}
	HookFind = func(fn func() error) error {
		return fn()
	}
	HookInsert = func(fn func() (int64, error)) (int64, error) {
		return fn()
	}
	HookDelete = func(fn func() (int64, error)) (int64, error) {
		return fn()
	}
	HookQuery = func(fn func() ([]map[string][]byte, error)) ([]map[string][]byte, error) {
		return fn()
	}
	HookUpdate = func(fn func() (int64, error)) (int64, error) {
		return fn()
	}
}
