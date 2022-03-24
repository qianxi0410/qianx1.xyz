export interface Post {
  id: string;
  title: string;
  tags?: string;
  create_date: number;
  update_date?: number;
  content?: string;
  pre: string;
  next: string;
  display_id: string;
}
