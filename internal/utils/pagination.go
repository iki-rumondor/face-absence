package utils

import (
	"fmt"
	"math"
)

func CalculatePagination(url string, totalRows, page, limit int) (map[string]interface{}, error) {

	var totalPages, fromRow, toRow int
	var previousPage, nextPage, firstPage, lastPage string

	if limit == 0 {
		totalPages = 0 
	} else {
		totalPages = int(math.Ceil(float64(totalRows) / float64(limit))) - 1
	}

	if page == 0 {
		fromRow = 1
		toRow = limit
	}

	if page != 0 && page <= totalPages {
		fromRow = (page-1)*limit + 1
		toRow = page * limit
	}

	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	firstPage = fmt.Sprintf("%s?limit=%d&page=%d", url, limit, 0)
	lastPage = fmt.Sprintf("%s?limit=%d&page=%d", url, limit, totalPages)

	if page > 0 {
		previousPage = fmt.Sprintf("%s?limit=%d&page=%d", url, limit, page-1)
	}

	if page < totalPages {
		nextPage = fmt.Sprintf("%s?limit=%d&page=%d", url, limit, page+1)
	}

	if page >= totalPages {
		nextPage = ""
	}

	pagination := map[string]interface{}{
		"limit":         limit,
		"page":          page,
		"total_rows":    totalRows,
		"total_pages":   totalPages,
		"from_row":      fromRow,
		"to_row":        toRow,
		"first_page":    firstPage,
		"last_page":     lastPage,
		"previous_page": previousPage,
		"next_page":     nextPage,
	}

	return pagination, nil
}
