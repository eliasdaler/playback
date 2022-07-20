package effect

import (
	"fmt"

	"github.com/gotracker/playback/format/it/layout/channel"
	"github.com/gotracker/playback/player/intf"
)

// FinePortaDown defines an fine portamento down effect
type FinePortaDown channel.DataEffect // 'EFx'

// Start triggers on the first tick, but before the Tick() function is called
func (e FinePortaDown) Start(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	cs.ResetRetriggerCount()
	cs.UnfreezePlayback()

	mem := cs.GetMemory()
	y := mem.PortaDown(channel.DataEffect(e)) & 0x0F

	return doPortaDown(cs, float32(y), 4, mem.Shared.LinearFreqSlides)
}

func (e FinePortaDown) String() string {
	return fmt.Sprintf("E%0.2x", channel.DataEffect(e))
}