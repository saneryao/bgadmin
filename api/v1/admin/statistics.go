package admin

import (
	"errors"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/validators"
	"strings"
)

// 定义一个数量统计API（用于获取系统中统计的一些数量）
type StatisticsApi struct {
	BaseApi
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处查询统计的一些数据
 * 参数：空
 * 返回值：空
 */
func (this *StatisticsApi) Get() {
	// 包装并处理返回结果
	data := make(map[string]interface{})
	others := make(map[string]interface{})
	defer func() {
		this.PackResultData(data, others)
	}()

	// 从URL获取Entry
	var strEntry string
	if strEntry, this.Error = validators.ParseEntryFromUrl(&this.Controller, ":entry"); this.Error != nil {
		return
	}

	// 获取数据
	switch {
	case strings.EqualFold(strEntry, "amounts"):
		this.Error = service.GetAmounts(data, nil)
	case strings.EqualFold(strEntry, "devices"):
		this.Error = service.GetDevices(data, nil)
	case strings.EqualFold(strEntry, "distributions"):
		this.Error = service.GetDistributions(data, others)
	case strings.EqualFold(strEntry, "hits"):
		this.Error = service.GetHits(data, nil)
	default:
		this.Error = errors.New(this.Tr("entry_not_found") + strEntry)
	}
}
