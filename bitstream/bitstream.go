/**
shamelessly adapted from https://github.com/dgryski/go-tsz/blob/master/bstream.go

But modified to have public methods and few other addons.

Note: this is not a thread safe implementation, locking should be done externally
*/

package bitstream

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Bit a single bit type
type Bit bool

const (
	// ZeroBit
	ZeroBit Bit = false
	// OneBit
	OneBit Bit = true
)

// BitStream is a stream of bits
type BitStream struct {
	// the data stream
	stream []byte

	// how many bits are valid in current byte
	count uint8

	bitsWritten int
	bitsRead    int64
}

// NewBReader new bit stream reader
func NewBReader(b []byte) *BitStream {
	return &BitStream{stream: b, count: 8, bitsWritten: 0}
}

// NewBWriter new bit stream writer
func NewBWriter(size int) *BitStream {
	return &BitStream{stream: make([]byte, 0, size), count: 0, bitsWritten: 0}
}

// SetStream reset a bitstream from some other byte slice
func (b *BitStream) SetStream(bs []byte) {
	b.stream = bs
}

// Len length in Bytes of the stream (not thread safe)
func (b *BitStream) Len() int {
	return int(b.bitsWritten / 8)
}

// Clone a BitStream in to a new one
func (b *BitStream) Clone() *BitStream {
	d := make([]byte, len(b.stream))
	copy(d, b.stream)
	return &BitStream{stream: d, count: b.count, bitsWritten: len(d)}
}

// Bytes the stream as it stands
func (b *BitStream) Bytes() []byte {
	return b.stream
}

// WriteBit write a single bit
func (b *BitStream) WriteBit(bit Bit) {

	if b.count == 0 {
		b.stream = append(b.stream, 0)
		b.count = 8
	}

	i := len(b.stream) - 1

	if bit {
		b.stream[i] |= 1 << (b.count - 1)
	}
	b.bitsWritten++
	b.count--
}

// WriteBytes write a byte slice
func (b *BitStream) WriteBytes(bs []byte) int {
	c := 0
	for _, by := range bs {
		b.WriteByte(by)
		c++
	}
	return c
}

// WriteByte write a byte to the stream
func (b *BitStream) WriteByte(byt byte) error {

	if b.count == 0 {
		b.stream = append(b.stream, 0)
		b.count = 8
	}

	i := len(b.stream) - 1

	// fill up b.b with b.count bits from byt
	b.stream[i] |= byt >> (8 - b.count)

	b.stream = append(b.stream, 0)
	i++
	b.stream[i] = byt << b.count
	b.bitsWritten += 8
	return nil
}

// WriteBits write nbits to the stream
func (b *BitStream) WriteBits(u uint64, nbits int) {
	u <<= (64 - uint(nbits))
	for nbits >= 8 {
		byt := byte(u >> 56)
		b.WriteByte(byt)
		u <<= 8
		nbits -= 8
	}

	for nbits > 0 {
		b.WriteBit((u >> 63) == 1)
		u <<= 1
		nbits--
	}
	b.bitsWritten += nbits
}

// ReadBit read a single bit
func (b *BitStream) ReadBit() (Bit, error) {

	if len(b.stream) == 0 {
		return false, io.EOF
	}

	if b.count == 0 {
		b.stream = b.stream[1:]
		// did we just run out of stuff to read?
		if len(b.stream) == 0 {
			return false, io.EOF
		}
		b.count = 8
	}

	b.count--
	d := b.stream[0] & 0x80
	b.stream[0] <<= 1
	b.bitsRead++
	return d != 0, nil
}

// ReadBytes read n bytes from the stream
func (b *BitStream) ReadBytes(n uint8) ([]byte, error) {

	if len(b.stream) == 0 {
		return nil, io.EOF
	}
	if len(b.stream) < int(n) {
		return nil, io.EOF
	}

	byts := make([]byte, n)
	var err error
	for i := uint8(0); i < n; i++ {
		byts[i], err = b.ReadByte()
		if err != nil {
			return nil, err
		}
	}
	return byts, nil

}

// ReadByte read a byte from the stream
func (b *BitStream) ReadByte() (byte, error) {

	if len(b.stream) == 0 {
		return 0, io.EOF
	}

	if b.count == 0 {
		b.stream = b.stream[1:]

		if len(b.stream) == 0 {
			return 0, io.EOF
		}

		b.count = 8
	}

	if b.count == 8 {
		b.count = 0
		b.bitsRead += 8
		return b.stream[0], nil
		//b.stream = b.stream[1:]
		//return byt, nil
	}

	byt := b.stream[0]
	b.stream = b.stream[1:]

	if len(b.stream) == 0 {
		return 0, io.EOF
	}

	byt |= b.stream[0] >> b.count
	b.stream[0] <<= (8 - b.count)

	b.bitsRead += 8
	return byt, nil
}

// ReadBits read nbits from the stream
func (b *BitStream) ReadBits(nbits int) (uint64, error) {

	var u uint64

	for nbits >= 8 {
		byt, err := b.ReadByte()
		if err != nil {
			return 0, err
		}

		u = (u << 8) | uint64(byt)
		nbits -= 8
	}

	var err error
	for nbits > 0 && err != io.EOF {
		byt, err := b.ReadBit()
		if err != nil {
			return 0, err
		}
		u <<= 1
		if byt {
			u |= 1
		}
		nbits--
	}
	b.bitsRead += int64(nbits)

	return u, nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface
func (b *BitStream) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, b.count)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, b.stream)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface
func (b *BitStream) UnmarshalBinary(bIn []byte) error {
	buf := bytes.NewReader(bIn)
	err := binary.Read(buf, binary.BigEndian, &b.count)
	if err != nil {
		return err
	}
	b.stream = make([]byte, buf.Len())
	err = binary.Read(buf, binary.BigEndian, &b.stream)
	if err != nil {
		return err
	}
	return nil
}
