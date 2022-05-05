import { getPost, getPostByDisplayId } from "../api/post";

export const usePost = async (id: string) => {
  // query by display id
  if (id.match(/^[a-zA-Z\-]+$/)) {
    const {
      data: { data: post },
    } = await getPostByDisplayId(id);

    return post;
  }

  const {
    data: { data: post },
  } = await getPost(id);

  return post;
};
