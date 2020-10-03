package lib

import "errors"

//音乐库
type MusicManager struct {
	musics []MusicEntry
}

//声明一个音乐库
func NewMusicManager() *MusicManager {

	return &MusicManager{make([]MusicEntry, 0)}

}

//返回音乐库的音乐数目
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//根据索引获取指定的音乐
func (m *MusicManager) Get(index int) ( music *MusicEntry, err error) {
	if index < 0 || index > len(m.musics) {
		return nil, errors.New("Index out of range.")
	}

	return &m.musics[index],nil

}

//根据音乐名查找音乐
func (m *MusicManager) Find(name string) (music *MusicEntry) {

	if len(m.musics) == 0 {
		return nil
	}

	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}

	return nil
}

//添加
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//删除
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removeMusic := &m.musics[index]

	m.musics = append(m.musics[:index], m.musics[index+1:]...)

	return removeMusic
}

