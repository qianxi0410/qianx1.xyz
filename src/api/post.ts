import { AxiosResponse } from "axios";
import { Post } from "../types";
import { req } from "./axios";

export const getPost = (id: string) => {
  return req.get(`/blog/post/${id}`) as Promise<AxiosResponse<Post>>;
};

export const getPosts = (page: number, size: number) => {
  return req.get(`/blog/post/${page}/${size}`) as Promise<AxiosResponse<Post[]>>;
};

export const getTotal = () => {
  return req.get(`/blog/posts/count`) as Promise<AxiosResponse<number>>;
};
