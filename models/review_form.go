package models

type ReviewForm struct {
	TicketID int    `json:"ticket_id,omitempty" example:"1"`
	Rating   int    `json:"rating" example:"5"`
	Content  string `json:"content" example:"What's Enjoyable : Story\nWhat did you feel : Tension\nComments : Great Movie!"`
	Date     string `json:"date" example:"25/04/2024"`
}

//	{
//	  "data": {
//	    "ticket_id": 1,
//	    "rating": 5,
//	    "content": "What's Enjoyable : Story\nWhat did you feel : Tension\nComments : Great Movie!",
//	    "date": "25/04/2024"
//	  }
//	}

type ReviewFormBody struct {
	ReviewForm ReviewForm `json:"data"`
}
