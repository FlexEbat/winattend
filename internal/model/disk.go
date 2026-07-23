package model

type Disk struct {
	TargetDisk int

	Layout DiskLayout

	WipeDisk bool

	Recovery RecoveryPartition

	Assertions DiskAssertions
}

type RecoveryPartition struct {
	Enabled bool

	SizeMB int
}

type DiskAssertions struct {
	MinSizeGB int

	MaxSizeGB int

	RequireEmpty bool

	RequireFixed bool
}
