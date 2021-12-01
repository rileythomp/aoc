# aoc
Advent of Code Solutions.

Usage:\
NB: Must have environment variable `AOC_COOKIE` set to your AoC API cookie.

Running\
`go run autocurday.go`\
will wait until a new problem is released and then generate\
`aoc/<year>/day<day>/main.go`: An AoC boilerplate Go file.
`aoc/<year>/day<day>/input.txt`: The problem input from the AoC API.
`aoc/<year>/day<day>/main.txt`: An empty file for test input.

Running\
`go run mkday.go <year> <day>`\
will do the same as `go run autocurday.go` but for the specified year and day.