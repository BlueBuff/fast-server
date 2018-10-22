package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
	"fmt"
	"time"
)

type DB_RESOURCE string

var DataResource Resource

type Resource interface {
	GetDB(name DB_RESOURCE) (*gorm.DB, bool)
	PutDB(name DB_RESOURCE, db *gorm.DB) bool
	Size() int
}

type DBResource struct {
	pool map[DB_RESOURCE]*gorm.DB
	once sync.Once
}

func NewDBResource() Resource {
	resource := new(DBResource)
	resource.pool = make(map[DB_RESOURCE]*gorm.DB)
	return resource
}

func (resource *DBResource) GetDB(name DB_RESOURCE) (*gorm.DB, bool) {
	if db, ok := resource.pool[name]; !ok {
		return nil, false
	} else {
		return db, true
	}
}

func (resource *DBResource) PutDB(name DB_RESOURCE, db *gorm.DB) bool {
	if db == nil {
		return false
	}
	if _, ok := resource.pool[name]; ok {
		return false
	} else {
		resource.pool[name] = db
		return true
	}
}

func (resource *DBResource) Size() int {
	return len(resource.pool)
}

func init() {
	resource := NewDBResource()
	if ConfigurationContext.DBConfigs != nil && len(ConfigurationContext.DBConfigs) != 0 {
		logger.Infof("init dbResource ... ")
		for _, dbConfig := range ConfigurationContext.DBConfigs {
			dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", dbConfig.Config.UserName, dbConfig.Config.Password, dbConfig.Config.Host, dbConfig.Config.Port, dbConfig.Config.DataBaseName)
			logger.Infof("connect... dbURL:%s", dbURL)
			db, err := gorm.Open(dbConfig.Config.Driver, dbURL)
			if err != nil {
				panic(err)
			}
			db.LogMode(dbConfig.Config.Mode)
			db.DB().SetConnMaxLifetime(time.Minute * dbConfig.Config.ConnMaxLifetime)
			db.DB().SetMaxOpenConns(dbConfig.Config.MaxOpenNum)
			db.DB().SetMaxIdleConns(dbConfig.Config.MaxIdleNum)
			ok := resource.PutDB(DB_RESOURCE(dbConfig.Name), db)
			if !ok {
				logger.Fatalf("the db add failed ...dbName:", dbConfig.Name)
			}
		}
		logger.Infof("dbResource Size:%d", resource.Size())
		if resource.Size() == 0 {
			logger.Errorf("检测到数据源为空,请检查数据源配置是否正确,如正确,可忽略,模板%s", defaultDBConfigTemplate)
		}
	}
	DataResource = resource
}
