% Writing Command Line Tools in Go
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
* 覚えておくべきテクニック

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

# 覚えておくべきテクニック

* `io.Reader` と `io.Writer` を使いこなす
