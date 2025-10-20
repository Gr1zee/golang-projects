package message_parser

import (
	"strings"
	"time"
	"unicode/utf8"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func GetTasks(text string, user *string, status *string) []Ticket {
	var result []Ticket
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "_")

		if len(parts) < 4 {
			continue
		}

		ticketID := parts[0]
		if utf8.RuneCountInString(ticketID) < 6 || ticketID[:6] != "TICKET" {
			continue
		}

		date, err := time.Parse("2006-01-02", parts[3])
		if err != nil {
			continue
		}

		if parts[2] != "Готово" &&
			parts[2] != "В работе" &&
			parts[2] != "Не будет сделано" {
			continue
		}

		matchesUser := user == nil || *user == parts[1]
		matchesStatus := status == nil || *status == parts[2]

		if matchesUser && matchesStatus {
			result = append(result, Ticket{
				Ticket: ticketID,
				User:   parts[1],
				Status: parts[2],
				Date:   date,
			})
		}
	}

	return result
}
