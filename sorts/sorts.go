/*

The missing "sort" types that golang, for some reason, decided to leave out.

*/

package sorts

type Int64List []int64

func (p Int64List) Len() int { return len(p) }
func (p Int64List) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Int64List) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Int32List []int32

func (p Int32List) Len() int { return len(p) }
func (p Int32List) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Int32List) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Int16List []int16

func (p Int16List) Len() int { return len(p) }
func (p Int16List) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Int16List) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Int8List []int8

func (p Int8List) Len() int { return len(p) }
func (p Int8List) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Int8List) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Uint64List []uint64

func (p Uint64List) Len() int { return len(p) }
func (p Uint64List) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Uint64List) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Uint32List []uint32

func (p Uint32List) Len() int { return len(p) }
func (p Uint32List) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Uint32List) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Uint16List []uint16

func (p Uint16List) Len() int { return len(p) }
func (p Uint16List) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Uint16List) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Uint8List []uint8

func (p Uint8List) Len() int { return len(p) }
func (p Uint8List) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Uint8List) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
