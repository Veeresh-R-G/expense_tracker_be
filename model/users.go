package model

type Users struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
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
