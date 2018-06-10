package page

import "math"

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, rows int, total int64) map[string]interface{} {

	var firstPage int //前一页地址
	var lastPage int  //后一页地址

	if page <= 0 {
		page = 1
	}

	if rows <= 0 {
		rows = 30
	}

	offset := (page - 1) * rows

	//根据total总数，和rows每页数量 生成分页总数
	totalPages := int(math.Ceil(float64(total) / float64(rows))) //page总数
	if totalPages == 0 {
		totalPages = 1
	}

	if page > totalPages {
		page = totalPages
	}

	var pages []int

	switch {
	case page >= totalPages-5 && totalPages > 5: //最后5页
		start := totalPages - 5 + 1
		firstPage = page - 1
		lastPage = int(math.Min(float64(totalPages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalPages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstPage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstPage = page - 1
		lastPage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalPages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstPage = int(math.Max(float64(1), float64(page-1)))
		lastPage = page + 1
		//fmt.Println(pages)
	}

	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalPages"] = totalPages
	paginatorMap["firstPage"] = firstPage
	paginatorMap["lastPage"] = lastPage
	paginatorMap["page"] = page
	paginatorMap["offset"] = offset
	paginatorMap["rows"] = rows
	paginatorMap["total"] = total

	return paginatorMap
}
