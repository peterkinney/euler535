package sequenceGenerator

import (
	"testing"
)

func TestLastGroup(tt *testing.T) {
	cases := [][]uint64 {
					[]uint64 {2,0,1},
					[]uint64 {3,1,2},
					[]uint64 {4,1,2},
					[]uint64 {5,2,4},
					[]uint64 {8,2,4},
					[]uint64 {9,3,8},
					[]uint64 {17,3,8},
					[]uint64 {18,4,17},
					[]uint64 {1000,7,867},
					[]uint64 {1e9,10,709584267},
					[]uint64 {1e18,11,12574247427901},
				}
	for _, vv := range cases {
		want1, want2 := vv[1],vv[2]
		got1,got2 := LastGroup(vv[0])
		if want1 != uint64(got1) || want2 != got2 {
			tt.Errorf("LastGroup(%d) Failed: want %d,%d, got %d,%d",vv[0],want1,want2,got1,got2)
		}
	}
}

func TestPeek(tt *testing.T){
	gen := NewGenerator()
	want := []uint64 {1,2,3,5,10}
	for ii,val := range want {
		got,_,_ := gen.groups[ii].peek()
		if got != val {
			tt.Errorf("groups[%d].peek() Failed: want %d got %d\n",ii,val,got)
		}
	}
}

func TestNextValue(tt *testing.T){
	
}

func TestLenOfNextCirInGroup(tt *testing.T) {
	gen := NewGenerator()
	cases := [][]uint64 {
					[]uint64 {1,1,1,2},
					[]uint64 {1,2,2,5},
					[]uint64 {1,3,4,16},
					[]uint64 {1,4,10,91},
					[]uint64 {2,2,1,3},
					[]uint64 {2,3,2,8},
					[]uint64 {2,4,5,41},
				}
	for _, vv := range cases {
		want1, want2 := vv[2],vv[3]
		got1,got2 := gen.LenOfNextCirInGroup(int(vv[0]),int(vv[1]))
		gen.UndoUpTo(int(vv[1]))
		if want1 != got1 || want2 != got2 {
			tt.Errorf("gen.LenOfNextCirInGroup(%d,%d) Failed: want %d,%d, got %d,%d",vv[0], vv[1],want1,want2,got1,got2)
		}
	}
	gen.LenOfNextCirInGroup(1,4)
	cases = [][]uint64 {
					[]uint64 {2,2,1,4},
					[]uint64 {2,3,3,19},
					[]uint64 {2,4,9,130},
				}
	for _, vv := range cases {
		want1, want2 := vv[2],vv[3]
		got1,got2 := gen.LenOfNextCirInGroup(int(vv[0]),int(vv[1]))
		gen.UndoUpTo(int(vv[1]))
		if want1 != got1 || want2 != got2 {
			tt.Errorf("gen.LenOfNextCirInGroup(%d,%d) Failed: want %d,%d, got %d,%d",vv[0], vv[1],want1,want2,got1,got2)
		}
	}
}
