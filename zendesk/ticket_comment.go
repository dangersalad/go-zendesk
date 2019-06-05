package zendesk

import (
	"time"
)

// TicketComment is struct for ticket comment payload
//
// https://developer.zendesk.com/rest_api/docs/support/ticket_comments
type TicketComment struct {
	ID          int64        `json:"id,omitempty"`
	Type        string       `json:"type,omitempty"`
	Body        string       `json:"body"`
	HTMLBody    string       `json:"html_body,omitempty"`
	PlainBody   string       `json:"plain_body,omitempty"`
	Public      bool         `json:"public,omitempty"`
	AuthorID    int64        `json:"author_id"`
	Attachments []Attachment `json:"attachments,omitempty"`
	// TODO: Via
	// TODO: metadata: comment flags
	CreatedAt time.Time `json:"created_at"`
}
