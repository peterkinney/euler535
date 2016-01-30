package main

import (
    "fmt"
    "github.com/peterkinney/euler535/sequenceGenerator"
	"github.com/peterkinney/moduloMath"
)

func sequenceSum(nn uint64) (sum uint64){
	gen := sequenceGenerator.NewGenerator()
	//get the target group and length up to that point
	targetGroup, len := sequenceGenerator.LastGroup(nn)
	nRemain := nn - len
	fmt.Printf("the %dth value of the sequence is the %dth value of group %d\n", nn, nRemain, targetGroup)
	sum = sequenceGenerator.SumFirstNGroups(targetGroup)
	for ii := 1; ii < targetGroup; ii++ {
		numTests := 0
		for gen.IsNextCircled(ii){
			numTests++
			len, partialSum := gen.LenOfNextCirInGroup(ii, targetGroup)
			if len > nRemain {
				//fmt.Printf("next value from group %d comes after the target\n", ii)
				gen.UndoUpTo(targetGroup)
				break
			} else {
				nRemain -= len
				//fmt.Printf("next value from group %d comes before the target with %d values accounted for and %d remaining\n", ii, len, nRemain)
				sum = moduloMath.Sum(1e9,sum, partialSum)
				if nRemain == 0 {
					return
				}
			}
		}
		fmt.Printf("moving onto group %d after %d tests\n",ii+1, numTests)
	}
	finalSumStart := gen.NextCircled(targetGroup)
	finalSum := moduloMath.SumRange(1e9,finalSumStart,nRemain)
	fmt.Printf("final group sum over %d sequential values starting at %d\n",nRemain, finalSumStart)
	sum = moduloMath.Sum(1e9,sum,finalSum)
	return
}

func main(){
	xx := uint64(1e18)
    answer := sequenceSum(xx)
	fmt.Printf("sum of first %d values in sequence = %d\n", xx, answer)
}


