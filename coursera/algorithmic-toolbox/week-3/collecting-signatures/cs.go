package collecting_signatures

import "sort"

type schedule struct {
	start  int
	finish int
}

type scheduleSort []schedule

func (s scheduleSort) Len() int           { return len(s) }
func (s scheduleSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s scheduleSort) Less(i, j int) bool { return s[i].finish <= s[j].finish }

func cs(schedules []schedule) (int, schedule) {

	sort.Sort(scheduleSort(schedules))

	count := 1
	scd := schedule{schedules[0].finish, 0}
	finish := scd.start

	for _, sc := range schedules[1:] {
		if sc.start > finish {
			scd.finish, finish = sc.finish, sc.finish
			count++
		}
		// fmt.Println(sc.start, sc.finish)
	}

	return count, scd
}
