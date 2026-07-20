package index

import (
	"slices"
	"time"
)

type FileMetadata struct {
	Ctime  time.Time
	Mtime  time.Time
	Device uint64
	Inode  uint64
	Mode   string
	Uid    uint32
	Gid    uint32
	Size   int64
	Flags  []string
	Path   string
}

func (fm FileMetadata) Equal(other FileMetadata) bool {
	if !time.Time.Equal(fm.Ctime, other.Ctime) {
		return false
	}
	if !time.Time.Equal(fm.Mtime, other.Mtime) {
		return false
	}
	if fm.Device != other.Device {
		return false
	}
	if fm.Inode != other.Inode {
		return false
	}
	if fm.Mode != other.Mode {
		return false
	}
	if fm.Uid != other.Uid {
		return false
	}
	if fm.Gid != other.Gid {
		return false
	}
	if fm.Size != other.Size {
		return false
	}
	if !slices.Equal(fm.Flags, other.Flags) {
		return false
	}
	if fm.Path != other.Path {
		return false
	}
	return true
}

type IndexEntry struct {
	FileMetadata
	Hash  string
	Stage int
}

func (ie IndexEntry) Equal(other IndexEntry) bool {
	if !ie.FileMetadata.Equal(other.FileMetadata) {
		return false
	}
	if ie.Hash != other.Hash {
		return false
	}
	if ie.Stage != other.Stage {
		return false
	}
	return true
}

type Index struct {
	Entries      []IndexEntry
	CurrentStage int
}

func GetIndex() *Index {
	i, err := Read()
	if err != nil {
		i = NewIndex()
	}
	return i
}

func NewIndex() *Index {
	return &Index{Entries: []IndexEntry{}, CurrentStage: 0}
}
