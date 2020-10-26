/*===============================================================
*   Copyright (C) 2020 All rights reserved.
*
*   FileName：hook.go
*   Author：WuGuoFu
*   Date： 2020-10-26
*   Description：
*
================================================================*/
package torm

import (
	"context"

	logger "github.com/tal-tech/loggerX"

	"xorm.io/xorm/contexts"
)

type TormHook struct {
}

func (hook *TormHook) BeforeProcess(ctx *contexts.ContextHook) (context.Context, error) {
	logger.Dx(ctx.Ctx, "[torm]", "torm before process")
	return ctx.Ctx, nil
}

func (hook *TormHook) AfterProcess(ctx *contexts.ContextHook) error {
	logger.Dx(ctx.Ctx, "[torm]", "torm after process")
	return nil
}
