package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/asynccnu/food_service/model"

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
func SegWord(str interface{}) (sql string) {
	//如果已经成功加载字典
	var wdslice []string
	if Segmenter.Dictionary() != nil {
		wds := sego.SegmentsToString(Segmenter.Segment([]byte(fmt.Sprintf("%v", str))), true)
		slice := strings.Split(wds, " ")
		for _, wd := range slice {
			w := strings.Split(wd, "/")[0]
			if i, _ := strconv.Atoi(w); i == 0 && w != "0" { //如果为0，则表示非数字
				w = "'%" + w + "%'"
				_ = model.AddNewSearchRecord(w)
				wdslice = append(wdslice, w)
			}
		}
		sql = strings.Join(wdslice, " or name like ")
	}
	return
}
