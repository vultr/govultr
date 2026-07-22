package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const ticketsPath = "/v2/tickets"

// TicketService is the interface to interact with the tickets endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/tickets
type TicketService interface {
	Create(ctx context.Context, ticketReq *TicketReq) (*Ticket, *http.Response, error)
	Get(ctx context.Context, ticketID string) (*Ticket, *http.Response, error)
	List(ctx context.Context, options *ListOptions) ([]Ticket, *Meta, *http.Response, error)
	Close(ctx context.Context, ticketID string) error

	CreateReply(ctx context.Context, ticketID string, ticketReplyReq *TicketReplyReq) error
	ListReplies(ctx context.Context, ticketID string) ([]TicketReply, *http.Response, error)
	RateReply(ctx context.Context, ticketID string, replyIdx int, ratingReq *TicketReplyRatingReq) error
	GetReplyAttachment(ctx context.Context, ticketID string, replyIdx int, attachmentIdx int) (*TicketAttachment, *http.Response, error)

	RequestSMTPUnblock(ctx context.Context, smtpReq *TicketSMTPUnblockReq) (string, *http.Response, error)
	RequestTaxExemption(ctx context.Context, taxReq *TicketTaxExemptionReq) (string, *http.Response, error)
}

// TicketServiceHandler handles interaction with the server methods for the Vultr API
type TicketServiceHandler struct {
	client *Client
}

// Ticket represents a customer support ticket
type Ticket struct {
	Subject             string               `json:"subject"`
	DateCreated         string               `json:"date_created"`
	DateLastEntry       string               `json:"date_last_entry"`
	LastReplyName       string               `json:"last_reply_name"`
	LinkedSubscriptions []TicketSubscription `json:"linked_subscriptions"`
	Reference           string               `json:"reference"`
	Status              string               `json:"entryStatus"`
}

// ticketBase represents the base response for a create ticket request
type ticketBase struct {
	Ticket *Ticket `json:"ticket"`
}

// ticketsBase represents the base response for a list tickets request
type ticketsBase struct {
	Tickets []Ticket `json:"tickets"`
	Meta    *Meta    `json:"meta"`
}

// TicketAttachment represents a file attached to a ticket reply
type TicketAttachment struct {
	ContextType string `json:"context_type"`
	File        string `json:"file"`
	FileName    string `json:"filename"`
	FileSize    int    `json:"filesize"`
}

// ticketAttachmentBase represents the base response for a getting a ticket reply attachment request
type ticketAttachmentBase struct {
	Attachment *TicketAttachment `json:"attachment"`
}

// TicketSubscription represents a subscription linked to the ticket
type TicketSubscription struct {
	Description string `json:"description"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	UUID        string `json:"uuid"`
}

// TicketReply represents a reply to a customer ticket
type TicketReply struct {
	Index         int                     `json:"index"`
	Age           string                  `json:"age"`
	Date          string                  `json:"date"`
	Description   string                  `json:"description"`
	FromEmail     string                  `json:"from_email"`
	FromName      string                  `json:"from_name"`
	FromType      string                  `json:"from_type"`
	ReviewComment string                  `json:"review_comment"`
	ReviewDate    string                  `json:"review_date"`
	ReviewRating  string                  `json:"review_rating"`
	Reviewable    string                  `json:"reviewable"`
	Attachments   []TicketReplyAttachment `json:"attachments"`
}

// ticketRepliesBase represents the base structure for listing replies to a ticket
type ticketRepliesBase struct {
	Replies []TicketReply `json:"replies"`
}

// TicketReplyAttachment represents a file attached to a reply
type TicketReplyAttachment struct {
	Index       int    `json:"index"`
	ContextType string `json:"context_type"`
	FileName    string `json:"filename"`
}

// TicketReq is used for creating a new ticket
type TicketReq struct {
	Subject          string                `json:"subject"`
	Description      string                `json:"description"`
	SubscriptionUUID string                `json:"sub-uuid,omitempty"`
	Category         string                `json:"category,omitempty"`
	Attachments      []TicketAttachmentReq `json:"attachments,omitempty"`
}

// TicketReplyReq is used for creating a new reply to a ticket
type TicketReplyReq struct {
	Description string                `json:"description"`
	Attachments []TicketAttachmentReq `json:"attachments,omitempty"`
}

// TicketReplyRatingReq is used for rating a ticket reply
type TicketReplyRatingReq struct {
	Comment string `json:"comment"`
	Rating  int    `json:"rating"`
}

// TicketAttachmentReq represents a file to attach to a ticket
type TicketAttachmentReq struct {
	FileName string `json:"filename"`
	File     string `json:"file"`
}

// TicketSMTPUnblockReq is used for creating a request to unblock SMTP
type TicketSMTPUnblockReq struct {
	Nature    string `json:"nature"`
	Volume    string `json:"volume"`
	Type      string `json:"type"`
	Signature string `json:"signature"`
}

// TicketTaxExemptionReq is used for creating a tax exemption request ticket
type TicketTaxExemptionReq struct {
	TaxID       string                `json:"tax_id"`
	Signature   string                `json:"signature"`
	Attachments []TicketAttachmentReq `json:"attachments"`
}

// Create a new ticket
func (c *TicketServiceHandler) Create(ctx context.Context, ticketReq *TicketReq) (*Ticket, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, ticketsPath, ticketReq)
	if err != nil {
		return nil, nil, err
	}

	ticket := new(ticketBase)
	resp, err := c.client.DoWithContext(ctx, req, ticket)
	if err != nil {
		return nil, resp, err
	}

	return ticket.Ticket, resp, nil
}

// Get information about a ticket
func (c *TicketServiceHandler) Get(ctx context.Context, ticketID string) (*Ticket, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", ticketsPath, ticketID)

	req, err := c.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	ticket := new(ticketBase)
	resp, err := c.client.DoWithContext(ctx, req, ticket)
	if err != nil {
		return nil, resp, err
	}

	return ticket.Ticket, resp, nil
}

// List all customer tickets on the account.
func (c *TicketServiceHandler) List(ctx context.Context, options *ListOptions) ([]Ticket, *Meta, *http.Response, error) { //nolint:dupl
	req, err := c.client.NewRequest(ctx, http.MethodGet, ticketsPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	tickets := new(ticketsBase)
	resp, err := c.client.DoWithContext(ctx, req, tickets)
	if err != nil {
		return nil, nil, resp, err
	}

	return tickets.Tickets, tickets.Meta, resp, nil
}

// Close an open ticket
func (c *TicketServiceHandler) Close(ctx context.Context, ticketID string) error {
	uri := fmt.Sprintf("%s/%s", ticketsPath, ticketID)
	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = c.client.DoWithContext(ctx, req, nil)
	return err
}

// CreateReply creates a new reply to an existing ticket
func (c *TicketServiceHandler) CreateReply(ctx context.Context, ticketID string, ticketReplyReq *TicketReplyReq) error {
	uri := fmt.Sprintf("%s/%s/replies", ticketsPath, ticketID)
	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, ticketReplyReq)
	if err != nil {
		return err
	}

	_, err = c.client.DoWithContext(ctx, req, nil)
	return err
}

// ListReplies lists all replies to a ticket
func (c *TicketServiceHandler) ListReplies(ctx context.Context, ticketID string) ([]TicketReply, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/replies", ticketsPath, ticketID)
	req, err := c.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	ticketReplies := new(ticketRepliesBase)
	resp, err := c.client.DoWithContext(ctx, req, ticketReplies)
	if err != nil {
		return nil, resp, err
	}

	return ticketReplies.Replies, resp, nil
}

// RateReply rates a ticket reply from Vultr
func (c *TicketServiceHandler) RateReply(ctx context.Context, ticketID string, replyIdx int, ratingReq *TicketReplyRatingReq) error {
	uri := fmt.Sprintf("%s/%s/replies/%d/review", ticketsPath, ticketID, replyIdx)
	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, ratingReq)
	if err != nil {
		return err
	}

	_, err = c.client.DoWithContext(ctx, req, nil)
	return err
}

// GetReplyAttachment gets a ticket reply attachment
func (c *TicketServiceHandler) GetReplyAttachment(ctx context.Context, ticketID string, replyIdx, attachmentIdx int) (*TicketAttachment, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/replies/%d/attachments/%d", ticketsPath, ticketID, replyIdx, attachmentIdx)

	req, err := c.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	attachment := new(ticketAttachmentBase)
	resp, err := c.client.DoWithContext(ctx, req, attachment)
	if err != nil {
		return nil, resp, err
	}

	return attachment.Attachment, resp, nil
}

// RequestSMTPUnblock submits a request to unblock SMTP port 25 on your account.
func (c *TicketServiceHandler) RequestSMTPUnblock(ctx context.Context, smtpReq *TicketSMTPUnblockReq) (string, *http.Response, error) {
	uri := fmt.Sprintf("%s/smtp", ticketsPath)
	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, smtpReq)
	if err != nil {
		return "", nil, err
	}

	ticket := new(Ticket)
	resp, err := c.client.DoWithContext(ctx, req, ticket)
	if err != nil {
		return "", resp, err
	}

	return ticket.Reference, resp, nil
}

func (c *TicketServiceHandler) RequestTaxExemption(ctx context.Context, taxReq *TicketTaxExemptionReq) (string, *http.Response, error) {
	uri := fmt.Sprintf("%s/tax", ticketsPath)
	req, err := c.client.NewRequest(ctx, http.MethodPost, uri, taxReq)
	if err != nil {
		return "", nil, err
	}

	ticket := new(Ticket)
	resp, err := c.client.DoWithContext(ctx, req, ticket)
	if err != nil {
		return "", resp, err
	}

	return ticket.Reference, resp, nil
}
