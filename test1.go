package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Test interface {
	Tester()
}
type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}
func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}
func main01() {
	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Tester()
}

func main02() {
	str := "1ajkdals"
	ok, err := regexp.MatchString(`^[a-zA-Z]`, str)
	if err != nil {

	}
	log.Println(str, "是否为字母开头:", ok)
}

func main03() {
	localseting := make(map[string]string)
	fs, err := os.Open("tsconfig.json")
	if err != nil {
		log.Printf("open config json error:%v\n:", err)
	}
	defer fs.Close()
	data, err := ioutil.ReadAll(fs)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(data, &localseting)

	mk := make([]string, len(localseting))
	i := 0
	for k, _ := range localseting {
		mk[i] = k
		i++
	}
	sort.Strings(mk)
	for _, v := range mk {
		fmt.Printf("%s = %s\n", v, localseting[v])
	}
}

func main04() {
	type test struct {
		Name string
		Age  int
	}
	teststr := test{
		Name: "123",
		Age:  123,
	}
	//t := reflect.TypeOf(teststr)
	//v := reflect.ValueOf(teststr)
	//for k := 0; k < t.NumField(); k++ {
	//	log.Printf("%s--------%v\n",t.Field(k).Name,v.Field(k).Interface())
	//}
	data,_:=json.Marshal(teststr)
	log.Println(string(data))
}
func main05()  {
	type ipInfo struct {
		ID      int    `json:"id"`
		IP      string `json:"ip"`
		Name    string `json:"name"`
		TxSpeed string `json:"txspeed"`
		RxSpeed string `json:"rxspeed"`
	}
	var (
		//jsonData struct {
		//	TemplateTaskID           int64  `json:"templatetaskid"`
		//	TemplateTaskAgencyUserID int64  `json:"templatetaskagencyuserid"`
		//	AgencyUserID             int64  `json:"agencyuserid"`
		//	ImageFileUrl             string `json:"imagefileurl"`
		//	ImageFileMD5             string `json:"imagefilemd5"`
		//	SourceFileUrl            string `json:"sourcefileurl"`
		//	SourceFileMD5            string `json:"sourcefilemd5"`
		//}

		ipRespInfo struct {
			//"error":0,"errormsg":"success","extramsg":null,"id":0,"data":[]
			Error    int         `json:"error"`
			ErrorMsg string      `json:"errormsg"`
			ExtraMsg interface{} `json:"extramsg"`
			ID       int         `json:"id"`
			Data     []ipInfo    `json:"data"`
		}
	)
	serverIp := ""
	getServerIPUrl := "http://www.speedyrender.cn/api/v1/kuaixuan/serverinfo/"
	ipResp, err := http.Post(getServerIPUrl, "application/json", nil)
	if err != nil {
		log.Println( "can't get server ip info,post url: "+getServerIPUrl+",error: ", err)
		return
	}
	ipRespData, err := ioutil.ReadAll(ipResp.Body)
	if err != nil {
		log.Println( "can't get server ip info,can't ready data from ip resp data,error: ", err)
		return
	}
	err = json.Unmarshal(ipRespData, &ipRespInfo)
	if err != nil {
		log.Println( "can't get server ip info,json unmarshal error: ", err)
		return
	}
	for _, v := range ipRespInfo.Data {
		rxSpeedInt, _ := strconv.Atoi(v.RxSpeed)
		if rxSpeedInt < 5 {
			serverIp = v.IP
			break
		}
	}
	log.Println("server ip :",serverIp)
}
func main() {
	var resp struct {
		Code string      `json:"code"`
		Info string      `json:"info"`
		Body interface{} `json:"body"`
	}
	resp.Code = "code"
	resp.Info = "info"
	resp.Body = "body"
	log.Println(resp)

}