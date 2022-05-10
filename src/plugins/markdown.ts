import hljs from 'highlight.js'
import { marked } from 'marked'

const renderer: Partial<marked.Renderer> = {
  heading: (text: string, level: number) => {
    return `<h${level} id="${text}"  class="text-h${level + 1}"><span class="title">${text}</span></h${level}>`
  },
  paragraph: (text: string) => {
    return `<p class="text-h6">${text}</p>`
  },
  link: (href: string, title: string, text: string) => {
    return `<a href="${href}" target="_blank" class="link">${text}</a>`
  },
  strong: (text: string) => {
    return `<strong class="strong">${text}</strong>`
  },
  del: (text: string) => {
    return `<del title="你知道的太多啦" class="del">${text}</del>`
  },
  blockquote: (text: string) => {
    return `<blockquote class="blockquote">${text}</blockquote>`
  },
  codespan: (code: string) => {
    return `<code class="codespan">${code}</code>`
  },
  image: (href: string, title: string, text: string) => {
    return `<img src="${href}" alt="${text}" class="image shadow-17" />`
  },
}

marked.use({ renderer })

marked.setOptions({
  breaks: true,
  gfm: true,
  pedantic: false,
  sanitize: false,
  smartLists: true,
  smartypants: false,
  highlight: (code, lang) => {
    const language = hljs.getLanguage(lang) ? lang : 'plaintext'

    return hljs.highlight(language, code).value
  },
  langPrefix: 'hljs language-',
  headerIds: true,
})

export const md2HTML = (text: string) => {
  return marked.parse(text)
}
