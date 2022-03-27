package cache

import "fmt"

func PostCacheKey(id string) string {
	return fmt.Sprintf("post:%s", id)
}

func PostsCacheKey(page, size string) string {
	return fmt.Sprintf("posts:%s:%s", page, size)
}

func PostsCountCacheKey() string {
	return "posts:count"
}
