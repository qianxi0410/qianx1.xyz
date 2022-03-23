export interface Blog {
  title: string;
  tags?: string[];
  create_date: number;
  update_date?: number;
  content?: string;
}
