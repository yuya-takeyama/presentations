% Go で I/O に依存したプログラムのユニットテスト
% @yuya_takeyama
% 2014/11/30@GoCon 2014 Autumn

# At first

* Mr. Pike, thank you for coming to Japan
* And thank you for great language
* 主催者の皆様、素晴らしいイベントと発表の機会をありがとうございます
* 楽天さん、会場をありがとうございます

# このスライドについて

* 発表に含まれるコードは [GitHub リポジトリ中](https://github.com/yuya-takeyama/presentations/tree/master/2014/11/30/gocon_2014_autumn) にあります
* Pandoc はじめて使いました。便利な運用についてご存知の方教えてください...

# こんにちは

* 普段は PHP で Web アプリ運用してます
* コマンドラインツールは Ruby で書いてます
* コマンドラインツールは徐々に Go に寄せつつあります
    * [db2yaml](https://github.com/yuya-takeyama/db2yaml)
    * [dbyaml2md](https://github.com/yuya-takeyama/dbyaml2md)
    * [envjson](https://github.com/yuya-takeyama/envjson)

# 本日のお題

I/O に依存したプログラムのユニットテスト

# in Ruby

`StringIO` を使う

~~~~ {.ruby}
def test_double
  stdin  = StringIO.new("hoge\n", "r")
  stdout = StringIO.new("", "w")

  double(stdin, stdout)

  assert_equal("hoge\nhoge\n", stdout.string)
end
~~~~

Python にも同名で同じような用途のクラスが存在

# in Go

`bytes.Buffer` を使う

# bytes.Buffer

* `[]byte` にいろんなメソッドが生えたような構造体
* 任意の `[]byte` や `string` を元に生成することができる
* `io.ReadWriter` インターフェイスを満たす
* だいたいファイルのように振る舞う

# io.ReadWriter

~~~~ {.go}
type ReadWriter interface {
	Reader
	Writer
}
~~~~

# io.Reader と io.Writer

~~~~ {.go}
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
~~~~

`os.File`, `net.TCPConn`, `net.UnixConn` などはこれらを満たしている。

# I/O に依存したプログラムを書いてみる

~~~~
$ echo hoge | go run double.go
hoge
hoge
~~~~

# main()

~~~~ {.go}
func main() {
	Double(os.Stdin, os.Stdout)
}
~~~~

# Double()

~~~~ {.go}
func Double(stdin io.Reader, stdout io.Writer) {
	buf, _ := ioutil.ReadAll(stdin)

	stdout.Write(buf)
	stdout.Write(buf)
}
~~~~

# ユニットテストを書いてみる

~~~~ {.go}
func TestDouble(t *testing.T) {
	stdin := bytes.NewBufferString("hoge\n")
	stdout := new(bytes.Buffer)

	Double(stdin, stdout)

	if stdout.String() != "hoge\nhoge\n" {
		t.Fatal("Not matched")
	}
}
~~~~

# Ruby のユニットテストと比較してみる

~~~~ {.ruby}
def test_double
  stdin  = StringIO.new("hoge\n", "r")
  stdout = StringIO.new("", "w")

  double(stdin, stdout)

  assert_equal("hoge\nhoge\n", stdout.string)
end
~~~~

# ポイント

* I/O に依存した部分を関数として切り出す
* `*os.File` ではなく `io.Reader` や `io.Writer` に依存する
* `Buffer.Bytes()` や `Buffer.String()` で中身を取り出して検査する
    * 検査それ自体は `string` なり `[]byte` なりの流儀に従って比較を行う

# bytes.Buffer 応用編

`exec.Cmd` の標準入出力を差し替えてコマンドの E2E テスト

~~~~ {.go}
func TestCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "double.go")
	cmd.Stdin = bytes.NewBufferString("hoge\n")
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout

	_ = cmd.Run()

	if stdout.String() != "hoge\nhoge\n" {
		t.Fatal("Not matched")
	}
}
~~~~

# ご静聴ありがとうございました
