package activity

import (
	"fmt"
	"log"
	"strconv"
)

const
(
	// 记录每一题的正确总人数
	cacheCorrectOfAll = "questions:correct:questionNum_%d"
	// 记录每一题的错误总人数
	cacheErrorOfAll = "questions:error:questionNum_%d"
	// 记录每一题的回答总人数
	cacheAll = "questions:all:questionNum_%d"
)

type AnswerResult struct {
	AllPerson     int     `json:"all_person"`
	CorrectPerson int     `json:"correct_person"`
	WrongPerson   int     `json:"wrong_person"`
	CorrectRate   float64 `json:"correct_rate"`
	Result        bool    `json:"result"`
}

var (
	allUser  map[int]bool
	UserChan = map[int]chan AnswerResult{}
)

func processAnswer(uId, qId int, result bool) {
	rds.incr(fmt.Sprintf(cacheAll, qId))
	if result {
		rds.incr(fmt.Sprintf(cacheCorrectOfAll, qId))
	} else {
		rds.incr(fmt.Sprintf(cacheErrorOfAll, qId))
	}
	allUser[uId] = result
	log.Println("处理用户", uId, "答题结果完成")
}

func pushSolutionInfo(qId int) {
	log.Println("开始推送结果到用户")
	all := getCountOfQuestion(qId, cacheAll)
	errors := getCountOfQuestion(qId, cacheErrorOfAll)
	correct := getCountOfQuestion(qId, cacheCorrectOfAll)
	// push Result
	for uId, channel := range UserChan {
		go func() {
			result := AnswerResult{
				all,
				correct,
				errors,
				correctRate(all, correct),
				allUser[uId],
			}
			log.Println("用户", uId, "推送结果是：", result)
			channel <- result
		}()
	}
}

func RegisterChan(uId int) {
	UserChan[uId] = make(chan AnswerResult)
}

func ReleaseChan(uId int) {
	close(UserChan[uId])
	delete(UserChan, uId)
}

func resetReport(qId int) {
	allUser = make(map[int]bool)
	rds.set(fmt.Sprintf(cacheAll, qId), "0")
	rds.set(fmt.Sprintf(cacheCorrectOfAll, qId), "0")
	rds.set(fmt.Sprintf(cacheErrorOfAll, qId), "0")
}

func correctRate(all, correct int) float64 {
	if float64(all) == 0 {
		return 0
	}
	return two(float64(correct) / float64(all))
}

func two(v float64) float64 {
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v), 4)
	return f
}

func getCountOfQuestion(qId int, key string) int {
	count, e := strconv.Atoi(rds.get(fmt.Sprintf(key, qId)))
	fmt.Println(count)
	if e != nil {
		count = 0
	}
	return count
}
