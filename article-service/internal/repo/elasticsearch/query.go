package elasticsearch

import (
	"encoding/json"
)

func buildKeywordQuery(keyword string, limit int, offset int) ([]byte, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": []string{"category", "title"},
			},
		},
		"from": offset,
		"size": limit,
		"sort": []map[string]interface{}{
			{"published_date.keyword": map[string]string{"order": "desc"}},
		},
	}
	return json.Marshal(query)
}

func buildMatchAllQuery(limit int, offset int) ([]byte, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"from": offset,
		"size": limit,
		"sort": []map[string]interface{}{
			{"published_date.keyword": map[string]string{"order": "desc"}},
		},
	}
	return json.Marshal(query)
}

func buildCountQuery(keyword string, category string) ([]byte, error) {
	if keyword == "" && category == "" {
		return json.Marshal(map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
		})
	}

	must := []map[string]interface{}{}
	if keyword != "" {
		must = append(must, map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": []string{"title", "description", "content"},
			},
		})
	}
	if category != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"category": category,
			},
		})
	}

	return json.Marshal(map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": must,
			},
		},
	})
}

func buildCategoryQuery(category string, limit int, offset int) ([]byte, error) {
	var query map[string]interface{}
	query = map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match": map[string]interface{}{"category": category}},
				},
			},
		},
		"from": offset,
		"size": limit,
		"sort": []map[string]interface{}{
			{"published_date.keyword": map[string]string{"order": "desc"}},
		},
	}
	return json.Marshal(query)
}

func buildKeywordAndCategoryQuery(keyword string, category string, limit int, offset int) ([]byte, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match": map[string]interface{}{"category": category}},
					{
						"multi_match": map[string]interface{}{
							"query":  keyword,
							"fields": []string{"title", "category"},
						},
					},
				},
			},
		},
		"from": offset,
		"size": limit,
		"sort": []map[string]interface{}{
			{"published_date.keyword": map[string]string{"order": "desc"}},
		},
	}
	return json.Marshal(query)
}
