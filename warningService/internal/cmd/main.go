package main

import (
	"fmt"
	"log"
	"time"
	"warningService/internal/clients/client"
	"warningService/internal/models"
	"warningService/internal/pkg"
)

func main() {
	for {
		notifyUserForCPU := false
		notifyUserForMEM := false
		// 拿告警配置，若没有告警配置，则不执行后续操作
		alarmRule, alarmRuleCode := client.GetAlarmRule()
		if alarmRuleCode != 200 {
			log.Println("no rules are set.")
			time.Sleep(time.Minute)
			continue
		} else {
			fmt.Println(alarmRule.ContactEmail)
			log.Println("receive rule success, start to get aggregateMsg")
		}
		// 拿到告警配置后，从数据分析模块拿目标聚合值
		Granularity := alarmRule.Granularity
		AggregateFunction := alarmRule.AggregateFunction
		aggregateMsg, aggregateValueCode := client.GetAggregateValue(Granularity, AggregateFunction)

		if aggregateValueCode != 200 {
			log.Println("no aggregate value find.")
			continue
		} else {
			fmt.Println(aggregateMsg)
			log.Println("receive aggregateMsg success, start to compare")
		}
		alarmMsg := models.AlarmMsg{}
		// 对比告警配置和目标聚合值，若目标聚合值>=告警配置的提示阈值，则发送邮件通知用户，邮件内容根据超过阈值的水平而定
		for i := 0; i < len(aggregateMsg.HostIDs); i++ {
			if aggregateMsg.CpuRateAggregateValues[i] / 100 > alarmRule.CpuNoteworthyThreshold {
				// cpu 告警
				notifyUserForCPU = true
				alarmMsg.CpuInDangerHostIDs = append(alarmMsg.CpuInDangerHostIDs, aggregateMsg.HostIDs[i])
				alarmMsg.CpuInDangerValues = append(alarmMsg.CpuInDangerValues, aggregateMsg.CpuRateAggregateValues[i] / 100)
			}

			if aggregateMsg.MemRateAggregateValues[i] / 100 > alarmRule.MemNoteworthyThreshold {
				// memory
				notifyUserForMEM = true
				alarmMsg.MemInDangerHostIDs = append(alarmMsg.MemInDangerHostIDs, aggregateMsg.HostIDs[i])
				alarmMsg.MemInDangerValues = append(alarmMsg.MemInDangerValues, aggregateMsg.MemRateAggregateValues[i] / 100)
			}
		}

		// 拼接邮件内容，邮件内容需要填充的部分有：告警级别 - 主机ID - 告警指标
		if notifyUserForCPU || notifyUserForMEM {
			// log
			log.Printf("HOST IN DANGER! GONA SENT AN EMAIL TO USER!")

			// ready to send email
			emailContent := pkg.WarningEmailBuild(alarmRule, alarmMsg)
			log.Println(emailContent)
			emailAddr := alarmRule.ContactEmail
			// notifyPeriod内只告警一次
			err := pkg.DeliverEmail(emailContent, emailAddr)
			if err != nil {
				return
			}
		} else {
			log.Println("everything is alright in this period")
		}

		// 60s执行一次检查
		if notifyUserForCPU || notifyUserForMEM {
			time.Sleep(time.Minute * 10)
		} else {
			time.Sleep(time.Minute)
		}
	}
}
