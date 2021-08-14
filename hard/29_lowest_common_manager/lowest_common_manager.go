package lowest_common_manager
// important question

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

type OrgInfo struct {
	lowestCommonManager *OrgChart
	numImportantReports int
}

// GetLowestCommonManager
// O(n) time | O(d) space - where n is the number of people
// in the org and d is the depth (height) of the org chart
func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	return getOrgInfo(org, reportOne, reportTwo).lowestCommonManager
}

func getOrgInfo(manager, reportOne, reportTwo *OrgChart) OrgInfo {
	numImportantReports := 0

	for _, directReport := range manager.DirectReports {
		orgInfo := getOrgInfo(directReport, reportOne, reportTwo)

		if orgInfo.lowestCommonManager != nil {
			return orgInfo
		}

		numImportantReports += orgInfo.numImportantReports
	}

	if manager == reportOne || manager == reportTwo {
		numImportantReports++
	}

	var lowestCommonManager *OrgChart
	if numImportantReports == 2 {
		lowestCommonManager = manager
	}

	return OrgInfo{
		lowestCommonManager: lowestCommonManager,
		numImportantReports: numImportantReports,
	}
}

// topManger Node A
// reportOne Node E
// reportTwo Node I

//               A
//          /         \
//         B             C
//      /  |  \       /  |  \
//     D   H   E     F   Y   G
//  /  |  \  \     /     |     \
// H   X   I  U   Z      V      S

// node B

func getLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	if reportOne == org || reportTwo == org {
		return org
	}

	nodesToMangers := map[string]*OrgChart{}
	populateNodesToMangers(org, nodesToMangers)

	return findLowestCommonManager(reportOne, reportTwo, nodesToMangers)
}

func findLowestCommonManager(reportOne, reportTwo *OrgChart, nodesToMangers map[string]*OrgChart) *OrgChart{
	reportOneDirectManger := nodesToMangers[reportOne.Name]
	reportTwoDirectManger := nodesToMangers[reportTwo.Name]

	if reportOneDirectManger == reportTwoDirectManger {
		return reportOneDirectManger
	}

	if isMangerOfTargetNode(reportOne, reportTwo) {
		return reportOne
	}

	if isMangerOfTargetNode(reportTwo, reportOne) {
		return reportTwo
	}

	if isMangerOfTargetNode(reportOneDirectManger, reportTwo) {
		return reportOneDirectManger
	}

	if isMangerOfTargetNode(reportTwoDirectManger, reportOne) {
		return reportTwoDirectManger
	}

	return findLowestCommonManager(reportOneDirectManger, reportTwoDirectManger,  nodesToMangers)
}

func isMangerOfTargetNode(manger, target *OrgChart) bool {
	for _, directReportNode := range manger.DirectReports {
		if directReportNode == target {
			return true
		}

		if isMangerOfTargetNode(directReportNode, target) {
			return true
		}
	}

	return false
}

func populateNodesToMangers(manger *OrgChart, nodesToMangers map[string]*OrgChart) {
	for _, directReportNode := range manger.DirectReports {
		nodesToMangers[directReportNode.Name] = manger
		populateNodesToMangers(directReportNode, nodesToMangers)
	}
}
