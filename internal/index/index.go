package index

import (
	"encoding/gob"
	"fmt"
	"os"
	"slices"
	"syscall"
	"time"

	"github.com/pkg/xattr"
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

func GenerateEntry(path, hash string, stage int) (*IndexEntry, error) {
	fi, err := os.Stat(path)
	stats := fi.Sys().(*syscall.Stat_t)
	if err != nil {
		return nil, fmt.Errorf("could not stat %s: %w", path, err)
	}
	flags, err := xattr.List(path)
	if err != nil {
		return nil, fmt.Errorf("could not get %s flags: %w", path, err)
	}

	ie := IndexEntry{
		FileMetadata: FileMetadata{
			Ctime:  time.Unix(stats.Ctim.Sec, stats.Ctim.Nsec),
			Mtime:  fi.ModTime().UTC(),
			Device: stats.Dev,
			Inode:  stats.Ino,
			Mode:   fi.Mode().String(),
			Uid:    stats.Uid,
			Gid:    stats.Gid,
			Size:   fi.Size(),
			Flags:  flags,
			Path:   path,
		},
		Hash:  hash,
		Stage: stage,
	}

	return &ie, nil
}

func AddFile(i *Index, path, hash string) error {
	ie, err := GenerateEntry(path, hash, i.CurrentStage)
	if err != nil {
		return err
	}
	if !EntryInIndex(ie, i.Entries) {
		i.Entries = append(i.Entries, *ie)
		i.CurrentStage = i.CurrentStage + 1 // increment Index current stage
	}
	return nil
}

func EntryInIndex(to_add *IndexEntry, entries []IndexEntry) bool {
	return slices.ContainsFunc(entries, func(ie IndexEntry) bool {
		return ie.Equal(*to_add)
	})
}

func PathInIndex(i *Index, path string) bool {
	return slices.ContainsFunc(i.Entries, func(ie IndexEntry) bool {
		return ie.Path == path
	})
}

func GetEntry(i *Index, path string) *IndexEntry {
	ie_idx := slices.IndexFunc(i.Entries, func(ie IndexEntry) bool {
		return ie.Path == path
	})

	if ie_idx != -1 {
		return &i.Entries[ie_idx]
	}
	return nil
}

func IsChanged(ie *IndexEntry, path string) bool {
	fi, _ := os.Stat(path)
	stats := fi.Sys().(*syscall.Stat_t)

	if !time.Time.Equal(fi.ModTime(), ie.Mtime) {
		return true
	}
	if stats.Dev != ie.Device {
		return true
	}
	if stats.Ino != ie.Inode {
		return true
	}
	if fi.Size() != ie.Size {
		return true
	}
	return false
}

func Read() (*Index, error) {
	file, err := os.Open(".gogit/index")
	if err != nil {
		return nil, fmt.Errorf("could not open index file: %w", err)
	}
	encoder := gob.NewDecoder(file)
	var i Index
	if err := encoder.Decode(&i); err != nil {
		return nil, err
	}
	return &i, nil
}

func Write(i *Index) error {
	file, err := os.OpenFile(".gogit/index", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("could not open index file: %w", err)
	}
	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(i); err != nil {
		return err
	}
	return nil
}
