# msds301-wk5

### Setup (Windows 11 + Git Bash)
- Clone repo with `git clone git@github.com:jeremycruzz/msds301-wk5.git`

### Building executable
- Run `go build -o scrapewiki.exe ./cmd/scrapewiki`

### Running Go executable
- Run `./scrapewiki.exe {threads: OPTIONAL}`
- results should be saved to `./results/corpus_{threads}.json`

### Importing sets
- add import
- `import github.com/jeremycruzz/msds301-wk5/sets`
- use `sets.Stopwords` in code

### Results

While I didn't grab the runtime for the python program, the Go program was significantly faster. It is important to note that the python program also saved the wikipages but even taking that into consideration Go was very fast. I did expect the 12 concurrent threads to run faster than the rest but not significantly faster since concurrency is pretty hard to predict. As the company data scientist I would highly recommend that crawlers be written in Go over python.

I got stuck for a really long time since I didn't call `colly.Wait()` and I couldn't figure it out. After that I figured I'd make a stopword package as some sort of utility that I can use in my other Go projects. There is more work to be done with parsing the text better. A lot of `\n` and `\t` appear in the text field.

<details> 
<summary> See results </summary>

| Concurrent Requests | Time (ns)    |
|---------------------|--------------|
| Go 1                   | 202749600   |
| Go 2                   | 202327000   |
| Go 3                   | 200905700   |
| Go 4                   | 219443600   |
| Go 5                   | 199313700   |
| Go 6                   | 202382800   |
| Go 7                   | 202605000   |
| Go 8                   | 197906400   |
| Go 9                   | 202593700   |
| Go 10                  | 199665100   |
| Go 11                  | 202349200   |
| Go 12                  | 198064600   |
| Python 1                | over 5 seconds   |

</details>