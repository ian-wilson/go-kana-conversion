package kana_conversion

import "testing"
import "fmt"


func Test(t *testing.T) {
    convert := new(KanaConversion)
    convert.Init()

    r2k := convert.Romaji_to_katakana("IAN")
    r2z := convert.Romaji_to_zenkaku("IAN")

    fmt.Println(r2k)
    fmt.Println(r2z)
}
