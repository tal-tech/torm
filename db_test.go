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
	"fmt"
	"testing"
	"time"
)

type GolangStats struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Metric     string    `json:"metric" xorm:"not null default '' comment('指标名') VARCHAR(100)"`
	EndPoint   string    `json:"end_point" xorm:"not null default '' comment('机器名') VARCHAR(100)"`
	Tags       string    `json:"tags" xorm:"not null default '' comment('服务名') VARCHAR(100)"`
	Ctime      int       `json:"ctime" xorm:"not null default 0 comment('时间戳') INT(11)"`
	Value      int       `json:"value" xorm:"not null default 0 comment('值') INT(11)"`
	Ctype      int       `json:"ctype" xorm:"not null default 1 comment('统计类型,1:计数 2:耗时') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"not null default '0001-01-01 00:00:00' comment('提交时间') DATETIME"`
}

type GolangStatsDao struct {
	DbBaseDao
	ctx context.Context
}

func NewGolangStatsDao(ctx context.Context, v ...interface{}) *GolangStatsDao {
	this := new(GolangStatsDao)
	this.ctx = ctx
	if ins := GetDbInstance("test", "writer"); ins != nil {
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
	this.InitSession(this.ctx)

	this.BuildQuery(mId, "id")

	err = this.Session.Find(&ret)
	return
}
func (this *GolangStatsDao) GetLimit(mId Param, pn, rn int) (ret []GolangStats, err error) {
	ret = make([]GolangStats, 0)
	this.InitSession(this.ctx)

	this.BuildQuery(mId, "id")

	err = this.Session.Limit(rn, pn).Find(&ret)
	return
}
func (this *GolangStatsDao) GetCount(mId Param) (ret int64, err error) {
	this.InitSession(this.ctx)

	this.BuildQuery(mId, "id")

	ret, err = this.Session.Count(new(GolangStats))
	return
}

func TestSession(t *testing.T) {
	session := GetDbInstance("test", "writer").GetSession()

	//insert
	gs := &GolangStats{
		Metric:     "dbdao",
		EndPoint:   "wuguofu-test",
		Tags:       "dbdao-test",
		Ctime:      55555,
		Value:      1,
		Ctype:      1,
		CreateTime: time.Now(),
	}
	cnt, err := session.Insert(gs)
	if err != nil {
		fmt.Println(err)
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
	cnt, err = session.Id(3).Update(gs)
	if err != nil {
		t.Fatal("update failed,", err)
	}
	t.Log("update success:", cnt)

	//delete
	cnt, err = session.ID(4).Delete(GolangStats{})
	if err != nil {
		t.Fatal("delete failed,", err)
	}
	t.Log("delete success:", cnt)
}

func TestEngine(t *testing.T) {
	engine := GetDbInstance("test", "writer").Engine

	//insert
	gs := &GolangStats{
		Metric:     "dbdao",
		EndPoint:   "wuguofu-test-engine",
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
	cnt, err = engine.ID(5).Update(gs)
	if err != nil {
		t.Fatal("update failed,", err)
	}
	t.Log("update success:", cnt)

	//delete
	cnt, err = engine.ID(3).Delete(GolangStats{})
	if err != nil {
		t.Fatal("delete failed,", err)
	}
	t.Log("delete success:", cnt)
}

func TestDbdao(t *testing.T) {
	dao := NewGolangStatsDao(context.Background())

	//insert
	gs := &GolangStats{
		Metric:     "dbdao",
		EndPoint:   "wuguofu-test-dbdao",
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
	mId := CastToParamIn([]int{5, 6, 7})
	gss, err := dao.GetLimit(mId, 0, 5)
	if err != nil {
		t.Fatal("get failed")
	}
	t.Log("select success:", len(gss))

	//update
	gs.Id = 6
	gs.Value = 6
	gs.Tags = "dbdao-update-6"
	cnt, err = dao.Update(gs)
	if err != nil {
		t.Fatal("update failed,", err)
	}
	t.Log("update success:", cnt)

	//delete
	item := &GolangStats{Id: 7}
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
	dao := NewGolangStatsDao(context.Background())
	//insert
	gs := &GolangStats{
		Metric:     "dbdao",
		EndPoint:   "wuguofu-test-dbdao",
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
