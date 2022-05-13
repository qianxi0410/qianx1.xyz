import { AxiosResponse } from 'axios'
import { Post, R } from '../types'
import { req as request } from './axios'

export const getPost = (id: string) => {
  return request.get(`/post/${id}`) as Promise<AxiosResponse<R<Post>>>
}

export const getPostByDisplayId = (id: string) => {
  return request.get(`/post/share/${id}`) as Promise<AxiosResponse<R<Post>>>
}

export const getPosts = () => {
  return request.get('/posts/all') as Promise<AxiosResponse<R<Post[]>>>
}

export const getTotal = () => {
  return request.get('/posts/count') as Promise<AxiosResponse<R<number>>>
}
