// import "sort"
type Pair struct {
    NextValue int
    Length int
}
func longestConsecutive(nums []int) int {
    if len(nums)<=1{
        return len(nums)
    }
    // sort.Slice(nums,func(i,j int)bool{
    //    return nums[i]<nums[j]
    // })
    //  ans :=1
    //  tempAns:=1
    // for i:=0;i<len(nums)-1;i++{
    //     if nums[i]+1 == nums[i+1]{
    //         tempAns++
    //     } else if nums[i] == nums[i+1]{
    //        continue
    //     } else {
    //         tempAns = 1
    //     }
    //     if tempAns > ans{
    //         ans = tempAns
    //     }
    // }

    fmap:= make(map[int]Pair)
    for i:=0;i<len(nums);i++{
        fmap[nums[i]] = Pair{
            NextValue: 1,
            Length: 1,
        }
    }

    ans :=1
  
  for key,_:= range fmap {
      tempAns:=RecursiveFindLength(key,fmap)
      if tempAns > ans{
        ans = tempAns
      }
      
  }

    
    return ans
}

func RecursiveFindLength(key int, fmap map[int]Pair) int{
    // fmt.Println(key,fmap)
    if fmap[key].Length>1{
        return fmap[key].Length
    }
    if value,exist := fmap[key+1]; exist{
        fmap[key] = Pair{
            NextValue:key+1,
            Length: fmap[key].Length,
        }
        if value.Length>1{
            fmap[key] = Pair{
            NextValue:key+1,
            Length: fmap[key].Length+value.Length,
        }
        } else{
         length:=   RecursiveFindLength(key+1,fmap)
         fmap[key] = Pair{
            NextValue:key+1,
            Length: fmap[key].Length+length,
        }

        }
    }
    return fmap[key].Length
}