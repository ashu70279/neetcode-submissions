/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

type Pair struct {
	time int
	val int
}
func minMeetingRooms(intervals []Interval) int {
     lineSweep :=[]Pair{}
	 for i:=0;i<len(intervals);i++{
		lineSweep = append(lineSweep,Pair{
			time:intervals[i].start,
			val:1,
		})

		lineSweep = append(lineSweep,Pair{
			time:intervals[i].end,
			val:-1,
		})
	 }

	 sort.Slice(lineSweep,func(i,j int)bool{
		if lineSweep[i].time == lineSweep[j].time{
			return lineSweep[i].val<lineSweep[j].val
		}
		return lineSweep[i].time<lineSweep[j].time
	 })

	 tempAns:=0
	 ans:=0
	 for i:=0;i<len(lineSweep);i++{
		tempAns+=lineSweep[i].val
		ans = max(ans,tempAns)
	 }

	 return ans

}
