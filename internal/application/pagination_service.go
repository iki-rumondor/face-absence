package application

import (
	"fmt"
	"math"

	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

func GeneratePages(urlPath string, page *domain.Pagination) *domain.Pagination {
	totalPages := int(math.Ceil(float64(page.TotalRows)/float64(page.Limit))) - 1
	page.TotalPages = totalPages

	var fromRow, toRow int

	if page.Page == 0 {
		fromRow = 1
		toRow = page.Limit
	} else {
		if page.Page <= totalPages {
			fromRow = page.Page*page.Limit + 1
			toRow = page.Page + 1*page.Limit
		}
	}

	if toRow > page.TotalRows {
		toRow = page.TotalRows
	}

	page.FromRow = fromRow
	page.ToRow = toRow
	
	page.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, page.Limit, 0)
	page.LastPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, page.Limit, page.TotalPages)

	if page.Page > 0 {
		page.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, page.Limit, page.Page-1)
	}

	if page.Page < page.TotalPages {
		page.NextPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, page.Limit, page.Page+1)
	}

	if page.Page > page.TotalPages {
		page.PreviousPage = ""
	}


	return page
}
