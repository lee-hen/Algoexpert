package lowest_common_manager

import (
	"github.com/stretchr/testify/require"
	"testing"
)


func (chart *OrgChart) addDirectReports(reports ...*OrgChart) {
	chart.DirectReports = append(chart.DirectReports, reports...)
}

func getOrgCharts() map[rune]*OrgChart {
	orgCharts := map[rune]*OrgChart{}
	for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		orgCharts[r] = &OrgChart{
			Name:          string(r),
			DirectReports: []*OrgChart{},
		}
	}
	return orgCharts
}

func TestCase1(t *testing.T) {
	orgCharts := getOrgCharts()
	orgCharts['A'].addDirectReports(orgCharts['B'], orgCharts['C'])
	orgCharts['B'].addDirectReports(orgCharts['D'], orgCharts['E'])
	orgCharts['C'].addDirectReports(orgCharts['F'], orgCharts['G'])
	orgCharts['D'].addDirectReports(orgCharts['H'], orgCharts['I'])
	lcm := GetLowestCommonManager(orgCharts['A'], orgCharts['E'], orgCharts['I'])
	require.Equal(t, lcm, orgCharts['B'], lcm)
}
