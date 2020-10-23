/*===============================================================
*   Copyright (C) 2020 All rights reserved.
*
*   FileName：db_test.go
*   Author：WuGuoFu
*   Date： 2020-06-15
*   Description：
*
================================================================*/

package torm

import (
	"context"
	"testing"
	"time"
)

type GolangStats struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Metric     string    `xorm:"not null default '' comment('指标名') VARCHAR(100)"`
	Endpoint   string    `xorm:"not null default '' comment('机器名') VARCHAR(100)"`
	Tags       string    `xorm:"not null default '' comment('服务名') VARCHAR(100)"`
	Ctime      int       `xorm:"not null default 0 comment('时间戳') index INT(11)"`
	Value      int       `xorm:"not null default 0 comment('值') INT(11)"`
	Ctype      int       `xorm:"not null default 1 comment('统计类型,1:计数 2:耗时') index TINYINT(2)"`
	CreateTime time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('提交时间') DATETIME"`
}

type GolangStatsDao struct {
	DbBaseDao
}

func NewGolangStatsDao(v ...interface{}) *GolangStatsDao {
	this := new(GolangStatsDao)
	if ins := GetDbInstance("default", "writer"); ins != nil {
		this.UpdateEngine(ins.Engine)
	} else {
		return nil
	}
	if len(v) != 0 {
		this.UpdateEngine(v...)
	}
	return this
}

func (this *GolangStatsDao) Get(mId Param) (ret []GolangStats, err error) {
	ret = make([]GolangStats, 0)
	this.InitSession(context.Background())

	this.BuildQuery(mId, "id")

	err = this.Session.Find(&ret)
	return
}
func (this *GolangStatsDao) GetLimit(mId Param, pn, rn int) (ret []GolangStats, err error) {
	ret = make([]GolangStats, 0)
	this.InitSession(context.Background())

	this.BuildQuery(mId, "id")

	err = this.Session.Limit(rn, pn).Find(&ret)
	return
}

func TestSession(t *testing.T) {
	session := GetDbInstance("default", "writer").GetSession()

	//insert
	gs := &GolangStats{
		Metric:     "dbdao",
		Endpoint:   "wuguofu-test",
		Tags:       "dbdao-test",
		Ctime:      55555,
		Value:      1,
		Ctype:      1,
		CreateTime: time.Now(),
	}
	cnt, err := session.Insert(gs)
	if err != nil {
		t.Fatal("insert failed")
	}
	t.Log("insert success :", cnt)

	//select
	gss := make([]GolangStats, 0, 5)
	err = session.SQL("select * from  golang_stats").Limit(5).Desc("create_time").Find(&gss)
	if err != nil {
		t.Fatal("select failed,", err)
	}
	t.Log("select success:", len(gss))

	//update
	gs.Id = 52
	gs.Value = 52
	gs.Tags = "dbdao-update"
	cnt, err = session.Id(52).Update(gs)
	if err != nil {
		t.Fatal("update failed,", err)
	}
	t.Log("update success:", cnt)

	//delete
	cnt, err = session.ID(53).Delete(GolangStats{})
	if err != nil {
		t.Fatal("delete failed,", err)
	}
	t.Log("delete success:", cnt)
}

func TestEngine(t *testing.T) {
	engine := GetDbInstance("default", "writer").Engine

	//insert
	gs := &GolangStats{
		Metric:     "dbdao",
		Endpoint:   "wuguofu-test-engine",
		Tags:       "dbdao-test-engine",
		Ctime:      666666,
		Value:      1,
		Ctype:      1,
		CreateTime: time.Now(),
	}
	cnt, err := engine.Insert(gs)
	if err != nil {
		t.Fatal("insert failed")
	}
	t.Log("insert success :", cnt)

	//select
	gss := make([]GolangStats, 0, 5)
	err = engine.SQL("select * from  golang_stats").Limit(5).Desc("create_time").Find(&gss)
	if err != nil {
		t.Fatal("select failed,", err)
	}
	t.Log("select success:", len(gss))

	//update
	gs.Id = 64
	gs.Value = 64
	gs.Tags = "dbdao-update-64"
	cnt, err = engine.ID(64).Update(gs)
	if err != nil {
		t.Fatal("update failed,", err)
	}
	t.Log("update success:", cnt)

	//delete
	cnt, err = engine.ID(63).Delete(GolangStats{})
	if err != nil {
		t.Fatal("delete failed,", err)
	}
	t.Log("delete success:", cnt)
}

func TestDbdao(t *testing.T) {
	dao := NewGolangStatsDao()

	//insert
	gs := &GolangStats{
		Metric:     "dbdao",
		Endpoint:   "wuguofu-test-dbdao",
		Tags:       "dbdao-test-dbdao",
		Ctime:      88888888,
		Value:      1,
		Ctype:      1,
		CreateTime: time.Now(),
	}
	cnt, err := dao.Create(gs)
	if err != nil {
		t.Fatal("insert failed")
	}
	t.Log("insert success :", cnt)

	//select
	mId := CastToParamIn([]int{51, 64, 58})
	gss, err := dao.GetLimit(mId, 0, 5)
	if err != nil {
		t.Fatal("get failed")
	}
	t.Log("select success:", len(gss))

	//update
	gs.Id = 60
	gs.Value = 60
	gs.Tags = "dbdao-update-60"
	cnt, err = dao.Update(gs)
	if err != nil {
		t.Fatal("update failed,", err)
	}
	t.Log("update success:", cnt)

	//delete
	item := &GolangStats{Id: 61}
	cnt, err = dao.Delete(item)
	if err != nil {
		t.Fatal("delete failed,", err)
	}
	t.Log("delete success:", cnt)
}

func TestAddMysqlWithoutConf(t *testing.T) {
	options := make([]MysqlOption, 1)
	writerOption := MysqlOption{}
	writerOption.Cluster = "default.writer"
	writerOption.Hosts = append(writerOption.Hosts, "root:123456@tcp(127.0.0.1:3306)/test")
	options = append(options, writerOption)
	AddMysql(options)
	dao := NewGolangStatsDao()
	//insert
	gs := &GolangStats{
		Metric:     "dbdao",
		Endpoint:   "wuguofu-test-dbdao",
		Tags:       "dbdao-test-dbdao",
		Ctime:      88888888,
		Value:      1,
		Ctype:      1,
		CreateTime: time.Now(),
	}
	cnt, err := dao.Create(gs)
	if err != nil {
		t.Fatal("insert failed")
	}
	t.Log(cnt)
}
