package effect

import (
	"fmt"

	"github.com/gotracker/playback/format/s3m/layout/channel"
	"github.com/gotracker/playback/player/intf"
)

// FinePortaDown defines an fine portamento down effect
type FinePortaDown ChannelCommand // 'EFx'

// Start triggers on the first tick, but before the Tick() function is called
func (e FinePortaDown) Start(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	cs.ResetRetriggerCount()
	cs.UnfreezePlayback()

	y := channel.DataEffect(e) & 0x0F

	return doPortaDown(cs, float32(y), 4)
}

func (e FinePortaDown) String() string {
	return fmt.Sprintf("E%0.2x", channel.DataEffect(e))
}