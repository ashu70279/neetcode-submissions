/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	sort.Slice(intervals,func(i,j int)bool{
		return intervals[i].start<intervals[j].start
	})
    takenRoom:=[]int{}
	for i:=0;i<len(intervals);i++{
		exist:=checkRoomAvaliablity(intervals[i],takenRoom)
		if !exist{
			takenRoom = append(takenRoom,intervals[i].end)
		}
	}

	return len(takenRoom)
}

func checkRoomAvaliablity(interval Interval, takenRoom []int)bool{
	for i:=0;i<len(takenRoom);i++{
		if interval.start >= takenRoom[i]{
			takenRoom[i] = interval.end
			return true
		}
	}
	return false
}
