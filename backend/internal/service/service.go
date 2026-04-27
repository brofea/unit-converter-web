package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type LengthRequest struct {
	Value float64 `json:"value" binding:"required"`
	From  string  `json:"from" binding:"required,oneof=m cm mm inch ft"`
	To    string  `json:"to" binding:"required,oneof=m cm mm inch ft"`
}
func ConvertLength(c *gin.Context) {
	// 解析请求参数
	var req LengthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 执行转换逻辑
	result, err := performLengthConversion(req.Value, req.From, req.To)
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
func performLengthConversion(val float64, from, to string) (float64, error) {
	toMeter := map[string]float64{
		"m":    1.0,
		"cm":   0.01,
		"mm":   0.001,
		"inch": 0.0254,
		"ft":   0.3048,
	}
	meters := val * toMeter[from]
	result := meters / toMeter[to]
	return result, nil
}

