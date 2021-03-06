package dbio

import (
	"fmt"
	"time"

	tcm "github.com/gurbos/tcmodels"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Configure(host string, port string, user string,
	pass string, name string, maxOpenConns int, maxIdleConns int,
	maxConnLifetime time.Duration, maxConnIdleTime time.Duration) {

	DataSource = new(DataSourceName)
	DataSource.Init(host, port, user, pass, name)
	dbConn = DBConnection()

	// Database connection pool configuration
	sqlConn, err := dbConn.DB()
	if err != nil {
		panic("gorm.DB(): " + err.Error())
	}
	sqlConn.SetMaxOpenConns(maxOpenConns)
	sqlConn.SetMaxIdleConns(maxIdleConns)
	sqlConn.SetConnMaxLifetime(maxConnLifetime)
	sqlConn.SetConnMaxIdleTime(maxConnIdleTime)
}

// DataSourceName attributes hold informaition about a specific database.
// The information is used to connect to said database.
type DataSourceName struct {
	connStr string
}

// Init
func (dsn *DataSourceName) Init(
	host string, port string, user string, pass string, name string) {
	format := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn.connStr = fmt.Sprintf(format, user, pass, host, port, name)
}

// DSNString returns a database connection string
func (dsn *DataSourceName) ConnString() string {
	return dsn.connStr
}

var DataSource *DataSourceName
var gormConfig gorm.Config
var dbConn *gorm.DB // Connection pool handle

func DBConnection() *gorm.DB {
	conn, err := gorm.Open(mysql.Open(DataSource.ConnString()), &gorm.Config{})
	if err != nil {
		panic("gorm.Open(): " + err.Error())
	}
	return conn
}

func QueryProductLines(plNames []string) ([]tcm.ProductLine, error) {
	var productLine []tcm.ProductLine // ProductLine query results
	var dbErr error

	dbconn := DBConnection() // Get database connection handle
	if dbconn.Error == nil {
		switch {
		case len(plNames) > 0:
			dbErr = dbconn.Model(tcm.ProductLine{}).Where("name IN (?)", plNames).Find(&productLine).Error
		default:
			dbErr = dbconn.Model(tcm.ProductLine{}).Find(&productLine).Error
		}
	} else {
		dbErr = dbconn.Error
	}
	return productLine, dbErr
}

/** NOTE: Work on **/
func QuerySets(productLineIDs []int64, setNames []string) ([]tcm.SetInfo, error) {
	var setInfos []tcm.SetInfo // Database query results
	var qErr error             // Database query error

	qm := make(map[string]interface{}) // Used to filter records in sql WHERE clause, e.g. WHERE key IN value

	if len(productLineIDs) > 0 {
		qm["product_line_id"] = &productLineIDs
	}
	if len(setNames) > 0 {
		qm["url_name"] = &setNames
	}
	qErr = dbConn.Model(tcm.SetInfo{}).Where(qm).Find(&setInfos).Error

	return setInfos, qErr
}

func QueryCards(plIDList []int64, setIDList []int64, offset int64, length int64) ([]tcm.YuGiOhCardInfo, error) {
	var dbErr error
	var cards []tcm.YuGiOhCardInfo

	dbconn := DBConnection()
	if dbconn.Error == nil {
		dbErr = dbconn.Model(&tcm.YuGiOhCardInfo{}).
			Preload("ProductLine").
			Preload("SetInfo").
			Where("product_line_id IN ?", plIDList).
			Where("set_id IN ?", setIDList).
			Offset(int(offset)).
			Limit(int(length)).
			Find(&cards).Error
	} else {
		dbErr = dbconn.Error
	}
	return cards, dbErr
}
