package day8

import (
	"container/list"
	"slices"
	"sort"
)

func JunctionBoxesP2(fp string) int {
	jb := readFileAndParse(fp)

	return findShortestDistanceP2(jb)
}

func calculateGroupSizesP2(jb []JB) bool {

	cmap := map[int][]int{}
	for _, val := range jb {
		if len(val.connected) > 0 {
			cmap[val.id] = val.connected
		}
	}

	if len(cmap) != len(jb) {
		return false
	}

	groups := make([][]int, len(jb))
	idx := 0

	for k := range cmap {
		queue := list.New()
		queue.PushBack(k)

		for queue.Len() > 0 {
			node := queue.Front()
			nv := node.Value.(int)
			if !slices.Contains(groups[idx], nv) {
				groups[idx] = append(groups[idx], nv)
			}

			// get the list of children and add them to the queue
			children, ok := cmap[nv]
			// after we have the children we can delete the key to stop cycles
			delete(cmap, nv)
			if ok {
				for _, child := range children {
					queue.PushBack(child)
				}
			}

			queue.Remove(node)
		}
		idx++
	}

	for _, group := range groups {
		if len(group) == len(jb) {
			return true
		}
	}

	return false
}

func findShortestDistanceP2(jb []JB) int {

	mappy := make(map[Pair]int)
	for i := range jb {
		for j := range jb {
			if i != j {
				_, ok := mappy[Pair{jb[i].id, jb[j].id}]
				_, ok1 := mappy[Pair{jb[j].id, jb[i].id}]
				if !ok1 && !ok {
					distance := calculateDistance(jb[i], jb[j])
					mappy[Pair{jb[i].id, jb[j].id}] = distance
					mappy[Pair{jb[j].id, jb[i].id}] = distance
				}
			}
		}
	}

	dictSlice := make([]KV, 0)
	for k, v := range mappy {
		dictSlice = append(dictSlice, KV{k, v})
	}

	sort.Slice(dictSlice, func(i, j int) bool {
		return dictSlice[i].second < dictSlice[j].second
	})

	for i, k := range dictSlice {
		pair := k.first
		jb[pair.first].connected = append(jb[pair.first].connected, pair.second)
		if i > len(jb)*2+1 {
			if calculateGroupSizesP2(jb) {
				return jb[pair.first].x * jb[pair.second].x
			}
		}
	}
	return 0
}
