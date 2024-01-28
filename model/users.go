package model

type Users struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID      string `json:"uuid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Total_Amt []struct {
		Month            string `json:"month"`
		Year             string `json:"year"`
		Amount_Threshold int    `json:"amount"`
	} `json:"total_amt"`
	Spends []struct {
		Date         string `json:"date"`
		Month        string `json:"month"`
		Year         string `json:"year"`
		Amount_Spent int    `json:"amount_spent"`
		Description  string `json:"desc"`
	} `json:"spends"`
}
