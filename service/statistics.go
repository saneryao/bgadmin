package service

import (
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"math/rand"
	"strconv"
	"time"
)

/* 功能：查询统计量（用户数、角色数、资源数、文章数、评论数）
 * 参数：data 表示查询到的数据集，other表示其他的额外参数
 * 返回值：err表示查询失败时的错误信息
 */
func GetAmounts(data, others map[string]interface{}) (err error) {
	// 定义变量
	var totalUsers int64
	var totalRoles int64
	var totalResources int64
	var totalArticles int64
	var totalComments int64

	// 查询用户总数
	db := orm.NewOrm()
	qsUser := db.QueryTable(new(models.User))
	if totalUsers, err = qsUser.Count(); err != nil {
		return
	}

	// 查询角色总数
	qsRole := db.QueryTable(new(models.Role))
	if totalRoles, err = qsRole.Count(); err != nil {
		return
	}

	// 查询资源总数
	qsResource := db.QueryTable(new(models.Menu))
	if totalResources, err = qsResource.Count(); err != nil {
		return
	}

	// // 查询文章总数
	// qsArticle := db.QueryTable(new(models.Article))
	// if totalArticles, this.Error = qsArticle.Count(); err != nil {
	// 	return
	// }

	// // 查询评论总数
	// qsComment := db.QueryTable(new(models.Comment))
	// if totalComments, this.Error = qsComment.Count(); err != nil {
	// 	return
	// }

	// 封装返回数据
	if data != nil {
		data["totalUsers"] = totalUsers
		data["totalRoles"] = totalRoles
		data["totalResources"] = totalResources
		data["totalArticles"] = totalArticles
		data["totalComments"] = totalComments
	}
	return
}

/* 功能：查询设备信息（CPU使用率、内存使用率、硬盘使用率）
 * 参数：data 表示查询到的数据集，other表示其他的额外参数
 * 返回值：err表示查询失败时的错误信息
 */
func GetDevices(data, others map[string]interface{}) (err error) {
	// 定义变量
	var cpuUsed int
	var memoryUsed int
	var diskUsed int

	// 查询CPU使用率
	cpuUsed = 21

	// 查询内存使用率
	memoryUsed = 35

	// 查询硬盘使用率
	diskUsed = 42

	// 封装返回数据
	if data != nil {
		data["cpuUsed"] = cpuUsed
		data["memoryUsed"] = memoryUsed
		data["diskUsed"] = diskUsed
	}
	return
}

/* 功能：查询文章分布情况（各顶级分类下的文章总数）
 * 参数：data 表示查询到的数据集，other表示其他的额外参数
 * 返回值：err表示查询失败时的错误信息
 */
func GetDistributions(data, others map[string]interface{}) (err error) {
	// 定义变量
	var totalFavorites int64
	var totalPraises int64

	// 查询各分类下的文章数量
	if data != nil {
		rand.Seed(time.Now().UnixNano())
		for i := 1; i <= 5; i++ {
			if i == 5 {
				data["others"] = rand.Intn(10000)
			} else {
				data["Group"+strconv.Itoa(i)] = rand.Intn(10000)
			}
		}
	}

	// 查询收藏数
	totalFavorites = rand.Int63n(10000)

	// 查询点赞数
	totalPraises = rand.Int63n(10000)

	// 封装返回数据
	if others != nil {
		others["totalFavorites"] = totalFavorites
		others["totalPraises"] = totalPraises
	}
	return
}

/* 功能：查询近30天的访问量
 * 参数：data 表示查询到的数据集，other表示其他的额外参数
 * 返回值：err表示查询失败时的错误信息
 */
func GetHits(data, others map[string]interface{}) (err error) {
	// 查询访问量
	if data != nil {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 30; i++ {
			data[strconv.Itoa(i)] = rand.Intn(10000)
		}
	}
	return
}
