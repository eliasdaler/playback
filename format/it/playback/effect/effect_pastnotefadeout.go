package effect

import (
	"fmt"

	"github.com/gotracker/playback/format/it/layout/channel"
	"github.com/gotracker/playback/player/intf"
	"github.com/gotracker/playback/song/note"
)

// PastNoteFade defines a past note fadeout effect
type PastNoteFade channel.DataEffect // 'S72'

// Start triggers on the first tick, but before the Tick() function is called
func (e PastNoteFade) Start(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	cs.DoPastNoteEffect(note.ActionFadeout)
	return nil
}

func (e PastNoteFade) String() string {
	return fmt.Sprintf("S%0.2x", channel.DataEffect(e))
}