package effect

import (
	"fmt"

	xmVolume "github.com/gotracker/playback/format/xm/conversion/volume"
	"github.com/gotracker/playback/format/xm/layout/channel"
	"github.com/gotracker/playback/player/intf"
)

// SetVolume defines a volume slide effect
type SetVolume channel.DataEffect // 'C'

// Start triggers on the first tick, but before the Tick() function is called
func (e SetVolume) Start(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	cs.ResetRetriggerCount()

	xx := xmVolume.XmVolume(e)

	cs.SetActiveVolume(xx.Volume())
	return nil
}

func (e SetVolume) String() string {
	return fmt.Sprintf("C%0.2x", channel.DataEffect(e))
}