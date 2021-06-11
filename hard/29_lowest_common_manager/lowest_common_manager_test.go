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
	require.Equal(t, orgCharts['B'], lcm)
}

func TestCase2(t *testing.T) {
	orgCharts := getOrgCharts()
	orgCharts['A'].addDirectReports(orgCharts['B'], orgCharts['C'], orgCharts['D'], orgCharts['E'], orgCharts['F'])
	orgCharts['B'].addDirectReports(orgCharts['G'], orgCharts['H'], orgCharts['I'])
	orgCharts['C'].addDirectReports(orgCharts['J'])
	orgCharts['D'].addDirectReports(orgCharts['K'], orgCharts['L'])
	orgCharts['F'].addDirectReports(orgCharts['M'], orgCharts['N'])
	orgCharts['H'].addDirectReports(orgCharts['O'], orgCharts['P'], orgCharts['Q'], orgCharts['R'])
	orgCharts['K'].addDirectReports(orgCharts['S'])
	orgCharts['P'].addDirectReports(orgCharts['T'], orgCharts['U'])
	orgCharts['R'].addDirectReports(orgCharts['V'])
	orgCharts['V'].addDirectReports(orgCharts['W'], orgCharts['X'], orgCharts['Y'])
	orgCharts['X'].addDirectReports(orgCharts['Z'])

	lcm := getLowestCommonManager(orgCharts['A'], orgCharts['T'], orgCharts['H'])
	require.Equal(t, orgCharts['H'], lcm)
}
