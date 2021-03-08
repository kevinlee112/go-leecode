package _01_300

import "sort"

/*

给定一系列的会议时间间隔，包括起始和结束时间[[s1,e1]，[s2,e2]，…(si < ei)，确定一个人是否可以参加所有会议。
 */

type Interval struct {
	     Start, End int
	 }
func canAttendMeetings (intervals []*Interval) bool {
	res := true
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start > intervals[j].Start
	})
	max := intervals[0].Start
	for i:=0;i<len(intervals);i++ {
		if intervals {

		}
	}

}

