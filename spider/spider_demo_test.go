package spider

import (
	"strconv"
	"testing"
)

func TestGetMovieTopList(t *testing.T) {
	url := "https://movie.douban.com/top250?start="
	for i := 0; i < 10; i++ {
		s := i * 25
		l, err := GetTopList(url + strconv.Itoa(s))
		if err != nil {
			t.Fatal("Error get top list:", err)
		}
		for _, v := range l {
			err = GetMovie(v)
			if err != nil {
				t.Fatal("Error get movie:", err)
			}
		}
	}
}