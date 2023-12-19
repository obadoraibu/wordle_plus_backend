package repository

func (r *Repository) GetNewWordFromStorage(length int) (string, error) {
	return "apple", nil
}

func (r *Repository) GetDailyWordFromStorage() (string, error) {
	r.Mutex.Lock()
	word := r.DailyWord
	r.Mutex.Unlock()
	return word, nil
}
