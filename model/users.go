package model

type Users struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Total_Amt []struct {
		Month            string `json:"month"`
		Year             string `json:"year"`
		Amount_Threshold int    `json:"amount"`
	}
	Spends []struct {
		Date         string `json:"date"`
		Month        string `json:"month"`
		Year         string `json:"year"`
		Amount_Spent int    `json:"amount_spent"`
	}
}
