package util

import (
	"time"
	"strconv"
)

const (
	FORMAT_DATE     = "2006-01-02"
	FORMAT_DATETIME = "2006-01-02 15:04:05"
)

func Today() string {
	return time.Now().Format(FORMAT_DATE)
}

func Now() string {
	return time.Now().Format(FORMAT_DATETIME)
}

func SinceMinutes(fromtime string) int {
	fromTime, err := time.ParseInLocation(FORMAT_DATETIME, fromtime, time.Local)
	if err != nil {
		return -1
	}
	return int(time.Since(fromTime).Minutes())
}

//日期，兼容java习惯
type DateTime struct {
	t *time.Time
}

func (this *DateTime) FromDateTime(dt string) {
	t,_:=time.ParseInLocation(FORMAT_DATETIME,dt,time.Local)
	this.t=&t
}

func (this *DateTime) FromTimeMillis(long int64) {
	t:=time.Unix(long/1000,long*1e6)
    this.t=&t
}

func (this DateTime) CurrentTimeMillis() string{
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}


func (this *DateTime) GetTimeMillis() string{
	if this.t==nil {
		return ""
	}
	return strconv.FormatInt(this.t.UnixNano()/1e6, 10)
}

func (this *DateTime) GetDate() string {
	if this.t==nil {
		return ""
	}
	return this.t.Format(FORMAT_DATE)
}


