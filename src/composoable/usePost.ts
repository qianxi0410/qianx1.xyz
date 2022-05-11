import { getPost, getPostByDisplayId } from '../api/post'
import { useToc } from './useToc'

export const usePost = async (id: string) => {
  // clear tocs
  const tocs = useToc()
  tocs.length = 0

  // query by display id
  if (/^[A-Za-z\-]+$/.test(id)) {
    const {
      data: { data: post },
    } = await getPostByDisplayId(id)

    return post
  }

  const {
    data: { data: post },
  } = await getPost(id)

  return post
}
