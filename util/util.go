package util

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/huichen/sego"
	"github.com/teris-io/shortid"
)

var (
	Segmenter   sego.Segmenter
	BasePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
)

func init() {
	//加载分词字典
	go func() {
		Segmenter.LoadDictionary(BasePath + "/dictionary/dictionary.txt")
	}()
}

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestID, ok := v.(string); ok {
		return requestID
	}
	return ""
}

//SegWord 分词
//param            str         需要分词的文字
func SegWord(str string) []string {
	var wdslice []string
	//如果已经成功加载字典
	if Segmenter.Dictionary() != nil {
		wds := sego.SegmentsToSlice(Segmenter.Segment([]byte(str)), true)
		for _, w := range wds {
			if i, _ := strconv.Atoi(w); i == 0 && w != "0" { //如果为0，则表示非数字
				//w = "'%" + w + "%'"
				wdslice = append(wdslice, w)
			}
		}
		//sql = strings.Join(wdslice, " or name like ")
	}

	return wdslice
}
