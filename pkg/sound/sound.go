package sound

import (
	"fmt"
	"gounchpad/pkg/common/chk"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// 文字に対してなにを鳴らすか?
var soundKeyMap map[rune]string = map[rune]string{
	// piano
	'a': "./sound/effect/c1.mp3",
	's': "./sound/effect/d1.mp3",
	'd': "./sound/effect/e1.mp3",
	'f': "./sound/effect/f1.mp3",
	'g': "./sound/effect/g1.mp3",
	'h': "./sound/effect/a1.mp3",
	'j': "./sound/effect/b1.mp3",
	// drum
	'z': "./sound/drum/bass.mp3",
	// skrillex
	'c': "./sound/skrillex/sound_1_v2.mp3",
	'b': "./sound/skrillex/sound_2.mp3",
	// 'f': "./sound/skrillex/sound_3.mp3",
}

// オンメモリ再生
var soundBufferMap map[rune]*beep.Buffer

func init() {
	// 初期化
	err := speaker.Init(44100, 256)
	chk.SE(err)

	soundBufferMap = make(map[rune]*beep.Buffer, len(soundKeyMap))
	for r, filePath := range soundKeyMap {

		f, err := os.Open(filePath)
		chk.SE(err)

		streamer, format, err := mp3.Decode(f)
		chk.SE(err)

		buffer := beep.NewBuffer(format)
		buffer.Append(streamer)

		// buffer mapに追加
		soundBufferMap[r] = buffer

		streamer.Close()
	}

}

// Sound 対応する音をならす
func Sound(key rune) {

	buffer, exists := soundBufferMap[key]
	if !exists {
		fmt.Println("not found sound...")
		return
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(buffer.Streamer(0, buffer.Len()), beep.Callback(
		func() {
			done <- true
		},
	)))
	<-done
}
