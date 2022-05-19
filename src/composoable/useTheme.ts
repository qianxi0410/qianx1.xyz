
export const useTheme = (theme: 'dark'|'light') => {
  let link = document.querySelector('#code-theme') as HTMLLinkElement
  if (link) {
    link.href = `/src/styles/${theme}.css`
    return
  }

  link = document.createElement('link')
  link.rel = 'stylesheet'
  link.id = 'code-theme'
  link.href = `/src/styles/${theme}.css`
  const head = document.querySelectorAll('head')[0]
  // if already exists, remove it
  head.append(link)
}
