package effect

import (
	"fmt"

	xmPanning "github.com/gotracker/playback/format/xm/conversion/panning"
	"github.com/gotracker/playback/format/xm/layout/channel"
	"github.com/gotracker/playback/player/intf"
)

// SetCoarsePanPosition defines a set pan position effect
type SetCoarsePanPosition channel.DataEffect // 'E8x'

// Start triggers on the first tick, but before the Tick() function is called
func (e SetCoarsePanPosition) Start(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	cs.ResetRetriggerCount()

	xy := channel.DataEffect(e)
	y := xy & 0x0F

	cs.SetPan(xmPanning.PanningFromXm(uint8(y) << 4))
	return nil
}

func (e SetCoarsePanPosition) String() string {
	return fmt.Sprintf("E%0.2x", channel.DataEffect(e))
}