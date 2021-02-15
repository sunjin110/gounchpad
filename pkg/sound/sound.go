package sound

import (
	"fmt"
	"gounchpad/pkg/common/chk"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var soundKeyMap map[rune]string = map[rune]string{
	'a': "./sound/effect/rion.mp3",
}

// Sound 対応する音をならす
func Sound(key rune) {

	filePath, exists := soundKeyMap[key]
	if !exists {
		fmt.Println("not found sound path...")
		return
	}

	f, err := os.Open(filePath)
	chk.SE(err)

	streamer, format, err := mp3.Decode(f)
	chk.SE(err)

	defer streamer.Close()
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	chk.SE(err)

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(
		func() {
			done <- true
		},
	)))

	<-done

}
