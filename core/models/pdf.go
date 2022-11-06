package models

import "core/internal/types"

type Pdf struct {
	Title            string           `json:"title"`
	Subject          string           `json:"subject"`
	GroupLeaderName  string           `json:"groupLeaderName"`
	GroupMemberName  string           `json:"groupMemberName"`
	Classes          string           `json:"classes"`
	Teacher          string           `json:"teacher"`
	Company          string           `json:"company"`
	LeaderWorkDivide []*types.Content `json:"leaderWorkDivide"`
	MemberWorkDivide []*types.Content `json:"memberWorkDivide"`
	Requirement      []*types.Content `json:"requirement"`
	DemandAnalysis   []*types.Content `json:"demandAnalysis"`
	OutlineDesign    []*types.Content `json:"outlineDesign"`
	SourceCode       []*types.Content `json:"sourceCode"`
	TestAndResult    []*types.Content `json:"testAndResult"`
	Question         []*types.Content `json:"question"`
	Summary          []*types.Content `json:"summary"`
}

//type Content struct {
//	Type        string
//	TextContent []string
//}
