package hash

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

func GenerateMd5String() string {
	now := time.Now().Unix()
	h := md5.New()
	_, _ = io.WriteString(h, strconv.FormatInt(now, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
