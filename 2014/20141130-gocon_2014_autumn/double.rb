def double(stdin, stdout)
  buf = stdin.read

  stdout.write buf
  stdout.write buf
end

if $0 == __FILE__
  double(STDIN, STDOUT)
end
