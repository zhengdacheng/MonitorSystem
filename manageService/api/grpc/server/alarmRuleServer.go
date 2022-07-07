package server

import (
	"context"
	pb "manageService/api/grpc/proto"
	"manageService/internal/models"
)

type AlarmRuleServer struct {
}

func (server *AlarmRuleServer) AlarmRuleQuery(ctx context.Context, req *pb.AlarmRuleQueryReq) (*pb.AlarmRuleQueryResp, error) {
	// If there are any alarm rules inside mysql, then return latest alarm rule to client as well as CODE = 200
	// if there isn't any alarm rules inside mysql, then return code = 400

	// 1. search rules from MySQL
	rule := models.AlarmRule{}
	Code, latestRule := rule.FindLatest()
	var resp = pb.AlarmRuleQueryResp{}
	if Code == 200 {
		// rule exits
		resp.Code = Code
		resp.CpuNoteworthyThreshold = latestRule.CpuNoteworthyThreshold
		resp.CpuSeriousThreshold = latestRule.CpuSeriousThreshold
		resp.CpuDeadlyThreshold = latestRule.CpuDeadlyThreshold
		resp.MemNoteworthyThreshold = latestRule.MemNoteworthyThreshold
		resp.MemSeriousThreshold = latestRule.MemSeriousThreshold
		resp.MemDeadlyThreshold = latestRule.MemDeadlyThreshold
		resp.Granularity = latestRule.Granularity
		resp.AggregateFunction = latestRule.AggregateFunction
		resp.ContactEmail = latestRule.ContactEmail
	} else if Code == 400 {
		resp.Code = Code
	}
	return &resp, nil
}
