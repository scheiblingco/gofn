package typetools_test

type CustomIntType int

const CustomInt = CustomIntType(123)

type CustomInt32Type int32

const CustomInt32 = CustomInt32Type(123)

type CustomInt64Type int64

const CustomInt64 = CustomInt64Type(123)

type CustomUintType uint

const CustomUint = CustomUintType(123)

type CustomUint32Type uint32

const CustomUint32 = CustomUint32Type(123)

type CustomUint64Type uint64

const CustomUint64 = CustomUint64Type(123)

type CustomFloat32Type float32

const CustomFloat32 = CustomFloat32Type(123.45)

type CustomFloat64Type float64

const CustomFloat64 = CustomFloat64Type(123.45)

type CustomUint8Type uint8

const CustomUint8 = CustomUint8Type(123)

type CustomUint16Type uint16

const CustomUint16 = CustomUint16Type(123)

type CustomUintptrType uintptr

const CustomUintptr = CustomUintptrType(123)

type CustomByteType byte

const CustomByte = CustomByteType(123)

type CustomRuneType rune

const CustomRune = CustomRuneType('a')

type CustomStringType string

const CustomString = CustomStringType("abc")

type CustomBytesliceType []byte

var CustomByteslice = CustomBytesliceType([]byte("abc"))

type CustomRunesliceType []rune

var CustomRuneslice = CustomRunesliceType([]rune("abc"))

type CustomBoolType bool

const CustomBool = CustomBoolType(true)
