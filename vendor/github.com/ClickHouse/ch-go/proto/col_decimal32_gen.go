// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColDecimal32 represents Decimal32 column.
type ColDecimal32 []Decimal32

// Compile-time assertions for ColDecimal32.
var (
	_ ColInput  = ColDecimal32{}
	_ ColResult = (*ColDecimal32)(nil)
	_ Column    = (*ColDecimal32)(nil)
)

// Rows returns count of rows in column.
func (c ColDecimal32) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDecimal32) Reset() {
	*c = (*c)[:0]
}

// Type returns ColumnType of Decimal32.
func (ColDecimal32) Type() ColumnType {
	return ColumnTypeDecimal32
}

// Row returns i-th row of column.
func (c ColDecimal32) Row(i int) Decimal32 {
	return c[i]
}

// Append Decimal32 to column.
func (c *ColDecimal32) Append(v Decimal32) {
	*c = append(*c, v)
}

// LowCardinality returns LowCardinality for Decimal32 .
func (c *ColDecimal32) LowCardinality() *ColLowCardinality[Decimal32] {
	return &ColLowCardinality[Decimal32]{
		index: c,
	}
}

// Array is helper that creates Array of Decimal32.
func (c *ColDecimal32) Array() *ColArr[Decimal32] {
	return &ColArr[Decimal32]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable(Decimal32).
func (c *ColDecimal32) Nullable() *ColNullable[Decimal32] {
	return &ColNullable[Decimal32]{
		Values: c,
	}
}

// NewArrDecimal32 returns new Array(Decimal32).
func NewArrDecimal32() *ColArr[Decimal32] {
	return &ColArr[Decimal32]{
		Data: new(ColDecimal32),
	}
}
