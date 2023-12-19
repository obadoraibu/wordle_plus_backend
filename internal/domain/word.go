package domain

type NewWordRequest struct {
	Length int
}

type NewWordResponse struct {
	Word       string `json:"word" binding:"required"`
	Definition string `json:"definition" binding:"required"`
}

type CheckWordRequest struct {
	Word string `json:"word" binding:"required"`
}

type DailyWordResponse struct {
	Word       string `json:"word" binding:"required"`
	Definition string `json:"definition" binding:"required"`
}

type DictionaryResponse []struct {
	Meanings []struct {
		Definitions []struct {
			Definition string `json:"definition"`
		} `json:"definitions"`
	} `json:"meanings"`
}
