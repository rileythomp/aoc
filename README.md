# aoc
Advent of Code Solutions.

### **Usage:**
NB: Must have environment variable `AOC_COOKIE` set to your AoC API cookie.

**Automated Submission**\
Running\
`go run submit.go <level> <input> <year> <day>`\
will automatically submit the output of\
`go run main.go <level> <input>`\
to the AoC API. It then generates the submission response locally and opens it with\
`open aoc/<year>/day<day>/submission.html`

The default values are:
- level: 1
- input: input.txt
- day:   current day
- year:  current year

---

**Automated Problem Fetching**\
Running\
`go run mkday.go <year> <day>`\
will generate\
`aoc/<year>/day<day>/problem.html`: The problem specification from the AoC API.\
`aoc/<year>/day<day>/input.txt`: The problem input from the AoC API.\
`aoc/<year>/day<day>/test.txt`: An empty file for test input.\
`aoc/<year>/day<day>/main.go`: An AoC boilerplate Go file.\
and open the problem specification with\
`open aoc/<year>/day<day>/problem.html`

The default values are:
- day:   next day
- year:  current year

So if a day is not given, it will wait until a new problem is released at midnight and then generate the files.

![midnight](imgs/midnight.png)

![output](imgs/output.png)
