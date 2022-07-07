package pkg

import (
	"fmt"
	"strings"
	"warningService/internal/models"
)

func WarningEmailBuild(rule models.AlarmRule, msg models.AlarmMsg) string {
	// Email Content
	var emailContent strings.Builder
	// CPU Part
	if len(msg.CpuInDangerHostIDs) > 0 {
		emailContent.WriteString("HostID   -   CpuRate -   WarningLevel")
		emailContent.WriteString("<br>")
		for i := 0; i < len(msg.CpuInDangerHostIDs); i++ {
			cpuWarnLevel := GetWarnLevel(rule, "CPU_Rate", msg.CpuInDangerValues[i])
			emailContent.WriteString(
				fmt.Sprintf(`%s - %.2f - %s`, msg.CpuInDangerHostIDs[i], msg.CpuInDangerValues[i], cpuWarnLevel))
			emailContent.WriteString("<br>")
		}
	}
	// Memory Part
	if len(msg.MemInDangerHostIDs) > 0 {
		emailContent.WriteString("HostID - MEM_Rate - Warning_Level")
		emailContent.WriteString("<br>")
		for i := 0; i < len(msg.MemInDangerHostIDs); i++ {
			memWarnLevel := GetWarnLevel(rule, "MEM_Rate", msg.MemInDangerValues[i])
			emailContent.WriteString(
				fmt.Sprintf(`%s - %.2f - %s`, msg.MemInDangerHostIDs[i], msg.MemInDangerValues[i], memWarnLevel))
			emailContent.WriteString("<br>")
		}
	}

	return emailContent.String()
}

func GetWarnLevel(rule models.AlarmRule, metricsType string, metricsValue float32) string {
	if metricsType == "CPU_Rate" {
		switch {
		case metricsValue < rule.CpuNoteworthyThreshold:
			return "FINE"
		case metricsValue >= rule.CpuNoteworthyThreshold && metricsValue < rule.CpuSeriousThreshold:
			return "NOTICE"
		case metricsValue >= rule.CpuSeriousThreshold && metricsValue < rule.CpuDeadlyThreshold:
			return "SERIOUS"
		case metricsValue >= rule.CpuDeadlyThreshold:
			return "DEADLY"
		}
	} else if metricsType == "MEM_Rate" {
		switch {
		case metricsValue < rule.MemNoteworthyThreshold:
			return "FINE"
		case metricsValue >= rule.MemNoteworthyThreshold && metricsValue < rule.CpuSeriousThreshold:
			return "NOTICE"
		case metricsValue >= rule.MemSeriousThreshold && metricsValue < rule.CpuDeadlyThreshold:
			return "SERIOUS"
		case metricsValue >= rule.MemDeadlyThreshold:
			return "DEADLY"
		}
	} else {
		panic("metricsType mismatch")
	}
	return ""
}
