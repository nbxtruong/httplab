package main

type Beast struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	OtherNames  []string `json:"otherNames"`
}
