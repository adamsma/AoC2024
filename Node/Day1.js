const fs = require("node:fs")

const samplePath = "../data/Sample_Day1.txt"
const inputPath = "../data/Day1.txt"

const CreateLists = (data) => {
  const l1 = []
  const l2 = []

  const rows = data.trim().split("\r\n")
  rows.forEach((row) => {
    let [x, y] = row.split("   ").slice(0, 2)
    l1.push(parseInt(x))
    l2.push(parseInt(y))
  })

  return [l1, l2]
}

// Calculate total distances between list
const Part1Calc = (l1, l2) => {
  let dist = 0

  l1.forEach((x, i) => {
    dist += Math.abs(x - l2[i])
  })

  return dist
}

// Calculate similarity score
const Part2Calc = (l1, l2) => {
  const counts = new Map()
  let sim = 0

  l1.forEach((x) => counts.set(x, 0))
  l2.forEach((x) => {
    if (counts.has(x)) counts.set(x, counts.get(x) + 1)
  })

  l1.forEach((val, i) => (sim += val * counts.get(val)))

  return sim
}

const sData = fs.readFileSync(samplePath, "utf8")
const data = fs.readFileSync(inputPath, "utf8")

let [sl1, sl2] = CreateLists(sData)
sl1.sort()
sl2.sort()
let s1 = Part1Calc(sl1, sl2)
let s2 = Part2Calc(sl1, sl2)

let [l1, l2] = CreateLists(data)
l1.sort()
l2.sort()
let p1 = Part1Calc(l1, l2)
let p2 = Part2Calc(l1, l2)

module.exports = { s1, s2, p1, p2 }
