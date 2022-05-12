package cache

import "fmt"

func PostCacheKey(id string) string {
	return fmt.Sprintf("post:%s", id)
}

func PostsCacheKey() string {
	return "posts:all"
}

func PostsCountCacheKey() string {
	return "posts:count"
}
