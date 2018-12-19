package activity

import (
	"log"
	"time"
)

type Activity struct {
	Questions  []Card
	CardState  State
	Processing bool
}

type State struct {
	Qid      int
	Card     Card
	CanReply bool
}

var (
	QuestionChan    chan Question
	ExplanationChan chan Explanation
	EndChan         chan interface{}
	CrtActivity     Activity
)

func init() {
	CrtActivity = Activity{}
	log.Println("do init activity.", CrtActivity)
}

func (ac *Activity) Build(questions []Card) {
	ac.Questions = questions
}

func (ac *Activity) Begin() {
	QuestionChan = make(chan Question)
	ExplanationChan = make(chan Explanation)
	EndChan = make(chan interface{})
	ac.Processing = true
	time.Sleep(5 * time.Second)
	for _, card := range ac.Questions {
		resetReport(card.Question.Id)
		ac.setState(card)
		ac.pushQuestion(card.Question)
		ac.pushExplanation(card.Explanation)
		ac.pushSolution(card.Solution)
	}
	ac.end()
}

func (ac *Activity) setState(card Card) {
	ac.CardState.Qid = card.Question.Id
	ac.CardState.Card = card
}

func (ac *Activity) end() {
	close(QuestionChan)
	close(ExplanationChan)
	close(EndChan)
	CrtActivity = Activity{Processing: false}
}

func (ac *Activity) pushQuestion(question Question) {
	QuestionChan <- question
	ac.WaitAnswer()
}

func (ac *Activity) pushExplanation(explanation Explanation) {
	ExplanationChan <- explanation
	ac.WaitGenerate()
}

func (ac *Activity) pushSolution(solution Solution) {
	pushSolutionInfo(ac.CardState.Qid)
	ac.WaitGenerate()
}

func (ac *Activity) WaitAnswer() {
	CrtActivity.CardState.CanReply = true
	time.Sleep(15 * time.Second)
	CrtActivity.CardState.CanReply = false
}

func (ac *Activity) WaitGenerate() {
	time.Sleep(15 * time.Second)
}

func (ac *Activity) Answer(uId int, option Option) {
	log.Println(ac.CardState.Card.Solution.Key)
	log.Println(option.Key)
	result := ac.CardState.Card.Solution.Key == option.Key
	log.Println(result)
	processAnswer(uId, option.Qid, result)
}
