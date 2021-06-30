package dbio

import (
	"fmt"

	tcm "github.com/gurbos/tcmodels"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
}

var DBConfig gorm.Config
var DataSource *DataSourceName

// DataSourceName attributes hold informaition about a specific database.
// The information is used to connect to said database.
type DataSourceName struct {
	connStr string
}

func (dsn *DataSourceName) Init(host string, port string, user string, passwd string, name string) {
	format := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn.connStr = fmt.Sprintf(format, user, passwd, host, port, name)
}

// DSNString returns a database connection string
func (dsn *DataSourceName) ConnString() string {
	return dsn.connStr
}

func DBConnection() *gorm.DB {
	conn, _ := gorm.Open(mysql.Open(DataSource.ConnString()), &DBConfig)
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

	qm := make(map[string]interface{})
	dbconn := DBConnection() // Get database connection handle
	if dbconn.Error == nil {
		switch {
		case len(productLineIDs) > 0:
			qm["product_line_id"] = &productLineIDs
		case len(setNames) > 0:
			qm["url_name"] = &setNames
		}
		qErr = dbconn.Model(tcm.SetInfo{}).Where(qm).Find(&setInfos).Error
	} else {
		qErr = dbconn.Error
	}

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
