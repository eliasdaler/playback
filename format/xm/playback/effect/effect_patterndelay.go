package effect

import (
	"fmt"

	"github.com/gotracker/playback/format/xm/layout/channel"
	effectIntf "github.com/gotracker/playback/format/xm/playback/effect/intf"
	"github.com/gotracker/playback/player/intf"
)

// PatternDelay defines a pattern delay effect
type PatternDelay channel.DataEffect // 'SEx'

// PreStart triggers when the effect enters onto the channel state
func (e PatternDelay) PreStart(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	m := p.(effectIntf.XM)
	return m.SetPatternDelay(int(channel.DataEffect(e) & 0x0F))
}

// Start triggers on the first tick, but before the Tick() function is called
func (e PatternDelay) Start(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	cs.ResetRetriggerCount()
	return nil
}

func (e PatternDelay) String() string {
	return fmt.Sprintf("S%0.2x", channel.DataEffect(e))
}