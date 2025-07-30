package utils

import (
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// 数字でソートするための比較関数
func NumericFileSort(files []string) {
	sort.Slice(files, func(i, j int) bool {
		// ファイル名のみ取得（例: "10.json"）
		baseI := filepath.Base(files[i])
		baseJ := filepath.Base(files[j])

		// 拡張子除去（例: "10"）
		nameI := strings.TrimSuffix(baseI, filepath.Ext(baseI))
		nameJ := strings.TrimSuffix(baseJ, filepath.Ext(baseJ))

		// 数字に変換（失敗したら0にする）
		numI, err1 := strconv.Atoi(nameI)
		numJ, err2 := strconv.Atoi(nameJ)
		if err1 != nil {
			numI = 0
		}
		if err2 != nil {
			numJ = 0
		}

		return numI < numJ
	})
}
