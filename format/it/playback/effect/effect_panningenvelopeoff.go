package effect

import (
	"fmt"

	"github.com/gotracker/playback/format/it/layout/channel"
	"github.com/gotracker/playback/player/intf"
)

// PanningEnvelopeOff defines a panning envelope: off effect
type PanningEnvelopeOff channel.DataEffect // 'S79'

// Start triggers on the first tick, but before the Tick() function is called
func (e PanningEnvelopeOff) Start(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	cs.ResetRetriggerCount()

	cs.SetPanningEnvelopeEnable(false)
	return nil
}

func (e PanningEnvelopeOff) String() string {
	return fmt.Sprintf("S%0.2x", channel.DataEffect(e))
}