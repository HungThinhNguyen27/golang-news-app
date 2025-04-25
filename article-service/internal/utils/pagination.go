package utils

func NormalizePagination(limit, page, maxLimit int) (int, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	} else if limit > maxLimit {
		limit = maxLimit
	}
	offset := (page - 1) * limit
	return limit, offset
}

func CalculateTotalPages(totalArticle, limit int) int {
	if limit == 0 {
		return 0
	}
	return (totalArticle + limit - 1) / limit
}
