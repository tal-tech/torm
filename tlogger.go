/*===============================================================
*   Copyright (C) 2020 All rights reserved.
*
*   FileName：tlog.go
*   Author：WuGuoFu
*   Date： 2020-10-21
*   Description：
*
================================================================*/
package torm

import (
	"context"
	"fmt"

	logger "github.com/tal-tech/loggerX"
	"xorm.io/xorm/log"
)

// TormLogger is the default implment of ILogger
type TormLogger struct {
	level   log.LogLevel
	showSQL bool
	ctx     context.Context
}

// Error implement ILogger
func (s *TormLogger) Error(v ...interface{}) {
	if s.level <= log.LOG_ERR {
		logger.Ex(s.ctx, "[torm]", fmt.Sprintln(v...))
	}
	return
}

// Errorf implement ILogger
func (s *TormLogger) Errorf(format string, v ...interface{}) {
	if s.level <= log.LOG_ERR {
		logger.Ex(s.ctx, "[torm]", fmt.Sprintf(format, v...))
	}
	return
}

// Debug implement ILogger
func (s *TormLogger) Debug(v ...interface{}) {
	if s.level <= log.LOG_DEBUG {
		logger.Dx(s.ctx, "[torm]", fmt.Sprintln(v...))
	}
	return
}

// Debugf implement ILogger
func (s *TormLogger) Debugf(format string, v ...interface{}) {
	if s.level <= log.LOG_DEBUG {
		logger.Dx(s.ctx, "[torm]", fmt.Sprintf(format, v...))
	}
	return
}

// Info implement ILogger
func (s *TormLogger) Info(v ...interface{}) {
	if s.level <= log.LOG_INFO {
		logger.Ix(s.ctx, "[torm]", fmt.Sprintln(v...))
	}
	return
}

// Infof implement ILogger
func (s *TormLogger) Infof(format string, v ...interface{}) {
	if s.level <= log.LOG_INFO {
		logger.Ix(s.ctx, "[torm]", fmt.Sprintf(format, v...))
	}
	return
}

// Warn implement ILogger
func (s *TormLogger) Warn(v ...interface{}) {
	if s.level <= log.LOG_WARNING {
		logger.Wx(s.ctx, "[torm]", fmt.Sprintln(v...))
	}
	return
}

// Warnf implement ILogger
func (s *TormLogger) Warnf(format string, v ...interface{}) {
	if s.level <= log.LOG_WARNING {
		logger.Wx(s.ctx, "[torm]", fmt.Sprintf(format, v...))
	}
	return
}

// Level implement ILogger
func (s *TormLogger) Level() log.LogLevel {
	return s.level
}

// SetLevel implement ILogger
func (s *TormLogger) SetLevel(l log.LogLevel) {
	s.level = l
	return
}

// ShowSQL implement ILogger
func (s *TormLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL implement ILogger
func (s *TormLogger) IsShowSQL() bool {
	return s.showSQL
}
