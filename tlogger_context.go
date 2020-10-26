/*===============================================================
*   Copyright (C) 2020 All rights reserved.
*
*   FileName：tlogger_context.go
*   Author：WuGuoFu
*   Date： 2020-10-21
*   Description：
*
================================================================*/
package torm

import (
	"fmt"

	"xorm.io/xorm/log"
)

// TormLoggerAdapter wraps a Logger interface as LoggerContext interface
type TormLoggerAdapter struct {
	logger TormLogger
}

// NewTormLoggerAdapter creates an adapter for old xorm logger interface
func NewTormLoggerAdapter(logger TormLogger) log.ContextLogger {
	return &TormLoggerAdapter{
		logger: logger,
	}
}

// BeforeSQL implements ContextTormLogger
func (l *TormLoggerAdapter) BeforeSQL(ctx log.LogContext) {
	l.logger.ctx = ctx.Ctx
}

// AfterSQL implements ContextTormLogger
func (l *TormLoggerAdapter) AfterSQL(ctx log.LogContext) {
	var sessionPart string
	v := ctx.Ctx.Value(log.SessionIDKey)
	if key, ok := v.(string); ok {
		sessionPart = fmt.Sprintf(" [%s]", key)
	}
	if ctx.ExecuteTime > 0 {
		l.logger.Infof("[SQL]%s %s %v - %v", sessionPart, ctx.SQL, ctx.Args, ctx.ExecuteTime)
	} else {
		l.logger.Infof("[SQL]%s %s %v", sessionPart, ctx.SQL, ctx.Args)
	}
}

// Debugf implements ContextTormLogger
func (l *TormLoggerAdapter) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

// Errorf implements ContextTormLogger
func (l *TormLoggerAdapter) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

// Infof implements ContextTormLogger
func (l *TormLoggerAdapter) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

// Warnf implements ContextTormLogger
func (l *TormLoggerAdapter) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

// Level implements ContextTormLogger
func (l *TormLoggerAdapter) Level() log.LogLevel {
	return l.logger.Level()
}

// SetLevel implements ContextTormLogger
func (l *TormLoggerAdapter) SetLevel(lv log.LogLevel) {
	l.logger.SetLevel(lv)
}

// ShowSQL implements ContextTormLogger
func (l *TormLoggerAdapter) ShowSQL(show ...bool) {
	l.logger.ShowSQL(show...)
}

// IsShowSQL implements ContextTormLogger
func (l *TormLoggerAdapter) IsShowSQL() bool {
	return l.logger.IsShowSQL()
}
