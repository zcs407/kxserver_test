## 任务接收信息
|字段|类型|必选|注释|
|:----    |:------- |:--- |------      |
|ID               |int64     |是   |唯一id|
|Name			  |string	 |是	  |任务名称|
|Owner            |int64     |是   |用户在我方库中的id|
|OwnerName        |string    |是   |用户名|
|AgencyUserID     |int64     |是   |在经销商库中的id|
|CreateTime		  |int64	 |是	  |任务创建时间|
|Type             |int       |是   | 0: video, 1: picture|
|Status           |int       |是   |渲染状态|
|RenderOutputFiles|[]fileObject |否 |输出文件|
|FrameCount		  |int		 |否	|帧数|
|Gross            |int64     |是   |帧的费用|
|FrameCompleted	  |int		 |否	|完成帧数|
|StartTime		  |int64	 |是   |任务开始渲图时间|
|CompletedTime	  |int64     |是	  |任务完成时间|
|RenderTimeThreadTotal |int64     |否   |线程时|
|LightJobGross    |int64     |是| 小光子任务最终费用|
|LightJobRenderTimeThread|int64|否 | 小光子渲染的线程时|
|RenderZipFile	  |string	 |是	 |任务源文件|
|Camera           |string     |是 | 相机名称|
|Action           |int        |否 | 程序同步的操作 0创建任务,1 更新任务状态,2任务完成并扣费,3任务删除操作,4经销商同步操作

- FileObject结构体
```
type FileObject struct {
	//文件的名称,可用于下载已完成文件,下载完成文件可参考:快渲经销商Api文档-任务已完成文件下载接口
	Name  string `json:"name"`
	//文件大小,不常用
	Size  int64  `json:"size"`
	 //文件时间,基本不用
	Time  int64  `json:"time"`
	 //文件是否为目录,基本不用
	IsDir bool   `json:"isdir"`
}
```
## 渲染帧接收信息
|字段|类型|必填|注释|
|:----    |:-------    |:--- |------|
|ID               |int64        |是|唯一id|
|Owner            |int64        |是|用户在我方库中的id|
|OwnerName        |string       |是|用户名|
|AgencyUserID     |int64        |是|在经销商库中的id|
|RenderTaskID     |int64        |是|所属任务id|
|FrameIndex       |int          |是|帧索引|
|Status           |int          |是|渲染状态|
|RenderTimeThread |int64        |是|线程时|
|Gross            |int64        |是|帧的费用|
|Files            |[]fileObject |是|输出文件|
|Type             |int          |是| 0: video, 1: picture|
|Action           |int          |是|同步数据的操作: 0创建任务,1 更新任务状态,2任务完成并扣费,3任务删除操作|
|LightJobGross           |int64 |是| 小光子任务最终费用|
|LightJobRenderTimeThread|int64 |是| 小光子渲染的线程时|
|LightJobStatus          |int   |是| 小光子任务状态|


## 经销商需要提供API地址
**用途**
用于拼接请求的完整api地址
例如
```
http://kxapi.cloudsx.com/api/task/
http://kxapi.cloudsx.com/api/frame/
```
**简要描述：**

## - 任务创建通知接口

**请求URL：**
- `http://kxapi.cloudsx.com/api/task/createinfo `
  
**请求方式：**
- POST

**参数：**

|字段|类型|必填|注释|
|:----    |:------- |:--- |------      |
|ID               |int64     |是   |唯一id|
|Name			  |string	 |是	  |任务名称|
|Owner            |int64     |是   |用户在我方库中的id|
|OwnerName        |string    |是   |用户名|
|AgencyUserID     |int64     |是   |在经销商库中的id|
|CreateTime		  |int64	 |是	  |任务创建时间|
|Type             |int       |是   | 0: video, 1: picture|
|Status           |int       |是   |渲染状态|
|RenderOutputFiles|[]fileObject |否 |输出文件|
|FrameCount		  |int		 |否	|帧数|
|Gross            |int64     |是   |帧的费用|
|FrameCompleted	  |int		 |否	|完成帧数|
|StartTime		  |int64	 |是   |任务开始渲图时间|
|CompletedTime	  |int64     |是	  |任务完成时间|
|RenderTimeThreadTotal |int64     |否   |线程时|
|LightJobGross    |int64     |是| 小光子任务最终费用|
|LightJobRenderTimeThread|int64|否 | 小光子渲染的线程时|
|RenderZipFile	  |string	 |是	 |任务源文件|
|Camera           |string     |是 | 相机名称|
|Action           |int        |否 | 程序同步的操作 0创建任务,1 更新任务状态,2任务完成并扣费,3任务删除操作,4经销商同步操作

**发送请求示例**
```
{
    "id": 67721,
    "name": "PrimerFrame_Box_modeling",
    "gross": 0,
    "owner": 1601,
    "agencyuserid":3,
	"type":1,
    "status": 0,
    "action": 0,
    "ownername": "test2",
    "starttime": 0,
    "createtime": 1574318124,
    "framecount": 0,
    "completedtime": 0,
    "lightjobgross": 0,
    "renderzipfile": "",
    "rendertimethreadtotal": 0,
    "framecompleted": 0,
    "renderoutputfiles": [],
    "lightjobtimethreadtotal": 0,
	"camera":"Camera01",
}
```
 **返回示例**
```
  {
    "error_code": 0,
    "error_msg": "成功",
  }
```

 **返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|error_code |int   |成功状态，0：成功；1：失败  |
|error_nsg  |string  |成功,失败

**简要描述：** 

## - 任务更新通知接口
所有任务的状态变更都是用此接口接收,我方直接传送状态变更后的所有任务状态信息,
- 其中包含的任务更新内容如下:
status 1:	渲染中
status 2:	任务完成
status 3:	任务暂停
status 4:	任务失败
status 5:	等待其他依赖任务完成后继续任务
status 6:	上传文件中
status 7:	处理上传后的文件,提供任务使用
status 8:	任务队列
**请求URL：** 
- `http://kxapi.cloudsx.com/api/task/updateinfo`
  
**请求方式：**
- POST 

**参数：** 

|字段|类型|必填|注释|
|:----    |:------- |:--- |------      |
|ID               |int64     |是   |唯一id|
|Name			  |string	 |是	  |任务名称|
|Owner            |int64     |是   |用户在我方库中的id|
|OwnerName        |string    |是   |用户名|
|AgencyUserID     |int64     |是   |在经销商库中的id|
|CreateTime		  |int64	 |是	  |任务创建时间|
|Type             |int       |是   | 0: video, 1: picture|
|Status           |int       |是   |渲染状态|
|RenderOutputFiles|[]fileObject |否 |输出文件|
|FrameCount		  |int		 |否	|帧数|
|Gross            |int64     |是   |帧的费用|
|FrameCompleted	  |int		 |否	|完成帧数|
|StartTime		  |int64	 |是   |任务开始渲图时间|
|CompletedTime	  |int64     |是	  |任务完成时间|
|RenderTimeThreadTotal |int64     |否   |线程时|
|LightJobGross    |int64     |是| 小光子任务最终费用|
|LightJobRenderTimeThread|int64|否 | 小光子渲染的线程时|
|RenderZipFile	  |string	 |是	 |任务源文件|
|Camera           |string     |是 | 相机名称|
|Action           |int        |否 | 程序同步的操作 0创建任务,1 更新任务状态,2任务完成并扣费,3任务删除操作,4经销商同步操作

**发送请求示例**
```
{
    "id": 67721,
    "name": "PrimerFrame_Box_modeling",
    "gross": 38,
    "owner": 1601,
    "agencyuserid":3,
	"type":1,
    "status": 2,
    "action": 1,
    "ownername": "test2",
    "starttime": 1574318212,
    "createtime": 1574318124,
    "framecount": 2,
    "completedtime": 1574318325,
    "lightjobgross": 2,
    "renderzipfile": "/2019-11-21/751278c5-a13b-4371-a4fc-2d955d9d3b02.zip",
    "rendertimethreadtotal": 10512,
    "framecompleted": 2,
    "renderoutputfiles": [
        {
            "name": "/2019-11-21/67721-PrimerFrame_Box_modeling-frames/000_kx_vray_framebuf.Alpha.0000.TGA",
            "size": 55477,
            "time": 0,
            "isdir": False
        },
        {
            "name": "/2019-11-21/67721-PrimerFrame_Box_modeling-frames/000_kx_vray_framebuf.RGB_color.0000.TGA",
            "size": 55477,
            "time": 0,
            "isdir": False
        },
        {
            "name": "/2019-11-21/67721-PrimerFrame_Box_modeling-frames/A-0000.TGA",
            "size": 55477,
            "time": 0,
            "isdir": False
        }
    ],
    "lightjobtimethreadtotal": 480,
	"camera":"Camera01",
}
```
 **返回示例**
``` 
  {
    "error_code": 0,
    "error_msg": "成功",
  }
```

 **返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|error_code |int   |成功状态，0：成功；1：失败  |
|error_nsg  |string  |成功,失败

**简要描述：** 

## - 任务删除通知接口
用户或快渲后台管理员删除任务时,将删除的任务状态同步到经销商
**请求URL：** 
- `http://kxapi.cloudsx.com/api/task/deleteinfo`
  
**请求方式：**
- POST 

**参数：** 

|字段|类型|必填|注释|
|:----    |:------- |:--- |------      |
|ID               |int64     |是   |唯一id|
|Name			  |string	 |是	  |任务名称|
|Owner            |int64     |是   |用户在我方库中的id|
|OwnerName        |string    |是   |用户名|
|AgencyUserID     |int64     |是   |在经销商库中的id|
|CreateTime		  |int64	 |是	  |任务创建时间|
|Type             |int       |是   | 0: video, 1: picture|
|Status           |int       |是   |渲染状态|
|RenderOutputFiles|[]fileObject |否 |输出文件|
|FrameCount		  |int		 |否	|帧数|
|Gross            |int64     |是   |帧的费用|
|FrameCompleted	  |int		 |否	|完成帧数|
|StartTime		  |int64	 |是   |任务开始渲图时间|
|CompletedTime	  |int64     |是	  |任务完成时间|
|RenderTimeThreadTotal |int64     |否   |线程时|
|LightJobGross    |int64     |是| 小光子任务最终费用|
|LightJobRenderTimeThread|int64|否 | 小光子渲染的线程时|
|RenderZipFile	  |string	 |是	 |任务源文件|
|Camera           |string     |是 | 相机名称|
|Action           |int        |否 | 程序同步的操作 0创建任务,1 更新任务状态,2任务完成并扣费,3任务删除操作,4经销商同步操作

**发送请求示例**
```
{
    "id": 67721,
    "name": "PrimerFrame_Box_modeling",
    "gross": 38,
    "owner": 1601,
    "agencyuserid":3,
	"type":1,
    "status": 2,
    "action": 3,
    "ownername": "test2",
    "starttime": 1574318212,
    "createtime": 1574318124,
    "framecount": 2,
    "completedtime": 1574318325,
    "lightjobgross": 2,
    "renderzipfile": "/2019-11-21/751278c5-a13b-4371-a4fc-2d955d9d3b02.zip",
    "rendertimethreadtotal": 10512,
    "framecompleted": 2,
    "renderoutputfiles": [
        {
            "name": "/2019-11-21/67721-PrimerFrame_Box_modeling-frames/000_kx_vray_framebuf.Alpha.0000.TGA",
            "size": 55477,
            "time": 0,
            "isdir": False
        },
        {
            "name": "/2019-11-21/67721-PrimerFrame_Box_modeling-frames/000_kx_vray_framebuf.RGB_color.0000.TGA",
            "size": 55477,
            "time": 0,
            "isdir": False
        },
        {
            "name": "/2019-11-21/67721-PrimerFrame_Box_modeling-frames/A-0000.TGA",
            "size": 55477,
            "time": 0,
            "isdir": False
        }
    ],
    "lightjobtimethreadtotal": 480,
	"camera":"Camera01",
}
```
 **返回示例**
``` 
  {
    "error_code": 0,
    "error_msg": "成功",
  }
```

 **返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|error_code |int   |成功状态，0：成功；1：失败  |
|error_msg  |string  |成功,失败

## - 任务扣费通知接口
用户或快渲后台管理员删除任务时,将删除的任务状态同步到经销商
**请求URL：** 
- `http://kxapi.cloudsx.com/api/task/paymentinfo`
  
**请求方式：**
- POST 

**参数：** 

|字段|类型|必填|注释|
|:----    |:------- |:--- |------      |
|ID               |int64     |是   |唯一id|
|Owner            |int64     |是   |用户在我方库中的id|
|OwnerName        |string    |是   |用户名|
|AgencyUserID     |int64     |是   |在经销商库中的id|
|CreateTime		  |int64	 |是	|任务创建时间|
|RenderTaskID     |int64     |是   |所属任务id|
|FrameCount		  |int		 |否	|帧数|
|FrameIndex       |int       |否   |帧索引|
|Status           |int       |是   |渲染状态|
|RenderTimeThread |int64     |否   |线程时|
|Gross            |int64     |是   |帧的费用|
|RenderOutputFiles|[]fileObject |否 |输出文件|
|Type             |int       |是   | 0: video, 1: picture|
|Deleted          |int       |是   | 1： 任务已经删除|
|FrameCompleted	  |int		 |否	|完成帧数|
|LightJobGross    |int64     |是| 小光子任务最终费用|
|LightJobRenderTimeThread|int64|否 | 小光子渲染的线程时|
|LightJobStatus          |int  |是 | 小光子任务状态|
|Camera           |string     |是 | 相机名称|
|Action           |int        |否 | 程序同步的操作 0创建任务,1 更新任务状态,2任务完成并扣费,3任务删除操作,4经销商同步操作

**发送请求示例**
```
{
    "id": 67458,
    "name": "",
    "gross": 1000,
    "owner": 0,
    "agencyuserid":1,
	"type":1,
    "status": 0,
    "ownername": "",
    "createtime": 0,
    "framecount": 0,
    "lightjobgross": 0,
    "rendertimethreadtotal": 0,
    "framecompleted": 0,
    "lightjobtimethreadtotal": 0,
	"camera":"Camera01",
}
```
 **返回示例**
``` 
  {
    "price": 1010,
    "balance": "1000",
  }
```

 **返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|----- |
|price |int64   |云尚渲扣除用户的费用  |
|balance  |int64  |云尚渲用户余额|


## - 渲染帧创建通知接口

**请求URL：**
- `http://kxapi.cloudsx.com/api/task/createrenderframeinfo `
  
**请求方式：**
- POST

**参数：**

|字段|类型|必填|注释|
|:----    |:------- |:--- |------      |
|ID               |int64        |是|唯一id|
|Owner            |int64        |是|用户在我方库中的id|
|OwnerName        |string       |是|用户名|
|AgencyUserID     |int64        |是|在经销商库中的id|
|RenderTaskID     |int64        |是|所属任务id|
|FrameIndex       |int          |是|帧索引|
|Status           |int          |是|渲染状态|
|RenderTimeThread |int64        |是|线程时|
|Gross            |int64        |是|帧的费用|
|Files            |[]fileObject |否|输出文件|
|Type             |int          |是| 0: video, 1: picture|
|Action           |int          |是|同步数据的操作: 0创建任务,1 更新任务状态,2任务完成并扣费,3任务删除操作|
|LightJobGross           |int64 |是| 小光子任务最终费用|
|LightJobRenderTimeThread|int64 |是| 小光子渲染的线程时|
|LightJobStatus          |int   |是| 小光子任务状态|

**发送请求示例**
```
{
	"action": 1,
	"agencyuserid": 0,
	"files": [],
	"frameindex": 0,
	"gross": 0,
	"id": 295718,
	"lightjobgross": 0,
	"lightjobrendertimethread": 0,
	"lightjobstatus": 0,
	"owner": 0,
	"ownername": "18330748166",
	"rendertaskid": 67678,
	"rendertimethread": 0,
	"status": 0,
	"type": 0
}
```
 **返回示例**
```
  {
    "error_code": 0,
    "error_msg": "成功",
  }
```
 **返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|error_code |int   |成功状态，0：成功；1：失败  |
|error_msg  |string  |成功,失败
## - 渲染帧更新通知接口

**请求URL：**
- `http://kxapi.cloudsx.com/api/task/updaterenderframeinfo `
  
**请求方式：**
- POST

**参数：**

|字段|类型|必填|注释|
|:----    |:------- |:--- |------      |
|ID               |int64        |是|唯一id|
|Owner            |int64        |是|用户在我方库中的id|
|OwnerName        |string       |是|用户名|
|AgencyUserID     |int64        |是|在经销商库中的id|
|RenderTaskID     |int64        |是|所属任务id|
|FrameIndex       |int          |是|帧索引|
|Status           |int          |是|渲染状态|
|RenderTimeThread |int64        |是|线程时|
|Gross            |int64        |是|帧的费用|
|Files            |[]fileObject |否|输出文件|
|Type             |int          |是| 0: video, 1: picture|
|Action           |int          |是|同步数据的操作: 0创建任务,1 更新任务状态,2任务完成并扣费,3任务删除操作|
|LightJobGross           |int64 |是| 小光子任务最终费用|
|LightJobRenderTimeThread|int64 |是| 小光子渲染的线程时|
|LightJobStatus          |int   |是| 小光子任务状态|

**发送请求示例**
```
{
	"action": 4,
	"agencyuserid": 0,
	"files": [
		{
			"isdir": false,
			"name": "/2019-11-15/67678-03-Ik-joe--Spline-IK--frames/000_kx_vray_framebuf.Alpha.0002.TGA",
			"size": 12539,
			"time": 0
		},
		{
			"isdir": false,
			"name": "/2019-11-15/67678-03-Ik-joe--Spline-IK--frames/000_kx_vray_framebuf.RGB_color.0002.TGA",
			"size": 673215,
			"time": 0
		},
		{
			"isdir": false,
			"name": "/2019-11-15/67678-03-Ik-joe--Spline-IK--frames/A-0002.TGA",
			"size": 672352,
			"time": 0
		}
            ],
	"frameindex": 2,
	"gross": 2,
	"id": 295718,
	"lightjobgross": 0,
	"lightjobrendertimethread": 0,
	"lightjobstatus": 0,
	"owner": 0,
	"ownername": "18330748166",
	"rendertaskid": 67678,
	"rendertimethread": 288,
	"status": 5,
	"type": 0
},
```
 **返回示例**
```
  {
    "error_code": 0,
    "error_msg": "成功",
  }
```
 **返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|error_code |int   |成功状态，0：成功；1：失败  |
|error_msg  |string  |成功,失败