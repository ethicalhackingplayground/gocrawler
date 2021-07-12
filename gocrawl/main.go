package main

import (
	"bufio"
	"flag"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/ethicalhackingplayground/gocrawler/gocrawl/crawler"
	"github.com/gocolly/colly"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

var (
	threads int
	depth   int
)

func printBanner() {

	// Banner
	const banner = `

   ______      ______                    __         
  / ____/___  / ____/________ __      __/ /__  _____
 / / __/ __ \/ /   / ___/ __  / | /| / / / _ \/ ___/
/ /_/ / /_/ / /___/ /  / /_/ /| |/ |/ / /  __/ /    
\____/\____/\____/_/   \__,_/ |__/|__/_/\___/_/     
                                                    
 `

	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)

	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msg("\t\tgithub.com/ethicalhackingplayground\n\n")

	gologger.Info().Msg("Use with caution. You are responsible for your actions\n")
	gologger.Info().Msg("Developers assume no liability and are not responsible for any misuse or damage.\n\n")
}

func main() {

	// Print the ascii banner
	printBanner()
	flag.IntVar(&threads, "t", 5, "the number of concurrent threads")
	flag.IntVar(&depth, "d", 5, "the crawl depth")

	// Parse the arguments
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	hosts := make(chan string)
	var wg sync.WaitGroup

	// Crawling is enabled
	c := colly.NewCollector(
		colly.MaxDepth(depth),
	)

	// Use parallelism to speed up the processing
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			for url := range hosts {
				// Scan using burps XML file
				crawler.Crawl(c, url)
			}
			wg.Done()
		}()
	}

	// Iterate over Stdin and parse the parameters to test.
	uScanner := bufio.NewScanner(os.Stdin)
	for uScanner.Scan() {

		hosts <- uScanner.Text()
	}
	close(hosts)
	wg.Wait()
}
