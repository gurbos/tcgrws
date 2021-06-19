package dbio

import (
	"fmt"
	"strings"

	tcm "github.com/gurbos/tcmodels"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataSource DataSourceName
var DBConfig gorm.Config

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

// queryAllProductLines returns all product line info on the database
func QueryAllProductLines() ([]tcm.ProductLine, error) {
	var productLines []tcm.ProductLine // Database query results

	dbconn := DBConnection() // Get database connection handle
	if dbconn.Error != nil {
		return productLines, dbconn.Error
	}
	tx := dbconn.Model(tcm.ProductLine{}).Find(&productLines) // Query database for all product line records
	return productLines, tx.Error
}

func QueryProductLines(plName []string) ([]tcm.ProductLine, error) {
	var productLine []tcm.ProductLine // ProductLine query results
	var dbErr error
	dbconn := DBConnection() // Get database connection handle
	if dbconn.Error == nil {
		plStr := strings.Join(plName, ",") // String of product line names

		// Query database for product line info
		dbErr = dbconn.Raw("SELECT * FROM product_lines WHERE url_name IN (?)", plStr).
			Find(&productLine).Error
	} else {
		dbErr = dbconn.Error
	}
	return productLine, dbErr
}

// queryAllSets returns all card set info in the database
func QueryAllSets() ([]tcm.SetInfo, error) {
	var sets []tcm.SetInfo

	dbconn := DBConnection()
	if dbconn.Error != nil {
		return sets, dbconn.Error
	}

	tx := dbconn.Model(tcm.SetInfo{}).Find(&sets)
	return sets, tx.Error
}

/** NOTE: Work on **/
func QuerySets(productLineIDs []int64, setNames []string) ([]tcm.SetInfo, error) {
	var setInfos []tcm.SetInfo // Database query results
	var qErr error             // Database query error

	dbconn := DBConnection() // Get database connection handle
	if dbconn.Error == nil {
		if len(setNames) == 0 { // If set name list is empty, then query all sets
			qErr = dbconn.Model(tcm.SetInfo{}).
				Where("product_line_id IN = ?", productLineIDs).
				Find(&setInfos).Error
		} else { // If set name list not empty, then query specified sets only
			qErr = dbconn.Model(tcm.SetInfo{}).
				Where("product_line_id IN ?", productLineIDs).
				Where("url_name IN ?", setNames).
				Find(&setInfos).Error
		}
	} else {
		qErr = dbconn.Error
	}

	return setInfos, qErr
}

func QueryCards(plIDList []int64, setIDList []int64, offset int64, size int64) ([]tcm.YuGiOhCardInfo, error) {
	var dbErr error
	var cards []tcm.YuGiOhCardInfo

	dbconn := DBConnection()
	if dbconn.Error == nil {
		dbErr = dbconn.Model(&tcm.YuGiOhCardInfo{}).
			Preload("ProductLine").
			Preload("SetInfo").
			Where("product_line_id IN ?", plIDList).
			Where("set_id IN ?", setIDList).
			Find(&cards).Error
	} else {
		dbErr = dbconn.Error
	}
	return cards, dbErr
}
