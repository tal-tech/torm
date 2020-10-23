package torm

import "context"

type DbBaseDao struct {
	Engine  *Engine
	Session *Session
}

type Param interface{}
type ParamNil struct{}
type ParamDesc bool
type ParamIn []interface{}
type ParamRange struct {
	Min interface{}
	Max interface{}
}
type ParamInDesc ParamIn
type ParamRangeDesc ParamRange

//cast input to []interface{} type.
func CastToParamIn(input interface{}) ParamIn {
	params := make(ParamIn, 0)
	switch v := input.(type) {
	case []interface{}:
		for _, param := range v {
			params = append(params, param)
		}
	case []int64:
		for _, param := range v {
			params = append(params, param)
		}
	case []int:
		for _, param := range v {
			params = append(params, param)
		}
	case []int32:
		for _, param := range v {
			params = append(params, param)
		}
	case []int8:
		for _, param := range v {
			params = append(params, param)
		}
	case []uint64:
		for _, param := range v {
			params = append(params, param)
		}
	case []uint:
		for _, param := range v {
			params = append(params, param)
		}
	case []uint32:
		for _, param := range v {
			params = append(params, param)
		}
	case []uint8:
		for _, param := range v {
			params = append(params, param)
		}
	case []string:
		for _, param := range v {
			params = append(params, param)
		}
	default:
		params = append(params, 0)
	}
	return params
}

//cast input to ParamIn type.
func CastToParamInDesc(input interface{}) ParamInDesc {
	return ParamInDesc(CastToParamIn(input))
}

//Mysql instance set session.
func (this *DbBaseDao) InitSession(ctx context.Context) {
	if this.Session == nil {
		this.Session = this.Engine.Context(ctx)
	}
}

//SetTable can specify a table name.
func (this *DbBaseDao) SetTable(ctx context.Context, tableName string) {
	if this.Session == nil {
		this.InitSession(ctx)
	}

	this.Session.Table(tableName)
}

// Construct sql query statement and execute.
func (this *DbBaseDao) BuildQuery(input Param, name string) {
	name = this.Engine.Quote(name)

	switch val := input.(type) {
	case ParamDesc:
		if val {
			this.Session = this.Session.Desc(name)
		}
	case ParamIn:
		if len(val) == 1 {
			this.Session = this.Session.And(name+"=?", val[0])
		} else {
			this.Session = this.Session.In(name, val)
		}
	case ParamInDesc:
		if len(val) == 1 {
			this.Session = this.Session.And(name+"=?", val[0])
		} else {
			this.Session = this.Session.In(name, val)
		}
		this.Session = this.Session.Desc(name)
	case ParamRange:
		if val.Min != nil {
			this.Session = this.Session.And(name+">=?", val.Min)
		}
		if val.Max != nil {
			this.Session = this.Session.And(name+"<?", val.Max)
		}
	case ParamRangeDesc:
		if val.Min != nil {
			this.Session = this.Session.And(name+">=?", val.Min)
		}
		if val.Max != nil {
			this.Session = this.Session.And(name+"<?", val.Max)
		}
		this.Session = this.Session.Desc(name)
	case ParamNil:
	case nil:
	default:
		this.Session = this.Session.And(name+"=?", val)
	}
}

// Update MySQL execution engine.
func (this *DbBaseDao) UpdateEngine(v ...interface{}) {
	if len(v) == 0 {
		this.Engine = GetDefault("reader").Engine
		this.Session = nil
	} else if len(v) == 1 {
		param := v[0]
		if engine, ok := param.(*Engine); ok {
			this.Engine = engine
			this.Session = nil
		} else if session, ok := param.(*Session); ok {
			this.Session = session
		} else if tpe, ok := param.(bool); ok {
			cluster := "reader"
			if tpe == true {
				cluster = "writer"
			}
			this.Engine = GetDefault(cluster).Engine
			this.Session = nil
		}
	}
}
