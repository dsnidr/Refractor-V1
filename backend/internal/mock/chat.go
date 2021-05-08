package mock

import (
	"github.com/sniddunc/refractor/refractor"
)

type mockChatRepo struct {
	messages map[int64]*refractor.ChatMessage
}

func NewMockChatRepo(mockMessages map[int64]*refractor.ChatMessage) refractor.ChatRepository {
	return &mockChatRepo{
		messages: mockMessages,
	}
}

func (r *mockChatRepo) Create(message *refractor.ChatMessage) (*refractor.ChatMessage, error) {
	newID := int64(len(r.messages) + 1)

	r.messages[newID] = message

	message.MessageID = newID

	return message, nil
}

func (r *mockChatRepo) FindByID(id int64) (*refractor.ChatMessage, error) {
	for _, message := range r.messages {
		if message.MessageID == id {
			return message, nil
		}
	}

	return nil, refractor.ErrNotFound
}

func (r *mockChatRepo) FindMany(args refractor.FindArgs) ([]*refractor.ChatMessage, error) {
	var messages []*refractor.ChatMessage

	for _, message := range r.messages {
		if args["MessageID"] != nil && args["MessageID"].(int64) != message.MessageID {
			continue
		}

		if args["PlayerID"] != nil && args["PlayerID"].(int64) != message.PlayerID {
			continue
		}

		if args["ServerID"] != nil && args["ServerID"].(int64) != message.ServerID {
			continue
		}

		var startDate int64 = 0
		var endDate int64 = 0

		if args["StartDate"] != nil {
			startDate = args["StartDate"].(int64)
		}

		if args["EndDate"] != nil {
			startDate = args["EndDate"].(int64)
		}

		if message.DateRecorded < startDate || message.DateRecorded > endDate {
			continue
		}

		messages = append(messages, message)
	}

	// If no matches were found, return ErrNotFound
	if len(messages) < 1 {
		return nil, refractor.ErrNotFound
	}

	// Otherwise return the matches
	return messages, nil
}

func (r *mockChatRepo) Search(args refractor.FindArgs, limit int, offset int, getPlayerName refractor.PlayerNameGetter) (int, []*refractor.ChatMessage, error) {
	panic("implement me")
}
