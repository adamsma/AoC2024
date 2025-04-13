const CURRENT_DAY = 1
var { s1, s2, p1, p2 } = require(`./Day${CURRENT_DAY}.js`)

const PrintSolutions = (sample1, part1, sample2, part2) => {
  console.log(`Part 1 Sample Answer: ${sample1}`)
  console.log(`Part 1 Answer: ${part1}`)
  console.log(`Part 2 Sample Answer: ${sample2}`)
  console.log(`Part 2 Answer: ${part2}`)
}

PrintSolutions(s1, p1, s2, p2)
