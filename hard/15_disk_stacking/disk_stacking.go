package disk_stacking
// important question

import (
	"math"
	"sort"
)

type Disk []int
type Disks []Disk

func (disks Disks) Len() int           { return len(disks) }
func (disks Disks) Swap(i, j int)      { disks[i], disks[j] = disks[j], disks[i] }
func (disks Disks) Less(i, j int) bool { return disks[i][2] < disks[j][2] }

func DiskStacking(input [][]int) [][]int {
	disks := make(Disks, len(input))
	for i, disk := range input {
		disks[i] = disk
	}

	sort.Sort(disks)
	heights := make([]int, len(disks))
	sequences := make([]int, len(disks))

	for i := range disks {
		heights[i] = disks[i][2]
		sequences[i] = -1
	}

	for i := 1; i < len(disks); i++ {
		disk := disks[i]
		for j := 0; j < i; j++ {
			other := disks[j]

			if areValidDimensions(other, disk) {
				if heights[i] <= disk[2]+heights[j] {
					heights[i] = disk[2] + heights[j]
					sequences[i] = j
				}
			}
		}
	}

	maxIndex := 0
	for i, height := range heights {
		if height >= heights[maxIndex] {
			maxIndex = i
		}
	}

	sequence := buildSequence(disks, sequences, maxIndex)
	return sequence
}

func areValidDimensions(o Disk, c Disk) bool {
	return o[0] < c[0] && o[1] < c[1] && o[2] < c[2]
}

func buildSequence(array []Disk, sequences []int, index int) [][]int {
	out := make([][]int, 0)
	for index != -1 {
		out = append(out, array[index])
		index = sequences[index]
	}

	reverse(out)
	return out
}

// [2, 1, 2] [3, 2, 3] [2, 2, 8] [2, 3, 4] [2, 2, 1] [4, 4, 5]
// [2, 1, 2] [3, 2, 3] [4, 4, 5]

//     0         1         2         3         4         5
// [2, 1, 2] [3, 2, 3] [2, 2, 8] [2, 3, 4] [2, 2, 1] [4, 4, 5]
//     2         5         8         4         1         10
//    -1         0        -1        -1        -1         1

//     0          1          2          3          4          5          6
// [2, 1, 2], [3, 2, 3], [2, 2, 8], [2, 3, 4], [1, 2, 1], [4, 4, 5], [1, 1, 4]
//     2          3          8          4          1          5          4

// my solution
func diskStacking(disks [][]int) [][]int {
	sort.Slice(disks, func(i, j int) bool{
		if disks[i][0] > disks[j][0] && disks[i][1] > disks[j][1] && disks[i][2] > disks[j][2] {
			return false
		}
		return true
	})

	diskIndexWithHeight := make([][]int, len(disks))
	for i := 0; i < len(disks); i++ {
		diskIndexWithHeight[i] = []int{-1, disks[i][2]}
	}

	maxIdx, maxHeight := -1, math.MinInt32
	for i, disk := range disks {
		maxValue := disks[i][2]
		for j := 0; j < i; j++ {
			otherDisk := disks[j]
			if disk[0] > otherDisk[0] && disk[1] > otherDisk[1] && disk[2] > otherDisk[2] {
				if maxValue < disks[i][2] + diskIndexWithHeight[j][1] {
					maxValue = disks[i][2] + diskIndexWithHeight[j][1]
					diskIndexWithHeight[i][0] = j
				}
			}
		}

		diskIndexWithHeight[i][1] = maxValue
		if maxHeight < diskIndexWithHeight[i][1] {
			maxHeight = diskIndexWithHeight[i][1]
			maxIdx = i
		}
	}

	if maxIdx != -1 {
		diskStaking := make([][]int, 0)
		currIdx := maxIdx
		for currIdx != -1 {
			diskStaking = append(diskStaking, disks[currIdx])
			currIdx = diskIndexWithHeight[currIdx][0]
		}
		reverse(diskStaking)
		return diskStaking
	}

	return [][]int{}
}

func reverse(numbers [][]int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}
