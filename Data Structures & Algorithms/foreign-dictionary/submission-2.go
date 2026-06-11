func foreignDictionary(words []string) string {

	adjList:=make(map[int][]int)
    for i:=1;i<len(words);i++{
		index1,index2:=getDiffCharacters(i-1,i,words)
		if index1 == 0 && index2 ==0{
            continue
		}
		if index1 == -1 && index2 ==-1{
            return ""
		}
        adjList[index1] = append(adjList[index1],index2)
	}
	for i:=0;i<len(words);i++{
		for j:=0;j<len(words[i]);j++{
			if _,exist:=adjList[int(words[i][j])-97];!exist{
				adjList[int(words[i][j])-97]=[]int{}
			}
		}
	}

	order:=topologicalSort(adjList)
	if len(adjList) == len(order){
	return string(order)
	}
	return ""
	
}

func getDiffCharacters(i,j int, words []string)(int,int){
	start1:=0
	start2:=0
	end1:=len(words[i])
	end2:=len(words[j])

	for start1<end1 && start2<end2{
		if words[i][start1]!=words[j][start2]{
			return int(words[i][start1])-97,int(words[j][start2])-97
		}
		start1++
		start2++
	}

	if start1<end1 && start2 == end2{
		return -1,-1
	}

	return 0,0

}

func topologicalSort(adjList map[int][]int)[]byte{
    indegree:=make(map[int]int)
	for key,_:=range adjList{
		indegree[key]=0
	}

    queue := []int{}
	for _,val:=range adjList{
		for j:=0;j<len(val);j++{
			indegree[val[j]]++
		}
	}

	for key,val:=range indegree{
		if val ==0{
          queue = append(queue,key)
		}
	}

	toplogicalSort:=[]byte{}

	for len(queue)>0{
		item:=queue[0]
		queue = queue[1:]
        toplogicalSort = append(toplogicalSort,byte(item+'a'))
		for i:=0;i<len(adjList[item]);i++{
			indegree[adjList[item][i]]--
			if indegree[adjList[item][i]] == 0{
				queue = append(queue,adjList[item][i])
			}
		}
	}

  return toplogicalSort


}
