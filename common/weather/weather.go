// Package weather 定义了对小米天气的请求和解析接口
package weather

import (
	"fmt"
	"strconv"
	"strings"
)

// 天气
type WeatherResp struct {
	Current        *Current        `json:"current"`        // 当前天气预报
	ForecastDaily  *ForecastDaily  `json:"forecastDaily"`  // 未来15日天气预报(含今日)
	ForecastHourly *ForecastHourly `json:"forecastHourly"` // 未来24h天气预报(不含当前小时)
	AQI            *AQI            `json:"aqi"`            // 当前空气质量

	errCode string
	errDesc string
}

// 当前天气
type Current struct {
	Weather     string         `json:"weather"`     // 天气码，对应信息见constant.go
	Temperature *ModuleCurrent `json:"temperature"` // 温度
	Humidity    *ModuleCurrent `json:"humidity"`    // 湿度
	Pressure    *ModuleCurrent `json:"pressure"`    // 气压
	PubTime     string         `json:"pubTime"`     // 发布时间
}

/**
【天气预报】
明天阴转小雨
气温3°C到12°C
东南风3级，湿度 64%
❄️天气寒冷，注意防寒❄️
各位工友出门注意保暖

【天气预报】
明天小雨
气温7°C到14°C
北风1级，湿度 84%
各位工友出门注意防雨防滑
*/

func (c *Current) String() string {
	return fmt.Sprintf("【天气预报】\n当前天气: %s\n当前温度: %s%s\n当前湿度: %s%s\n当前气压: %s%s\n",
		GetWeatherCodeDesc(c.Weather),
		c.Temperature.Value, c.Temperature.Unit,
		c.Humidity.Value, c.Humidity.Unit,
		c.Pressure.Value, c.Pressure.Unit,
		//strings.Join(strings.Split(strings.TrimSuffix(c.PubTime, "+08:00"), "T"), " "),
	)
}

// 当前天气的模块，包含体感、湿度、气压、温度
type ModuleCurrent struct {
	Unit  string // 单位
	Value string // 值
}

// 15天天气预报(含今日)
type ForecastDaily struct {
	PubTime     string       `json:"pubTime"`     // 发布时间
	Temperature *ModuleDaily `json:"temperature"` // 15天天气预报的温度模块
	Weather     *ModuleDaily `json:"weather"`     // 15天天气预报的天气模块
	Wind        *Wind        `json:"wind"`        // 15 天天气预报的风信息模块
}

type Wind struct {
	Direction *ModuleDaily `json:"direction"` // 15天风向模块
	Speed     *ModuleDaily `json:"speed"`     // 15天风速模块
}

// 15天天气预报的模块，包含温度和天气
type ModuleDaily struct {
	Unit  string                      `json:"unit"`  // 单位
	Value []Pair `json:"value"` // 值，一般有15个
}

type Pair struct {
	From, To string
}

// 未来24小时天气预报(不含当前小时)
type ForecastHourly struct {
	Temperature *ModuleHourly `json:"temperature"` // 未来24小时天气预报的温度模块
	Weather     *ModuleHourly `json:"weather"`     // 未来24小时天气预报的天气模块

	AQI *ModuleHourly `json:"aqi"` // 未来24小时天气预报的AQI模块
}

// 未来24小时天气预报的模块，包含温度、天气
type ModuleHourly struct {
	PubTime string `json:"pubTime"` // 发布时间
	Value   []int  `json:"value"`   // 23个值
}

// 当前的空气质量
type AQI struct {
	Aqi     string `json:"aqi"`     // 空气质量
	CO      string `json:"co"`      // 一氧化碳
	NO2     string `json:"no2"`     // 二氧化氮
	O3      string `json:"o3"`      // 臭氧
	PM10    string `json:"pm10"`    // 10微米以下可吸入颗粒
	PM25    string `json:"pm25"`    // 25微米以下可吸入颗粒
	SO2     string `json:"so2"`     // 二氧化硫
	PubTime string `json:"pubTime"` // 发布时间
	Suggest string `json:"suggest"` // 发布时间
}

// func (w *WeatherResp) GetTomorrowWeatherInfoAttachCare() string {
// 	// 关心工友
// 	const (
// 		rainWorkerDesc = "请注意防雨防滑\n"
// 		coldWorkerDesc = "❄️天气寒冷，注意防寒❄️\n各位工友出门注意保暖"
// 	)

// 	const (
// 		MAX_TEMP_THRESHOLD = 15
// 		MIN_TEMP_THRESHOLD = 10
// 	)

// 	tip := w.GetTomorrowWeatherInfo()

// 	forcastSet := initTomorrowForcastSet(w)
// 	tRainCare := ""
// 	for _, weather := range forcastSet {
// 		iweather, err := strconv.Atoi(weather)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if (3 <= iweather && iweather <= 12) || (21 <= iweather && iweather <= 28) {
// 			tRainCare = rainWorkerDesc
// 		}

// 	}

// 	// temperature
// 	tempSet := initTomorrowTempSet(w)
// 	maxTemp, err := strconv.Atoi(tempSet[1])
// 	if err != nil {
// 		panic(err)
// 	}

// 	minTemp, err := strconv.Atoi(tempSet[0])
// 	if err != nil {
// 		panic(err)
// 	}

// 	tTempCare := ""
// 	if maxTemp < MAX_TEMP_THRESHOLD || minTemp < MIN_TEMP_THRESHOLD {
// 		tTempCare = coldWorkerDesc
// 	}

// 	// care := ""
// 	// if len(tTempCare) > 0 && len(tRainCare) > 0 {
// 	// 	care = coldWorkerDesc
// 	// }

// 	return tip + tRainCare + tTempCare

// }

func initTomorrowTempSet(w *WeatherResp) []string {
	ret := make([]string, 2)
	minTemp := w.ForecastDaily.Temperature.Value[1].To
	maxTemp := w.ForecastDaily.Temperature.Value[1].From

	ret[0] = minTemp
	ret[1] = maxTemp
	return ret
}

func initTomorrowForcastSet(w *WeatherResp) []string {
	slice := make([]string, 2)
	from := w.ForecastDaily.Weather.Value[1].From
	to := w.ForecastDaily.Weather.Value[1].To

	fmt.Println("FROM:", from, to)
	slice[0] = from
	slice[1] = to
	return slice
}

// 获取天气信息中的当前天气信息
func (w *WeatherResp) GetTomorrowWeatherInfo() string {
	weatherTomorrow := w.ForecastDaily.Weather.Value[1]
	tempTomorrow := w.ForecastDaily.Temperature.Value[1]
	tempUnit := w.ForecastDaily.Temperature.Unit
	windDireTomorrow := w.ForecastDaily.Wind.Direction.Value[1]
	windSpeedTomorrow := w.ForecastDaily.Wind.Speed.Value[1]
	// windSpeedUnit := w.ForecastDaily.Wind.Speed.Unit

	
	avg:= convStrAndGetAvgFloat(windSpeedTomorrow.From, windSpeedTomorrow.To)
	windOp := windOptional(&windDireTomorrow, GetWindLevel(avg))

	weather1 := GetWeatherCodeDesc(weatherTomorrow.From)
	weather2 := GetWeatherCodeDesc(weatherTomorrow.To)
	return fmt.Sprintf(
		"【天气预报】\n"+
			"明天%s\n"+
			"%s"+
			"气温：%s到%s%s\n"+
			"%s"+
			"%s",
		GetWeatherStr(weather1, weather2),
		rainCare(w),
		tempTomorrow.To, tempTomorrow.From, tempUnit,
		tempCare(w),
		windOp)
}

func windOptional(wind *Pair, i int) string {
	s1 := GetWindDesc(wind.From)
	s2 := GetWindDesc(wind.To)
	if(s1 == s2) {
		return fmt.Sprintf("风向：%s风，风速%d级", s1, i)
	} else {
		return fmt.Sprintf("风向：%s风转%s风，风速%d级", s1, s2, i)
	}
}

func convStrAndGetAvgFloat(s1 string, s2 string) string {
	i1, err := strconv.ParseFloat(s1, 32)
	if err != nil {
		panic(err)
	}

	i2, err := strconv.ParseFloat(s2, 32)
	if err != nil {
		panic(err)
	}

	avg := (i1 + i2) / 2

	return fmt.Sprintf("%f", avg)

}

func tempCare(w *WeatherResp) string {

	const coldWorkerDesc = "❄️天气寒冷，注意防寒❄️\n各位工友出门注意保暖\n"
	const (
		MAX_TEMP_THRESHOLD = 8
		MIN_TEMP_THRESHOLD = 5
	)

	tempSet := initTomorrowTempSet(w)
	maxTemp, err := strconv.Atoi(tempSet[1])
	if err != nil {
		panic(err)
	}

	minTemp, err := strconv.Atoi(tempSet[0])
	if err != nil {
		panic(err)
	}

	tTempCare := ""
	if maxTemp < MAX_TEMP_THRESHOLD || minTemp < MIN_TEMP_THRESHOLD {
		tTempCare = coldWorkerDesc
	}

	return tTempCare
}

func rainCare(w *WeatherResp) string {
	const rainWorkerDesc = "请注意防雨防滑\n"
	// temperature
	forcastSet := initTomorrowForcastSet(w)
	tRainCare := ""
	for _, weather := range forcastSet {
		iweather, err := strconv.Atoi(weather)
		if err != nil {
			panic(err)
		}

		if (3 <= iweather && iweather <= 12) || (21 <= iweather && iweather <= 28) {
			tRainCare = rainWorkerDesc
		}

	}

	return tRainCare
}

func GetWindLevel(windSpeedTomorrow string) int {

	speed, err := strconv.ParseFloat(windSpeedTomorrow, 32)
	if err != nil {
		panic(err)
	}

	return windLevel(speed)
}

func windLevel(speed float64) int {
	switch {
	case speed < 1:
		return 0
	case speed < 6:
		return 1
	case speed < 12:
		return 2
	case speed < 20:
		return 3
	case speed < 29:
		return 4
	case speed < 39:
		return 5
	case speed < 50:
		return 6
	case speed < 61:
		return 7
	case speed < 74:
		return 8
	case speed < 88:
		return 9
	case speed < 102:
		return 10
	case speed < 117:
		return 11
	default:
		return 12
	}
}

// north: 0d
// east: 90d
// south: 180d
// west: 270d
// ref: https://dev.qweather.com/docs/resource/wind-info/
func GetWindDesc(s string) string {
	angle, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}

	ret := ""
	// north
	if (0 <= angle && angle <= 22.5) || (337.5 < angle && angle <= 360) {
		ret = "北"
	} else if 22.5 < angle && angle <= 67.5 {
		ret = "东北"
	} else if 67.5 < angle && angle <= 112.5 {
		ret = "东"
	} else if 112.5 < angle && angle <= 157.5 {
		ret = "东南"
	} else if 157.5 < angle && angle <= 202.5 {
		ret = "南"
	} else if 202.5 < angle && angle <= 247.5 {
		ret = "西南"
	} else if 247.5 < angle && angle <= 292.5 {
		ret = "西"
	} else if 292.5 < angle && angle <= 337.5 {
		ret = "西北"
	}

	return ret

}

// 获取天气信息中的当前天气信息
func (w *WeatherResp) GetCurrentWeatherInfo() string {
	next1 := strconv.Itoa(w.ForecastHourly.Weather.Value[0])
	next2 := strconv.Itoa(w.ForecastHourly.Weather.Value[1])
	next3 := strconv.Itoa(w.ForecastHourly.Weather.Value[2])
	weather1  := GetWeatherCodeDesc(w.ForecastDaily.Weather.Value[0].From)
	weather2 := GetWeatherCodeDesc(w.ForecastDaily.Weather.Value[0].To)
	return fmt.Sprintf(
		"【天气预报】\n"+
			"当前天气: %s\n"+
			"当前温度: %s%s\n"+
			"当前湿度%s%s\n"+
			"当前空气质量: %s\n"+
			"预期未来三小时天气: %s, %s, %s\n"+
			"今日气温: %s到%s\n"+
			"今日天气预期: %s\n",
			
		//"本次数据更新时间: %s",
		GetWeatherCodeDesc(w.Current.Weather),
		w.Current.Temperature.Value, w.Current.Temperature.Unit,
		w.Current.Humidity.Value, w.Current.Humidity.Unit,
		w.AQI.Aqi,
		GetWeatherCodeDesc(next1), GetWeatherCodeDesc(next2), GetWeatherCodeDesc(next3),
		w.ForecastDaily.Temperature.Value[0].To, w.ForecastDaily.Temperature.Value[0].From,
		GetWeatherStr(weather1, weather2),
		//strings.Join(strings.Split(strings.TrimSuffix(w.Current.PubTime, "+08:00"), "T"), " "),

	)
}

func GetWeatherStr(weather1, weather2 string) string {
	if weather1 == weather2 {
		return weather1
	} else {
		return weather1+"转"+ weather2 
	}
}

// 获取天气信息中的AQI空气质量信息
func (w *WeatherResp) GetAQIInfo() string {
	return fmt.Sprintf("当前空气质量: %s %s\n"+
		"PM2.5细颗粒物: %sμg/m³\n"+
		"PM10可吸入颗粒物: %sμg/m³\n"+
		"SO2二氧化硫: %sμg/m³\n"+
		"NO2二氧化氮: %sμg/m³\n"+
		"O3臭氧: %sμg/m³\n"+
		"CO一氧化碳: %smg/m³\n"+
		"本次数据更新时间: %s",
		w.AQI.Aqi, GetAQIQuality(w.AQI.Aqi),
		w.AQI.PM25,
		w.AQI.PM10,
		w.AQI.SO2,
		w.AQI.NO2,
		w.AQI.O3,
		w.AQI.CO,
		strings.Join(strings.Split(strings.TrimSuffix(w.Current.PubTime, "+08:00"), "T"), " "),
	)
}

// 获取AQI空气质量指标的描述
func AQIIndicesDesc() string {
	return strings.Join([]string{pm25Desc, pm10Desc, so2Desc, no2Desc, o3Desc, coDesc}, "\n")
}

// 模糊查询城市接口响应结果
type CityLikeResp struct {
	Data    map[string]string `json:"data"` // "城市id" : "省份, 城市, 区/县" 为 kv 的数据结果
	Message string            // 响应消息
	Status  int               // 响应状态码
}

// 获取各个城市及其对应的城市 id 映射表
func (r *CityLikeResp) GetCityLike() map[string]string {
	if len(r.Data) == 0 {
		return r.Data
	}
	reversedMap := make(map[string]string, len(r.Data))
	for k, v := range r.Data {
		v = strings.Join(strings.Split(v, ", "), "")
		reversedMap[v] = k
	}
	return reversedMap
}
