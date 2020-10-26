package torm

import (
	"context"

	"xorm.io/xorm"
	"xorm.io/xorm/contexts"
)

type Engine struct {
	*xorm.Engine
}

// Context creates a session with the context
func (engine *Engine) Context(ctx context.Context) *Session {
	session := engine.NewSession()
	return session.Context(ctx)
}

// SetDefaultContext set the default context
func (engine *Engine) SetDefaultContext(ctx context.Context) {
	engine.Engine.SetDefaultContext(ctx)
}

// PingContext tests if database is alive
func (engine *Engine) PingContext(ctx context.Context) error {
	session := engine.NewSession()
	return session.PingContext(ctx)
}

func (engine *Engine) NoCache() *Session {
	session := engine.NewSession()
	session = session.NoCache()
	return session
}

// NoCascade If you do not want to auto cascade load object
func (engine *Engine) NoCascade() *Session {
	session := engine.NewSession()
	session = session.NoCascade()
	return session
}

// NewSession New a session
func (engine *Engine) NewSession() *Session {
	session := &Session{Session: engine.Engine.NewSession(), engine: engine}
	return session
}

// Sql provides raw sql input parameter. When you have a complex SQL statement
// and cannot use Where, Id, In and etc. Methods to describe, you can use SQL.
//
// Deprecated: use SQL instead.
func (engine *Engine) Sql(querystring string, args ...interface{}) *Session {
	session := engine.NewSession()
	session.Session = engine.Engine.SQL(querystring, args...)
	return session
}

// SQL method let's you manually write raw SQL and operate
// For example:
//
//         engine.SQL("select * from user").Find(&users)
//
// This    code will execute "select * from user" and set the records to users
func (engine *Engine) SQL(query interface{}, args ...interface{}) *Session {
	session := engine.NewSession()
	session.Session = session.Session.SQL(query, args...)
	return session
}

// NoAutoTime Default if your struct has "created" or "updated" filed tag, the fields
// will automatically be filled with current time when Insert or Update
// invoked. Call NoAutoTime if you dont' want to fill automatically.
func (engine *Engine) NoAutoTime() *Session {
	session := engine.NewSession()
	session.Session = session.Session.NoAutoTime()
	return session
}

// NoAutoCondition disable auto generate Where condition from bean or not
func (engine *Engine) NoAutoCondition(no ...bool) *Session {
	session := engine.NewSession()
	session.Session = session.Session.NoAutoCondition(no...)
	return session
}

// Cascade use cascade or not
func (engine *Engine) Cascade(trueOrFalse ...bool) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Cascade(trueOrFalse...)
	return session
}

// Where method provide a condition query
func (engine *Engine) Where(query interface{}, args ...interface{}) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Where(query, args...)
	return session
}

// ID method provoide a condition as (id) = ?
func (engine *Engine) ID(id interface{}) *Session {
	session := engine.NewSession()
	session.Session = session.Session.ID(id)
	return session
}

// Before apply before Processor, affected bean is passed to closure arg
func (engine *Engine) Before(closures func(interface{})) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Before(closures)
	return session
}

// After apply after insert Processor, affected bean is passed to closure arg
func (engine *Engine) After(closures func(interface{})) *Session {
	session := engine.NewSession()
	session.Session = session.Session.After(closures)
	return session
}

// Charset set charset when create table, only support mysql now
func (engine *Engine) Charset(charset string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Charset(charset)
	return session
}

// StoreEngine set store engine when create table, only support mysql now
func (engine *Engine) StoreEngine(storeEngine string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.StoreEngine(storeEngine)
	return session
}

// Distinct use for distinct columns. Caution: when you are using cache,
// distinct will not be cached because cache system need id,
// but distinct will not provide id
func (engine *Engine) Distinct(columns ...string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Distinct(columns...)
	return session
}

// Select customerize your select columns or contents
func (engine *Engine) Select(str string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Select(str)
	return session
}

// Cols only use the parameters as select or update columns
func (engine *Engine) Cols(columns ...string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Cols(columns...)
	return session
}

// AllCols indicates that all columns should be use
func (engine *Engine) AllCols() *Session {
	session := engine.NewSession()
	session.Session = session.Session.AllCols()
	return session
}

// MustCols specify some columns must use even if they are empty
func (engine *Engine) MustCols(columns ...string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.MustCols(columns...)
	return session
}

// UseBool xorm automatically retrieve condition according struct, but
// if struct has bool field, it will ignore them. So use UseBool
// to tell system to do not ignore them.
// If no parameters, it will use all the bool field of struct, or
// it will use parameters's columns
func (engine *Engine) UseBool(columns ...string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.UseBool(columns...)
	return session
}

// Omit only not use the parameters as select or update columns
func (engine *Engine) Omit(columns ...string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Omit(columns...)
	return session
}

// Nullable set null when column is zero-value and nullable for update
func (engine *Engine) Nullable(columns ...string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Nullable(columns...)
	return session
}

// In will generate "column IN (?, ?)"
func (engine *Engine) In(column string, args ...interface{}) *Session {
	session := engine.NewSession()
	session.Session = session.Session.In(column, args...)
	return session
}

// NotIn will generate "column NOT IN (?, ?)"
func (engine *Engine) NotIn(column string, args ...interface{}) *Session {
	session := engine.NewSession()
	session.Session = session.Session.NotIn(column, args...)
	return session
}

// Incr provides a update string like "column = column + ?"
func (engine *Engine) Incr(column string, arg ...interface{}) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Incr(column, arg...)
	return session
}

// Decr provides a update string like "column = column - ?"
func (engine *Engine) Decr(column string, arg ...interface{}) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Decr(column, arg...)
	return session
}

// SetExpr provides a update string like "column = {expression}"
func (engine *Engine) SetExpr(column string, expression interface{}) *Session {
	session := engine.NewSession()
	session.Session = session.Session.SetExpr(column, expression)
	return session
}

// Table temporarily change the Get, Find, Update's table
func (engine *Engine) Table(tableNameOrBean interface{}) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Table(tableNameOrBean)
	return session
}

// Alias set the table alias
func (engine *Engine) Alias(alias string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Alias(alias)
	return session
}

// Limit will generate "LIMIT start, limit"
func (engine *Engine) Limit(limit int, start ...int) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Limit(limit, start...)
	return session
}

// Desc will generate "ORDER BY column1 DESC, column2 DESC"
func (engine *Engine) Desc(colNames ...string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Desc(colNames...)
	return session
}

// Asc will generate "ORDER BY column1,column2 Asc"
// This method can chainable use.
//
//        engine.Desc("name").Asc("age").Find(&users)
//        // SELECT * FROM user ORDER BY name DESC, age ASC
//
func (engine *Engine) Asc(colNames ...string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.Asc(colNames...)
	return session
}

// OrderBy will generate "ORDER BY order"
func (engine *Engine) OrderBy(order string) *Session {
	session := engine.NewSession()
	session.Session = session.Session.OrderBy(order)
	return session
}

// Prepare enables prepare statement
func (engine *Engine) Prepare() *Session {
	session := engine.NewSession()
	session = session.Prepare()
	return session
}

// Join the join_operator should be one of INNER, LEFT OUTER, CROSS etc - this will be prepended to JOIN
func (engine *Engine) Join(joinOperator string, tablename interface{}, condition string, args ...interface{}) *Session {
	session := engine.NewSession()
	session = session.Join(joinOperator, tablename, condition, args...)
	return session
}

// GroupBy generate group by statement
func (engine *Engine) GroupBy(keys string) *Session {
	session := engine.NewSession()
	session = session.GroupBy(keys)
	return session
}

// Having generate having statement
func (engine *Engine) Having(conditions string) *Session {
	session := engine.NewSession()
	session = session.Having(conditions)
	return session
}

func (engine *Engine) AddHook(hook contexts.Hook) {
	engine.Engine.AddHook(hook)
}
