package admin

import (
	"errors"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/validators"
	"strings"
)

// StatisticsAPI 定义一个数量统计API（用于获取系统中统计的一些数量）
type StatisticsAPI struct {
	baseAPI
}

// Get 执行http请求GET方法（beego定义的接口，查询统计的一些数据）
func (api *StatisticsAPI) Get() {
	// 包装并处理返回结果
	data := make(map[string]interface{})
	others := make(map[string]interface{})
	defer func() {
		api.PackResultData(data, others)
	}()

	// 从URL获取Entry
	var strEntry string
	if strEntry, api.Error = validators.ParseEntryFromURL(&api.Controller, ":entry"); api.Error != nil {
		return
	}

	// 获取数据
	switch {
	case strings.EqualFold(strEntry, "amounts"):
		api.Error = service.GetAmounts(data, nil)
	case strings.EqualFold(strEntry, "devices"):
		api.Error = service.GetDevices(data, nil)
	case strings.EqualFold(strEntry, "distributions"):
		api.Error = service.GetDistributions(data, others)
	case strings.EqualFold(strEntry, "hits"):
		api.Error = service.GetHits(data, nil)
	default:
		api.Error = errors.New(api.Tr("entry_not_found") + strEntry)
	}
}
