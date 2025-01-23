package cron

import (
	"github.com/robfig/cron/v3"
	"lyp-go/logger"
	"time"
)

var c *cron.Cron

func Init() *cron.Cron {
	c = cron.New()

	// 添加定时任务
	_, err := c.AddFunc("@every 1m", func() {
		logger.Debugf("定时任务执行: %s", time.Now())
	})
	if err != nil {
		logger.Debugf("添加定时任务失败: %v", err)
	}
	// 启动 cron
	c.Start()
	return c
}
