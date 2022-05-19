const dark = `
pre code.hljs {
  display: block;
  overflow-x: auto;
  background-color: rgb(32, 34, 32);
  border-radius: 0.5em;
}

.hljs {
  color: #a8acb3;
  background: #44474d;
  padding-left: 1em;
  padding-top: 1em;
  padding-bottom: 1em;
  padding-right: 1em;
}
.hljs-comment,
.hljs-quote {
  color: #5c6370;
  font-style: italic;
}
.hljs-doctag,
.hljs-formula,
.hljs-keyword {
  color: #38bb89;
}
.hljs-deletion,
.hljs-name,
.hljs-section,
.hljs-selector-tag,
.hljs-subst {
  color: #179254;
}
.hljs-literal {
  color: #289e67;
}
.hljs-addition,
.hljs-attribute,
.hljs-meta .hljs-string,
.hljs-regexp,
.hljs-string {
  color: #98c379;
}
.hljs-attr,
.hljs-number,
.hljs-selector-attr,
.hljs-selector-class,
.hljs-selector-pseudo,
.hljs-template-variable,
.hljs-type,
.hljs-variable {
  color: #64a11e;
}
.hljs-bullet,
.hljs-link,
.hljs-meta,
.hljs-selector-id,
.hljs-symbol,
.hljs-title {
  color: #0b8f1f;
}
.hljs-built_in,
.hljs-class .hljs-title,
.hljs-title.class_ {
  color: #0a753c;
}
.hljs-emphasis {
  font-style: italic;
}
.hljs-strong {
  font-weight: 700;
}
.hljs-link {
  text-decoration: underline;
}
`

const light = `
pre code.hljs {
  display: block;
  overflow-x: auto;
  border-radius: 0.5em;
  overflow: auto;
}

.hljs {
  background: rgb(251, 248, 248);
  overflow-block: none;
  color: gray;
  padding-left: 1em;
  padding-top: 1em;
  padding-bottom: 1em;
  padding-right: 1em;
}
.xml .hljs-meta {
  color: white;
}
.hljs-comment,
.hljs-quote {
  color: #0d950d;
}
.hljs-attribute,
.hljs-keyword,
.hljs-literal,
.hljs-name,
.hljs-selector-tag,
.hljs-tag {
  color: #4eb84e;
}
.hljs-template-variable,
.hljs-variable {
  color: #a7710f;
}
.hljs-code,
.hljs-meta .hljs-string,
.hljs-string {
  color: #a0730c6c;
}
.hljs-link,
.hljs-regexp {
  color: #5fe09b;
}
.hljs-bullet,
.hljs-number,
.hljs-symbol,
.hljs-title {
  color: #2d9750;
}
.hljs-meta,
.hljs-section {
  color: #1a794c;
}
.hljs-built_in,
.hljs-class .hljs-title,
.hljs-params,
.hljs-title.class_,
.hljs-type {
  color: #95500f;
}
.hljs-attr {
  color: #836c28;
}
.hljs-subst {
  color: #000;
}
.hljs-formula {
  background-color: rgb(255, 255, 255);
  font-style: italic;
}
.hljs-addition {
  background-color: #baeeba;
}
.hljs-deletion {
  background-color: #ffc8bd;
}
.hljs-selector-class,
.hljs-selector-id {
  color: #9b703f;
}
.hljs-doctag,
.hljs-strong {
  font-weight: 700;
}
.hljs-emphasis {
  font-style: italic;
}
`

export const useTheme = (theme: 'dark'|'light') => {
  let style = document.querySelector('#code-style') as HTMLStyleElement
  if (style) {
    style.innerHTML = theme === 'dark' ? dark : light
    return
  }

  style = document.createElement('style')
  style.type = 'text/css'
  style.id = 'code-style'
  style.innerHTML = theme === 'dark' ? dark : light
  const head = document.querySelectorAll('head')[0]
  head.append(style)
}
