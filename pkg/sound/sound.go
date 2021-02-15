package sound

import (
	"fmt"
	"gounchpad/pkg/common/chk"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var soundKeyMap map[rune]string = map[rune]string{
	// 'a': "./sound/effect/rion.mp3",
	'a': "./sound/effect/c1.mp3",
	's': "./sound/effect/d1.mp3",
	'd': "./sound/effect/e1.mp3",
	'f': "./sound/effect/f1.mp3",
	'g': "./sound/effect/g1.mp3",
	'h': "./sound/effect/a1.mp3",
	'j': "./sound/effect/b1.mp3",
}

// SoundStreamer .
// type SoundStreamer struct {
// 	Streamer beep.StreamCloser
// 	Format   beep.Format
// }

// var soundStreamerMap map[rune]*SoundStreamer

func init() {

	// log.Println("sound init...")
	//
	// soundStreamerMap = make(map[rune]*SoundStreamer, len(soundKeyMap))

	// // 先にロードしておく
	// for r, filePath := range soundKeyMap {
	// 	f, err := os.Open(filePath)
	// 	chk.SE(err)
	// 	streamer, format, err := mp3.Decode(f)
	// 	soundStreamerMap[r] = &SoundStreamer{
	// 		Streamer: streamer,
	// 		Format:   format,
	// 	}
	// }

	// err := speaker.Init(44100, bufferSize int)

	err := speaker.Init(44100, 441)
	chk.SE(err)

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

	streamer, _, err := mp3.Decode(f)
	// streamer, format, err := mp3.Decode(f)
	chk.SE(err)

	// soundStreamer, exists := soundStreamerMap[key]
	// if !exists {
	// fmt.Println("not found sound path...")
	// return
	// }

	defer streamer.Close()
	// format := soundStreamer.Format

	// log.Println("format is ", jsonutil.Marshal(format))
	// log.Println("sample rate is ", format.SampleRate)
	// log.Println("n is ", format.SampleRate.N(time.Second/100))
	// err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/100))
	// chk.SE(err)

	// streamer := soundStreamer.Streamer

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(
		func() {
			done <- true
		},
	)))
	fmt.Println("start")
	<-done
	fmt.Println("end")
}
