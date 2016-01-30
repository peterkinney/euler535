package sequenceGenerator

import (
	"fmt"
	"github.com/peterkinney/euler535/floorSqrt"
	"github.com/peterkinney/moduloMath"
)


var groupLengths [11]uint64

type group struct {
	cirNext, cirRemaining, nonCirNext uint64
	uCirNext, uCirRemaining, uNonCirNext uint64
	prevGroup *group
}

type Generator struct {
	groups [12]group
}

func init () {
	groupLengths[0] = 1
	for ii := 1; ii < 11; ii++ {
		groupLengths[ii] = floorSqrt.SumFSq(groupLengths[ii-1]) + groupLengths[ii-1]
	}
	/*for ii := 0; ii < 11; ii++ {
		fmt.Printf("group %d length %d\n",ii, groupLengths[ii])
	}*/
}

func NewGenerator () *Generator {
	gener := Generator{}
	gen := &gener
	gen.Reset()
	return gen
}

func (gen *Generator) Reset () {
	gen.groups[0] = group{nonCirNext: uint64(1)}
	gen.groups[0].SaveState()
	for ii := 1; ii < 12; ii++ {
		gen.groups[ii] = group{cirNext: groupLengths[ii-1]+1}
		gen.groups[ii].prevGroup = &(gen.groups[ii-1])
		gen.groups[ii].SaveState()
	}
}

func (gg *group) nextValue () (value uint64, circled, ok bool) {
	gg.SaveState()
	if gg.nonCirNext == 0 {
		gg.nonCirNext, _, ok = gg.prevGroup.nextValue()
		if !ok {
			return 0,false,false
		}
		gg.cirRemaining = floorSqrt.FloorSqrt(gg.nonCirNext)
	}
	if gg.cirRemaining > 0 {
		value = gg.cirNext
		circled = true
		ok = true
		gg.cirNext++
		gg.cirRemaining--
		return
	}	
	value = gg.nonCirNext
	circled = false
	ok = true
	gg.nonCirNext = 0
	return
}

func (gg *group) peek () (value uint64, circled, ok bool) {
	if gg.cirRemaining > 0 {
		value = gg.cirNext
		circled = true
		ok = true
		return
	}
	if gg.nonCirNext > 0 {
		value = gg.nonCirNext
		circled = false
		ok = true
		return
	}
	_, _, ok = gg.prevGroup.peek()
	if ok {
		value = gg.cirNext
		circled = true
	} else {
		value = 0
		circled = false
	}
	return
}

func (gg *group) skipCircled(lenToSkip uint64) {
	gg.SaveState()
	gg.cirNext += lenToSkip
	gg.cirRemaining = 0
	gg.nonCirNext = 0
}

func (gen *Generator) LenOfNextCirInGroup(start, target int) (len uint64, sum uint64){
	if start == target {
		sum, circled, ok := gen.groups[start].nextValue()
		if !circled || !ok {
			gen.groups[start].Undo()
			fmt.Printf("Group %d ran out of accessible circled numbers\n", start)
			return 0,0
		}
		return 1,sum
	}
	newestCircled, _, _ := gen.groups[start].nextValue()
	sum += newestCircled
	
	nonCircledCount, circledCount := uint64(1), floorSqrt.SumFSqRange(newestCircled,1)
	//TODO: handle the ok
	newestCircled, _ , _ = gen.groups[start+1].peek()
	gen.groups[start+1].skipCircled(circledCount)
	sum = moduloMath.Sum(1e9, sum, moduloMath.SumRange(1e9, newestCircled, circledCount))
	for ii := start + 2; ii <= target; ii++ {
		nonCircledCount += circledCount
		circledCount += floorSqrt.SumFSqRange(newestCircled,circledCount)
		//TODO: handle the ok
		newestCircled,_,_ = gen.groups[ii].peek()
		gen.groups[ii].skipCircled(circledCount)
		sum = moduloMath.Sum(1e9, sum, moduloMath.SumRange(1e9, newestCircled, circledCount))
	}
	len = nonCircledCount + circledCount
	return
}

func (gen *Generator) IsNextCircled (groupNum int) bool {
	_,circled,_ := gen.groups[groupNum].peek()
	return circled
}

func (gen *Generator) NextCircled (groupNum int) uint64 {
	value,_,_ := gen.groups[groupNum].peek()
	return value
}

func (gg *group) SaveState () {
	gg.uCirNext, gg.uCirRemaining, gg.uNonCirNext = gg.cirNext, gg.cirRemaining, gg.nonCirNext
}

func (gg *group) Undo () {
	gg.cirNext, gg.cirRemaining, gg.nonCirNext = gg.uCirNext, gg.uCirRemaining, gg.uNonCirNext
}

func (gen *Generator) UndoUpTo(index int) {
	for ii := 0; ii <= index; ii++ {
		gen.groups[ii].Undo()
	}
}

func (gen *Generator) UndoAll() {
	for index := 0; index < 12; index++ {
		gen.groups[index].Undo()
	}
}

func SequenceIndexOfGroup(groupNumber uint) (sum uint64){
	sum = uint64(1)
	for ii := uint(0); ii < groupNumber; ii++ {
		sum += groupLengths[ii]
	}
	return
}

func LastGroup (nn uint64) (groupIndex int, groupPosition uint64)  {
	if nn <= 1 {
		groupIndex = -1
		groupPosition = 0
		return
	}
	groupPosition = uint64(1)
	for ; groupIndex < 11; groupIndex++ {
		if nn <= groupPosition + groupLengths[groupIndex] {
			return
		}
		groupPosition += groupLengths[groupIndex]
	}
	return
}

func SumFirstNGroups (nn int) (sum uint64) {
	sum = 1
	for ii := 0; ii < nn; ii++ {
		sum = moduloMath.Sum(1e9,sum,moduloMath.SumRange(1e9,1,groupLengths[ii]))
	}
	return
}