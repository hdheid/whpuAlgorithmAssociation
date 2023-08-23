## 更新内容
### 1.测试自动化包部署
按照作者的视频教程操作

了解了这个项目的架构
### 2.每日算法打卡菜单
写了一个前端界面，以下简称DA

#### 每日打卡
- 前端的css
- vue组件的再利用，把原本弹出的的form提交窗口栏，变成了确认窗口，但是没有改变原有的提交数据的逻辑

#### 本月打卡日历
- 用gpt生成了一个日历，但是不太方便使用，暂时搁置

### 3.排行榜
目的很简单，根据打卡次数排序，做一个榜单，在“用户管理”菜单的基础之上制作。

后续还可以根据用户提供的刷题网站（例如LeetCode，codefores，acwing等）用户名，爬取刷题量，竞赛rating等信息排序

#### 准备环节
把user.vue中的内容完全拷贝到src\view\DailyAlgorithmRankBoard\RankBoard.vue

稍稍出手修改前端界面的内容更像一个排行榜

修改结构体和相关api，（貌似数据库的列信息会自动创建后加上的列，省了）
- server/api/v1/system/sys_user.go
- server/model/system/request/sys_operation_record.go
- server/model/system/sys_user.go
- server/service/system/sys_user.go
添加了一个QQ号码类型

成功在前端中显示，并且可以通过前端发送给后端消息，修改数据库中的内容

#### 统计打卡
在前端页面中中加上了一个打卡次数的列

和上面的内容类似，但是需要统计打卡次数并且做出修改

完全交给后端来操作，每天24点调用函数countDailyAlgorithmRank，统计打卡次数，直接更新数据库，免去了前端操作

- 函数countDailyAlgorithmRank的功能为
- 在dailyalgorithmrecord表中有user_name和date两个字段，统计user_name在30天内的打卡次数，返回一个键值对数组`map[string]int`，每天的打卡最多计算一次
- 先遍历从今天开始的三十天？？？感觉有更好的办法！
- 在有人提交记录时，就给该用户今天的打卡次数+1，直接更新数据库中的数据
- 每天只需要排除第31天之前的所有数据
- 成功优化了查询效率！！！
- 但是为了更具有鲁棒性，之前的函数可以保留，在出现错误时重新调用，可以保证排行榜的正确性

#### 原本数据库date
原本每日打卡记录表中date使用的是datetime类型，包含了时分秒，对数据库的效率有影响

而且不方便更新记录，同一天的打卡记录会被覆盖，包含了时分秒需要特殊判断一下

还有就是数据库中已经自动创建了一个记录创建时间里，date的意义不大

所以就修改了数据库中的date的类型为date，原本date.Time类型也可以自动兼容

为了实现代码的覆盖，需要在路由方面再写一个coverDailyAlgorithmRecord函数，在每次判断edit还是create前调用

#### 前端后端消息传递和调用
**前端**
```vue
  const resCoverRecord = await coverDailyAlgorithmRecord({ date:formData.value.date })
  if (resCoverRecord.code === 0) {
    formData.value = resCoverRecord.data.reDAR
    type.value = 'update'
  }
  // console.log(type.value)


```
在save保存按钮前加上上面这一段代码，以更新type.value，改变数据的格式

调用了coverDailyAlgorithmRecord函数api，花括号内是json格式的数据

在这个示例中，await 关键字用于等待 coverDailyAlgorithmRecord 函数返回一个 Promise 对象的结果。这意味着在接收到结果之前，代码将暂停执行，并等待 Promise 对象的解决（即成功或失败）。

要特别注意的是，前端还设置了一个api调用的权限，需要设置这个api调用的权限

**后端**
```go
// CoverDailyAlgorithmRecord 用date查询DailyAlgorithmRecord
// @Tags DailyAlgorithmRecord
// @Summary  用date查询DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query DailyAlgorithm.DailyAlgorithmRecord true "用date查询DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DAR/findDailyAlgorithmRecord [get]
func (DARApi *DailyAlgorithmRecordApi) CoverDailyAlgorithmRecord(c *gin.Context) {
	var DAR DailyAlgorithm.DailyAlgorithmRecord
	err := c.ShouldBindQuery(&DAR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if reDAR, err := DARService.GetCoverDailyAlgorithmRecord(*(DAR.Date)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDAR": reDAR}, c)
	}
}
```
这是一个用gin写的api，功能是用date查询DailyAlgorithmRecord数据表，和findDailyAlgorithmRecord非常类似

`ShouldBindQuery`函数用于将请求参数解析到结构体中

将调用下面的函数`GetCoverDailyAlgorithmRecord`，用gorm在数据库中查询
```go
// GetCoverDailyAlgorithmRecord 根据date获取DailyAlgorithmRecord记录
// Author [CFDDFC](https://github.com/cfddd)
func (DARService *DailyAlgorithmRecordService) GetCoverDailyAlgorithmRecord(date time.Time) (DAR DailyAlgorithm.DailyAlgorithmRecord, err error) {
	dateString := date.Format("2006-01-02")

	err = global.GVA_DB.Where("date = ?", dateString).First(&DAR).Error
	return
}
```

完成之后通过`response`返回给前端调用

#### 对于修改操作
在执行update之前，需要先find
find找到后，会给formData赋值，并且type赋值为update
接下来的updata操作只传递了需要更新的字段，有点不解
去后台看源码，发现gorm调用的save函数是自动根据根据primary字段查找的
找到则更新,没找到则创建