# gstr

> Package: `github.com/gogf/gf/v2/text/gstr`

```go
import "github.com/gogf/gf/v2/text/gstr"
```

## 概述

Package gstr provides functions for string handling.

## 源文件

- `gstr.go`
- `gstr_array.go`
- `gstr_case.go`
- `gstr_compare.go`
- `gstr_contain.go`
- `gstr_convert.go`
- `gstr_count.go`
- `gstr_create.go`
- `gstr_domain.go`
- `gstr_is.go`
- `gstr_length.go`
- `gstr_list.go`
- `gstr_parse.go`
- `gstr_pos.go`
- `gstr_replace.go`
- `gstr_similar.go`
- `gstr_slashes.go`
- `gstr_split_join.go`
- `gstr_sub.go`
- `gstr_trim.go`
- `gstr_upper_lower.go`
- `gstr_version.go`

## 类型定义

### CaseType

**类型**: type

CaseType is the type for Case.


## 函数

### SearchArray

```go
func SearchArray(a []string, s string) int
```

SearchArray searches string `s` in string slice `a` case-sensitively,
returns its index in `a`.
If `s` is not found in `a`, it returns -1.

### InArray

```go
func InArray(a []string, s string) bool
```

InArray checks whether string `s` in slice `a`.

### PrefixArray

```go
func PrefixArray(array []string, prefix string)
```

PrefixArray adds `prefix` string for each item of `array`.

Example:
PrefixArray(["a","b"], "gf_") -> ["gf_a", "gf_b"]

### CaseTypeMatch

```go
func CaseTypeMatch(caseStr string) CaseType
```

CaseTypeMatch matches the case type from string.

### CaseConvert

```go
func CaseConvert(s string, caseType CaseType) string
```

CaseConvert converts a string to the specified naming convention.
Use CaseTypeMatch to match the case type from string.

### CaseCamel

```go
func CaseCamel(s string) string
```

CaseCamel converts a string to CamelCase.

Example:
CaseCamel("any_kind_of_string") -> AnyKindOfString
CaseCamel("anyKindOfString")    -> AnyKindOfString

### CaseCamelLower

```go
func CaseCamelLower(s string) string
```

CaseCamelLower converts a string to lowerCamelCase.

Example:
CaseCamelLower("any_kind_of_string") -> anyKindOfString
CaseCamelLower("AnyKindOfString")    -> anyKindOfString

### CaseSnake

```go
func CaseSnake(s string) string
```

CaseSnake converts a string to snake_case.

Example:
CaseSnake("AnyKindOfString") -> any_kind_of_string

### CaseSnakeScreaming

```go
func CaseSnakeScreaming(s string) string
```

CaseSnakeScreaming converts a string to SNAKE_CASE_SCREAMING.

Example:
CaseSnakeScreaming("AnyKindOfString") -> ANY_KIND_OF_STRING

### CaseSnakeFirstUpper

```go
func CaseSnakeFirstUpper(word string, underscore ...string) string
```

CaseSnakeFirstUpper converts a string like "RGBCodeMd5" to "rgb_code_md5".
TODO for efficiency should change regexp to traversing string in future.

Example:
CaseSnakeFirstUpper("RGBCodeMd5") -> rgb_code_md5

### CaseKebab

```go
func CaseKebab(s string) string
```

CaseKebab converts a string to kebab-case.

Example:
CaseKebab("AnyKindOfString") -> any-kind-of-string

### CaseKebabScreaming

```go
func CaseKebabScreaming(s string) string
```

CaseKebabScreaming converts a string to KEBAB-CASE-SCREAMING.

Example:
CaseKebab("AnyKindOfString") -> ANY-KIND-OF-STRING

### CaseDelimited

```go
func CaseDelimited(s string, del byte) string
```

CaseDelimited converts a string to snake.case.delimited.

Example:
CaseDelimited("AnyKindOfString", '.') -> any.kind.of.string

### CaseDelimitedScreaming

```go
func CaseDelimitedScreaming(s string, del uint8, screaming bool) string
```

CaseDelimitedScreaming converts a string to DELIMITED.SCREAMING.CASE or delimited.screaming.case.

Example:
CaseDelimitedScreaming("AnyKindOfString", '.') -> ANY.KIND.OF.STRING

### Compare

```go
func Compare(a string, b string) int
```

Compare returns an integer comparing two strings lexicographically.
The result will be 0 if a==b, -1 if a < b, and +1 if a > b.

### Equal

```go
func Equal(a string, b string) bool
```

Equal reports whether `a` and `b`, interpreted as UTF-8 strings,
are equal under Unicode case-folding, case-insensitively.

### Contains

```go
func Contains(str string, substr string) bool
```

Contains reports whether `substr` is within `str`, case-sensitively.

### ContainsI

```go
func ContainsI(str string, substr string) bool
```

ContainsI reports whether substr is within str, case-insensitively.

### ContainsAny

```go
func ContainsAny(s string, chars string) bool
```

ContainsAny reports whether any Unicode code points in `chars` are within `s`.

### Chr

```go
func Chr(ascii int) string
```

Chr return the ascii string of a number(0-255).

Example:
Chr(65) -> "A"

### Ord

```go
func Ord(char string) int
```

Ord converts the first byte of a string to a value between 0 and 255.

Example:
Chr("A") -> 65

### OctStr

```go
func OctStr(str string) string
```

OctStr converts string container octal string to its original string,
for example, to Chinese string.

Example:
OctStr("\346\200\241") -> 怡

### Reverse

```go
func Reverse(str string) string
```

Reverse returns a string which is the reverse of `str`.

Example:
Reverse("123456") -> "654321"

### NumberFormat

```go
func NumberFormat(number float64, decimals int, decPoint string, thousandsSep string) string
```

NumberFormat formats a number with grouped thousands.
Parameter `decimals`: Sets the number of decimal points.
Parameter `decPoint`: Sets the separator for the decimal point.
Parameter `thousandsSep`: Sets the thousands' separator.
See http://php.net/manual/en/function.number-format.php.

Example:
NumberFormat(1234.56, 2, ".", "")  -> 1234,56
NumberFormat(1234.56, 2, ",", " ") -> 1 234,56

### Shuffle

```go
func Shuffle(str string) string
```

Shuffle randomly shuffles a string.
It considers parameter `str` as unicode string.

Example:
Shuffle("123456") -> "325164"
Shuffle("123456") -> "231546"
...

### HideStr

```go
func HideStr(str string, percent int, hide string) string
```

HideStr replaces part of the string `str` to `hide` by `percentage` from the `middle`.
It considers parameter `str` as unicode string.

### Nl2Br

```go
func Nl2Br(str string, isXhtml ...bool) string
```

Nl2Br inserts HTML line breaks(`br`|<br />) before all newlines in a string:
\n\r, \r\n, \r, \n.
It considers parameter `str` as unicode string.

### WordWrap

```go
func WordWrap(str string, width int, br string) string
```

WordWrap wraps a string to a given number of characters.
This function supports cut parameters of both english and chinese punctuations.
TODO: Enable custom cut parameter, see http://php.net/manual/en/function.wordwrap.php.

### Count

```go
func Count(s string, substr string) int
```

Count counts the number of `substr` appears in `s`.
It returns 0 if no `substr` found in `s`.

### CountI

```go
func CountI(s string, substr string) int
```

CountI counts the number of `substr` appears in `s`, case-insensitively.
It returns 0 if no `substr` found in `s`.

### CountWords

```go
func CountWords(str string) map[string]int
```

CountWords returns information about words' count used in a string.
It considers parameter `str` as unicode string.

### CountChars

```go
func CountChars(str string, noSpace ...bool) map[string]int
```

CountChars returns information about chars' count used in a string.
It considers parameter `str` as Unicode string.

### Repeat

```go
func Repeat(input string, multiplier int) string
```

Repeat returns a new string consisting of multiplier copies of the string input.

Example:
Repeat("a", 3) -> "aaa"

### IsSubDomain

```go
func IsSubDomain(subDomain string, mainDomain string) bool
```

IsSubDomain checks whether `subDomain` is sub-domain of mainDomain.
It supports '*' in `mainDomain`.

### IsNumeric

```go
func IsNumeric(s string) bool
```

IsNumeric tests whether the given string s is numeric.

### LenRune

```go
func LenRune(str string) int
```

LenRune returns string length of unicode.

### List2

```go
func List2(str string, delimiter string) (part1 string, part2 string)
```

List2 Split the `str` with `delimiter` and returns the result as two parts string.

### ListAndTrim2

```go
func ListAndTrim2(str string, delimiter string) (part1 string, part2 string)
```

ListAndTrim2 SplitAndTrim the `str` with `delimiter` and returns the result as two parts string.

### List3

```go
func List3(str string, delimiter string) (part1 string, part2 string, part3 string)
```

List3 Split the `str` with `delimiter` and returns the result as three parts string.

### ListAndTrim3

```go
func ListAndTrim3(str string, delimiter string) (part1 string, part2 string, part3 string)
```

ListAndTrim3 SplitAndTrim the `str` with `delimiter` and returns the result as three parts string.

### Parse

```go
func Parse(s string) (result map[string]any, err error)
```

Parse parses the string into map[string]any.

v1=m&v2=n           -> map[v1:m v2:n]
v[a]=m&v[b]=n       -> map[v:map[a:m b:n]]
v[a][a]=m&v[a][b]=n -> map[v:map[a:map[a:m b:n]]]
v[]=m&v[]=n         -> map[v:[m n]]
v[a][]=m&v[a][]=n   -> map[v:map[a:[m n]]]
v[][]=m&v[][]=n     -> map[v:[map[]]] // Currently does not support nested slice.
v=m&v[a]=n          -> error
a .[[b=c            -> map[a___[b:c]

### Pos

```go
func Pos(haystack string, needle string, startOffset ...int) int
```

Pos returns the position of the first occurrence of `needle`
in `haystack` from `startOffset`, case-sensitively.
It returns -1, if not found.

### PosRune

```go
func PosRune(haystack string, needle string, startOffset ...int) int
```

PosRune acts like function Pos but considers `haystack` and `needle` as unicode string.

### PosI

```go
func PosI(haystack string, needle string, startOffset ...int) int
```

PosI returns the position of the first occurrence of `needle`
in `haystack` from `startOffset`, case-insensitively.
It returns -1, if not found.

### PosIRune

```go
func PosIRune(haystack string, needle string, startOffset ...int) int
```

PosIRune acts like function PosI but considers `haystack` and `needle` as unicode string.

### PosR

```go
func PosR(haystack string, needle string, startOffset ...int) int
```

PosR returns the position of the last occurrence of `needle`
in `haystack` from `startOffset`, case-sensitively.
It returns -1, if not found.

### PosRRune

```go
func PosRRune(haystack string, needle string, startOffset ...int) int
```

PosRRune acts like function PosR but considers `haystack` and `needle` as unicode string.

### PosRI

```go
func PosRI(haystack string, needle string, startOffset ...int) int
```

PosRI returns the position of the last occurrence of `needle`
in `haystack` from `startOffset`, case-insensitively.
It returns -1, if not found.

### PosRIRune

```go
func PosRIRune(haystack string, needle string, startOffset ...int) int
```

PosRIRune acts like function PosRI but considers `haystack` and `needle` as unicode string.

### Replace

```go
func Replace(origin string, search string, replace string, count ...int) string
```

Replace returns a copy of the string `origin`
in which string `search` replaced by `replace` case-sensitively.

### ReplaceI

```go
func ReplaceI(origin string, search string, replace string, count ...int) string
```

ReplaceI returns a copy of the string `origin`
in which string `search` replaced by `replace` case-insensitively.

### ReplaceByArray

```go
func ReplaceByArray(origin string, array []string) string
```

ReplaceByArray returns a copy of `origin`,
which is replaced by a slice in order, case-sensitively.

### ReplaceIByArray

```go
func ReplaceIByArray(origin string, array []string) string
```

ReplaceIByArray returns a copy of `origin`,
which is replaced by a slice in order, case-insensitively.

### ReplaceByMap

```go
func ReplaceByMap(origin string, replaces map[string]string) string
```

ReplaceByMap returns a copy of `origin`,
which is replaced by a map in unordered way, case-sensitively.

### ReplaceIByMap

```go
func ReplaceIByMap(origin string, replaces map[string]string) string
```

ReplaceIByMap returns a copy of `origin`,
which is replaced by a map in unordered way, case-insensitively.

### ReplaceFunc

```go
func ReplaceFunc(origin string, search string, f func) string
```

ReplaceFunc returns a copy of the string `origin` in which each non-overlapping substring
that matches the given search string is replaced by the result of function `f` applied to that substring.
The function `f` is called with each matching substring as its argument and must return a string to be used
as the replacement value.

### ReplaceIFunc

```go
func ReplaceIFunc(origin string, search string, f func) string
```

ReplaceIFunc returns a copy of the string `origin` in which each non-overlapping substring
that matches the given search string is replaced by the result of function `f` applied to that substring.
The match is done case-insensitively.
The function `f` is called with each matching substring as its argument and must return a string to be used
as the replacement value.

### Levenshtein

```go
func Levenshtein(str1 string, str2 string, costIns int, costRep int, costDel int) int
```

Levenshtein calculates Levenshtein distance between two strings.
costIns: Defines the cost of insertion.
costRep: Defines the cost of replacement.
costDel: Defines the cost of deletion.
See http://php.net/manual/en/function.levenshtein.php.

### SimilarText

```go
func SimilarText(first string, second string, percent *float64) int
```

SimilarText calculates the similarity between two strings.
See http://php.net/manual/en/function.similar-text.php.

### Soundex

```go
func Soundex(str string) string
```

Soundex calculates the soundex key of a string.
See http://php.net/manual/en/function.soundex.php.

### AddSlashes

```go
func AddSlashes(str string) string
```

AddSlashes quotes with slashes `\` for chars: '"\.

### StripSlashes

```go
func StripSlashes(str string) string
```

StripSlashes un-quotes a quoted string by AddSlashes.

### QuoteMeta

```go
func QuoteMeta(str string, chars ...string) string
```

QuoteMeta returns a version of `str` with a backslash character (`\`).
If custom chars `chars` not given, it uses default chars: .\+*?[^]($)

### Split

```go
func Split(str string, delimiter string) []string
```

Split splits string `str` by a string `delimiter`, to an array.

### SplitAndTrim

```go
func SplitAndTrim(str string, delimiter string, characterMask ...string) []string
```

SplitAndTrim splits string `str` by a string `delimiter` to an array,
and calls Trim to every element of this array. It ignores the elements
which are empty after Trim.

### Join

```go
func Join(array []string, sep string) string
```

Join concatenates the elements of `array` to create a single string. The separator string
`sep` is placed between elements in the resulting string.

### JoinAny

```go
func JoinAny(array any, sep string) string
```

JoinAny concatenates the elements of `array` to create a single string. The separator string
`sep` is placed between elements in the resulting string.

The parameter `array` can be any type of slice, which be converted to string array.

### Explode

```go
func Explode(delimiter string, str string) []string
```

Explode splits string `str` by a string `delimiter`, to an array.
See http://php.net/manual/en/function.explode.php.

### Implode

```go
func Implode(glue string, pieces []string) string
```

Implode joins array elements `pieces` with a string `glue`.
http://php.net/manual/en/function.implode.php

### ChunkSplit

```go
func ChunkSplit(body string, chunkLen int, end string) string
```

ChunkSplit splits a string into smaller chunks.
Can be used to split a string into smaller chunks which is useful for
e.g. converting BASE64 string output to match RFC 2045 semantics.
It inserts end every chunkLen characters.
It considers parameter `body` and `end` as unicode string.

### Fields

```go
func Fields(str string) []string
```

Fields returns the words used in a string as slice.

### Str

```go
func Str(haystack string, needle string) string
```

Str returns part of `haystack` string starting from and including
the first occurrence of `needle` to the end of `haystack`.

This function performs exactly as function SubStr, but to implement the same function
as PHP: http://php.net/manual/en/function.strstr.php.

Example:
Str("av.mp4", ".") -> ".mp4"

### StrEx

```go
func StrEx(haystack string, needle string) string
```

StrEx returns part of `haystack` string starting from and excluding
the first occurrence of `needle` to the end of `haystack`.

This function performs exactly as function SubStrEx, but to implement the same function
as PHP: http://php.net/manual/en/function.strstr.php.

Example:
StrEx("av.mp4", ".") -> "mp4"

### StrTill

```go
func StrTill(haystack string, needle string) string
```

StrTill returns part of `haystack` string ending to and including
the first occurrence of `needle` from the start of `haystack`.

Example:
StrTill("av.mp4", ".") -> "av."

### StrTillEx

```go
func StrTillEx(haystack string, needle string) string
```

StrTillEx returns part of `haystack` string ending to and excluding
the first occurrence of `needle` from the start of `haystack`.

Example:
StrTillEx("av.mp4", ".") -> "av"

### SubStr

```go
func SubStr(str string, start int, length ...int) substr string
```

SubStr returns a portion of string `str` specified by the `start` and `length` parameters.
The parameter `length` is optional, it uses the length of `str` in default.

Example:
SubStr("123456", 1, 2) -> "23"

### SubStrRune

```go
func SubStrRune(str string, start int, length ...int) substr string
```

SubStrRune returns a portion of string `str` specified by the `start` and `length` parameters.
SubStrRune considers parameter `str` as unicode string.
The parameter `length` is optional, it uses the length of `str` in default.

Example:
SubStrRune("一起学习吧！", 2, 2) -> "学习"

### StrLimit

```go
func StrLimit(str string, length int, suffix ...string) string
```

StrLimit returns a portion of string `str` specified by `length` parameters, if the length
of `str` is greater than `length`, then the `suffix` will be appended to the result string.

Example:
StrLimit("123456", 3)      -> "123..."
StrLimit("123456", 3, "~") -> "123~"

### StrLimitRune

```go
func StrLimitRune(str string, length int, suffix ...string) string
```

StrLimitRune returns a portion of string `str` specified by `length` parameters, if the length
of `str` is greater than `length`, then the `suffix` will be appended to the result string.
StrLimitRune considers parameter `str` as unicode string.

Example:
StrLimitRune("一起学习吧！", 2)      -> "一起..."
StrLimitRune("一起学习吧！", 2, "~") -> "一起~"

### SubStrFrom

```go
func SubStrFrom(str string, need string) substr string
```

SubStrFrom returns a portion of string `str` starting from first occurrence of and including `need`
to the end of `str`.

Example:
SubStrFrom("av.mp4", ".") -> ".mp4"

### SubStrFromEx

```go
func SubStrFromEx(str string, need string) substr string
```

SubStrFromEx returns a portion of string `str` starting from first occurrence of and excluding `need`
to the end of `str`.

Example:
SubStrFromEx("av.mp4", ".") -> "mp4"

### SubStrFromR

```go
func SubStrFromR(str string, need string) substr string
```

SubStrFromR returns a portion of string `str` starting from last occurrence of and including `need`
to the end of `str`.

Example:
SubStrFromR("/dev/vda", "/") -> "/vda"

### SubStrFromREx

```go
func SubStrFromREx(str string, need string) substr string
```

SubStrFromREx returns a portion of string `str` starting from last occurrence of and excluding `need`
to the end of `str`.

Example:
SubStrFromREx("/dev/vda", "/") -> "vda"

### Trim

```go
func Trim(str string, characterMask ...string) string
```

Trim strips whitespace (or other characters) from the beginning and end of a string.
The optional parameter `characterMask` specifies the additional stripped characters.

### TrimStr

```go
func TrimStr(str string, cut string, count ...int) string
```

TrimStr strips all the given `cut` string from the beginning and end of a string.
Note that it does not strip the whitespaces of its beginning or end.

### TrimLeft

```go
func TrimLeft(str string, characterMask ...string) string
```

TrimLeft strips whitespace (or other characters) from the beginning of a string.

### TrimLeftStr

```go
func TrimLeftStr(str string, cut string, count ...int) string
```

TrimLeftStr strips all the given `cut` string from the beginning of a string.
Note that it does not strip the whitespaces of its beginning.

### TrimRight

```go
func TrimRight(str string, characterMask ...string) string
```

TrimRight strips whitespace (or other characters) from the end of a string.

### TrimRightStr

```go
func TrimRightStr(str string, cut string, count ...int) string
```

TrimRightStr strips all the given `cut` string from the end of a string.
Note that it does not strip the whitespaces of its end.

### TrimAll

```go
func TrimAll(str string, characterMask ...string) string
```

TrimAll trims all characters in string `str`.

### HasPrefix

```go
func HasPrefix(s string, prefix string) bool
```

HasPrefix tests whether the string s begins with prefix.

### HasSuffix

```go
func HasSuffix(s string, suffix string) bool
```

HasSuffix tests whether the string s ends with suffix.

### ToLower

```go
func ToLower(s string) string
```

ToLower returns a copy of the string s with all Unicode letters mapped to their lower case.

### ToUpper

```go
func ToUpper(s string) string
```

ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.

### UcFirst

```go
func UcFirst(s string) string
```

UcFirst returns a copy of the string s with the first letter mapped to its upper case.

### LcFirst

```go
func LcFirst(s string) string
```

LcFirst returns a copy of the string s with the first letter mapped to its lower case.

### UcWords

```go
func UcWords(str string) string
```

UcWords uppercase the first character of each word in a string.

### IsLetterLower

```go
func IsLetterLower(b byte) bool
```

IsLetterLower tests whether the given byte b is in lower case.

### IsLetterUpper

```go
func IsLetterUpper(b byte) bool
```

IsLetterUpper tests whether the given byte b is in upper case.

### IsGNUVersion

```go
func IsGNUVersion(version string) bool
```

IsGNUVersion checks and returns whether given `version` is valid GNU version string.

### CompareVersion

```go
func CompareVersion(a string, b string) int
```

CompareVersion compares `a` and `b` as standard GNU version.

It returns  1 if `a` > `b`.

It returns -1 if `a` < `b`.

It returns  0 if `a` = `b`.

GNU standard version is like:
v1.0
1
1.0.0
v1.0.1
v2.10.8
10.2.0
etc.

### CompareVersionGo

```go
func CompareVersionGo(a string, b string) int
```

CompareVersionGo compares `a` and `b` as standard Golang version.

It returns  1 if `a` > `b`.

It returns -1 if `a` < `b`.

It returns  0 if `a` = `b`.

Golang standard version is like:
1.0.0
v1.0.1
v2.10.8
10.2.0
v0.0.0-20190626092158-b2ccc519800e
v1.12.2-0.20200413154443-b17e3a6804fa
v4.20.0+incompatible
etc.

Docs: https://go.dev/doc/modules/version-numbers

## 内部依赖

- `errors/gcode`
- `errors/gerror`
- `internal/utils`
- `util/gconv`
- `util/grand`

## 外部依赖

- `golang.org/x/text/cases`
- `golang.org/x/text/language`

