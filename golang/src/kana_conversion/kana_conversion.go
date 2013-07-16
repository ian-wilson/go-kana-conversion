package kana_conversion

import "regexp"
import "strings"
import "strconv"
import "fmt"
import "sort"


type KanaConversion struct {
    // Dictionaries
    romaji_dict map[string] string
    romaji_keys map[string] string
    kana_dict map[string] string
    kana_keys map[string] string
    zenkaku_dict map[string] string
    zenkaku_keys map[string] string

    // regexes
    re_xu *regexp.Regexp
    re_ltu *regexp.Regexp
    re_er *regexp.Regexp
    re_n *regexp.Regexp
    re_oo *regexp.Regexp
    re_mba *regexp.Regexp
    re_xtu *regexp.Regexp
    re_a *regexp.Regexp
    re_hiragana_to_katakana *regexp.Regexp
    re_katakana_to_hiragana *regexp.Regexp
    re_romaji_to_katakana *regexp.Regexp
    re_romaji_to_zenkaku *regexp.Regexp
    re_katakana_to_romaji *regexp.Regexp
}

func (self *KanaConversion) build_conversion_dictionary(word_dict map[string]string, assist_dict map[string]string, is_reverse_keys bool) (map[string]string, []string) {

    // Create a word dictionary merged from the given tables with sorted keys list
    // :param word_dict:
    // :param assist_dict:
    // :param is_reverse_keys:
    // :return: merged dictionary, sorted list of dict keys

    var conversion_dict = make(map[string]string)

    for key, value := range word_dict {
        if is_reverse_keys {
            conversion_dict[value] = key
        } else {
            conversion_dict[key] = value
        }
    }

    for key, value := range assist_dict {
        if is_reverse_keys {
            conversion_dict[value] = key
        } else {
            conversion_dict[key] = value
        }
    }

    var conversion_keys = make([]string, len(conversion_dict))
    var key_index int = 0

    for key, _ := range conversion_dict {
        conversion_keys[key_index] = key
        key_index = key_index + 1
    }

    sort.Strings(conversion_keys)

    return conversion_dict, conversion_keys
}

func (self *KanaConversion) build_conversion_pattern(word_separator string, word_list []string) string {

    // Create a regex pattern of the given word dict
    // :param word_separator: word separator
    // :param word_list: words to concatenate
    // :return:

    var mapped_words = make([]string, len(word_list))

    for _, word := range word_list {
        escaped_word := strconv.Quote(word)
        mapped_words = append(mapped_words, escaped_word)
    }

    pattern := strings.Join(mapped_words, word_separator)

    return pattern
}

func (self *KanaConversion) convert_text(regex_pattern *regexp.Regexp, conversion_dict map[string] string, text string) string {
    /*
    Use a regex to substitute characters to convert
    :param regex_pattern: precompiled regex pattern
    :param conversion_dict: conversion word dictionary
    :param text: text to convert
    :return: converted text
    */
    var conversion string = text

    for key, _ := range conversion_dict {
        conversion = regex_pattern.ReplaceAllString(conversion, key)
        fmt.Println(conversion)
    }

    return conversion
}

func (self *KanaConversion) pre_process_romaji(text string) string {
    /*
    Pre processing for difficult words
    :param romaji_text:
    :return: pre processed romaji text
    */
    var preprocessed_text string = text

    preprocessed_text = strings.ToLower(preprocessed_text)
    preprocessed_text = self.re_mba.ReplaceAllString(preprocessed_text, MBA_SUB_PATTERN)
    preprocessed_text = self.re_xu.ReplaceAllString(preprocessed_text, XU__SUB_PATTERN)
    preprocessed_text = self.re_a.ReplaceAllString(preprocessed_text, A___SUB_PATTERN)

    return preprocessed_text
}

func (self *KanaConversion) post_process_romaji_text(text string) string {
    /*
    Post process for difficult words
    :param romaji_text:
    :return: post processed romaji text
    */
    var postprocessed_text string = text

    postprocessed_text = self.re_xtu.ReplaceAllString(postprocessed_text, XTU_SUB_PATTERN)
    postprocessed_text = self.re_ltu.ReplaceAllString(postprocessed_text, LTU_SUB_PATTERN)
    postprocessed_text = self.re_er.ReplaceAllString(postprocessed_text, ER__SUB_PATTERN)
    postprocessed_text = self.re_n.ReplaceAllString(postprocessed_text, N___SUB_PATTERN)
    postprocessed_text = self.re_oo.ReplaceAllString(postprocessed_text, OO__SUB_PATTERN)

    return postprocessed_text
}

func (self *KanaConversion) sort_by_key_length(word_map map[string] string) map[string] string {
    /*
    Post process for difficult words
    :param romaji_text:
    :return: post processed romaji text
    */
    var postprocessed_text map[string] string

    return postprocessed_text
}

func (self *KanaConversion) Init() {
    /*
    Init dictionaries, regexs
    */

    // Build romaji dictionaries and keys
    var romaji_keys [] string
    var kana_keys [] string
    var zenkaku_keys [] string
    var hiragana_keys [] string
    var katakana_keys [] string

    self.romaji_dict, romaji_keys = self.build_conversion_dictionary(romaji, romaji_asist, false)
    self.kana_dict, kana_keys = self.build_conversion_dictionary(romaji, kana_asist, true)
    self.zenkaku_dict, zenkaku_keys = self.build_conversion_dictionary(zenkaku, zenkaku_assist, false)

    for key, _ := range hiragana {
        hiragana_keys = append(hiragana_keys, key)
    }

    for key, _ := range katakana {
        katakana_keys = append(katakana_keys, key)
    }

    // Convert keys to string patterns
    hiragana_to_katakana_pattern := self.build_conversion_pattern(JOIN_CHAR, hiragana_keys)
    katakana_to_hiragana_pattern := self.build_conversion_pattern(JOIN_CHAR, katakana_keys)
    romaji_to_katakana_pattern   := self.build_conversion_pattern(JOIN_CHAR, romaji_keys)
    romaji_to_zenkaku_pattern    := self.build_conversion_pattern(JOIN_CHAR, zenkaku_keys)
    katakana_to_romaji_pattern   := self.build_conversion_pattern(JOIN_CHAR, kana_keys)

    // Precompile regexes
    self.re_hiragana_to_katakana, _ = regexp.Compile(hiragana_to_katakana_pattern)
    self.re_katakana_to_hiragana, _ = regexp.Compile(katakana_to_hiragana_pattern)
    self.re_romaji_to_katakana, _   = regexp.Compile(romaji_to_katakana_pattern)
    self.re_romaji_to_zenkaku, _    = regexp.Compile(romaji_to_zenkaku_pattern)
    self.re_katakana_to_romaji, _   = regexp.Compile(katakana_to_romaji_pattern)

    self.re_xu, _  = regexp.Compile(XU__PATTERN)
    self.re_ltu, _ = regexp.Compile(LTU_PATTERN)
    self.re_er, _  = regexp.Compile(ER__PATTERN)
    self.re_n, _   = regexp.Compile(N___PATTERN)
    self.re_oo, _  = regexp.Compile(OO__PATTERN)
    self.re_mba, _ = regexp.Compile(MBA_PATTERN)
    self.re_xtu, _ = regexp.Compile(XTU_PATTERN)
    self.re_a, _   = regexp.Compile(A___PATTERN)
}

func (self *KanaConversion) Hiragana_to_katakana(text string) string {
    /*
    Example conversion -> いあん　ー＞　イアン
    :param text:
    :return: unicode string
    */
    conversion := self.convert_text(self.re_hiragana_to_katakana, hiragana, text)
    return conversion
}

func (self *KanaConversion) Katakana_to_hiragana(text string) string {
    /*
    Example conversion -> イアン　ー＞　いあん
    :param text:
    :return: unicode string
    */
    conversion := self.convert_text(self.re_katakana_to_hiragana, katakana, text)
    return conversion
}

func (self *KanaConversion) Romaji_to_katakana(text string) string {
    /*
    Example conversion -> ian　ー＞　イアン
    :param text:
    :return: unicode string
    */

    pre_processed_text := self.pre_process_romaji(text)
    conversion := self.convert_text(self.re_romaji_to_katakana, self.romaji_dict, pre_processed_text)
    return conversion
}

func (self *KanaConversion) Romaji_to_hiragana(text string) string {
    /*
    Example conversion -> ian　ー＞　いあん
    :param text:
    :return: unicode string
    */
    conversion := self.Romaji_to_katakana(text)
    conversion  = self.Katakana_to_hiragana(conversion)

    return conversion
}

func (self *KanaConversion) Romaji_to_zenkaku(text string) string {
    /*
    Example conversion -> ian　ー＞　ｉａｎ
    :param text:
    :return: unicode string
    */
    conversion := self.convert_text(self.re_romaji_to_zenkaku, self.zenkaku_dict, text)
    return conversion
}

func (self *KanaConversion) Katakana_to_romaji(text string) string {
    /*
    Example conversion -> イアン　ー＞　ian
    :param text:
    :return: unicode string
    */
    conversion := self.convert_text(self.re_katakana_to_romaji, self.kana_dict, text)
    post_processed_text := self.post_process_romaji_text(conversion)

    return post_processed_text
}

func (self *KanaConversion) Hiragana_to_romaji(text string) string {
    /*
    Example conversion -> いあん　ー＞　ian
    :param text:
    :return: unicode string
    */
    conversion := self.Hiragana_to_katakana(text)
    conversion  = self.Katakana_to_romaji(conversion)

    return conversion
}
