
export const useDate = (time: number) => {
  const date = new Date(Number(time) * 1000)
  const array = date.toDateString().slice(4).split(' ')

  let res = `${array[0]} ${array[1]},${array[2]}`
  if (date.getFullYear() === new Date().getFullYear()) res = res.slice(0, Math.max(0, res.length - 5))
  return res
}
