package controllers

import (
	"fmt"
	"math"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"false.kr/Monitor-Web/database"
	"github.com/gofiber/fiber"
)

func Index(c *fiber.Ctx) error {
	db := database.DBConn

	var webDataList []WebData
	type WebDataIndex struct {
		Idx       int
		Name      string
		Url       string
		UrlShort  string
		Lastcheck string
		Uptimeper float64
	}
	var webDataIndexList []WebDataIndex

	if dbCheckTable(db, "WEB") {
		webDataList = dbQueryWebSELECT(db, "")

		for _, webData := range webDataList {
			var webDataIndex WebDataIndex
			webDataIndex.Idx = webData.idx
			webDataIndex.Name = webData.name
			webDataIndex.Url = webData.url

			u, err := url.Parse(webData.url)
			if err == nil {
				if host, port, err := net.SplitHostPort(u.Host); err == nil {
					webDataIndex.UrlShort = host + ":" + port
				} else {
					webDataIndex.UrlShort = u.Host
				}
			} else {
				webDataIndex.UrlShort = webData.url
			}

			webDataIndex.Uptimeper = webData.uptimeper

			if webData.alarm != 0 {
				if webData.lastcheck {
					webDataIndex.Lastcheck = "wnorm"
				} else {
					webDataIndex.Lastcheck = "unnorm"
				}
			} else {
				webDataIndex.Lastcheck = "disable"
			}

			if webDataIndexList == nil {
				webDataIndexList = []WebDataIndex{webDataIndex}
			} else {
				webDataIndexList = append(webDataIndexList, webDataIndex)
			}
		}
	}

	data := fiber.Map{"WebDataList": webDataIndexList}
	return c.Render("index", data)
}

func Additional(c *fiber.Ctx) error {
	return c.Render("additional", "")
}

func AdditionalData(c *fiber.Ctx) error {
	name := c.FormValue("name")
	url := c.FormValue("url")
	chkcon := c.FormValue("chkcon")
	rcmdtrs := c.FormValue("recommend")
	mailList := c.FormValue("mail[]")
	tlscheck := c.FormValue("tlscheck")
	statcheck := c.FormValue("statcheck")
	timeout := c.FormValue("timeout")
	alarm := c.FormValue("alarm")

	db := database.DBConn
	dbQueryIU(db, "INSERT INTO WEB (NAME, URL, CHKCON, RCMDTRS, MAIL, TLSCHECK, STATCHECK, TIMEOUT, ALARM, USERIDX) VALUES('"+name+"', '"+url+"', '"+chkcon+"', '"+rcmdtrs+"', '"+strings.ReplaceAll(string(mailList), ",", " ")+"', "+tlscheck+", "+statcheck+", "+timeout+", "+alarm+", 0)")
	return c.Redirect("/index")
}

func Checker(c *fiber.Ctx) error {
	if _, err := strconv.Atoi(c.Params("idx")); err != nil {
		return c.Redirect("/index")
	}

	db := database.DBConn

	type WebDataChecker struct {
		Idx           int
		Name          string
		Url           string
		UrlShort      string
		Chkcon        string
		Rcmdtrs       string
		Lastresult    bool
		Laststatus    int
		Lastcheck     string
		Lasttime      string
		SSLexpiredate string
		SSLexpireday  string
		Uptimeper     float64
		Uptime7       float64
		Uptime30      float64
		Uptime90      float64
		Alarm         string
		Timeout       int
	}
	var webDataCheckerList []WebDataChecker

	if dbCheckTable(db, "WEB") {
		webDataList := dbQueryWebSELECT(db, "WHERE IDX="+c.Params("idx"))

		if webDataList == nil {
			return c.Redirect("/index")
		}

		for _, webData := range webDataList {
			var webDataChecker WebDataChecker
			webDataChecker.Idx = webData.idx
			webDataChecker.Name = webData.name
			webDataChecker.Url = webData.url
			webDataChecker.Chkcon = webData.chkcon
			webDataChecker.Rcmdtrs = webData.rcmdtrs
			webDataChecker.Lastresult = webData.lastresult
			webDataChecker.Laststatus = webData.laststatus
			webDataChecker.Lasttime = webData.lasttime
			webDataChecker.Timeout = webData.timeout

			u, err := url.Parse(webData.url)
			if err == nil {
				if host, port, err := net.SplitHostPort(u.Host); err == nil {
					webDataChecker.UrlShort = host + ":" + port
				} else {
					webDataChecker.UrlShort = u.Host
				}
			} else {
				webDataChecker.UrlShort = webData.url
			}

			webDataChecker.Uptimeper = webData.uptimeper

			if webData.alarm != 0 {
				if webData.lastcheck {
					webDataChecker.Lastcheck = "wnorm"
				} else {
					webDataChecker.Lastcheck = "unnorm"
				}

				if webData.alarm == 1 {
					webDataChecker.Alarm = "지속 알람 발생"
				} else if webData.alarm == 2 {
					webDataChecker.Alarm = "지속 알람 발생 (해소 알람 포함)"
				} else if webData.alarm == 3 {
					webDataChecker.Alarm = "일회성 알람 발생"
				} else {
					webDataChecker.Alarm = "알람 끄기"
				}
			} else {
				webDataChecker.Lastcheck = "disable"
				webDataChecker.Alarm = "검사 대상 제외"
			}

			if webData.sslexpire != nil {
				expireDate, err := time.Parse("2006-01-02 15:04:05", *webData.sslexpire)
				if err == nil {
					webDataChecker.SSLexpireday = "D-" + strconv.Itoa(int(time.Until(expireDate).Hours()/24))
					webDataChecker.SSLexpiredate = expireDate.Format("2006-01-02(Mon) 15:04:05")
				} else {
					webDataChecker.SSLexpiredate = "None"
				}
			} else {
				webDataChecker.SSLexpiredate = "None"
			}

			if webDataCheckerList == nil {
				webDataCheckerList = []WebDataChecker{webDataChecker}
			} else {
				webDataCheckerList = append(webDataCheckerList, webDataChecker)
			}
		}
	} else {
		return c.Redirect("/index")
	}

	if dbCheckTable(db, "WEBUPTIME") {
		webUptimeDataList := dbQueryUptimeSELECT(db, "WHERE URLIDX="+c.Params("idx"))
		if webUptimeDataList == nil {
			webDataCheckerList[0].Uptime7 = webDataCheckerList[0].Uptimeper
			webDataCheckerList[0].Uptime30 = webDataCheckerList[0].Uptimeper
			webDataCheckerList[0].Uptime90 = webDataCheckerList[0].Uptimeper
		} else {
			var uptime7, uptime30, uptime90 float64

			for _, webUptimeData := range webUptimeDataList {
				checkDate, err := time.Parse("2006-01-02", webUptimeData.checkday)
				if err == nil {
					nowDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
					dayDiff := nowDate.Sub(checkDate).Hours() / 24

					switch {
					case dayDiff <= 7:
						uptime7 += webUptimeData.uptimeper
						fallthrough
					case dayDiff <= 30:
						uptime30 += webUptimeData.uptimeper
						fallthrough
					case dayDiff <= 90:
						uptime90 += webUptimeData.uptimeper
					}
				} else {
					continue
				}
			}

			if len(webUptimeDataList) < 7 {
				webDataCheckerList[0].Uptime7 = math.Floor((uptime7 / float64(len(webUptimeDataList)) * 100)) / 100
				webDataCheckerList[0].Uptime30 = webDataCheckerList[0].Uptime7
				webDataCheckerList[0].Uptime90 = webDataCheckerList[0].Uptime7
			} else if len(webUptimeDataList) < 30 {
				webDataCheckerList[0].Uptime7 = math.Floor((uptime7/7.0)*100) / 100
				webDataCheckerList[0].Uptime30 = math.Floor((uptime30 / float64(len(webUptimeDataList)) * 100)) / 100
				webDataCheckerList[0].Uptime90 = webDataCheckerList[0].Uptime30
			} else if len(webUptimeDataList) < 90 {
				webDataCheckerList[0].Uptime7 = math.Floor((uptime7/7.0)*100) / 100
				webDataCheckerList[0].Uptime30 = math.Floor((uptime30/30.0)*100) / 100
				webDataCheckerList[0].Uptime90 = math.Floor((uptime90 / float64(len(webUptimeDataList)) * 100)) / 100
			} else {
				webDataCheckerList[0].Uptime7 = math.Floor((uptime7/7.0)*100) / 100
				webDataCheckerList[0].Uptime30 = math.Floor((uptime30/30.0)*100) / 100
				webDataCheckerList[0].Uptime90 = math.Floor((uptime90/90.0)*100) / 100
			}

			if math.IsNaN(webDataCheckerList[0].Uptime7) {
				webDataCheckerList[0].Uptime7 = 0
			}
			if math.IsNaN(webDataCheckerList[0].Uptime30) {
				webDataCheckerList[0].Uptime30 = 0
			}
			if math.IsNaN(webDataCheckerList[0].Uptime90) {
				webDataCheckerList[0].Uptime90 = 0
			}
		}
	} else {
		webDataCheckerList[0].Uptime7 = webDataCheckerList[0].Uptimeper
		webDataCheckerList[0].Uptime30 = webDataCheckerList[0].Uptimeper
		webDataCheckerList[0].Uptime90 = webDataCheckerList[0].Uptimeper
	}

	data := fiber.Map{"WebDataList": webDataCheckerList}
	return c.Render("checker", data)
}

func Modify(c *fiber.Ctx) error {
	if _, err := strconv.Atoi(c.Params("idx")); err != nil {
		return c.Redirect("/index")
	}

	db := database.DBConn

	type WebDataChecker struct {
		Idx       int
		Name      string
		Url       string
		UrlShort  string
		Chkcon    string
		Rcmdtrs   string
		Mail      string
		MailList  []string
		MailCnt   int
		TLScheck  int
		Statcheck int
		Alarm     int
		Timeout   int
	}
	var webDataCheckerList []WebDataChecker

	if dbCheckTable(db, "WEB") {
		webDataList := dbQueryWebSELECT(db, "WHERE IDX="+c.Params("idx"))

		if webDataList == nil {
			return c.Redirect("/index")
		}

		for _, webData := range webDataList {
			var webDataChecker WebDataChecker
			webDataChecker.Idx = webData.idx
			webDataChecker.Name = webData.name
			webDataChecker.Url = webData.url
			webDataChecker.Chkcon = webData.chkcon
			webDataChecker.Rcmdtrs = webData.rcmdtrs
			webDataChecker.Name = webData.name
			webDataChecker.Alarm = webData.alarm
			webDataChecker.Timeout = webData.timeout

			u, err := url.Parse(webData.url)
			if err == nil {
				if host, port, err := net.SplitHostPort(u.Host); err == nil {
					webDataChecker.UrlShort = host + ":" + port
				} else {
					webDataChecker.UrlShort = u.Host
				}
			} else {
				webDataChecker.UrlShort = webData.url
			}

			if webData.tlscheck {
				webDataChecker.TLScheck = 1
			} else {
				webDataChecker.TLScheck = 0
			}

			if webData.statcheck {
				webDataChecker.Statcheck = 1
			} else {
				webDataChecker.Statcheck = 0
			}

			webDataChecker.MailCnt = len(strings.Split(webData.mail, " "))
			if webDataChecker.MailCnt == 1 {
				webDataChecker.Mail = webData.mail
			} else {
				webDataChecker.Mail = strings.Split(webData.mail, " ")[0]
				for num, mailL := range strings.Split(webData.mail, " ") {
					if num != 0 {
						if webDataChecker.MailList == nil {
							webDataChecker.MailList = []string{mailL}
						} else {
							webDataChecker.MailList = append(webDataChecker.MailList, mailL)
						}

					}
				}
			}

			if webDataCheckerList == nil {
				webDataCheckerList = []WebDataChecker{webDataChecker}
			} else {
				webDataCheckerList = append(webDataCheckerList, webDataChecker)
			}
		}
	} else {
		return c.Redirect("/index")
	}

	data := fiber.Map{"WebDataList": webDataCheckerList}
	return c.Render("modify", data)
}

func ModifyData(c *fiber.Ctx) error {
	idx := c.FormValue("idx")

	if _, err := strconv.Atoi(c.FormValue("idx")); err != nil {
		return c.Redirect("/checker/" + idx)
	}

	name := c.FormValue("name")
	url := c.FormValue("url")
	chkcon := c.FormValue("chkcon")
	rcmdtrs := c.FormValue("recommend")
	mailList := c.FormValue("mail[]")
	tlscheck := c.FormValue("tlscheck")
	statcheck := c.FormValue("statcheck")
	timeout := c.FormValue("timeout")
	alarm := c.FormValue("alarm")

	db := database.DBConn
	dbQueryIU(db, "UPDATE WEB SET NAME='"+name+"', URL='"+url+"', CHKCON='"+chkcon+"', RCMDTRS='"+rcmdtrs+"', MAIL='"+strings.ReplaceAll(string(mailList), ",", " ")+"', TLSCHECK="+tlscheck+", STATCHECK="+statcheck+", TIMEOUT="+timeout+", ALARM="+alarm+" WHERE IDX="+idx)

	return c.Redirect("/checker/" + idx)
}

func DeleteCheck(c *fiber.Ctx) error {
	if _, err := strconv.Atoi(c.Params("idx")); err != nil {
		return c.Redirect("/index")
	}

	db := database.DBConn

	if dbCheckTable(db, "WEB") {
		webDataList := dbQueryWebSELECT(db, "WHERE IDX="+c.Params("idx"))

		if webDataList == nil {
			return c.Redirect("/")
		}
		dbQueryIU(db, "DELETE FROM WEB WHERE IDX="+c.Params("idx"))
		dbQueryIU(db, "DELETE FROM WEBUPTIME WHERE URLIDX="+c.Params("idx"))
		dbQueryIU(db, "DELETE FROM CHKRESULT WHERE URLIDX="+c.Params("idx"))

	}

	return c.Redirect("/")
}

func GroupManage(c *fiber.Ctx) error {
	db := database.DBConn

	var webGroupList []WebGroupData

	if dbCheckTable(db, "WEBGROUP") {
		webGroupList = dbQueryGroupSELECT(db, "")
	} else {
		return c.Redirect("/index")
	}

	data := fiber.Map{"WebGroupList": webGroupList}
	return c.Render("groupmanage", data)
}

func GroupAdd(c *fiber.Ctx) error {
	db := database.DBConn

	type WebDataChecker struct {
		Idx      int
		Name     string
		UrlShort string
	}

	var webDataCheckerList []WebDataChecker

	if dbCheckTable(db, "WEB") {
		webDataList := dbQueryWebSELECT(db, "")

		for _, webData := range webDataList {
			var webDataChecker WebDataChecker
			webDataChecker.Idx = webData.idx
			webDataChecker.Name = webData.name

			u, err := url.Parse(webData.url)
			if err == nil {
				if host, port, err := net.SplitHostPort(u.Host); err == nil {
					webDataChecker.UrlShort = host + ":" + port
				} else {
					webDataChecker.UrlShort = u.Host
				}
			} else {
				webDataChecker.UrlShort = webData.url
			}

			if webDataCheckerList == nil {
				webDataCheckerList = []WebDataChecker{webDataChecker}
			} else {
				webDataCheckerList = append(webDataCheckerList, webDataChecker)
			}
		}

	} else {
		return c.Redirect("/index")
	}

	data := fiber.Map{"WebDataList": webDataCheckerList}
	return c.Render("groupadd", data)
}

func GroupAddData(c *fiber.Ctx) error {
	name := c.FormValue("name")
	webList := c.FormValue("web[]")

	db := database.DBConn

	if dbCheckTable(db, "WEBGROUP") {
		webGroupList := dbQueryGroupSELECT(db, "WHERE NAME='"+name+"'")

		if len(webGroupList) != 0 {
			type ReturnData struct {
				Url     string
				Message string
			}

			var returnData ReturnData

			returnData.Url = "groupadd"
			returnData.Message = "이미 존재하는 이름입니다!"

			data := fiber.Map{"ReturnData": returnData}

			return c.Render("returndata", data)
		}
	} else {
		return c.Redirect("/groupmanage")
	}

	var lenWeb string

	if string(webList) == "" {
		lenWeb = "0"
	} else {
		lenWeb = strconv.Itoa(len(strings.Split(string(webList), ",")))
	}

	dbQueryIU(db, "INSERT INTO WEBGROUP VALUES('"+name+"', '"+strings.ReplaceAll(string(webList), ",", " ")+"', "+lenWeb+")")
	return c.Redirect("/groupmanage")
}

func GroupCheck(c *fiber.Ctx) error {
	name := c.Params("name")

	db := database.DBConn
	where := ""

	if dbCheckTable(db, "WEBGROUP") {
		webGroupList := dbQueryGroupSELECT(db, "WHERE NAME='"+name+"'")
		if len(webGroupList) != 1 {
			return c.Redirect("/groupmanage")
		}

		if webGroupList[0].Count == 0 {
			where = "WHERE IDX=0"
		} else {
			where = "WHERE "
			firstC := true
			for _, idx := range strings.Split(webGroupList[0].Member, " ") {
				if firstC {
					where = where + "IDX=" + idx + " "
					firstC = false
				}
				where = where + "OR IDX=" + idx + " "
			}
		}
	}

	type WebDataIndex struct {
		Idx       int
		Name      string
		Url       string
		UrlShort  string
		Lastcheck string
		Uptimeper float64
	}
	var webDataIndexList []WebDataIndex
	var webDataList []WebData

	if dbCheckTable(db, "WEB") {
		webDataList = dbQueryWebSELECT(db, where)

		for _, webData := range webDataList {
			var webDataIndex WebDataIndex
			webDataIndex.Idx = webData.idx
			webDataIndex.Name = webData.name
			webDataIndex.Url = webData.url

			u, err := url.Parse(webData.url)
			if err == nil {
				if host, port, err := net.SplitHostPort(u.Host); err == nil {
					webDataIndex.UrlShort = host + ":" + port
				} else {
					webDataIndex.UrlShort = u.Host
				}
			} else {
				webDataIndex.UrlShort = webData.url
			}

			webDataIndex.Uptimeper = webData.uptimeper

			if webData.alarm != 0 {
				if webData.lastcheck {
					webDataIndex.Lastcheck = "wnorm"
				} else {
					webDataIndex.Lastcheck = "unnorm"
				}
			} else {
				webDataIndex.Lastcheck = "disable"
			}

			if webDataIndexList == nil {
				webDataIndexList = []WebDataIndex{webDataIndex}
			} else {
				webDataIndexList = append(webDataIndexList, webDataIndex)
			}
		}
	}

	data := fiber.Map{"GroupName": name, "WebDataList": webDataIndexList}
	return c.Render("groupcheck", data)
}

func GroupDelete(c *fiber.Ctx) error {
	db := database.DBConn
	name := c.Params("name")

	if dbCheckTable(db, "WEBGROUP") {
		webGroupList := dbQueryGroupSELECT(db, "WHERE NAME='"+name+"'")
		if len(webGroupList) != 1 {
			return c.Redirect("/groupmanage")
		}

		dbQueryIU(db, "DELETE FROM WEBGROUP WHERE NAME='"+name+"'")
	}

	return c.Redirect("/groupmanage")
}

func GroupModify(c *fiber.Ctx) error {
	db := database.DBConn
	name := c.Params("name")
	var checked []int

	if dbCheckTable(db, "WEBGROUP") {
		webGroupList := dbQueryGroupSELECT(db, "WHERE NAME='"+name+"'")
		if len(webGroupList) != 1 {
			return c.Redirect("/groupmanage")
		}

		if webGroupList[0].Count != 0 {
			for _, data := range strings.Split(webGroupList[0].Member, " ") {
				dataint, err := strconv.Atoi(data)
				if err == nil {
					if checked == nil {
						checked = []int{dataint}
					} else {
						checked = append(checked, dataint)
					}
				} else {
					fmt.Println("TEST")
					return c.Redirect("/groupmanage")
				}
			}
		}
	}

	type WebDataChecker struct {
		Idx      int
		Name     string
		UrlShort string
		Checked  string
	}

	var webDataCheckerList []WebDataChecker

	if dbCheckTable(db, "WEB") {
		webDataList := dbQueryWebSELECT(db, "")

		for _, webData := range webDataList {
			var webDataChecker WebDataChecker
			webDataChecker.Idx = webData.idx
			webDataChecker.Name = webData.name
			webDataChecker.Checked = "unchecked"

			u, err := url.Parse(webData.url)
			if err == nil {
				if host, port, err := net.SplitHostPort(u.Host); err == nil {
					webDataChecker.UrlShort = host + ":" + port
				} else {
					webDataChecker.UrlShort = u.Host
				}
			} else {
				webDataChecker.UrlShort = webData.url
			}

			if webDataCheckerList == nil {
				webDataCheckerList = []WebDataChecker{webDataChecker}
			} else {
				webDataCheckerList = append(webDataCheckerList, webDataChecker)
			}
		}
	} else {
		return c.Redirect("/index")
	}

	for _, idx := range checked {
		webDataCheckerList[idx-1].Checked = "checked"
	}

	data := fiber.Map{"GroupName": name, "WebDataList": webDataCheckerList}
	return c.Render("groupmodify", data)
}

func GroupModData(c *fiber.Ctx) error {
	name := c.FormValue("name")
	webList := c.FormValue("web[]")

	db := database.DBConn

	if dbCheckTable(db, "WEBGROUP") {
		webGroupList := dbQueryGroupSELECT(db, "WHERE NAME='"+name+"'")

		if len(webGroupList) != 1 {
			return c.Redirect("/groupmanage")
		}
	} else {
		return c.Redirect("/groupmanage")
	}

	var lenWeb string

	if string(webList) == "" {
		lenWeb = "0"
	} else {
		lenWeb = strconv.Itoa(len(strings.Split(string(webList), ",")))
	}

	dbQueryIU(db, "UPDATE WEBGROUP SET MEMBER='"+strings.ReplaceAll(string(webList), ",", " ")+"', COUNT="+lenWeb+" WHERE NAME='"+name+"'")
	return c.Redirect("/groupcheck/" + name)
}
