package utils

import "strconv"

func ConverPageAndSize(page, size string) (int, int, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, err
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return 0, 0, err
	}

	return pageInt, sizeInt, nil
}
