import { getPost, getPosts } from '../../api/post'
import { useDate } from '../../composoable/useDate'
import { md2HTML } from '../../plugins/markdown'
import { Post } from '../../types'

const commands = new Set(['ls', 'll', 'whoami', 'cat', 'clear', 'help'])

const cache = [] as Post[]

const ls = async (blocks: string[]) => {
  if (cache.length > 0) {
    const titles = cache.map(post => `${post.display_id.trim()}.md`)
    blocks.push(`${titles.join('&nbsp;&nbsp;')}`)
    return
  }
  const { data: { data: posts } } = await getPosts()
  const titles = posts.map(post => `${post.display_id.trim()}.md`)
  cache.push(...posts)
  blocks.push(`${titles.join('&nbsp;&nbsp;')}`)
}

const ll = async (blocks: string[]) => {
  if (cache.length > 0) {
    const titles = cache.map(post => `-rw-r--r--&nbsp;qianxi&nbsp;qianxi&nbsp;${useDate(post.create_time)}&nbsp;${useDate(post.update_time)}&nbsp;${post.display_id}.md`)
    blocks.push(`${titles.join('<br>')}`)
    return
  }

  const { data: { data: posts } } = await getPosts()
  const titles = posts.map(post => `-rw-r--r--&nbsp;qianxi&nbsp;qianxi&nbsp;${useDate(post.create_time)}&nbsp;${useDate(post.update_time)}&nbsp;${post.display_id}.md`)
  cache.push(...posts)
  blocks.push(`${titles.join('<br>')}`)
}

const whoami = (blocks: string[]) => {
  const info = {
    name: 'qianxi',
    GitHub: 'https://github.com/qianxi0410',
    Email: 'bGl1eWloYW4wNDEwQGdtYWlsLmNvbQo=',
  } as Record<string, string>

  const keys = Object.keys(info)
  // eslint-disable-next-line security/detect-object-injection
  const pairs = keys.map(key => `${key}: ${info[key]}`)

  blocks.push(`${pairs.join('<br>')}`)
}

const help = (blocks: string[]) => {
  const help = [
    'ls: list files',
    'll: list files with details',
    'whoami: show user info',
    'cat: show file content',
    'clear: clear screen',
    'help: show help',
  ]

  blocks.push(`${help.join('<br>')}`)
}

const cat = async (file: string, blocks: string[]) => {
  if (cache.length > 0) {
    const post = cache.find(post => post.display_id === file.trim())
    if (post) {
      const { data: { data: res } } = await getPost(post.id)
      blocks.push(md2HTML(res.content!))
      return
    }
    blocks.push(`${file}.md file not found`)
    return
  }

  const { data: { data: posts } } = await getPosts()
  cache.push(...posts)
  const post = posts.find(post => post.display_id === file.trim())
  if (post) {
    const { data: { data: res } } = await getPost(post.id)
    blocks.push(md2HTML(res.content!))
    return
  }
  blocks.push(`${file}.md file not found`)
}

export const handle = async (cmd: string, blocks: string[], history: string[]) => {
  const command = cmd.trim().replace(/\s+/g, ' ')

  if (command.startsWith('cat')) {
    history.push(command)
    let file = command.trim().slice(3).trim()

    if (file.endsWith('.md'))
      file = file.slice(0, -3)

    await cat(file, blocks)
    return
  }

  if (!commands.has(command)) {
    history.push(command)
    blocks.push(`${command}: command not found`)
    return
  }

  switch (command) {
    case 'ls':
      history.push(command)
      await ls(blocks)
      break
    case 'll':
      history.push(command)
      await ll(blocks)
      break
    case 'whoami':
      history.push(command)
      whoami(blocks)
      break
    case 'help':
      history.push(command)
      help(blocks)
      break
    case 'clear':
      blocks.length = 0
      history.length = 0
  }
}
