package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTicketServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "ticket": {
            "subject": "Ticket subject",
            "date_created": "2024-06-05T10:13:31+00:00",
            "date_last_entry": "2024-06-05T10:13:31+00:00",
            "last_reply_name": "John Doe",
            "entryStatus": "open",
            "linked_subscriptions": [
              {
                "description": "65536.00 MB Regular Cloud Compute",
                "status": "active",
                "type": "vps",
                "uuid": "8481b17d-12f0-4c62-92c3-4ee8c2752761"
              }
            ]
          }
		}`
		fmt.Fprint(writer, response)
	})

	req := &TicketReq{
		Subject:          "Ticket subject",
		SubscriptionUUID: "8481b17d-12f0-4c62-92c3-4ee8c2752761",
	}

	ticket, _, err := client.Ticket.Create(ctx, req)
	if err != nil {
		t.Errorf("Ticket.Create returned %+v", err)
	}

	expected := &Ticket{
		Subject:       "Ticket subject",
		DateCreated:   "2024-06-05T10:13:31+00:00",
		DateLastEntry: "2024-06-05T10:13:31+00:00",
		LastReplyName: "John Doe",
		Status:        "open",
		LinkedSubscriptions: []TicketSubscription{
			{
				Description: "65536.00 MB Regular Cloud Compute",
				Status:      "active",
				Type:        "vps",
				UUID:        "8481b17d-12f0-4c62-92c3-4ee8c2752761",
			},
		},
	}

	if !reflect.DeepEqual(ticket, expected) {
		t.Errorf("Ticket.Create returned %+v, expected %+v", ticket, expected)
	}
}

func TestTicketServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets/ABC-12345", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "ticket": {
            "subject": "Ticket subject",
            "date_created": "2024-06-05T10:13:31+00:00",
            "date_last_entry": "2024-06-05T10:13:31+00:00",
            "last_reply_name": "John Doe",
            "entryStatus": "open",
            "linked_subscriptions": [
              {
                "description": "65536.00 MB Regular Cloud Compute",
                "status": "active",
                "type": "vps",
                "uuid": "8481b17d-12f0-4c62-92c3-4ee8c2752761"
              }
            ]
          }
		}`
		fmt.Fprint(writer, response)
	})

	ticket, _, err := client.Ticket.Get(ctx, "ABC-12345")
	if err != nil {
		t.Errorf("Ticket.Get returned %+v", err)
	}

	expected := &Ticket{
		Subject:       "Ticket subject",
		DateCreated:   "2024-06-05T10:13:31+00:00",
		DateLastEntry: "2024-06-05T10:13:31+00:00",
		LastReplyName: "John Doe",
		Status:        "open",
		LinkedSubscriptions: []TicketSubscription{
			{
				Description: "65536.00 MB Regular Cloud Compute",
				Status:      "active",
				Type:        "vps",
				UUID:        "8481b17d-12f0-4c62-92c3-4ee8c2752761",
			},
		},
	}

	if !reflect.DeepEqual(ticket, expected) {
		t.Errorf("Ticket.Get returned %+v, expected %+v", ticket, expected)
	}
}

func TestTicketServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "tickets": [
			{
			  "subject": "Ticket subject",
			  "date_created": "2024-06-05T10:13:31+00:00",
			  "date_last_entry": "2024-06-05T10:13:31+00:00",
			  "last_reply_name": "John Doe",
			  "entryStatus": "open",
			  "linked_subscriptions": [
				{
				  "description": "65536.00 MB Regular Cloud Compute",
				  "status": "active",
				  "type": "vps",
				  "uuid": "8481b17d-12f0-4c62-92c3-4ee8c2752761"
				}
			  ]
			}
		  ],
		  "meta": {
			"total": 1,
			"links": {
			  "next": "",
			  "prev": ""
			}
		  }
		}`
		fmt.Fprint(writer, response)
	})

	tickets, meta, _, err := client.Ticket.List(ctx, nil)
	if err != nil {
		t.Errorf("Ticket.List returned %+v", err)
	}

	expected := []Ticket{
		{
			Subject:       "Ticket subject",
			DateCreated:   "2024-06-05T10:13:31+00:00",
			DateLastEntry: "2024-06-05T10:13:31+00:00",
			LastReplyName: "John Doe",
			Status:        "open",
			LinkedSubscriptions: []TicketSubscription{
				{
					Description: "65536.00 MB Regular Cloud Compute",
					Status:      "active",
					Type:        "vps",
					UUID:        "8481b17d-12f0-4c62-92c3-4ee8c2752761",
				},
			},
		},
	}

	if !reflect.DeepEqual(tickets, expected) {
		t.Errorf("Ticket.List returned %+v, expected %+v", tickets, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
		Links: &Links{
			Next: "",
			Prev: "",
		},
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Ticket.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestTicketServiceHandler_Close(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets/ABC-12345", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "")
	})

	err := client.Ticket.Close(ctx, "ABC-12345")
	if err != nil {
		t.Errorf("Ticket.Close returned %+v", err)
	}
}

func TestTicketServiceHandler_CreateReply(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets/ABC-12345/replies", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	req := &TicketReplyReq{
		Description: "Reply message",
		Attachments: []TicketAttachmentReq{
			{
				FileName: "upload_file.txt",
				File:     []byte{},
			},
		},
	}

	err := client.Ticket.CreateReply(ctx, "ABC-12345", req)
	if err != nil {
		t.Errorf("Ticket.CreateReply returned %+v", err)
	}
}

func TestTicketServiceHandler_ListReplies(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets/ABC-12345/replies", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "replies": [
			{
			  "index": 0,
			  "age": "6d",
			  "attachments": [
				{
				  "context_type": "text/plain",
				  "filename": "uploaded_file.txt",
				  "index": 0
				}
			  ],
			  "date": "2024-06-05T10:13:31+00:00",
			  "description": "Reply message",
			  "from_email": "user@example.com",
			  "from_name": "Example User",
			  "from_type": "user",
			  "review_comment": null,
			  "review_date": null,
			  "review_rating": "10",
			  "reviewable": "0"
			}
		  ]
		}`
		fmt.Fprint(writer, response)
	})

	replies, _, err := client.Ticket.ListReplies(ctx, "ABC-12345")
	if err != nil {
		t.Errorf("Ticket.ListReplies returned %+v", err)
	}

	expected := []TicketReply{
		{
			Index: 0,
			Age:   "6d",
			Attachments: []TicketReplyAttachment{
				{
					ContextType: "text/plain",
					FileName:    "uploaded_file.txt",
					Index:       0,
				},
			},
			Date:         "2024-06-05T10:13:31+00:00",
			Description:  "Reply message",
			FromEmail:    "user@example.com",
			FromName:     "Example User",
			FromType:     "user",
			ReviewRating: "10",
			Reviewable:   "0",
		},
	}

	if !reflect.DeepEqual(replies, expected) {
		t.Errorf("Ticket.ListReplies returned %+v, expected %+v", replies, expected)
	}
}

func TestTicketServiceHandler_RateReply(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets/ABC-12345/replies/123/review", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	req := &TicketReplyRatingReq{
		Comment: "Reply message",
		Rating:  1,
	}

	err := client.Ticket.RateReply(ctx, "ABC-12345", 123, req)
	if err != nil {
		t.Errorf("Ticket.RateReply returned %+v", err)
	}
}

func TestTicketServiceHandler_GetReplyAttachment(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets/ABC-12345/replies/0/attachments/0", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "attachment": {
		    "context_type": "File type",
		    "filename": "uploaded_file.txt",
		    "filesize": 16
		  }
		}`
		fmt.Fprint(writer, response)
	})

	attachment, _, err := client.Ticket.GetReplyAttachment(ctx, "ABC-12345", 0, 0)
	if err != nil {
		t.Errorf("Ticket.GetReplyAttachment returned %+v", err)
	}

	expected := &TicketAttachment{
		ContextType: "File type",
		FileName:    "uploaded_file.txt",
		FileSize:    16,
	}

	if !reflect.DeepEqual(attachment, expected) {
		t.Errorf("Ticket.GetReplyAttachment returned %+v, expected %+v", attachment, expected)
	}
}

func TestTicketServiceHandler_RequestSMTPUnblock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets/smtp", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "reference": "ABC-12345"
		}`
		fmt.Fprint(writer, response)
	})

	req := &TicketSMTPUnblockReq{
		Nature:    "Transactional emails for our SaaS application",
		Volume:    "Approximately 10,000 emails per day",
		Type:      "No, we do not send marketing promotions, newsletters or coupons",
		Signature: "John Doe",
	}

	ticket, _, err := client.Ticket.RequestSMTPUnblock(ctx, req)
	if err != nil {
		t.Errorf("Ticket.RequestSMTPUnblock returned %+v", err)
	}

	expected := "ABC-12345"

	if !reflect.DeepEqual(ticket, expected) {
		t.Errorf("Ticket.RequestSMTPUnblock returned %+v, expected %+v", ticket, expected)
	}
}

func TestTicketServiceHandler_RequestTaxExemption(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/tickets/tax", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
		  "reference": "ABC-12345"
		}`
		fmt.Fprint(writer, response)
	})

	req := &TicketTaxExemptionReq{
		TaxID:     "123456789",
		Signature: "John Doe",
		Attachments: []TicketAttachmentReq{
			{
				FileName: "tax-exempt.pdf",
				File:     []byte{},
			},
		},
	}

	ticket, _, err := client.Ticket.RequestTaxExemption(ctx, req)
	if err != nil {
		t.Errorf("Ticket.RequestTaxExemption returned %+v", err)
	}

	expected := "ABC-12345"

	if !reflect.DeepEqual(ticket, expected) {
		t.Errorf("Ticket.RequestTaxExemption returned %+v, expected %+v", ticket, expected)
	}
}
