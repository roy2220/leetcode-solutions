/**
 * Definition for an interval.
 * type Interval struct {
 *         Start int
 *         End   int
 * }
 */
import "sort"

type intervals []Interval

func (is intervals) Less(i, j int) bool {
    return is[i].Start < is[j].Start
}

func (is intervals) Swap(i, j int) {
    is[i], is[j] = is[j], is[i]
}

func (is intervals) Len() int {
    return len(is)
}

func merge(is []Interval) []Interval {
    result := []Interval{}

    if len(is) == 0 {
        return result
    }

    sort.Sort(intervals(is))
    start := is[0].Start
    end := is[0].End

    for i := 1; i < len(is); i++ {
        if is[i].Start <= end {
            if is[i].End > end {
                end = is[i].End
            }
        } else {
            result = append(result, Interval{Start: start, End: end})
            start = is[i].Start
            end = is[i].End
        }
    }

    result = append(result, Interval{Start: start, End: end})
    return result
}

func insert(intervals []Interval, newInterval Interval) []Interval {
    intervals = append(intervals, newInterval)
    return merge(intervals)
}
