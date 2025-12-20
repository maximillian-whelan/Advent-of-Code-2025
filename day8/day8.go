package day8

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type JB struct {
	x         int
	y         int
	z         int
	id        int
	connected []int
}

func (jb *JB) Equal(ajb JB) bool {
	return jb.x == ajb.x && jb.y == ajb.y && jb.z == ajb.z && jb.id == ajb.id
}

func JunctionBoxes(fp string, pairs int) int {
	jb := readFileAndParse(fp)

	findShortestDistance(jb, pairs)

	return calculateGroupSizes(jb)
}

func calculateGroupSizes(jb []JB) int {

	cmap := map[int][]int{}
	for _, val := range jb {
		if len(val.connected) > 0 {
			cmap[val.id] = val.connected
		}
	}
	fmt.Printf("%v\n\n", cmap)

	groups := make([][]int, len(cmap))
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

	sort.SliceStable(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})

	for _, group := range groups {
		fmt.Printf("%v\n", group)
	}

	return len(groups[0]) * len(groups[1]) * len(groups[2])
}

type Pair struct {
	first  int
	second int
}

type KV struct {
	first  Pair
	second int
}

func findShortestDistance(jb []JB, pairs int) {

	mappy := make(map[Pair]int)
	for i := range jb {
		for j := range jb {
			if i != j {
				distance := calculateDistance(jb[i], jb[j])
				mappy[Pair{jb[i].id, jb[j].id}] = distance
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

	fmt.Printf("%v\n", dictSlice)

	for i, k := range dictSlice {
		if i >= pairs * 2{
			break
		}
		// gets us the two ids that are connected
		pair := k.first 
		jb[pair.first].connected = append(jb[pair.first].connected, pair.second)
	}
}

func calculateDistance(first, second JB) int {
	inner := math.Pow(float64(first.x)-float64(second.x), 2) +
		math.Pow(float64(first.y)-float64(second.y), 2) +
		math.Pow(float64(first.z)-float64(second.z), 2)
	return int(math.Sqrt(inner))
}

func readFileAndParse(fp string) []JB {
	junctionBoxes := make([]JB, 0)

	fi, err := os.Open(fp)
	if err != nil {
		panic("unable to open file")
	}

	s := bufio.NewScanner(fi)
	id := 0
	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		junctionBoxes = append(junctionBoxes, JB{x, y, z, id, []int{}})
		id++

	}
	return junctionBoxes
}
