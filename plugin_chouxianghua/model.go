package chouxianghua

type pinyin struct {
	Word   string `db:"word"`
	Pronun string `db:"pronunciation"`
}
type emoji struct {
	Pronun string `db:"pronunciation"`
	Emoji  string `db:"emoji"`
}

func getPinyinByWord(word string) string {
	var p pinyin
	db.Find("pinyin", &p, "where word = '"+word+"'")
	return p.Pronun
}

func getPronunByDWord(w0, w1 rune) string {
	return getPinyinByWord(string(w0)) + getPinyinByWord(string(w1))
}

func getEmojiByPronun(pronun string) string {
	var e emoji
	db.Find("emoji", &e, "where pronunciation = '"+pronun+"'")
	return e.Emoji
}
