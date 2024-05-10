package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var changesNum int
	var originalSticker string
	fmt.Fscanln(in, &originalSticker)
	fmt.Fscanln(in, &changesNum)

	for i := 0; i < changesNum; i++ {
		originalStickerArr := strings.Split(originalSticker, "")
		line, _ := in.ReadString('\n')
		changes := strings.Fields(line)
		from, err := strconv.Atoi(changes[0])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(changes[1])
		if err != nil {
			panic(err)
		}

		originalSticker = strings.Join(originalStickerArr[0:from-1], "") + changes[2] + strings.Join(originalStickerArr[to:], "")
	}
	fmt.Fprintln(out, originalSticker)
}
