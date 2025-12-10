package ftpclient

import (
	"strconv"
	"strings"
)

type FtpEntry struct {
	Perms string
	Links int
	UID   int
	GID   int
	Size  int64
	Month string
	Day   int
	Time  string
	Name  string
	IsDir bool
}

func ParseDir(data string) ([]*FtpEntry, error) {
	var items []*FtpEntry

	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)

		parts := strings.Fields(line)

		if len(parts) < 9 {
			continue
		}

		perms := parts[0]
		isDir := strings.HasPrefix(perms, "d")
		links, _ := strconv.Atoi(parts[1])
		uid, _ := strconv.Atoi(parts[2])
		gid, _ := strconv.Atoi(parts[3])
		size, _ := strconv.ParseInt(parts[4], 10, 64)
		day, _ := strconv.Atoi(parts[6])

		name := strings.Join(parts[8:], " ")

		items = append(items, &FtpEntry{
			Perms: perms,
			Links: links,
			UID:   uid,
			GID:   gid,
			Size:  size,
			Month: parts[5],
			Day:   day,
			Time:  parts[7],
			Name:  name,
			IsDir: isDir,
		})
	}

	return items, nil
}
