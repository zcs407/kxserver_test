package controller

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	prvkey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAz0yVlwDwX9eskudCnKTt2fDcfT3oGlWLEwNXhQb7ekBJW0Ze
UwbPyE2jXVb0dQ+wP8u5/Nfd4bNd48HaSmHhwiXWZPG+Wl1a28rDPfIr3mcOE+4n
3KULUfgrGh6LNua1qGJn07D9QmFocdVdDnkA9q3n/0xqKl1xMV9oDi/yhE+L+BQs
4jJ3EGJLRILECg78KM1By+eiowSAYL20bqz0vAti3sbvpkQbbR749qStvgW56As4
/sY3gDHVP+jW6693ZHzbw6bhR4OsQ3d2GS3MU0fYwS3Iyq/PfqNN1RzZLsMxwZQt
GitXsa/9lj0EpqeZue1zsvTTogeyAqvJLYQqpQIDAQABAoIBAQDJBTsKc56lUj/H
NPsja8w6y5cE3EN3RfzXMyZrmZnDsxNvr41IdhKH6sHAgdIMsmn3c1eoGKzRcV5a
vmEwQDrsSkTdHo+4kU7KVJWAPJbN1KGSMh/1lxajJkSlz3iwhIkkAEkuvzLYbB0c
Rgs3PZ/xljKjHzbUXkil2B7Poy9JnRrUcgJkmlrNJu5llRqWovBTC2UtFisgdtw+
1nZhc8uWsagzXJAy+xfgjm5plF7RTWS8mGbXglK57HQEF0P+gK0VzaZB78AB8ME+
FoNbos1+6g/tyXLOnslFPCKh179Lxb/FAglqysV01M/Qu0PiGO8qLtSjYay5nLCy
irvH+2sBAoGBAN67ey/hG4A0QZRUcj5H2T84Ylh1PRhFCmFdBWPKITNDmENEmHpG
nxsMSsmjpRFbUsRYXbkMXWnM4NuAiWNtHI1LhHtJw8BlcLqPa6OvXDKXVIiMxDoA
nWhS1rvhhMc5VZf6GHsbsCdU8CF/LqWITtp4tHDtBGL9KrhADZqOJeLnAoGBAO5C
/bKqWf7zUr8jGIOFHsPwsVLPmGWJ7LqDIQbxdakx0oP/OMGfBNxumJJadq7LtIvS
M+5aI4qvtdmVkCTGvYjQjF9h0p8B2GjgzvllzEdb0sUm534xYXPqjYFUzxtauZtI
YNNT/bOWIlKHahEkrPrLiv9ZL0eLC3fD7OI8qiCTAoGADs6+CNfZYTOYLIlUswlJ
yycvepwIvMVSRFjP0+uLO4JB7C7ySCbLyxuNGUy026uLnBwX1waYa0FArbck6yRE
4qvjmeK0jeTwkqaYTGCLK53d89oP7Z8+18GyHvmGP0xzgVASMpULqAHAmSmAa2bd
fy7JKDzJrt8P6QHxJZZPtH8CgYEA2gerqQeCe2+m1QoEsLXsxVlIq4MU7jYcz0CX
xIbJKR9SiT/QbD5ccGs0axklaic2/IxKwV7zD0Jjoszerwi/AKf3DIGz/5Xst2yh
ek/Rc6tvYMKNLEl76FtHSoaVT27iUlsVX82IaAKHPgZ05WMueAIzHCA8x7dRszMz
XoQtGskCgYEAtNfDz8exqblsApHEaQ9p2uoHO9UGqYmh1bQXSMtsovLoOpkQfp7y
VTXDscOL6NDJh3pmnA36FozwIVGGADUHQBZXbydg0LFUk5kAqjhIZYElWq/11iKl
RjXn8ak2KQ8Rchpb0S8K8APeaPW2rr0YctE3BatS7DyTH8LezDyREaY=
-----END RSA PRIVATE KEY-----`)
	agencyid = "M4JUWRUQ"
	sign     = []byte("")
	jsonData = make(map[string]string)
)

type user struct {
	Id       int64  `json:"id"`
	AgencyId string `json:"agencyid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Vip      string `json:"vip"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Balance  int64  `json:"balance"`
}

type taskInfo struct {
	//更新时不动
	ID           int64  `json:"id"`           //任务id
	Name         string `json:"name"`         //任务名称
	Owner        int64  `json:"owner"`        //经销商用户在我方库中的用户id
	OwnerName    string `json:"ownername"`    //经销商用户名
	AgencyUserID int64  `json:"agencyuserid"` //经销商用户id
	CreateTime   int64  `json:"createtime"`   //任务创建时间
	//更新时变动
	Type                    int          `json:"type"`                    //0: video, 1: picture
	Status                  int          `json:"status"`                  //任务状态
	RenderOutputFiles       []fileObject `json:"renderoutputfiles"`       //渲染输出文件切片(多文件)
	FrameCount              int          `json:"framecount"`              //帧数
	Gross                   int64        `json:"gross"`                   //最终费用
	FrameCompleted          int          `json:"framecompleted"`          //完成帧数
	StartTime               int64        `json:"starttime"`               //渲染的开始时间
	CompletedTime           int64        `json:"completedtime"`           //渲染的完成时间
	RenderTimeThreadTotal   int64        `json:"rendertimethreadtotal"`   //渲染线程耗时总和
	LightJobGross           int64        `json:"lightjobgross"`           //小光子任务总费用
	LightJobTimeThreadTotal int64        `json:"lightjobtimethreadtotal"` //小光子渲染的线程时
	RenderZipFile           string       `json:"renderzipfile"`           //渲染源文件zip包
	Camera                  string       `json:"camera"`
	Action                  int          `json:"action"`
}
type renderFrame struct {
	ID               int64        `json:"id"`
	Owner            int64        `json:"owner"`
	OwnerName        string       `json:"ownername"`
	AgencyUserID     int64        `json:"agencyuserid"`
	RenderTaskID     int64        `json:"rendertaskid"` // RenderTask DB ID
	FrameIndex       int          `json:"frameindex"`   // deadline Task frame index: 1000, max: 9999
	Status           int          `json:"status"`
	RenderTimeThread int64        `json:"rendertimethread"` // 线程时
	Gross            int64        `json:"gross"`
	Files            []fileObject `json:"files"`
	Type             int          `json:"type"`    // 0: video, 1: picture
	Deleted          int          `json:"deleted"` // 1： 任务已经删除
	Action           int          `json:"action"`  //操作: 0创建任务,1 更新任务状态,2任务完成并扣费,3任务删除操作
	// 光子相关信息，只对动画任务有意义
	LightJobGross            int64 `json:"lightjobgross"`            // 小光子任务最终费用
	LightJobRenderTimeThread int64 `json:"lightjobrendertimethread"` // 小光子渲染的线程时
	LightJobStatus           int   `json:"lightjobstatus"`           // 小光子任务状态
}
type grossInfo struct {
	ID           int64  `json:"id"`           //任务id
	Name         string `json:"name"`         //任务名称
	Owner        int64  `json:"owner"`        //经销商用户在我方库中的用户id
	OwnerName    string `json:"ownername"`    //经销商用户名
	AgencyUserID int64  `json:"agencyuserid"` //经销商用户id
	CreateTime   int64  `json:"createtime"`   //任务创建时间
	//更新时变动
	Type                    int          `json:"type"`                    //0: video, 1: picture
	Status                  int          `json:"status"`                  //任务状态
	RenderOutputFiles       []fileObject `json:"renderoutputfiles"`       //渲染输出文件切片(多文件)
	FrameCount              int          `json:"framecount"`              //帧数
	Gross                   int64        `json:"gross"`                   //最终费用
	FrameCompleted          int          `json:"framecompleted"`          //完成帧数
	StartTime               int64        `json:"starttime"`               //渲染的开始时间
	CompletedTime           int64        `json:"completedtime"`           //渲染的完成时间
	MachineLimit            int          `json:"machinelimit"`            //核数同步数据时需要处理为24倍
	Services                int          `json:"services"`                //任务分配的服务器数量
	RenderTimeThreadTotal   int64        `json:"rendertimethreadtotal"`   //渲染线程耗时总和
	LightJobGross           int64        `json:"lightjobgross"`           //小光子任务总费用
	LightJobTimeThreadTotal int64        `json:"lightjobtimethreadtotal"` //小光子渲染的线程时
	RenderZipFile           string       `json:"renderzipfile"`           //渲染源文件zip包
	Camera                  string       `json:"camera"`
	Action                  int          `json:"action"`
	FrameIndex              int          `json:"frameindex"`
	FrameType               int          `json:"frametype"`
}
type fileObject struct {
	//文件的名称,可用于下载已完成文件,下载完成文件可参考:快渲经销商Api文档-任务已完成文件下载接口
	Name string `json:"name"`
	//文件大小,不常用
	Size int64 `json:"size"`
	//文件时间,基本不用
	Time int64 `json:"time"`
	//文件是否为目录,基本不用
	IsDir bool `json:"isdir"`
}

func TaskCreateInfo(ctx *gin.Context) {
	var jsonTaskInfo taskInfo
	err := ctx.BindJSON(&jsonTaskInfo)
	if err != nil {
		log.Println("kuaixuan create task info sync error:", err)
		resp(ctx, "4000", "kuaixuan create task info sync error:", err)
	}
	//jsonMap := taskJsonMapHandler(jsonTaskInfo)
	//log.Println("kuaixuan create task info sync success, the task info is", jsonMap)
	resp(ctx, "200", "kuaixuan create task info sync success, the task info is", jsonTaskInfo)
}
func TaskUpdateInfo(ctx *gin.Context) {
	var jsonTaskInfo taskInfo
	err := ctx.BindJSON(&jsonTaskInfo)
	if err != nil {
		log.Println("kuaixuan update task info sync error:", err)
		resp(ctx, "4000", "kuaixuan update task info sync error:", err)
	}
	//jsonMap := taskJsonMapHandler(jsonTaskInfo)
	//log.Println("kuaixuan update task info sync success, the task info is", jsonMap)
	resp(ctx, "200", "kuaixuan update task info sync success, the task info is", jsonTaskInfo)
}
func TaskDeleteInfo(ctx *gin.Context) {
	var jsonTaskInfo taskInfo
	err := ctx.BindJSON(&jsonTaskInfo)
	if err != nil {
		log.Println("kuaixuan delete task info sync error:", err)
		resp(ctx, "4000", "kuaixuan delete task info sync error:", err)
	}
	jsonMap := taskJsonMapHandler(jsonTaskInfo)
	log.Println("kuaixuan delete task info sync success, the task info is", jsonMap)
	resp(ctx, "200", "kuaixuan delete task info sync success, the task info is", jsonTaskInfo)
}
func TaskGrossInfo(ctx *gin.Context) {
	var jsonTaskInfo grossInfo
	err := ctx.BindJSON(&jsonTaskInfo)
	if err != nil {
		log.Println("kuaixuan task gross info sync error:", err)
		resp(ctx, "4000", "kuaixuan task gross info sync error:", err)
	}
	jsonMap := taskGrossJsonMapHandler(jsonTaskInfo)
	log.Println("kuaixuan  task gross info sync success, the task info is", jsonMap)
	resp(ctx, "200", "kuaixuan task gross info sync success, the task info is", jsonTaskInfo)
}
func RenderFrameCreateInfo(ctx *gin.Context) {
	var jsonFrameInfo renderFrame
	err := ctx.BindJSON(&jsonFrameInfo)
	if err != nil {
		log.Println("kuaixuan create render frame info sync error:", err)
		resp(ctx, "4000", "kuaixuan create render frame info sync error:", err)
	}
	jsonMap := frameJsonMapHandler(jsonFrameInfo)
	log.Println("kuaixuan create render frame info sync success\n the task info is\n", jsonMap)
	resp(ctx, "200", "kuaixuan create render frame info sync success, the render frame info is", jsonFrameInfo)
}
func RenderFrameUpdateInfo(ctx *gin.Context) {
	var jsonFrameInfo renderFrame
	err := ctx.BindJSON(&jsonFrameInfo)
	if err != nil {
		log.Println("kuaixuan update render frame info sync error:", err)
		resp(ctx, "4000", "kuaixuan update render frame info sync error:", err)
	}
	jsonMap := frameJsonMapHandler(jsonFrameInfo)
	log.Println("kuaixuan update render frame info sync success\n the task info is\n", jsonMap)
	resp(ctx, "200", "kuaixuan update render frame info sync success, the render frame info is", jsonFrameInfo)
}
func UserRegister(ctx *gin.Context) {
	var userList = []user{}
	err := ctx.BindJSON(&userList)
	if err != nil {
		log.Println("agency id: "+agencyid+" add user info from kx error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" add user info from kx error:", err)
		return
	}
	dataByte, err := json.Marshal(userList)
	if err != nil {
		log.Println("添加用户系列化错误: ", err)
		resp(ctx, "4000", "添加用户系列化错误", err)
		return
	}
	sign, err = RsaSignWithSha256(dataByte, prvkey)
	if err != nil {
		log.Println("签名失败", err)
		resp(ctx, "4000", "agency id: "+agencyid+" add user info from kx sign failed:", err)
		return
	}
	log.Println("签名成功,", hex.EncodeToString(sign))
	jsonData["agencyid"] = agencyid
	jsonData["data"] = string(dataByte)
	jsonData["sign"] = hex.EncodeToString(sign)
	jsona, err := json.Marshal(jsonData)
	if len(userList) == 0 || len(agencyid) == 0 || len(sign) == 0 {
		log.Println("agency id: "+agencyid+" add user info from kx error:", "userlist agencyid sign can't be null")
		resp(ctx, "4000", "agency id: "+agencyid+" add user info from kx error:", "userlist agencyid sign can't be null")
		return
	}
	res, err := http.Post("http://127.0.0.1:8891/api/v1/agency/aaa/user/add", "application/json", bytes.NewBuffer(jsona))
	if err != nil {
		respBody, _ := ioutil.ReadAll(res.Body)
		log.Println("快渲回复的信息", respBody)
		respmap := make(map[string]interface{})
		err = json.Unmarshal(respBody, &respmap)
		if err != nil {
			log.Println("连接错误,反序列化错误  error:", err)
			resp(ctx, "4000", "agency id: "+agencyid+" add user info from kx", err)
			return
		}
	}
	respBody, _ := ioutil.ReadAll(res.Body)
	respmap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respmap)
	if err != nil {
		log.Println("发送数据成功,反序列化出错 error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" add user info from kx", err)
		return
	}
	log.Println("agency id: " + agencyid + " add user info from kx success:")
	resp(ctx, "2000", "agency id: "+agencyid+" add user info from kx", respmap)
}
func UserUpdate(ctx *gin.Context) {
	var userInfo = user{}
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		log.Println("agency id: "+agencyid+" update user info from kx error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" update user info from kx error:", err)
		return
	}
	dataByte, err := json.Marshal(userInfo)
	if err != nil {
		log.Println("更新用户系列化错误: ", err)
		resp(ctx, "4000", "更新用户系列化错误", err)
		return
	}
	sign, err = RsaSignWithSha256(dataByte, prvkey)
	if err != nil {
		log.Println("签名失败", err)
		resp(ctx, "4000", "agency id: "+agencyid+" update user info from kx sign failed:", err)
		return
	}
	log.Println("签名成功,", hex.EncodeToString(sign))
	jsonData["agencyid"] = agencyid
	jsonData["data"] = string(dataByte)
	jsonData["sign"] = hex.EncodeToString(sign)
	jsona, err := json.Marshal(jsonData)
	if len(agencyid) == 0 || len(sign) == 0 {
		log.Println("agency id: "+agencyid+" update user info from kx error:", "userlist agencyid sign can't be null")
		resp(ctx, "4000", "agency id: "+agencyid+" update user info from kx error:", "userlist agencyid sign can't be null")
		return
	}
	res, err := http.Post("http://127.0.0.1:8891/api/v1/agency/aaa/user/update", "application/json", bytes.NewBuffer(jsona))
	if err != nil {
		log.Println("http post error: ",err)
		return
	}
	respBody, _ := ioutil.ReadAll(res.Body)
	respmap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respmap)
	if err != nil {
		log.Println("发送数据成功,反序列化出错 error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" update user info from kx", err)
		return
	}
	log.Println("agency id: " + agencyid + " update user info from kx success:")
	resp(ctx, "2000", "agency id: "+agencyid+" update user info from kx", respmap)
}
func UserDelete(ctx *gin.Context) {
	type jsonuser struct {
		AgencyUserID int64  `json:"id"`
		AgencyId     string `json:"agencyid"` //经销商ID
	}
	var userInfo jsonuser
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		log.Println("agency id: "+agencyid+" delete user info from kx error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" delete user info from kx error:", err)
		return
	}
	dataByte, err := json.Marshal(userInfo)
	if err != nil {
		log.Println("删除用户系列化错误: ", err)
		resp(ctx, "4000", "删除用户系列化错误", err)
		return
	}
	sign, err = RsaSignWithSha256(dataByte, prvkey)
	if err != nil {
		log.Println("签名失败", err)
		resp(ctx, "4000", "agency id: "+agencyid+" delete user info from kx sign failed:", err)
		return
	}
	log.Println("签名成功,", hex.EncodeToString(sign))
	jsonData["agencyid"] = agencyid
	jsonData["data"] = string(dataByte)
	jsonData["sign"] = hex.EncodeToString(sign)
	jsona, err := json.Marshal(jsonData)
	if len(agencyid) == 0 || len(sign) == 0 {
		log.Println("agency id: "+agencyid+" delete user info from kx error:", "userlist agencyid sign can't be null")
		resp(ctx, "4000", "agency id: "+agencyid+" delete user info from kx error:", "user agencyid sign can't be null")
		return
	}
	res, err := http.Post("http://127.0.0.1:8891/api/v1/agency/aaa/user/delete", "application/json", bytes.NewBuffer(jsona))
	if err != nil {
		respBody, _ := ioutil.ReadAll(res.Body)
		log.Println("快渲回复的信息", respBody)
		respmap := make(map[string]interface{})
		err = json.Unmarshal(respBody, &respmap)
		if err != nil {
			log.Println("连接错误,反序列化错误  error:", err)
			resp(ctx, "4000", "agency id: "+agencyid+" delete user info from kx", err)
			return
		}
	}
	respBody, _ := ioutil.ReadAll(res.Body)
	respmap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respmap)
	if err != nil {
		log.Println("发送数据成功,反序列化出错 error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" update user info from kx", err)
		return
	}
	log.Println("agency id: " + agencyid + " update user info from kx success:")
	resp(ctx, "2000", "agency id: "+agencyid+" update user info from kx", respmap)
}

func UserIdList(ctx *gin.Context) {
	type agencyData struct {
		AgencyId int64 `json:"agencyid"`
	}
	agencydata := agencyData{}
	err := ctx.BindJSON(&agencydata)
	if err != nil {
		log.Println("无法获取request中的时间戳")
		return
	}
	dataByte, err := json.Marshal(agencydata)
	if err != nil {
		log.Println("获取用户id列表失败,序列化错误: ", err)
		resp(ctx, "4000", "获取用户id列表失败,序列化错误", err)
		return
	}
	sign, err = RsaSignWithSha256(dataByte, prvkey)
	if err != nil {
		log.Println("签名失败", err)
		resp(ctx, "4000", "agency id: "+agencyid+" get user id list  from kx sign failed:", err)
		return
	}
	log.Println("签名成功,", hex.EncodeToString(sign))
	jsonData["agencyid"] = agencyid
	jsonData["data"] = string(dataByte)
	jsonData["sign"] = hex.EncodeToString(sign)
	jsona, err := json.Marshal(jsonData)
	if len(agencyid) == 0 || len(sign) == 0 {
		log.Println("agency id: "+agencyid+" get user id list from kx error:", " agencyid sign can't be null")
		resp(ctx, "4000", "agency id: "+agencyid+" get user id list from kx error:", " agencyid sign can't be null")
		return
	}
	res, err := http.Post("https://www.speedyrender.cn:8890/api/v1/agency/aaa/user/idlist", "application/json", bytes.NewBuffer(jsona))
	if err != nil {
		respBody, _ := ioutil.ReadAll(res.Body)
		log.Println("快渲回复的信息", respBody)
		respmap := make(map[string]interface{})
		err = json.Unmarshal(respBody, &respmap)
		if err != nil {
			log.Println("连接错误,反序列化错误  error:", err)
			resp(ctx, "4000", "agency id: "+agencyid+" get user id list from kx", err)
			return
		}
	}
	respBody, _ := ioutil.ReadAll(res.Body)
	respmap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respmap)
	if err != nil {
		log.Println("发送数据成功,反序列化出错 error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" get user id list from kx", err)
		return
	}
	log.Println("agency id: " + agencyid + " get user id list from kx success:")
	resp(ctx, "2000", "agency id: "+agencyid+" get user id list from kx", respmap)
}

func UserInfo(ctx *gin.Context) {
	type agencyData struct {
		Id int64 `json:"id"`
	}
	agencydata := agencyData{}
	err := ctx.BindJSON(&agencydata)
	if err != nil {
		log.Println("无法获取request中的时间戳")
		return
	}
	dataByte, err := json.MarshalIndent(agencydata,"","")
	if err != nil {
		log.Println("获取用户信息失败,序列化错误: ", err)
		resp(ctx, "4000", "获取用户信息失败,序列化错误", err)
		return
	}
	sign, err = RsaSignWithSha256(dataByte, prvkey)
	if err != nil {
		log.Println("签名失败", err)
		resp(ctx, "4000", "agency id: "+agencyid+" get user info  from kx sign failed:", err)
		return
	}
	log.Println("签名成功,", hex.EncodeToString(sign))
	jsonData["agencyid"] = agencyid
	jsonData["data"] = string(dataByte)
	jsonData["sign"] = hex.EncodeToString(sign)
	jsona, err := json.Marshal(jsonData)
	if agencydata.Id == 0 || len(agencyid) == 0 || len(sign) == 0 {
		log.Println("agency id: "+agencyid+" get user info from kx error:", " agencyid userid sign can't be null")
		resp(ctx, "4000", "agency id: "+agencyid+" get user info from kx error:", " agencyid sign can't be null")
		return
	}
	res, err := http.Post("https://www.speedyrender.cn:8890/api/v1/agency/aaa/user/info", "application/json", bytes.NewBuffer(jsona))
	if err != nil {
		respBody, _ := ioutil.ReadAll(res.Body)
		log.Println("快渲回复的信息", respBody)
		respmap := make(map[string]interface{})
		err = json.Unmarshal(respBody, &respmap)
		if err != nil {
			log.Println("连接错误,反序列化错误  error:", err)
			resp(ctx, "4000", "agency id: "+agencyid+" get user info from kx", err)
			return
		}
	}
	respBody, _ := ioutil.ReadAll(res.Body)
	respmap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respmap)
	if err != nil {
		log.Println("发送数据成功,反序列化出错 error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" get user info from kx", err)
		return
	}
	log.Println("agency id: " + agencyid + " get user info from kx success:")
	resp(ctx, "2000", "agency id: "+agencyid+" get user info from kx", respmap)
}

func GetTaskidsByCreatetime(ctx *gin.Context) {
	type agencyData struct {
		CreateTime int64 `json:"createtime"`
	}
	agencydata := agencyData{}
	err := ctx.BindJSON(&agencydata)
	createtime := agencydata.CreateTime
	if err != nil {
		log.Println("无法获取request中的时间戳")
		return
	}
	if createtime == 0 {
		log.Println("创建时间不可为0 ")
		return
	}
	//dataByte := []byte(strconv.FormatInt(createtime, 10))
	dataByte, err := json.Marshal(agencydata)
	if err != nil {
		log.Println("获取任务id列表失败,序列化错误: ", err)
		resp(ctx, "4000", "获取任务id列表失败,序列化错误", err)
		return
	}
	sign, err = RsaSignWithSha256(dataByte, prvkey)
	if err != nil {
		log.Println("签名失败", err)
		resp(ctx, "4000", "agency id: "+agencyid+" get taskids info from kx sign failed:", err)
		return
	}
	log.Println("签名成功,", hex.EncodeToString(sign))
	jsonData["agencyid"] = agencyid
	jsonData["data"] = string(dataByte)
	jsonData["sign"] = hex.EncodeToString(sign)
	jsona, err := json.Marshal(jsonData)
	if createtime == 0 || len(agencyid) == 0 || len(sign) == 0 {
		log.Println("agency id: "+agencyid+" get taskids info from kx error:", "createtime agencyid sign can't be null")
		resp(ctx, "4000", "agency id: "+agencyid+" get taskids info from kx error:", "createtime agencyid sign can't be null")
		return
	}
	res, err := http.Post("https://www.speedyrender.cn:8890/api/v1/kuaixuan/agency/task/idlist", "application/json", bytes.NewBuffer(jsona))
	if err != nil {
		respBody, _ := ioutil.ReadAll(res.Body)
		log.Println("快渲回复的信息", respBody)
		respmap := make(map[string]interface{})
		err = json.Unmarshal(respBody, &respmap)
		if err != nil {
			log.Println("连接错误,反序列化错误  error:", err)
			resp(ctx, "4000", "agency id: "+agencyid+" get taskids info from kx", err)
			return
		}
	}
	respBody, _ := ioutil.ReadAll(res.Body)
	respmap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respmap)
	if err != nil {
		log.Println("发送数据成功,反序列化出错 error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" get taskids info from kx", err)
		return
	}
	log.Println("agency id: " + agencyid + " get taskids info from kx success:")
	resp(ctx, "2000", "agency id: "+agencyid+" get taskids info from kx", respmap)
}

func GetTaskInfo(ctx *gin.Context) {
	type agencyData struct {
		TaskId int64 `json:"taskid"`
	}
	agencydata := agencyData{}
	err := ctx.BindJSON(&agencydata)
	if agencydata.TaskId == 0 {
		log.Println("任务id不可为0 ")
		return
	}
	dataByte, err := json.Marshal(agencydata)
	if err != nil {
		log.Println("获取任务信息失败,序列化错误: ", err)
		resp(ctx, "4000", "获取任务信息失败,序列化错误", err)
		return
	}
	sign, err = RsaSignWithSha256(dataByte, prvkey)
	if err != nil {
		log.Println("签名失败", err)
		resp(ctx, "4000", "agency id: "+agencyid+" get task info  from kx sign failed:", err)
		return
	}
	log.Println("签名成功,", hex.EncodeToString(sign))
	jsonData["agencyid"] = agencyid
	jsonData["data"] = string(dataByte)
	jsonData["sign"] = hex.EncodeToString(sign)
	jsona, err := json.Marshal(jsonData)
	if agencydata.TaskId == 0 || len(agencyid) == 0 || len(sign) == 0 {
		log.Println("agency id: "+agencyid+" get task info from kx error:", "createtime agencyid sign can't be null")
		resp(ctx, "4000", "agency id: "+agencyid+" get task info from kx error:", "createtime agencyid sign can't be null")
	}
	res, err := http.Post("https://www.speedyrender.cn:8890/api/v1/kuaixuan/agency/task/info", "application/json", bytes.NewBuffer(jsona))
	if err != nil {
		respBody, _ := ioutil.ReadAll(res.Body)
		log.Println("快渲回复的信息", respBody)
		respmap := make(map[string]interface{})
		err = json.Unmarshal(respBody, &respmap)
		if err != nil {
			log.Println("连接错误,反序列化错误  error:", err)
			resp(ctx, "4000", "agency id: "+agencyid+" get task info from kx", err)
			return
		}
	}
	respBody, _ := ioutil.ReadAll(res.Body)
	respmap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respmap)
	if err != nil {
		log.Println("发送数据成功,反序列化出错 error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" get task info from kx", err)
		return
	}
	log.Println("agency id: " + agencyid + " get task info from kx success:")
	resp(ctx, "2000", "agency id: "+agencyid+" get task info from kx", respmap)
}

func GetFrameInfo(ctx *gin.Context) {
	type agencyData struct {
		TaskId int64 `json:"taskid"`
	}

	agencydata := agencyData{}
	err := ctx.BindJSON(&agencydata)
	if agencydata.TaskId == 0 {
		log.Println("任务id不可为0 ")
		return
	}
	dataByte, err := json.Marshal(agencydata)
	if err != nil {
		log.Println("获取任务帧失败,序列化错误: ", err)
		resp(ctx, "4000", "获取任务帧信息失败,序列化错误", err)
		return
	}
	sign, err = RsaSignWithSha256(dataByte, prvkey)
	if err != nil {
		log.Println("签名失败", err)
		resp(ctx, "4000", "task id: "+strconv.FormatInt(agencydata.TaskId, 10)+" get frame info sign failed:", err)
		return
	}
	log.Println("签名成功,", hex.EncodeToString(sign))
	jsonData["agencyid"] = agencyid
	jsonData["data"] = string(dataByte)
	jsonData["sign"] = hex.EncodeToString(sign)
	jsona, err := json.Marshal(jsonData)
	if agencydata.TaskId == 0 || len(agencyid) == 0 || len(sign) == 0 {
		log.Println("agency id: "+agencyid+" get task info from kx error:", "createtime agencyid sign can't be null")
		resp(ctx, "4000", "agency id: "+agencyid+" get task info from kx error:", "createtime agencyid sign can't be null")
		return
	}
	res, err := http.Post("https://www.speedyrender.cn:8890/api/v1/kuaixuan/agency/task/frameinfo", "application/json", bytes.NewBuffer(jsona))
	if err != nil {
		respBody, _ := ioutil.ReadAll(res.Body)
		log.Println("快渲回复的信息", respBody)
		respmap := make(map[string]interface{})
		err = json.Unmarshal(respBody, &respmap)
		if err != nil {
			log.Println("连接错误,反序列化错误  error:", err)
			resp(ctx, "4000", "agency id: "+agencyid+" add user info from kx", err)
			return
		}
	}
	respBody, _ := ioutil.ReadAll(res.Body)
	respmap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respmap)
	if err != nil {
		log.Println("发送数据成功,反序列化出错 error:", err)
		resp(ctx, "4000", "agency id: "+agencyid+" add user info from kx", err)
		return
	}
	log.Println("agency id: " + agencyid + " add user info from kx success:")
	resp(ctx, "2000", "agency id: "+agencyid+" add user info from kx", respmap)
}

func GetFileUrl(ctx *gin.Context) {
	var jsonData = make(map[string]interface{})
	var data struct {
		Uid      int64  `json:"id"`
		FimeName string `json:"fimename"`
	}
	type sendData struct {
		Uid int64 `json:"id"`
	}
	senddata := sendData{}
	err := ctx.BindJSON(&data)
	if err != nil {
		log.Println("无法获取request中的用户id")
		return
	}
	if data.Uid == 0 || len(data.FimeName) == 0 {
		log.Println("用户id和文件名不可为空 ")
		return
	}
	senddata.Uid = data.Uid
	dataByte, err := json.Marshal(senddata)
	sign, err = RsaSignWithSha256(dataByte, prvkey)
	if err != nil {
		log.Println("签名失败", err)
		resp(ctx, "4000", "agency id: "+agencyid+" get download file url  from kx sign failed:", err)
	}
	log.Println("签名成功,", hex.EncodeToString(sign))
	jsonData["agencyid"] = agencyid
	jsonData["data"] = string(dataByte)
	jsonData["sign"] = hex.EncodeToString(sign)
	jsona, err := json.Marshal(jsonData)
	if data.Uid == 0 || len(agencyid) == 0 || len(sign) == 0 {
		log.Println("agency id: "+agencyid+" get download file url from kx error:", "createtime agencyid sign can't be null")
		resp(ctx, "4000", "agency id: "+agencyid+" get download file url from kx error:", "createtime agencyid sign can't be null")
	}
	//前序工作成功后创建本地文件,用于文件接收
	dstFile, err := os.Create("./aaa.zip")
	if err != nil {
		log.Println("无法创建接收下载的文件错误: ", err)
	}
	defer dstFile.Close()
	res, err := http.Post("http://123.55.235.42:8889/api/v1/kuaixuan/agency/task/download"+data.FimeName, "application/json", bytes.NewBuffer(jsona))
	if err != nil {
		respBody, _ := ioutil.ReadAll(res.Body)
		log.Println("快渲回复的信息", respBody)
	}
	log.Println("访问的路径为:", "http://123.55.233.97:8889/api/v1/kuaixuan/agency/task/download"+data.FimeName)
	_, err = io.Copy(dstFile, res.Body)
	if err != nil {
		log.Println("无法接收下载的文件内容,错误为: ", err)
	}

	log.Println("agency id: " + agencyid + " get download file url from kx success:")
	resp(ctx, "2000", "agency id: "+agencyid+" get download file url from kx", nil)
}

func resp(ctx *gin.Context, code, info string, body interface{}) {
	var resp struct {
		Code string      `json:"code"`
		Info string      `json:"info"`
		Body interface{} `json:"body"`
	}
	resp.Code = code
	resp.Info = info
	resp.Body = body
	ctx.JSON(200, resp)
}

func taskJsonMapHandler(jsonTaskInfo taskInfo) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	jsonMap["任务id"] = jsonTaskInfo.ID
	jsonMap["任务名"] = jsonTaskInfo.Name
	jsonMap["本库用户id"] = jsonTaskInfo.Owner
	jsonMap["用户名"] = jsonTaskInfo.OwnerName
	jsonMap["用户在经销商库中的id"] = jsonTaskInfo.AgencyUserID
	jsonMap["创建时间"] = jsonTaskInfo.CreateTime
	switch jsonTaskInfo.Type {
	case 0:
		jsonMap["任务类型"] = "动画类型"
	case 1:
		jsonMap["任务类型"] = "效果图类型"
	default:
		jsonMap["任务类型"] = "错误的任务类型"
	}
	//汉化任务状态
	switch jsonTaskInfo.Status {
	case 0:
		jsonMap["任务状态"] = "等待文件上传"
	case 1:
		jsonMap["任务状态"] = "渲染中"
	case 2:
		jsonMap["任务状态"] = "任务完成"
	case 3:
		jsonMap["任务状态"] = "任务暂停"
	case 4:
		jsonMap["任务状态"] = "任务失败"
	case 5:
		jsonMap["任务状态"] = "等待其他依赖任务完成后继续任务"
	case 6:
		jsonMap["任务状态"] = "上传文件中"
	case 7:
		jsonMap["任务状态"] = "处理上传后的文件,提供任务使用"
	case 8:
		jsonMap["任务状态"] = "任务队列"
	default:
		jsonMap["任务状态"] = "错误的状态"
	}
	jsonMap["输出文件"] = jsonTaskInfo.RenderOutputFiles
	jsonMap["总帧数"] = jsonTaskInfo.FrameCount
	jsonMap["总费用"] = jsonTaskInfo.Gross
	jsonMap["完成帧数"] = jsonTaskInfo.FrameCompleted
	jsonMap["开始渲染时间"] = jsonTaskInfo.StartTime
	jsonMap["结束时间"] = jsonTaskInfo.CompletedTime
	jsonMap["渲染总耗时"] = jsonTaskInfo.RenderTimeThreadTotal
	jsonMap["小光子总费用"] = jsonTaskInfo.LightJobGross
	jsonMap["小光子总耗时"] = jsonTaskInfo.LightJobTimeThreadTotal
	jsonMap["渲染任务源文件"] = jsonTaskInfo.RenderZipFile
	jsonMap["相机"] = jsonTaskInfo.Camera
	switch jsonTaskInfo.Action {
	case 0:
		jsonMap["操作"] = "创建任务"
	case 1:
		jsonMap["操作"] = "更新任务状态"
	case 2:
		jsonMap["操作"] = "任务完成并扣费"
	case 3:
		jsonMap["操作"] = "删除任务"
	default:
		jsonMap["操作"] = "未知操作"
	}
	return jsonMap
}
func taskGrossJsonMapHandler(jsonTaskInfo grossInfo) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	jsonMap["任务id"] = jsonTaskInfo.ID
	jsonMap["任务名"] = jsonTaskInfo.Name
	jsonMap["本库用户id"] = jsonTaskInfo.Owner
	jsonMap["用户名"] = jsonTaskInfo.OwnerName
	jsonMap["用户在经销商库中的id"] = jsonTaskInfo.AgencyUserID
	jsonMap["创建时间"] = jsonTaskInfo.CreateTime
	jsonMap["帧位置"] = jsonTaskInfo.FrameIndex
	jsonMap["帧类型"] = jsonTaskInfo.FrameType
	switch jsonTaskInfo.Type {
	case 0:
		jsonMap["任务类型"] = "动画类型"
	case 1:
		jsonMap["任务类型"] = "效果图类型"
	default:
		jsonMap["任务类型"] = "错误的任务类型"
	}
	//汉化任务状态
	switch jsonTaskInfo.Status {
	case 0:
		jsonMap["任务状态"] = "等待文件上传"
	case 1:
		jsonMap["任务状态"] = "渲染中"
	case 2:
		jsonMap["任务状态"] = "任务完成"
	case 3:
		jsonMap["任务状态"] = "任务暂停"
	case 4:
		jsonMap["任务状态"] = "任务失败"
	case 5:
		jsonMap["任务状态"] = "等待其他依赖任务完成后继续任务"
	case 6:
		jsonMap["任务状态"] = "上传文件中"
	case 7:
		jsonMap["任务状态"] = "处理上传后的文件,提供任务使用"
	case 8:
		jsonMap["任务状态"] = "任务队列"
	default:
		jsonMap["任务状态"] = "错误的状态"
	}
	jsonMap["输出文件"] = jsonTaskInfo.RenderOutputFiles
	jsonMap["总帧数"] = jsonTaskInfo.FrameCount
	jsonMap["总费用"] = jsonTaskInfo.Gross
	jsonMap["完成帧数"] = jsonTaskInfo.FrameCompleted
	jsonMap["开始渲染时间"] = jsonTaskInfo.StartTime
	jsonMap["结束时间"] = jsonTaskInfo.CompletedTime
	jsonMap["渲染总耗时"] = jsonTaskInfo.RenderTimeThreadTotal
	jsonMap["小光子总费用"] = jsonTaskInfo.LightJobGross
	jsonMap["小光子总耗时"] = jsonTaskInfo.LightJobTimeThreadTotal
	jsonMap["渲染任务源文件"] = jsonTaskInfo.RenderZipFile
	jsonMap["相机"] = jsonTaskInfo.Camera
	switch jsonTaskInfo.Action {
	case 0:
		jsonMap["操作"] = "创建任务"
	case 1:
		jsonMap["操作"] = "更新任务状态"
	case 2:
		jsonMap["操作"] = "任务完成并扣费"
	case 3:
		jsonMap["操作"] = "删除任务"
	default:
		jsonMap["操作"] = "未知操作"
	}
	return jsonMap
}
func frameJsonMapHandler(jsonFrameInfo renderFrame) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	jsonMap["渲染任务id"] = jsonFrameInfo.ID
	jsonMap["所属任务id"] = jsonFrameInfo.RenderTaskID
	jsonMap["经销商用户id"] = jsonFrameInfo.AgencyUserID
	jsonMap["用户名"] = jsonFrameInfo.OwnerName
	jsonMap["用户在经销商库中的id"] = jsonFrameInfo.AgencyUserID
	jsonMap["状态"] = jsonFrameInfo.Status
	switch jsonFrameInfo.Type {
	case 0:
		jsonMap["任务类型"] = "动画类型"
	case 1:
		jsonMap["任务类型"] = "效果图类型"
	default:
		jsonMap["任务类型"] = "错误的任务类型"
	}
	jsonMap["输出文件"] = jsonFrameInfo.Files
	jsonMap["帧位置"] = jsonFrameInfo.FrameIndex
	jsonMap["总费用"] = jsonFrameInfo.Gross
	jsonMap["渲染总耗时"] = jsonFrameInfo.RenderTimeThread
	jsonMap["小光子总费用"] = jsonFrameInfo.LightJobGross
	jsonMap["小光子总耗时"] = jsonFrameInfo.LightJobRenderTimeThread
	//jsonMap["任务是否删除"] = jsonFrameInfo.Deleted
	switch jsonFrameInfo.Action {
	case 0:
		jsonMap["操作"] = "创建任务"
	case 1:
		jsonMap["操作"] = "完成任务并扣费"
	case 2:
		jsonMap["操作"] = "任务更新"
	default:
		jsonMap["操作"] = "未知操作"
	}
	return jsonMap
}
func RsaSignWithSha256(data []byte, keyBytes []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS1PrivateKey err", err)
		return nil, errors.New("private key error")
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
