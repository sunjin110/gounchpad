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

func init() {
	// 初期化
	err := speaker.Init(44100, 256)
	chk.SE(err)

	// https://game.criware.jp/learn/tutorial/unity/unity_tyukyu_03/
	// オンメモリ再生を実現したい
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

	// https://qiita.com/mfykmn/items/1dca6630f83f39582abf
	// 一つのstreamから複数回読み込みをしたい...

	// filePathからいちいち持ってくるのコスト掛かりそうだから、増えてきたらinitでstreamをmemoriyに上げて
	// 消費するときは、copyして使用するようにしたい

	streamer, _, err := mp3.Decode(f)
	streamer.Seek(0)

	chk.SE(err)
	defer streamer.Close()
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(
		func() {
			done <- true
		},
	)))
	// fmt.Printf("start:%s", filePath)
	<-done
	// fmt.Printf("end:%s", filePath)
}
