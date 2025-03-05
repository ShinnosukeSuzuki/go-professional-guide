# Goのbyte型と文字列の詳細解説

## byte型の基本

Goでは`byte`型は`uint8`のエイリアスで、0〜255の値を持つ8ビット（1バイト）の符号なし整数です。

```go
var b byte = 65
fmt.Printf("値: %d, 文字として: %c\n", b, b) // 出力: 値: 65, 文字として: A
```

## 文字列とbyte型の関係

Goの文字列は内部的にはバイトのシーケンス（`[]byte`）として格納されています。これらのバイトはUTF-8でエンコードされた文字を表します。

```go
str := "Hello"
fmt.Printf("文字列のバイト: % x\n", []byte(str)) // 出力: 文字列のバイト: 48 65 6c 6c 6f
```

## 相互変換

文字列と`[]byte`の間は簡単に変換できます：

```go
// 文字列からバイトスライスへ
str := "Hello, 世界"
bytes := []byte(str)

// バイトスライスから文字列へ
newStr := string(bytes)
```

注意点: 日本語などの非ASCII文字はUTF-8では複数バイトで表されます:

```go
str := "世界"
bytes := []byte(str)
fmt.Println(len(bytes)) // 出力: 6 (2文字ですが、UTF-8では各文字が3バイト)
```

## 効率性と性能

- 文字列とバイトスライス間の変換ではメモリコピーが発生します
- 大量の文字列操作を行う場合は、`bytes.Buffer`や`strings.Builder`を使用すると効率的です

```go
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteByte(byte('A' + i%26))
}
result := builder.String()
```

## バイト操作

バイトスライスには様々な操作が可能です：

```go
// バイトスライスの作成
bs := []byte{72, 101, 108, 108, 111} // "Hello"に相当

// バイトの置換
bs[0] = 74 // 'H'を'J'に変更

// スライスの追加
bs = append(bs, []byte(" World")...)

// 文字列に戻す
fmt.Println(string(bs)) // 出力: Jello World
```

## バイトと文字（rune）の違い

- `byte`はASCII文字には十分ですが、マルチバイト文字には対応できません
- マルチバイト文字を扱う場合は`rune`型（`int32`のエイリアス）を使用します

```go
s := "こんにちは"
fmt.Println("byteスライスの長さ:", len([]byte(s)))    // 出力: 15 (3バイト×5文字)
fmt.Println("runeスライスの長さ:", len([]rune(s)))    // 出力: 5 (文字数)
```

## 実用的な例

### 1. ファイル操作

```go
data, err := os.ReadFile("file.txt") // []byteとして読み込み
if err != nil {
    log.Fatal(err)
}
content := string(data) // 必要に応じて文字列に変換
```

### 2. HTTP通信

```go
resp, err := http.Get("https://example.com")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body) // []byteとして本文を読み込み
```

### 3. バイナリデータ処理

```go
var header [8]byte
_, err := io.ReadFull(reader, header[:]) // ヘッダー部分を読み込み
magic := binary.BigEndian.Uint32(header[0:4]) // 最初の4バイトを解釈
```

## メモリ効率

文字列は不変（イミュータブル）なので、変更が必要な場合は新しいメモリ領域が割り当てられます。一方、`[]byte`は可変（ミュータブル）なので、同じメモリ上で変更が可能です：

```go
// 文字列の場合（非効率）
s := "Hello"
s = s + " World" // 新しいメモリ割り当てが発生

// バイトスライスの場合（効率的）
b := []byte("Hello")
b = append(b, []byte(" World")...) // 可能であれば同じメモリを拡張
```

このような特性を理解することで、パフォーマンスが重要な場面で適切な選択ができるようになります。
