package load

import (
	"github.com/gotracker/playback/format/settings"
	"github.com/gotracker/playback/format/xm/layout"
	"github.com/gotracker/playback/format/xm/playback"
)

type readerFunc func(filename string, s *settings.Settings) (*layout.Song, error)

func load(filename string, reader readerFunc, s *settings.Settings) (*playback.Manager, error) {
	xmSong, err := reader(filename, s)
	if err != nil {
		return nil, err
	}

	m, err := playback.NewManager(xmSong)

	return m, err
}