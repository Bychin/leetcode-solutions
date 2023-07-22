package main

import (
	"fmt"
)

func largestWordCount(messages []string, senders []string) (maxSender string) {
	totalWordCounts := make(map[string]int, len(senders))
	maxWordCount := 0

	for i := range messages {
		wordCount := 1
		for _, c := range messages[i] {
			if c == ' ' {
				wordCount += 1
			}
		}

		totalWordCounts[senders[i]] += wordCount
		if totalWordCounts[senders[i]] > maxWordCount {
			maxWordCount = totalWordCounts[senders[i]]
			maxSender = senders[i]
		}
		if totalWordCounts[senders[i]] == maxWordCount && maxSender < senders[i] {
			maxSender = senders[i]
		}
	}

	fmt.Println(totalWordCounts)
	fmt.Println("EMoUs" < "FnZd")

	return
}

func main() {
	fmt.Println(largestWordCount(
		[]string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"},
		[]string{"Alice", "userTwo", "userThree", "Alice"},
	))
	fmt.Println(largestWordCount(
		[]string{"tP x M VC h lmD", "D X XF w V", "sh m Pgl", "pN pa", "C SL m G Pn v", "K z UL B W ee", "Yf yo n V U Za f np", "j J sk f qr e v t", "L Q cJ c J Z jp E", "Be a aO", "nI c Gb k Y C QS N", "Yi Bts", "gp No g s VR", "py A S sNf", "ZS H Bi De dj dsh", "ep MA KI Q Ou"},
		[]string{"OXlq", "IFGaW", "XQPeWJRszU", "Gb", "HArIr", "Gb", "FnZd", "FnZd", "HArIr", "OXlq", "IFGaW", "XQPeWJRszU", "EMoUs", "Gb", "EMoUs", "EMoUs"},
	))
}
