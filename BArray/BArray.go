package BArray

import (
	"encoding/binary"
	"math"
	"strconv"
)

type BArray struct {
	Data []byte
}

// func ReadStruct(b []byte, w reflect.Value) {
// 	index := 0
// 	for i := 0; i < w.Elem().NumField(); i++ {
// 		v := w.Elem().Field(i)
// 		switch v.Kind() {
// 		case reflect.Ptr:
// 			v = v.Elem()
// 		}
// 		//log.Println("getValue.Field(i)", getValue.Field(i))
// 		switch v.Kind() {
// 		case reflect.Uint8:
// 			r := ReadByte(b, &index)
// 			fmt.Println("Uint8:", r)
// 			v.SetUint(uint64(r))
// 		case reflect.Uint16:
// 			r := ReadUInt16(b, &index)
// 			fmt.Println("Uint16:", r)
// 			v.SetUint(uint64(r))
// 		case reflect.Uint32:
// 			r := ReadUInt32(b, &index)
// 			fmt.Println("Uint32:", r)
// 			v.SetUint(uint64(r))
// 		case reflect.Uint64:
// 			r := ReadUInt64(b, &index)
// 			fmt.Println("Uint64:", r)
// 			v.SetUint(uint64(r))
// 		// case reflect.Int8:

// 		case reflect.Int16:
// 			r := ReadInt16(b, &index)
// 			fmt.Println("int16:", r)
// 			v.SetInt(int64(r))
// 		case reflect.Int32:
// 			r := ReadInt32(b, &index)
// 			fmt.Println("int32:", r)
// 			v.SetInt(int64(r))
// 		case reflect.Int64:
// 			r := ReadInt64(b, &index)
// 			fmt.Println("int16:", r)
// 			v.SetInt(int64(r))
// 		case reflect.Bool:
// 			r := ReadBool(b, &index)
// 			fmt.Println("Bool:", r)
// 			v.SetBool(r)
// 		case reflect.String:
// 			r := ReadStringAndWordLen(b, &index)
// 			fmt.Println("String:", r)
// 			v.SetString(r)
// 		case reflect.Array:
// 			r := ReadByteArrayAndWordLen(b, &index)
// 			fmt.Println("Array:", r)
// 			v.SetBytes(r)
// 		case reflect.Slice:
// 			r := ReadByteArrayAndWordLen(b, &index)
// 			fmt.Println("Slice:", r)
// 			v.SetBytes(r)
// 		case reflect.Struct:
// 			// r := ReadByteArray(b, &index, len(b))
// 			// fmt.Println("Struct:", r)

// 			// ReadStruct(b, w.Elem().Field(7))
// 		case reflect.Ptr:
// 			r := ReadByteArray(b, &index, len(b))
// 			fmt.Println("Ptr:", r)
// 			ReadStruct(r, v)
// 		default:
// 			fmt.Println("非預期之型態:", w.Elem().Field(i).Kind())
// 		}
// 	}
// }

func ReadBool(b []byte, i *int) (r bool) {
	prei := *i
	newi := *i + 1
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r, _ = strconv.ParseBool(string(b[prei]))
	*i = newi
	return
}

func ReadByte(b []byte, i *int) (r byte) {
	newi := *i + 1
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r = b[*i]
	*i = newi
	return
}

func ReadUInt16(b []byte, i *int) (r uint16) {
	prei := *i
	newi := *i + 2
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r = binary.LittleEndian.Uint16(b[prei:newi])
	*i = newi
	return
}

func ReadUInt32(b []byte, i *int) (r uint32) {
	prei := *i
	newi := *i + 4
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r = binary.LittleEndian.Uint32(b[prei:newi])
	*i = newi
	return
}

func ReadUInt64(b []byte, i *int) (r uint64) {
	prei := *i
	newi := *i + 8
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r = binary.LittleEndian.Uint64(b[prei:newi])
	*i = newi
	return
}

func ReadInt16(b []byte, i *int) (r int16) {
	prei := *i
	newi := *i + 2
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r = int16(binary.LittleEndian.Uint16(b[prei:newi]))
	*i = newi
	return
}

func ReadInt32(b []byte, i *int) (r int32) {
	prei := *i
	newi := *i + 4
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r = int32(binary.LittleEndian.Uint32(b[prei:newi]))
	*i = newi
	return
}

func ReadInt64(b []byte, i *int) (r int64) {
	prei := *i
	newi := *i + 8
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r = int64(binary.LittleEndian.Uint64(b[prei:newi]))
	*i = newi
	return
}

func ReadFloat32(b []byte, i *int) (r float32) {
	prei := *i
	newi := *i + 4
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r = math.Float32frombits(binary.LittleEndian.Uint32(b[prei:newi]))
	*i = newi
	return
}

func ReadFloat64(b []byte, i *int) (r float64) {
	prei := *i
	newi := *i + 8
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	r = math.Float64frombits(binary.LittleEndian.Uint64(b[prei:newi]))
	*i = newi
	return
}

func ReadStringAndByteLen(b []byte, i *int) (r string) {
	prei := *i
	newi := *i + 1
	if b == nil {
		return
	} else if len(b) < newi {
		return
	}

	i8 := ReadByte(b, i)
	prei = newi
	newi = int(i8*2) + prei
	r = string(b[prei:newi])
	*i = newi
	return
}

func ReadStringAndWordLen(b []byte, i *int) (r string) {
	prei := *i
	newi := *i + 2
	if b == nil {
		return
	} else if len(b) <= newi {
		return
	}

	i16 := ReadUInt16(b, i)
	prei = newi
	newi = int(i16*2) + prei
	r = string(b[prei:newi])
	*i = newi
	return
}

func ReadStringAndIntLen(b []byte, i *int) (r string) {
	prei := *i
	newi := *i + 4
	if b == nil {
		return
	} else if len(b) <= newi {
		return
	}

	i32 := ReadUInt32(b, i)
	prei = newi
	newi = int(i32*2) + prei
	r = string(b[prei:newi])
	*i = newi
	return
}

func ReadByteArray(b []byte, i *int, s int) (b2 []byte) {
	if b == nil {
		return
	} else if len(b) <= 0 {
		return
	}

	b2 = b[*i:s]
	*i = *i + s
	return
}

func ReadByteArrayAndWordLen(b []byte, i *int) (b2 []byte) {
	prei := *i
	newi := *i + 2

	if b == nil {
		return
	} else if len(b) <= newi {
		return
	}

	b2 = b[prei:newi]
	*i = newi
	return
}

func WriteBool(b []byte, w bool) (r []byte) {
	if w {
		r = append(b, []byte{1}...)
	} else {
		r = append(b, []byte{0}...)
	}

	return
}

func WriteByte(b []byte, w byte) (r []byte) {
	r = append(b, w)
	return
}

func WriteUInt16(b []byte, w uint16) (r []byte) {
	r = make([]byte, 2)
	binary.LittleEndian.PutUint16(r, w)
	r = append(b, r...)
	return
}

func WriteUInt32(b []byte, w uint32) (r []byte) {
	r = make([]byte, 4)
	binary.LittleEndian.PutUint32(r, w)
	r = append(b, r...)
	return
}

func WriteUInt64(b []byte, w uint64) (r []byte) {
	r = make([]byte, 8)
	binary.LittleEndian.PutUint64(r, w)
	r = append(b, r...)
	return
}

func WriteInt16(b []byte, w int16) (r []byte) {
	r = make([]byte, 2)
	binary.LittleEndian.PutUint16(r, uint16(w))
	r = append(b, r...)
	return
}

func WriteInt32(b []byte, w int32) (r []byte) {
	r = make([]byte, 4)
	binary.LittleEndian.PutUint32(r, uint32(w))
	r = append(b, r...)
	return
}

func WriteInt64(b []byte, w int64) (r []byte) {
	r = make([]byte, 8)
	binary.LittleEndian.PutUint64(r, uint64(w))
	r = append(b, r...)
	return
}

func WriteFloat32(b []byte, w float32) (r []byte) {
	r = make([]byte, 4)
	binary.LittleEndian.PutUint32(b, math.Float32bits(w))
	r = append(b, r...)
	return
}

func WriteFloat64(b []byte, w float64) (r []byte) {
	r = make([]byte, 8)
	binary.LittleEndian.PutUint64(b, math.Float64bits(w))
	r = append(b, r...)
	return
}

func WriteStringAndByteLen(b []byte, w string) (r []byte) {
	S := []byte(w)
	r = WriteByte(b, byte(len(S)))
	r = append(b, S...)
	return
}

func WriteStringAndWordLen(b []byte, w string) (r []byte) {
	S := []byte(w)
	r = WriteInt16(b, int16(len(S)))
	r = append(b, S...)
	return
}

func WriteStringAndIntLen(b []byte, w string) (r []byte) {
	S := []byte(w)
	r = WriteInt32(b, int32(len(S)))
	r = append(b, S...)
	return
}

func WriteByteArray(b []byte, b2 []byte) (r []byte) {
	r = append(b, b2...)
	return
}

func WriteByteArrayAndWordLen(b []byte, b2 []byte) (r []byte) {
	r = WriteUInt16(b, uint16(2))
	r = WriteByteArray(r, b2)
	return
}
