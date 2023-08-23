package DailyAlgorithm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/DailyAlgorithm"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"time"
)

// CountDailyAlgorithmRank
// 统计user_name在前30天内的打卡次数，返回一个键值对数组`map[string]int`，每天的打卡最多计算一次
// 关于在定时器中调用，因为是多线程，就要考虑对数据库写入的互斥操作
// 所以下面两个函数必须在对数据库写入时
// 使用读写锁，仅在写入时互斥
//
// 当时只有新开辟的线程会使用这个函数，互斥应该在数据库层面上的互斥
// 这里互斥没有用

func CountDailyAlgorithmRank() (err error) {

	// 定义一个键值对数组，用于存储该月用户的打卡天数
	rank := make(map[string]int)

	// 获取当前日期
	now := time.Now()

	// 循环遍历过去30天的日期
	for i := 0; i < 30; i++ {
		// 计算当前日期的字符串表示
		date := now.AddDate(0, 0, -i).Format("2006-01-02")

		// 创建db
		db := global.GVA_DB.Model(&DailyAlgorithm.DailyAlgorithmRecord{})
		var DARs []DailyAlgorithm.DailyAlgorithmRecord

		db = db.Where("date = ?", date)

		err = db.Find(&DARs).Error

		//返回错误
		if err != nil {
			return nil
		}

		// 定义一个键值对数组，用于该日用户是否打卡
		isExist := make(map[string]bool)
		for _, record := range DARs {
			if isExist[record.User_name] == false {
				rank[record.User_name]++
			}
			isExist[record.User_name] = true
		}
	}

	for name, cnt := range rank {
		system.CoverDACount(name, cnt)
	}

	return nil
}

// RemoveCountedOutDate
// 返回哪些用户需要被-1的map
// - 在有人提交记录时，就给该用户今天的打卡次数+1，直接更新数据库中的数据
// - 每天只需要排除第31天之前的所有数据
// - 成功优化了查询效率！！！
// - 但是为了更具有鲁棒性，之前的函数可以保留，在出现错误时重新调用，可以保证排行榜的正确性
// 统计user_name在前30天内的打卡次数，返回一个键值对数组`map[string]int`，每天的打卡最多计算一次
func RemoveCountedOutDate() (err error) {

	// 获取当前日期
	now := time.Now()
	// 计算当前日期的字符串表示
	date := now.AddDate(0, 0, -31).Format("2006-01-02")

	// 创建db
	db := global.GVA_DB.Model(&DailyAlgorithm.DailyAlgorithmRecord{})
	var DARs []DailyAlgorithm.DailyAlgorithmRecord
	//查找
	db = db.Where("date = ?", date)

	//出错返回空map和错误
	err = db.Find(&DARs).Error
	if err != nil {
		return err
	}

	for _, record := range DARs {
		system.UpdateDACount(record.User_name, -1)
	}
	return nil
}

// Timer
// 定义一个名为timer的函数,新开一个goroutine，检测时间，每天更新每日打卡
func Timer() (err error) {

	//每次运行，调用一次更新30天内的函数，保证数据的可维护性
	err = CountDailyAlgorithmRank()
	if err != nil {
		global.GVA_LOG.Error("DACount Error")
		return err
	}
	global.GVA_LOG.Info("DACount Counted")

	// 获取当前时间
	now := time.Now()

	// 计算距离下一个00:00的持续时间
	nextMidnight := now.Add(time.Duration(24-now.Hour()) * time.Hour)
	nextMidnight = nextMidnight.Add(time.Duration(-now.Minute()) * time.Minute)
	nextMidnight = nextMidnight.Add(time.Duration(-now.Second()) * time.Second)

	// 计算下一个00:00的时间间隔
	interval := nextMidnight.Sub(now)

	// 创建一个Ticker，用于每天定时执行
	ticker := time.NewTicker(interval)

	// 在一个独立的goroutine中处理定时任务
	go func() {
		for {
			<-ticker.C
			global.GVA_LOG.Info("processed counted work")

			err = RemoveCountedOutDate()
			//每次运行报错，就调用获得最新的数据
			if err != nil {
				global.GVA_LOG.Error("process counted work occur to error ")
				CountDailyAlgorithmRank()
			}

			// 更新下一个00:00的时间间隔
			nextMidnight = nextMidnight.Add(24 * time.Hour)
			interval = nextMidnight.Sub(time.Now())
			ticker.Reset(interval)
		}
	}()

	// 让程序保持运行状态
	select {}
}

//多线程就要考虑对数据库写入的互斥操作
//chatgpt帮我写好了互斥在DA的service里面

func DATimer() {
	go Timer()
}
