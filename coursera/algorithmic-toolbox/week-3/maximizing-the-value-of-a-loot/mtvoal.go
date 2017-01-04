package maximizing_the_value_of_a_loot

import "sort"

type loot struct {
	value  int
	weight int
}

type lootSort []loot

func (l lootSort) Len() int           { return len(l) }
func (l lootSort) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l lootSort) Less(i, j int) bool { return l[i].value/l[i].weight >= l[j].value/l[j].weight }

func NativeSolution(loots []loot, capacity int) int {

	sort.Sort(lootSort(loots))

	return getMaxLootNative(loots, capacity)
}

func getMaxLootNative(loots []loot, capacity int) int {

	if len(loots) <= 0 || capacity <= 0 {
		return 0
	}

	l := loots[0]

	takenWeight := 0

	if capacity < l.weight {
		takenWeight = capacity * l.value / l.weight
		capacity = 0
	} else {
		takenWeight = l.value
		capacity = capacity - l.weight
	}

	return takenWeight + getMaxLootNative(loots[1:], capacity)
}
