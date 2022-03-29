package handlers

import (
	"api/cache"
	"api/dao"
	"api/r"
	"api/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
)

const defaultDuration = time.Hour

type PostHandler struct{}

func (p *PostHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, r.Error[string]("post id is required"))
		}

		cli := cache.New()
		defer cli.Close()

		if cli == nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string]("redis client is nil"))
		}

		var post dao.Post
		bytes, err := cli.Get(context.Background(), cache.PostCacheKey(id)).Bytes()

		if err != nil && err != redis.Nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
		} else if err != nil && err == redis.Nil {
			// find post from db
			post, err = dao.GetPost(id)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
			}
			// store post to cache
			bytes, _ := msgpack.Marshal(post)
			cli.Set(context.Background(), cache.PostCacheKey(id), bytes, defaultDuration)
		} else {
			msgpack.Unmarshal(bytes, &post)
		}

		ctx.JSON(http.StatusOK, r.Ok(post))
	}
}

func (p *PostHandler) PostWithDisplayId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("display_id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, r.Error[string]("post display_id is required"))
		}

		post, err := dao.GetPostWithDisplayId(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
		}

		ctx.JSON(http.StatusOK, r.Ok(post))
	}
}

func (p *PostHandler) Posts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, size := ctx.Param("page"), ctx.Param("size")
		if page == "" || size == "" {
			ctx.JSON(http.StatusBadRequest, r.Error[string]("page and size are required"))
		}

		cli := cache.New()
		defer cli.Close()

		if cli == nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string]("redis client is nil"))
		}

		var posts []dao.Post

		// get posts from cache
		bytes, err := cli.Get(context.Background(), cache.PostsCacheKey(page, size)).Bytes()
		if err != nil && err != redis.Nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
		} else if err != nil && err == redis.Nil {
			// convert page and size to int
			pageInt, sizeInt, err := utils.ConverPageAndSize(page, size)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, r.Error[string]("page and size must be int"))
			}

			offset := (pageInt - 1) * sizeInt

			// get posts from db
			posts, err = dao.GetPosts(offset, sizeInt)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
			}
			// store posts to cache
			bytes, _ := msgpack.Marshal(posts)
			cli.Set(context.Background(), cache.PostsCacheKey(page, size), bytes, defaultDuration)
		} else {
			msgpack.Unmarshal(bytes, &posts)
		}

		ctx.JSON(http.StatusOK, r.Ok(posts))
	}
}

func (p *PostHandler) Count() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cli := cache.New()
		defer cli.Close()

		if cli == nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string]("redis client is nil"))
		}

		count, err := cli.Get(context.Background(), cache.PostsCountCacheKey()).Int64()
		if err != nil && err != redis.Nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
		} else if err != nil && err == redis.Nil {
			// get count from db
			count, err = dao.GetPostsCount()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
			}

			// store count to cache
			cli.Set(context.Background(), cache.PostsCountCacheKey(), count, defaultDuration)
		}

		ctx.JSON(http.StatusOK, r.Ok(count))
	}
}

func (p *PostHandler) DeletePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, r.Error[string]("post id is required"))
		}

		err := dao.DeletePost(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
		}

		// invalid cache
		cli := cache.New()
		if cli == nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string]("redis client is nil"))
		}
		defer cli.Close()

		// list cache
		keys, _, err := cli.Scan(context.Background(), 0, "posts:*", 0).Result()
		cli.Del(context.Background(), keys...)

		// count cache
		// post cache
		cli.Del(context.Background(), cache.PostsCountCacheKey(), cache.PostCacheKey(id))

		ctx.JSON(http.StatusOK, r.Ok("dekete success"))
	}
}

func (p *PostHandler) UpdatePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var post dao.Post
		err := ctx.ShouldBind(&post)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string]("post bind failed"))
		}

		err = dao.UpdatePost(post)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
		}

		// invalid cache
		cli := cache.New()
		if cli == nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string]("redis client is nil"))
		}

		// post cache
		cli.Del(context.Background(), cache.PostCacheKey(post.ID))
		// list cache
		keys, _, err := cli.Scan(context.Background(), 0, cache.PostsCacheKey("*", "*"), 0).Result()
		cli.Del(context.Background(), keys...)

		ctx.JSON(http.StatusOK, r.Ok("update success"))
	}
}

func (p *PostHandler) CreatePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var post dao.Post
		err := ctx.ShouldBind(&post)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string]("post bind failed"))
		}

		err = dao.CreatePost(post)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string](err.Error()))
		}

		// invalid cache
		cli := cache.New()
		if cli == nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string]("redis client is nil"))
		}

		// list cache
		keys, _, err := cli.Scan(context.Background(), 0, cache.PostsCacheKey("*", "*"), 0).Result()
		cli.Del(context.Background(), keys...)

		// count cache
		cli.Del(context.Background(), cache.PostsCountCacheKey())

		ctx.JSON(http.StatusOK, r.Ok("create success"))
	}
}
