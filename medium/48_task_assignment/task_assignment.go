package task_assignment

import (
	"sort"
)

func TaskAssignment(k int, tasks []int) [][]int {
	pairedTasks := make([][]int, 0)
	taskDurationsToIndices := getTaskDurationsToIndices(tasks)

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] < tasks[j]
	})

	var task1Index, task2Index int
	for idx := 0; idx < k; idx++ {
		task1Duration := tasks[idx]
		indicesWithTask1Duration := taskDurationsToIndices[task1Duration]
		task1Index, taskDurationsToIndices[task1Duration] =
			indicesWithTask1Duration[len(indicesWithTask1Duration)-1], indicesWithTask1Duration[:len(indicesWithTask1Duration)-1]

		task2SortedIndex := len(tasks) - 1 - idx
		task2Duration := tasks[task2SortedIndex]
		indicesWithTask2Duration := taskDurationsToIndices[task2Duration]
		task2Index, taskDurationsToIndices[task2Duration] =
			indicesWithTask2Duration[len(indicesWithTask2Duration)-1], indicesWithTask2Duration[:len(indicesWithTask2Duration)-1]

		pairedTasks = append(pairedTasks, []int{task1Index, task2Index})
	}
	return pairedTasks
}

func getTaskDurationsToIndices(tasks []int) map[int][]int {
	taskDurationsToIndices := map[int][]int{}
	for idx := range tasks {
		taskDuration := tasks[idx]
		taskDurationsToIndices[taskDuration] = append(taskDurationsToIndices[taskDuration], idx)
	}
	return taskDurationsToIndices
}

func taskAssignment(k int, tasks []int) (assignments [][]int) {
	if len(tasks) == 0 {
		return [][]int{{}}
	}

	taskPosition := make(map[int][]int)
	for idx, task := range tasks {
		if _, found := taskPosition[task]; found {
			taskPosition[task] = append(taskPosition[task], idx)
		} else {
			taskPosition[task] = []int{idx}
		}
	}

	newTasks := make([]int, len(tasks))
	copy(newTasks, tasks)
	sort.Ints(newTasks)

	pairOfTask := make([][]int, 0)
	for i, j := 0, len(newTasks)-1; i < j; i, j = i+1, j-1 {
		pairOfTask = append(pairOfTask, []int{newTasks[i], newTasks[j]})
	}

	for _, pair := range pairOfTask {
		task1, task2 := pair[0], pair[1]
		indices1, _ := taskPosition[task1]
		indices2, _ := taskPosition[task2]

		if indices1[len(indices1)-1] != indices2[len(indices2)-1] {
			assignments = append(assignments, []int{indices1[len(indices1)-1], indices2[len(indices2)-1]})
			taskPosition[task1] = indices1[:len(indices1)-1]
			taskPosition[task2] = indices2[:len(indices2)-1]
		} else {
			assignments = append(assignments, []int{indices1[len(indices1)-1], indices2[len(indices2)-2]})
			taskPosition[task1] = indices1[:len(indices1)-2]
			taskPosition[task2] = taskPosition[task1]
		}
	}

	return assignments
}
