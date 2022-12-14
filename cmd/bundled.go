package main

import (
	"fyne.io/fyne/v2"
)

var resource10kVTxt = []fyne.StaticResource{
	{
		StaticName: "10kV柱上变压器停电操作步骤.txt",
		StaticContent: []byte(
			"现场核对线路名称、杆号所变编号\n设置安全围栏，并设置作业标志\n检查绝缘安全工具外观及安全试验合格证\n检查所变电源状态（三相电压指示）\n断开所变低压断路器\n检查所变低压断路器已断开\n断中相跌落式熔断器\n检查中相跌落式熔断器已断开\n断下风相跌落式熔断器\n检查下风相跌落式熔断器已断开\n断上风相跌落式熔断器\n检查上风相跌落式熔断器已断开\n取下绝缘操作杆，拆除安全围栏，并清理现场，作业任务结束"),
	},
	{
		StaticName: "10kV出线开关柜操作（冷备用转运行）步骤.txt",
		StaticContent: []byte(
			"听调度指令\n判断10KV配电房高压柜状态\n摇入10kV配电房馈线柜小车\n判断10kV配电房高压柜状态\n合10kV 配电房高压进线柜断路器操作开关\n合10kV 配电房高压馈线柜断路器操作开关\n再次检查10kV配电房高压柜断路器状态\n挂上运行警示牌\n检查任务结束状态"),
	},
	{
		StaticName: "10kV柱上变压器送电操作步骤.txt",
		StaticContent: []byte(
			"现场核对线路名称、杆号所变编号\n设置安全围栏，并设置作业标志\n检查绝缘安全工具外观及安全试验合格证\n送电前检查，检查即将带电设备、线路，确认其上没有人员、异物，接地线且所变低压侧断路器全部合上\n合上风相跌落式熔断器\n检查上风相跌落式熔断器已合上\n合下风相跌落式熔断器\n检查下风相跌落式熔断器已合上\n合中相跌落式熔断器\n检查中相跌落式熔断器已合上\n合上所变低压断路器, 观察所变负荷状态，确定其在运行状态\n取下绝缘操作杆，拆除安全围栏，并清理现场，作业任务结束"),
	},
	{

		StaticName: "10kV高压出线开关柜操作（运行转冷备用）步骤.txt",
		StaticContent: []byte(
			"听调度指令\n判断10KV配电房高压柜状态\n分10kV 配电房高压馈线柜断路器操作开关\n分10kV 配电房高压进线柜断路器操作开关\n检查10kV配电房高压柜断路器状态\n摇出10kV配电房馈线柜小车\n再次判断10kV配电房高压柜状态\n检查任务结束状态"),
	},
}
