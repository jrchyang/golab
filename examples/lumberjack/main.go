package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	filename   string
	maxsize    int
	maxage     int
	maxbackups int
	localtime  bool
	compress   bool
)

func init() {
	flag.StringVar(&filename, "f", "/var/log/golab/test-lumberjack.log", "Filename is the file to write logs to.")
	flag.IntVar(&maxsize, "s", 10, "MaxSize is the maximum size in megabytes of the log file before it gets. It defaults to 100 megabytes.")
	flag.IntVar(&maxage, "a", 0, "MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename.")
	flag.IntVar(&maxbackups, "b", 3, "MaxBackups is the maximum number of old log files to retain.")
	flag.BoolVar(&localtime, "l", true, "LocalTime determines if the time used for formatting the timestamps in backup files is the computer's local time.")
	flag.BoolVar(&compress, "c", true, "Compress determines if the rotated log files should be compressed using gzip.")
	flag.Parse()
}

func checkAndOutputArgs() error {
	// check log file dir
	dirPath := filepath.Dir(filename)
	info, err := os.Stat(dirPath)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return os.ErrInvalid
	}

	fmt.Printf("\n  filename : %s\n", filename)
	fmt.Printf("   maxsize : %d\n", maxsize)
	fmt.Printf("    maxage : %d\n", maxage)
	fmt.Printf("maxbackups : %d\n", maxbackups)
	fmt.Printf(" localtime : %t\n", localtime)
	fmt.Printf("  compress : %t\n\n", compress)

	return nil
}

func generateRandomLog(index int) string {
	adjectives := []string{"quick", "brown", "happy", "sad", "angry", "calm", "brave", "weak", "strong", "smart"}
	nouns := []string{"fox", "dog", "cat", "bird", "tree", "rock", "mountain", "river", "ocean", "sun"}
	verbs := []string{"runs", "jumps", "flies", "swims", "walks", "climbs", "eats", "sleeps", "thinks", "laughs"}
	adverbs := []string{"quickly", "slowly", "happily", "sadly", "angrily", "calmly", "bravely", "weakly", "strongly", "smartly"}

	// 生成不同结构的句子，确保日志内容不重复
	rand.Seed(time.Now().UnixNano() + int64(index))
	structureType := rand.Intn(3)

	switch structureType {
	case 0:
		return fmt.Sprintf("Log entry %d: The %s %s %s %s.",
			index,
			adjectives[rand.Intn(len(adjectives))],
			nouns[rand.Intn(len(nouns))],
			verbs[rand.Intn(len(verbs))],
			adverbs[rand.Intn(len(adverbs))])
	case 1:
		return fmt.Sprintf("Log entry %d: A %s %s %s near the %s.",
			index,
			adjectives[rand.Intn(len(adjectives))],
			nouns[rand.Intn(len(nouns))],
			verbs[rand.Intn(len(verbs))],
			nouns[rand.Intn(len(nouns))])
	default:
		return fmt.Sprintf("Log entry %d: Why did the %s %s %s? Because it was %s.",
			index,
			nouns[rand.Intn(len(nouns))],
			verbs[rand.Intn(len(verbs))],
			adverbs[rand.Intn(len(adverbs))],
			adjectives[rand.Intn(len(adjectives))])
	}
}

func main() {
	if err := checkAndOutputArgs(); err != nil {
		fmt.Println("failed to parse args, please check")
		os.Exit(1)
	}

	log.SetOutput(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,
		MaxAge:     maxage,
		MaxBackups: maxbackups,
		LocalTime:  localtime,
		Compress:   compress,
	})

	for i := 0; i < 10000; i++ {
		log.Println(generateRandomLog(i))
		if i%1000 == 0 && i > 0 {
			fmt.Printf("already written %d log entries ...\n", i)
		}
	}
}
