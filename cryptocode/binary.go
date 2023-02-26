package cryptocode

import "encoding/binary"

func Int64ToBytes(i int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, i)
	return buf[:n]
}

func BytesToInt64(b []byte) int64 {
	i, _ := binary.Varint(b)
	return i
}
