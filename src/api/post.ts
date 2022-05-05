import { AxiosResponse } from "axios";
import { Post, R } from "../types";
import { req } from "./axios";

export const getPost = (id: string) => {
  return req.get(`/post/${id}`) as Promise<AxiosResponse<R<Post>>>;
};

export const getPostByDisplayId = (id: string) => {
  return req.get(`/post/share/${id}`) as Promise<AxiosResponse<R<Post>>>;
};

export const getPosts = (page: number, size: number) => {
  return req.get(`/posts/${page}/${size}`) as Promise<AxiosResponse<R<Post[]>>>;
};

export const getTotal = () => {
  return req.get(`/posts/count`) as Promise<AxiosResponse<R<number>>>;
};
