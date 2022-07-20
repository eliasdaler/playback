package song

import (
	"github.com/gotracker/playback/song/index"
	"github.com/gotracker/playback/song/instrument"
	"github.com/gotracker/playback/song/note"
)

// Data is an interface to the song data
type Data interface {
	GetOrderList() []index.Pattern
	IsChannelEnabled(int) bool
	GetOutputChannel(int) int
	NumInstruments() int
	IsValidInstrumentID(instrument.ID) bool
	GetInstrument(instrument.ID) (*instrument.Instrument, note.Semitone)
	GetName() string
}

type PatternData[TChannelData any] interface {
	GetPattern(index.Pattern) Pattern[TChannelData]
}