package dbdao

import (
	"fmt"
	"time"

	"github.com/tal-tech/loggerX"
	"xorm.io/core"
)

//Introduce logger plugin.
var dbLogger = &dbLog{}

//Implement go-xorm core.ILogger interface method.
type dbLog struct{}

//Debug function can print debug log with a line feed.
func (d *dbLog) Debug(v ...interface{}) {
	logger.D("[dbdao]", fmt.Sprint(v...))
}

//Debugf function can print formatted debug log.
func (d *dbLog) Debugf(format string, v ...interface{}) {
	logger.D("[dbdao]", format, v...)
}

//Info function can print Info log with a line feed.
func (d *dbLog) Info(v ...interface{}) {
	logger.I("[dbdao]", fmt.Sprint(v...))
}

//Infof function can print formatted info log.
func (d *dbLog) Infof(format string, v ...interface{}) {
	if len(v) > 0 {
		duration, ok := v[len(v)-1].(time.Duration)
		if ok && duration < slowDuration {
			return
		}
	}
	logger.I("[dbdao]", format, v...)
}

//Warn function can print Warn log with a line feed.
func (d *dbLog) Warn(v ...interface{}) {
	logger.W("[dbdao]", fmt.Sprint(v...))
}

//Warnf function can print formatted Warn log.
func (d *dbLog) Warnf(format string, v ...interface{}) {
	logger.W("[dbdao]", format, v...)
}

//Error function can print Error log with a line feed.
func (d *dbLog) Error(v ...interface{}) {
	logger.E("[dbdao]", fmt.Sprint(v...))
}

//Errorf function can print formatted Error log.
func (d *dbLog) Errorf(format string, v ...interface{}) {
	logger.E("[dbdao]", format, v...)
}

func (d *dbLog) Level() core.LogLevel {
	return core.LOG_INFO
}

func (d *dbLog) SetLevel(l core.LogLevel) {
	return
}

func (d *dbLog) ShowSQL(show ...bool) {
	return
}

func (d *dbLog) IsShowSQL() bool {
	return showSql
}
