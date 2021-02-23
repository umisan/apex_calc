package main

const maxKill = 6
const worstRank = 14
const rankResoultion = 14

var rankPointTable []int = []int{0, 5, 5, 5, 10, 10, 20, 20, 30, 30, 40, 40, 60, 100}
var rankLeverageTable []int = []int{10, 10, 10, 10, 12, 12, 12, 12, 12, 15, 15, 20, 20, 25}
var classEntryFeeTable []int = []int{0, 12, 24, 36, 48, 60, 60}

type Cand struct {
	Rank int
	Kill int
}

type CandList = []Cand

func reverseCalculation(needPoint int, class int) CandList {
	result := CandList{}
	for tableIndex := 0; tableIndex < rankResoultion; tableIndex++ {
		for kill := 0; kill <= maxKill; kill++ {
			point := rankPointTable[tableIndex] + calcKillPoint(kill, rankLeverageTable[tableIndex]) - classEntryFeeTable[class]
			if point >= needPoint {
				result = append(result, Cand{Rank: mapTableIndexToRank(tableIndex), Kill: kill})
				break
			}
		}
	}
	return result
}

func calcKillPoint(kill, leverage int) int {
	if kill > maxKill {
		kill = maxKill
	}
	return kill * leverage
}

func mapTableIndexToRank(tableIndex int) int {
	return worstRank - tableIndex
}
