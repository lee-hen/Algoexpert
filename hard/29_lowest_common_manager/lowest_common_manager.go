package lowest_common_manager
type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}


func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	return nil
}
