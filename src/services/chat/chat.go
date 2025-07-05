package chat

import "fmt"

func (ch *Chat) GetChat(q string) {
	fmt.Printf("question: %s", q)
	answer, err := ch.repository.OAI.Ask(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf(" answer:  %+v\n", answer.Choices[0].Message.Content)
}
