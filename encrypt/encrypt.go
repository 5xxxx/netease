/*
 *
 * encrypt.go
 * encrpt
 *
 * Created by lintao on 2020/6/9 9:57 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package encrypt

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"math/big"
	"strconv"
)

func RandomString(length int) string {
	const base = 36
	size := big.NewInt(base)

	n := make([]byte, length)
	for i, _ := range n {
		c, _ := rand.Int(rand.Reader, size)
		str := strconv.FormatInt(c.Int64(), base)
		n[i] = str[0]
	}
	return string(n)
}

func GenerateCheckSum(nonce,appId, curTime string) string {
	var str = fmt.Sprintf("%s%s%s", appId, nonce, curTime)
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
