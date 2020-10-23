package torm

import (
	"context"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
	"github.com/tal-tech/xtools/confutil"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

type DBDao struct {
	Engine *Engine
	quiter chan struct{}
}

var (
	db_instance  map[string][]*DBDao //db instance collection
	curDbPoses   map[string]*uint64  //Currently selected database
	showSql      bool                //show sql switch
	showExecTime bool                //show sql executed time switch.
	logLevel     = xlog.LOG_INFO
	slowDuration time.Duration       //Slow query time config
	maxConn      int           = 100 //Maximum number of connections
	maxIdle      int           = 30  //Maximum number of idle connections
	lock         sync.RWMutex
)

//newDBDaoWithParams returns a dbdao instance specified by host.
func newDBDaoWithParams(host string, driver string) (Db *DBDao) {
	Db = new(DBDao)
	//Initialize a dbdao engine
	eng, err := xorm.NewEngine(driver, host)
	engine := &Engine{Engine: eng}

	Db.Engine = engine
	//TODO: 增加存活检查
	if err != nil {
		log.Fatal(err)
	}

	Db.Engine.SetMaxOpenConns(maxConn)
	Db.Engine.SetMaxIdleConns(maxIdle)
	//最大超时时间
	Db.Engine.SetConnMaxLifetime(time.Second * 3000)
	Db.Engine.ShowSQL(showSql)
	//set logger Plug-in
	logger := TormLogger{}
	logger.SetLevel(logLevel)
	logger.ShowSQL(true)
	Db.Engine.SetLogger(NewTormLoggerAdapter(logger))
	return
}

func GetDefault(cluster string) *DBDao {
	return GetDbInstance("default", cluster)
}

//Initialize mysql configuration...
func init() {
	db_instance = make(map[string][]*DBDao, 0)
	curDbPoses = make(map[string]*uint64)
	//idc := confdao.GetIDC()
	idc := ""
	showLog := confutil.GetConfStringMap("MysqlConfig")
	showSql = showLog["showSql"] == "true"
	if showLog["level"] != "" {
		logLevel = xlog.LogLevel(cast.ToInt(showLog["level"]))
	}
	showExecTime = showLog["showExecTime"] == "true"
	slowDuration = time.Duration(cast.ToInt(showLog["slowDuration"])) * time.Millisecond
	maxConnConfig := cast.ToInt(showLog["maxConn"])
	if maxConnConfig > 0 {
		maxConn = maxConnConfig
	}
	maxIdleConfig := cast.ToInt(showLog["maxIdle"])
	if maxIdleConfig > 0 {
		maxIdle = maxIdleConfig
	}
	if maxIdle > maxConn {
		maxIdle = maxConn
	}
	for cluster, hosts := range confutil.GetConfArrayMap("MysqlCluster") {
		items := strings.Split(cluster, ".")
		//必须包含 writer 和 reader
		if len(items) < 2 {
			continue
		}
		//过滤IDC
		if len(items) > 2 && items[2] != idc {
			continue
		}
		//such as default.writer or default.reader
		instance := items[0] + "." + items[1]
		dbs := make([]*DBDao, 0)
		for _, host := range hosts {
			dbs = append(dbs, newDBDaoWithParams(host, "mysql"))
		}
		db_instance[instance] = dbs
		curDbPoses[instance] = new(uint64)
	}
}

type MysqlOption struct {
	Cluster string
	Hosts   []string
}

//add mysql without config file.
func AddMysql(options []MysqlOption) {
	for _, option := range options {
		cluster := option.Cluster
		hosts := option.Hosts
		items := strings.Split(cluster, ".")
		//必须包含 writer 和 reader
		if len(items) < 2 {
			continue
		}
		//such as default.writer or default.reader
		instance := items[0] + "." + items[1]
		dbs := make([]*DBDao, 0)
		for _, host := range hosts {
			dbs = append(dbs, newDBDaoWithParams(host, "mysql"))
		}
		lock.Lock()
		db_instance[instance] = dbs
		curDbPoses[instance] = new(uint64)
		lock.Unlock()
	}
}

//GetDbInstance returns DBDao instance.
func GetDbInstance(db, cluster string) *DBDao {
	key := db + "." + cluster
	lock.RLock()
	defer lock.RUnlock()
	if instances, ok := db_instance[key]; ok {
		// round-robin选择数据库
		cur := atomic.AddUint64(curDbPoses[key], 1) % uint64(len(instances))
		return instances[cur]
	} else {
		return nil
	}
}

//GetDbInstanceWithCtx returns DBDao instance.
//If it's a stress test scenario,return benchmark_dbName.
func GetDbInstanceWithCtx(ctx context.Context, db, cluster string) *DBDao {
	bench := ctx.Value("IS_BENCHMARK")
	if cast.ToString(bench) == "1" {
		db = "benchmark_" + db
	}
	key := db + "." + cluster
	lock.RLock()
	defer lock.RUnlock()
	if instances, ok := db_instance[key]; ok {
		// round-robin选择数据库
		cur := atomic.AddUint64(curDbPoses[key], 1) % uint64(len(instances))
		return instances[cur]
	} else {
		return nil
	}
}

//GetSession returns a new Session
func (this *DBDao) GetSession() *Session {
	return this.Engine.NewSession()
}

//Close the engine
func (this *DBDao) Close() {
	this.Engine.Close()
}
