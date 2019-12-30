package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Option struct {
	UserName string `json:"username"`
	Balance  int    `json:"balance"`
}

type Funmod func(option *Option)
type foo struct {
	option Option
}

func new(funmod Funmod) *foo {
	op := Option{
		UserName: "datou",
		Balance:  100,
	}
	funmod(&op)
	return &foo{option: op}
}

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	log.Println(elem)
	return s
}

func main0() {
	s := NewSlice()
	defer func() { s.Add(1).Add(2) }()
	//http.HandleFunc("/", getvaluehandler)
	//http.ListenAndServe("127.0.0.1:8800", nil)
	//fo := new(func(option *Option) {
	//	option.Balance = 2000
	//})
	s.Add(3)

}
func getvaluehandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	value := query.Get("test_key")
	log.Println(reflect.TypeOf(value))
	log.Println(value)
}
func main1() {
	renderzipfile:="/2019-12-24/68179-测试-files/"
	zipFilePath := renderzipfile[len(time.Unix(time.Now().Unix(), 0).Format("2006-01-02"+"/"))+1:]

	basepath := "getTaskBasePath"                                                                                                                                                      //./files/1610/2019-01-11
	outputpath := basepath + "/" + time.Unix(1577170380, 0).Format("2006-01-02") + zipFilePath[len(time.Unix(time.Now().Unix(), 0).Format("2006-01-02"+"/")):]      //./files/1610/2019-01-11/68114-熊猫——动画-frames/
	renderSourceDir := basepath + "/" + time.Unix(1577170380, 0).Format("2006-01-02") + zipFilePath[len(time.Unix(time.Now().Unix(), 0).Format("2006-01-02"+"/")):]

	//log.Println("zipFilePath: ",zipFilePath,len(time.Unix(time.Now().Unix(), 0).Format("2006-01-02"+"/")))
	//renderZipFile := "/" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02"+"/") + zipFilePath
	//log.Println("renderZipFile: ",renderZipFile)
	//renderOutputDir := "/" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02"+"/") + strconv.FormatInt(68114, 10) + "-" + "熊猫——动画" + "-frames/"
	//renderSourceDir := "/" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02"+"/") + strconv.FormatInt(68114, 10) + "-" + "熊猫——动画" + "-files/"
	log.Println("renderOutputDir: ",outputpath," renderSourceDir: ",renderSourceDir)

}
func main2()  {
	logPreFix:="test"
	httpClient := http.Client{Timeout: time.Second * 3}
	err:=errors.New("aaa")
	//var resp *http.Response
	for i := 1; i < 4; i++ {
		_, err = httpClient.Post("imageFileUrl", "application/json",nil)
		if err != nil {
			log.Println(logPreFix, "can't get image url from agency url "+"timeout No:"+strconv.Itoa(i)+",error: ", err)
			time.Sleep(time.Second)
			continue
		}
	}
	if err != nil {
		for i := 1; i < 4; i++ {
			err =nil
			if err != nil {
				log.Println(logPreFix, "can't get source file from agency url "+"timeout No:"+strconv.Itoa(i)+",error: ", err)
				time.Sleep(time.Second)
				continue
			}
		}
	}
	if err != nil  {
		log.Println(logPreFix, "============can't get source file from agency url "+"or", ",error: ", err)
		return
	}
	log.Println("test ok!!!!")

}

func main3()  {
	file, err := os.Create("./test1.go")
	if err != nil {
		log.Println(err)
	}
	md5h := md5.New()
	defer  file.Close()
	io.Copy(md5h,file)
	abc := string(md5h.Sum([]byte("")))
	fileMd5:=fmt.Sprintf("%x",abc)
	log.Println(fileMd5)

}