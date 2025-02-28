# シリコンバレー一流プログラマーが教える Goプロフェッショナル大全
## Lesson1 Goの基本
### Section2 変数の作り方をマスターしよう
短縮変数宣言の`:=`をしようすると、データ型が自動的に設定される。


`fmt.Printf("%T", x)` の `%T` は、変数 `x` の 型（type） を出力します。  
例えば、以下のコードを実行すると：
```go
package main

import "fmt"

func main() {
    x := 42
    fmt.Printf("%T\n", x) // int

    y := 3.14
    fmt.Printf("%T\n", y) // float64

    z := "Hello"
    fmt.Printf("%T\n", z) // string

    a := []int{1, 2, 3}
    fmt.Printf("%T\n", a) // []int

    b := map[string]int{"a": 1, "b": 2}
    fmt.Printf("%T\n", b) // map[string]int
}
```
**varと短縮変数宣言の違い**  
varを使うと関数の外で定義することができ、複数の関数から呼び出せる。短縮変数宣言は関数の中でしか使用できない。  
```go
package main

import "fmt"

func main() {
    // ① var を使った変数宣言（関数の外でも使える）
    var a int
    a = 10
    fmt.Println("a:", a) // a: 10

    // ② 短縮変数宣言（関数内でのみ使用可能）
    b := 20
    fmt.Println("b:", b) // b: 20

    // ③ var では型を明示的に指定できる
    var c float64 = 3.14
    fmt.Println("c:", c) // c: 3.14

    // ④ 短縮変数宣言は型推論される
    d := "Hello"
    fmt.Println("d:", d) // d: Hello

    // ⑤ var はゼロ値を持つ（int のゼロ値は 0）
    var e int
    fmt.Println("e:", e) // e: 0

    // ⑥ 短縮変数宣言は再宣言できるが、var はできない
    x := 1
    {
        x := 2 // 新しいスコープで再宣言（別の変数として扱われる）
        fmt.Println("inner x:", x) // inner x: 2
    }
    fmt.Println("outer x:", x) // outer x: 1

    // ⑦ var は関数の外（パッケージレベル）でも使える
    fmt.Println(globalVar) // "I am global"
}

// ⑧ var は関数の外（パッケージレベル）で宣言できる
var globalVar string = "I am global"
```
不変変数(定数)は`const`を使って宣言する。宣言と同意に初期化を行う。  
定数は関数内でも定義できるが、基本的は関数外で定義する。他のファイルから呼び出す場合は大文字にする。

**定数のオーバーフロー挙動について**  
Goの定数のオーバーフロー挙動のポイントは以下の通り：

1. 型なし定数:  
非常に大きな精度（少なくとも256ビット）で表現
コンパイル時に評価され、コンパイラの内部表現の範囲内であればオーバーフローしない
変数への代入時に変数の型の範囲外ならコンパイルエラー


2. 型付き定数:  
宣言時に指定された型の範囲内に収まる必要がある
範囲外の値を割り当てるとコンパイルエラー


3. 実行時の演算:
整数演算のオーバーフローはラップアラウンド（桁あふれ）する
Go言語ではオーバーフローの検出や例外を標準で提供していない


4. 型変換:
定数から別の型への変換で、変換先の型の範囲外ならコンパイルエラー
実行時の型変換でもオーバーフローはラップアラウンドする

サンプル
```go
package main

import (
	"fmt"
)

func main() {
	// 型なし定数の例
	const bigUntyped = 1 << 100  // コンパイル時に評価され、問題なし
	fmt.Println("1<<100 (表現可能な範囲): ", bigUntyped)

	// 型なし定数の計算 - コンパイル時に計算され、オーバーフローしない
	const bigCalculation = 1<<30 * 1<<30  // 2^60
	fmt.Println("1<<30 * 1<<30 (コンパイル時計算): ", bigCalculation)

	// 型付き定数の例 - 範囲内に収まる必要がある
	const smallInt int8 = 127
	// const overflowInt int8 = 128  // コンパイルエラー: constant 128 overflows int8

	// 定数から変数への代入
	var i int8 = 100  // OK
	// var j int8 = 1000  // コンパイルエラー: constant 1000 overflows int8

	// 実行時のオーバーフロー
	var a int8 = 127
	a++
	fmt.Println("int8 overflow (127++): ", a)  // -128 が出力される

	var b uint8 = 255
	b++
	fmt.Println("uint8 overflow (255++): ", b)  // 0 が出力される

	// 型なし定数を使った計算の例
	const trillion = 1000000000000
	fmt.Println("trillion (型なし): ", trillion)

	// 実行時の演算とオーバーフロー
	var c int32 = 2147483647  // int32の最大値
	var d int32 = 1
	fmt.Println("int32 max + 1 (オーバーフロー): ", c+d)  // -2147483648 が出力される

	// コンパイル時に検出可能なオーバーフロー (コメントアウトしているのでコンパイルエラーにならない)
	// const overflowTest int64 = 9223372036854775807 + 1  // コンパイルエラー: constant 9223372036854775808 overflows int64

	// 型変換時のオーバーフロー
	var largeuint uint64 = 18446744073709551615  // uint64の最大値
	// var converted int64 = int64(largeuint)  // コンパイルエラー: constant 18446744073709551615 overflows int64
}
```

安全な整数演算を行いたい場合は、math/bigパッケージを使用するか、オーバーフローを自分で検出するコードを書く必要があります。

### Section3 データ型について学ぼう
文字列型からインデックスで指定した文字を取得した場合、GoではASCIIコードが出力されるため、文字として出力する場合は`string()`を使った型変換が必要。  
```go
import "fmt"

func main() {
	s := "Hello, World!"
	x := s[1]
	y := string(x)

	fmt.Printf("x: %T, %v\n", x, x) // x: uint8, 101
	fmt.Printf("y: %T, %v\n", y, y) // y: string, e
}
```

**文字列リテラル**  
Goではダブルクォートやバッククォートで囲んだ文字列を文字列リテラルという。配列とは異なるため、配列の要素に値を代入する形で、文字を代入することはできない。


`\`とバッククォートは`"`を文字列として、出力する際にも使う。

データ型を変換することを**cast**という。

### Section4 データ構造のしくみを学ぼう
**Goにおける配列とスライスの違い**

- **配列（Array）**
    1. 固定長：
       - 宣言時にサイズを指定し、後から変更できない
       - var a [5]int のように定義
    2. 値型：
       - 配列は値型（value type）
       - 関数に渡す際は値がコピーされる
       - 配列の代入は全要素のコピーが発生する
    3. メモリ割り当て：
       - コンパイル時に固定サイズのメモリが割り当てられる
       - スタック上に確保されることが多い（小さい配列の場合）
    4. 使用例：
       ```go
       var arr [5]int                    // 要素5個の配列を宣言
       arr2 := [3]string{"Go", "Java", "Python"} // 初期化と同時に宣言
       arr3 := [...]int{1, 2, 3, 4}      // サイズを自動的に決定（この場合は4）
       ```
- **スライス（Slice）**
    1. 可変長：
       - 長さが動的に変更可能
       - var s []int のように定義（サイズ指定なし）
    2. 参照型：
       - スライスは参照型（reference type）
       - 関数に渡す際は参照が渡される（実データはコピーされない）
       - スライスの代入は参照のコピーのみ
    3. 内部構造：
       - ポインタ、長さ(length)、容量(capacity)の3つの要素を持つ
       - 実際のデータは別の場所（ヒープ上）に保存
    4. 使用例：
       ```go
       var s []int                        // 空のスライスを宣言
       s = make([]int, 5)                 // 要素5個のスライスを作成
       s = make([]int, 3, 10)             // 長さ3、容量10のスライスを作成
       s2 := []string{"Go", "Java", "Python"} // 初期化と同時に宣言
       ```
長さ0のスライスを作成する方法はmake関数を使う方法と`var 変数名 []型`の2通りある。  
前者はメモリに確保されているが、後者はnilという状態でメモリに確保されていない。


## Lesson2 ステートメント


## Lesson3 ポインタ


## Lesson4 Structオリエンテッド


## Lesson5 ゴルーチン


## Lesson6 パッケージ


## Lesson7 Webアプリケーションの作成
