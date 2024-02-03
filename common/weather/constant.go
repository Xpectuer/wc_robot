package weather

import (
	"fmt"
	"strconv"
)

// AQI污染物描述
const (
	pm25Desc = "PM2.5指的是直径小于或等于2.5微米的颗粒物, 又称为细颗粒物"
	pm10Desc = "PM10指可吸入颗粒, 主要来源是建筑活动和从地表扬起的尘土, 含有氧化物矿物和其他成分"
	so2Desc  = "SO2指二氧化硫, 人为主要来源为家庭取暖, 发电和机动车而燃烧含有硫磺的矿物燃料，以及对含有硫磺的矿物的冶炼"
	no2Desc  = "NO2指二氧化氮, 短期浓度超过200微克/立方米时，是一种引起呼吸道严重发炎的有毒气体"
	o3Desc   = "O3指臭氧, 地面的臭氧主要由车辆和工业释放出的氧化氮等污染物以及由机动车、溶剂和工业释放的挥发性有机化合物与阳光反应而生成"
	coDesc   = "CO指一氧化碳, 八成来自汽车尾气, 交通高峰期时, 公路沿线产生的CO浓度会高于平常"
)

const (
	wSUNNY      = "0"
	wCLOUDY     = "1"
	wOVERCAST   = "2"
	wSHOWER     = "3"
	wTHUNDER    = "4"
	wLIGHTRAIN  = "7"
	wMEDIUMRAIN = "8"
	wHEAVYRAIN  = "9"
)

var mapWeatherDesc = map[string]string{
	"0":  "晴",
	"1":  "多云",
	"2":  "阴",
	"3":  "阵雨",
	"4":  "雷阵雨",
	"5":  "雷阵雨并伴有冰雹",
	"6":  "雨夹雪",
	"7":  "小雨",
	"8":  "中雨",
	"9":  "大雨",
	"10": "暴雨",
	"11": "大暴雨",
	"12": "特大暴雨",
	"13": "阵雪",
	"14": "小雪",
	"15": "中雪",
	"16": "大雪",
	"17": "暴雪",
	"18": "雾",
	"19": "冻雨",
	"20": "沙尘暴",
	"21": "小雨-中雨",
	"22": "中雨-大雨",
	"23": "大雨-暴雨",
	"24": "暴雨-大暴雨",
	"25": "大暴雨-特大暴雨",
	"26": "小雪-中雪",
	"27": "中雪-大雪",
	"28": "大雪-暴雪",
	"29": "浮沉",
	"30": "扬沙",
	"31": "强沙尘暴",
	"32": "飑",
	"33": "龙卷风",
	"34": "若高吹雪",
	"35": "轻雾",
	"53": "霾",
	"99": "未知",
}

// GetWeatherCodeDesc 根据天气码返回对应的天气信息
func GetWeatherCodeDesc(weather string) string {
	if desc, exist := mapWeatherDesc[weather]; exist {
		return desc
	}
	return fmt.Sprintf("未知的天气码: %s", weather)
}

// 返回空气质量级别
func GetAQIQuality(aqi string) string {
	a, _ := strconv.Atoi(aqi)
	if a < 51 {
		return "优"
	}
	if a < 101 {
		return "良"
	}
	if a < 151 {
		return "轻度污染"
	}
	if a < 201 {
		return "中度污染"
	}
	if a < 301 {
		return "重度污染"
	}
	return "严重污染"
}
