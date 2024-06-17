package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) GetQuestDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetQuestDataScRsp)
	rsp.QuestList = make([]*proto.Quest, 0)
	for _, questList := range gdconf.GetQuestDataMap() {
		quest := &proto.Quest{
			Progress:   0,
			Status:     proto.QuestStatus_QUEST_DOING,
			Id:         questList.QuestID,
			FinishTime: 1699688465,
		}
		rsp.QuestList = append(rsp.QuestList, quest)
	}

	g.Send(cmd.GetQuestDataScRsp, rsp)
}

func (g *GamePlayer) GetDailyActiveInfoCsReq(payloadMsg []byte) {
	dailyActiveQuestIdList := []uint32{2100132, 2100133, 2100134, 2100139, 2100150, 2100152, 2100153, 2100154}
	rsp := &proto.GetDailyActiveInfoScRsp{
		DailyActiveLevelList:   make([]*proto.DailyActiveLevel, 0),
		DailyActiveQuestIdList: dailyActiveQuestIdList,
		DailyActivePoint:       500,
	}

	for i := 1; i < 5; i++ {
		dailyActivityInfo := &proto.DailyActiveLevel{
			WorldLevel:            g.BasicBin.WorldLevel,
			Level:                 uint32(i),
			DailyActivePoint:      uint32(i * 100),
			HasTakenDailyActivity: true,
		}
		rsp.DailyActiveLevelList = append(rsp.DailyActiveLevelList, dailyActivityInfo)
	}

	g.Send(cmd.GetDailyActiveInfoScRsp, rsp)
}
