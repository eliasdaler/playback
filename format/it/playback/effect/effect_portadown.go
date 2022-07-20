package effect

import (
	"fmt"

	"github.com/gotracker/playback/format/it/layout/channel"
	"github.com/gotracker/playback/player/intf"
)

// PortaDown defines a portamento down effect
type PortaDown channel.DataEffect // 'E'

// Start triggers on the first tick, but before the Tick() function is called
func (e PortaDown) Start(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	cs.ResetRetriggerCount()
	cs.UnfreezePlayback()
	return nil
}

// Tick is called on every tick
func (e PortaDown) Tick(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback, currentTick int) error {
	mem := cs.GetMemory()
	xx := mem.PortaDown(channel.DataEffect(e))

	return doPortaDown(cs, float32(xx), 4, mem.Shared.LinearFreqSlides)
}

func (e PortaDown) String() string {
	return fmt.Sprintf("E%0.2x", channel.DataEffect(e))
}