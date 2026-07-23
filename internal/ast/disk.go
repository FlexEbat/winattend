package ast

type Disk struct {

	TargetDisk int

	Layout DiskLayout

	Recovery RecoveryPartition

	Assertions DiskAssertions
}

type DiskLayout string

const (

	GPT DiskLayout = "gpt"

	MBR DiskLayout = "mbr"

	Auto DiskLayout = "auto"
)

type RecoveryPartition struct {

	Enabled bool

	SizeMB int
}

type DiskAssertions struct {

	MinSizeGB int

	MaxSizeGB int

	Empty bool

	Fixed bool
}
