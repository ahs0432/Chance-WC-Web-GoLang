package controllers

import (
	"database/sql"
	"fmt"
)

// Table 존재 유무 확인 함수
func dbCheckTable(db *sql.DB, table string) bool {
	_, table_check := db.Query("SELECT * FROM " + table)

	if table_check != nil {
		//logAddLine(infoLogFile, "Database "+table+" not found")
		fmt.Println(table_check)
	}

	return table_check == nil
}

// Web 테이블 데이터 검색 함수 (WHERE 절의 경우 별도 입력)
func dbQueryWebSELECT(db *sql.DB, where string) []WebData {
	queryRow, err := db.Query("SELECT * FROM WEB " + where)
	if err != nil {
		//errCheck(err, "Database Web Select Query")
		fmt.Println(err)
	}

	defer queryRow.Close()

	var webDataList []WebData

	for queryRow.Next() {
		var webData WebData
		err := queryRow.Scan(&webData.idx, &webData.name, &webData.url, &webData.chkcon, &webData.rcmdtrs, &webData.mail, &webData.lastresult, &webData.laststatus, &webData.lastcheck, &webData.lasttime, &webData.sslexpire, &webData.uptimeper, &webData.tlscheck, &webData.statcheck, &webData.alarm, &webData.timeout, &webData.useridx)
		if err != nil {
			webDataList = nil
			//errCheck(err, "Database Web Select Query Scan")
			fmt.Println(err)
		} else {
			if webDataList != nil {
				webDataList = append(webDataList, webData)
			} else {
				webDataList = []WebData{webData}
			}
		}
	}

	return webDataList
}

// WEBUPTIME 테이블 데이터 검색 함수 (WHERE 절의 경우 별도 입력)
func dbQueryUptimeSELECT(db *sql.DB, where string) []WebUptimeData {
	queryRow, err := db.Query("SELECT * FROM WEBUPTIME " + where)
	if err != nil {
		//errCheck(err, "Database Web Select Query")
		fmt.Println(err)
	}

	defer queryRow.Close()

	var webUptimeDataList []WebUptimeData

	for queryRow.Next() {
		var webUptimeData WebUptimeData
		err := queryRow.Scan(&webUptimeData.idx, &webUptimeData.uptimeper, &webUptimeData.checkday)
		if err != nil {
			webUptimeDataList = nil
			//errCheck(err, "Database Web Select Query Scan")
			fmt.Println(err)
		} else {
			if webUptimeDataList != nil {
				webUptimeDataList = append(webUptimeDataList, webUptimeData)
			} else {
				webUptimeDataList = []WebUptimeData{webUptimeData}
			}
		}
	}

	return webUptimeDataList
}

// WEBGROUP 테이블 데이터 검색 함수 (WHERE 절의 경우 별도 입력)
func dbQueryGroupSELECT(db *sql.DB, where string) []WebGroupData {
	queryRow, err := db.Query("SELECT * FROM WEBGROUP " + where)
	if err != nil {
		//errCheck(err, "Database Web Select Query")
		fmt.Println(err)
	}

	defer queryRow.Close()

	var webGroupDataList []WebGroupData

	for queryRow.Next() {
		var webGroupData WebGroupData
		err := queryRow.Scan(&webGroupData.Name, &webGroupData.Member, &webGroupData.Count)
		if err != nil {
			webGroupDataList = nil
			//errCheck(err, "Database Web Select Query Scan")
			fmt.Println(err)
		} else {
			if webGroupDataList != nil {
				webGroupDataList = append(webGroupDataList, webGroupData)
			} else {
				webGroupDataList = []WebGroupData{webGroupData}
			}
		}
	}

	return webGroupDataList
}

// 데이터 추가/변경 함수
func dbQueryIU(db *sql.DB, query string) {
	queryRow, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	} else {
		_, err = queryRow.RowsAffected()

		if err != nil {
			fmt.Println(err)
		}
	}
}
