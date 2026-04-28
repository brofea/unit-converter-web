package service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Value float64 `json:"value" binding:"required"`
	From  string  `json:"from" binding:"required"`
	To    string  `json:"to" binding:"required"`
}

func Convert(c *gin.Context) {
	// 解析请求参数
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 执行转换逻辑
	result, err := performConversion(req.Value, req.From, req.To)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"original_value": req.Value,
		"from":           req.From,
		"to":             req.To,
		"result":         result,
	})
}
func performConversion(val float64, from, to string) (float64, error) {
	valtype := map[string]string{
		// 长度
		"m": "length", "cm": "length", "mm": "length", "km": "length",
		"inch": "length", "ft": "length", "yd": "length", "mi": "length",

		// 质量
		"kg": "weight", "g": "weight", "mg": "weight", "lb": "weight", "oz": "weight",

		// 面积
		"m2": "area", "cm2": "area", "mm2": "area", "km2": "area",
		"ac": "area", "ha": "area", "sqft": "area", "sqin": "area",

		// 体积
		"m3": "volume", "l": "volume", "ml": "volume", "gal": "volume", "qt": "volume",

		// 时间
		"s": "time", "min": "time", "h": "time", "d": "time",

		// 温度
		"cel": "temperature", "degf": "temperature", "kel": "temperature",
	}
	toBase := map[string]float64{
		// --- 长度 (Base: m) ---
		"m":    1.0,
		"cm":   0.01,
		"mm":   0.001,
		"km":   1000.0,
		"inch": 0.0254,
		"ft":   0.3048,
		"yd":   0.9144,
		"mi":   1609.344,

		// --- 质量 (Base: kg) ---
		"kg": 1.0,
		"g":  0.001,
		"mg": 0.000001,
		"lb": 0.45359237,
		"oz": 0.02834952,

		// --- 面积 (Base: m2) ---
		"m2":   1.0,
		"cm2":  0.0001,
		"mm2":  0.000001,
		"km2":  1000000.0,
		"ha":   10000.0,    // 公顷
		"ac":   4046.856,   // 英亩
		"sqft": 0.092903,   // 平方英尺
		"sqin": 0.00064516, // 平方英寸

		// --- 体积 (Base: m3) ---
		"m3":  1.0,
		"l":   0.001,      // 升
		"ml":  0.000001,   // 毫升
		"gal": 0.00378541, // 美制加仑
		"qt":  0.00094635, // 美制夸脱

		// --- 时间 (Base: s) ---
		"s":   1.0,
		"min": 60.0,
		"h":   3600.0,
		"d":   86400.0,
	}
	if valtype[from] != valtype[to] {
		return 0, errors.New("incompatible unit types")
	}
	switch valtype[from] {
	case "temperature":
		// 统一转换为开尔文
		var tempK float64
		switch from {
		case "cel":
			tempK = val + 273.15
		case "degf":
			tempK = (val-32)*5/9 + 273.15
		case "kel":
			tempK = val
		}
		// 转换为目标单位
		switch to {
		case "cel":
			return tempK - 273.15, nil
		case "degf":
			return (tempK-273.15)*9/5 + 32, nil
		case "kel":
			return tempK, nil
		}
	default:
		var baseVal, result float64
		baseVal = val * toBase[from]
		result = baseVal / toBase[to]
		return result, nil
	}
	return 0, errors.New("unsupported conversion")
}

func Exchange(c *gin.Context) {
}
