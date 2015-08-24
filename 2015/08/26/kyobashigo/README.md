% Go and UNIX Philosophy
% @yuya_takeyama
% 2015/08/26 Kyobashi.go #1

# こんにちは

* @yuya_takeyama
* 赤坂で PHP 書いてます
* 9 月から京橋で Ruby とか書いてると思います
* Go は主にコマンドラインツールを作るのに利用
* コマンドラインツールは Ruby でも書く (自作のイチオシは [jr](https://github.com/yuya-takeyama/jr))

# 今まで Go で作ってきたもの

* [ltsv2json](https://github.com/yuya-takeyama/ltsv2json): LTSV を入力として受け取り JSON として出力
* [hi](https://github.com/yuya-takeyama/hi): Web API の通信内容をターミナルに出力するプロキシ
* [numstat](https://github.com/yuya-takeyama/numstat): 入力の統計情報 (最小・最大・平均等) を出力
* [percentile](https://github.com/yuya-takeyama/percentile): 入力した数値のパーセンタイル値を出力
* などなど

# 今日話す内容

* Go でコマンドラインツールを書くモチベーション
* UNIX 哲学とは
* Go における UNIX 哲学

# 今日話さない内容

* Go の基本文法とかの説明
* Goroutine
* Channel

# 何故 Go で書くか

* コンパイルしてしまえばバイナリを置くだけで動く
    * Ruby 等の処理系をインストールする必要がない
    * 依存ライブラリをインストールする必要もない
* クロスコンパイルができる
    * Mac で開発して Linux で動かすのも簡単
    * Windows でも動かせる (個人的にはやったことないけど)

# その他 Go やってみて感じたこと

* シンプルであること、やり方が少ないことの良さ
    * Ruby/Perl みたいに TIMTOWTDI ではない
    * 継承も例外もイテレータもジェネレータもないジェネリクスもないけど
* 開発ツールが標準で充実しているので生産性が高い
    * go test / go fmt / go tool pprof
* 静的型付も悪くない、というかむしろ良い

# それでも Go を使わない場合

* 言語処理系自体の機能を使いたい場合
    * Ruby で言語内 DSL を実装したい場合
    * JavaScript でプラグイン機構を実装したい場合
* Go 標準のデータ型がマッチしない場合
    * map が順序を持ってないと困る場合
    * JSON の操作も割と面倒

# UNIX 哲学とは

# UNIX 哲学とは

<div style="text-align: center;"><img src="gancarz.png" alt="UNIXという考え方" width="40%"></div>

# マイク・ガンカーズの UNIX 哲学 (一部抜粋)

* 小さいものは美しい。
* 各プログラムが一つのことをうまくやるようにせよ。
* 全てのプログラムはフィルタとして振る舞うようにせよ。

# UNIX プログラムとは

~~~~
$ seq 100 | awk '$1 % 2 == 0' | tail
82
84
86
88
90
92
94
96
98
100
~~~~

# UNIX プログラムとは

* 小さなプログラムを組み合わせる
* プログラム同士はパイプでつなぐ
* パイプは UNIX における標準インターフェイス
* Go でもこういうプログラムを作っていきたい！！

# Go におけるインターフェイス

+ Go にはクラスがない
+ Go には継承もない
* でもインターフェイスがある！

# Go におけるインターフェイス

~~~~ {.go}
type Stringer interface {
        String() string
}
~~~~

* `Stringer` インターフェイスを満たせば `fmt.Println` できる
* クラスの継承改装等は関係ない (そもそもクラスがない)
* 特定の構造体である必要もない
* インターフェイスに定められたメソッドがあればよい (ダックタイピング的)

# コマンドラインツールを作るなら

* `io.Reader`
* `io.Writer`
* これが UNIX でいうパイプ、標準インターフェイス！

# `io.Reader` と `io.Writer`

~~~~ {.go}
type Reader interface {
        Read(p []byte) (n int, err error)
}

type Writer interface {
        Write(p []byte) (n int, err error)
}
~~~~
