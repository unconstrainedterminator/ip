package ip

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/unconstrainedterminator/os"
	"path/filepath"
	"strings"
)

type Address struct {
	addr   string
	region string
}

var (
	dbpath   = filepath.Join(os.GetCurrentPath(), "region.xdb")
	Searcher *xdb.Searcher
)

func init() {
	buff, err := xdb.LoadContentFromFile(dbpath)
	if err != nil {
		glog.Error(context.Background(), err)
		return
	}

	Searcher, err = xdb.NewWithBuffer(buff)
	if err != nil {
		glog.Error(context.Background(), err)
		return
	}
}

func New(addr string) *Address {
	a := &Address{}
	a.addr = addr
	a.Search()
	return a
}

func (a *Address) Search() {
	var err error

	if strings.Contains(a.addr, ",") {
		addr := strings.Split(a.addr, ",")
		if len(addr) > 0 {
			a.addr = addr[0]
		}
	}

	a.region, err = Searcher.SearchByStr(a.addr)
	if err != nil {
		glog.Error(context.Background(), err)
	}
}

func (a *Address) Country() string {
	if a.region != "" {
		info := strings.Split(a.region, "|")
		if len(info) == 5 {
			if info[0] == "0" && info[1] == "0" && info[2] == "0" {
				return info[4]
			}
		}
		return info[0]
	}
	return ""
}

func (a *Address) Province() string {
	if a.region != "" {
		info := strings.Split(a.region, "|")
		if len(info) == 5 && info[2] != "0" {
			return info[2]
		}
	}
	return ""
}

func (a *Address) City() string {
	if a.region != "" {
		info := strings.Split(a.region, "|")
		if len(info) == 5 && info[3] != "0" {
			return info[3]
		}
	}
	return ""
}

func (a *Address) Region(separator ...string) string {
	if a.region != "" {
		info := strings.Split(a.region, "|")
		info = a.unique(info)

		sep := ","
		if len(separator) == 1 {
			sep = separator[0]
		}
		return a.removeZero(info, sep)
	}
	return ""
}

func (a *Address) unique(src []string) (unique []string) {
	for _, v := range src {
		if !a.valueInArray(v, unique) {
			unique = append(unique, v)
		}
	}
	return
}

func (a *Address) valueInArray(src string, dst []string) bool {
	for _, v := range dst {
		if src == v {
			return true
		}
	}
	return false
}

func (a *Address) removeZero(src []string, separator string) string {
	var result []string
	for _, str := range src {
		if str != "0" {
			result = append(result, str)
		}
	}
	return strings.Join(result, separator)
}
