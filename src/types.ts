export interface Post {
  id: string
  title: string
  tags?: string
  create_time: number
  update_time: number
  content?: string
  prev: string
  next: string
  display_id: string
  issue_number: number
}

export interface R<T> {
  data: T
  err: string
}

export interface Toc {
  id: string
  level: number
}
