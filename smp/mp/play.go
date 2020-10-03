package mp

import "fmt"

//音乐播放模块
//音乐播放模块应该是很容易扩展的，不应该在每次增加一种新音乐文件类型支持时都需要大幅度调整代码
//没有直接将MusicEntry作为参数传入，是因为MusicEntry包含一些多余的信息，本着最小原则
//通过一批类型（比如MP3Player和WAVPlayer）实现这个接口，以达到尽量的架构灵活性
type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)

}






