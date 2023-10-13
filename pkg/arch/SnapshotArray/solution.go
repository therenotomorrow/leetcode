package SnapshotArray

type Snapshot struct {
	id  int
	val int
}

type SnapshotArray struct {
	snap [][]Snapshot
	id   int
}

func Constructor(length int) SnapshotArray {
	snap := make([][]Snapshot, length)

	for i := 0; i < length; i++ {
		snap[i] = []Snapshot{{}}
	}

	return SnapshotArray{snap: snap, id: 0}
}

func (array *SnapshotArray) Set(index int, val int) {
	snap := array.snap[index]
	if array.id == 0 {
		snap[0].val = val
	}

	array.snap[index] = append(array.snap[index], Snapshot{id: array.id, val: val})
}

func (array *SnapshotArray) Snap() int {
	array.id++
	return array.id - 1
}

func (array *SnapshotArray) Get(index int, snapID int) int {
	snap := array.snap[index]

	for i := len(snap) - 1; i >= 0; i-- {
		if snap[i].id == snapID {
			return snap[i].val
		}

		if snap[i].id < snapID {
			return snap[i].val
		}
	}

	return snap[len(array.snap[index])-1].val
}
