require './double'
require 'stringio'
require 'test/unit'

class DoubleTest < Test::Unit::TestCase
  def test_double
    stdin  = StringIO.new("hoge\n", "r")
    stdout = StringIO.new("", "w")

    double(stdin, stdout)

    assert_equal("hoge\nhoge\n", stdout.string)
  end
end
