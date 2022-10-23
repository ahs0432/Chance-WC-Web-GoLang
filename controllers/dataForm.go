package controllers

type WebData struct {
	idx        int
	name       string
	url        string
	chkcon     string
	rcmdtrs    string
	mail       string
	lastresult bool
	laststatus int
	lastcheck  bool
	lasttime   string
	sslexpire  *string
	uptimeper  float64
	tlscheck   bool
	statcheck  bool
	alarm      int
	timeout    int
	useridx    int
}

type WebUptimeData struct {
	idx       int
	uptimeper float64
	checkday  string
}

type WebGroupData struct {
	Name   string
	Member string
	Count  int
}
