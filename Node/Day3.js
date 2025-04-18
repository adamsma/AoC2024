const fs = require("node:fs")

const samplePath = "../data/Sample_Day3.txt"
const samplePathP2 = "../data/Sample_Day3p2.txt"
const inputPath = "../data/Day3.txt"

const D3_MUL_REGEX = /mul\(\d+,\d+\)/g

const ParseCode = (data) => {
  return [...data.matchAll(D3_MUL_REGEX)]
}

const ParseCodeV2 = (data) => {
  // assume at least one don't statement in text
  const sides = data.split("don't()")
  let matches = [...sides.shift().matchAll(D3_MUL_REGEX)]

  if (!matches) {
    matches = []
  }

  matches.push(...ParseRecurse(sides.join("don't()"), false, 0))
  return matches
}

const ParseRecurse = (data, capture, n) => {
  const matchWord = capture ? "don't()" : "do()"
  const sides = data.split(matchWord)

  let matches
  if (capture) {
    matches = [...sides.shift().matchAll(D3_MUL_REGEX)]

    if (!matches[0]) {
      matches = []
    }
  } else {
    sides.shift()
    matches = []
  }

  if (sides.length == 0) {
    return matches
  }

  // replace matchWord as it's removed when splitting
  matches.push(...ParseRecurse(sides.join(matchWord), !capture, n + 1))
  return matches
}

const ExecMul = (mul) => {
  let nums = [...mul[0].matchAll(/\d+/g)]
  return parseInt(nums[0][0]) * parseInt(nums[1][0])
}

const Part1Calc = (instructions) => {
  let sum = 0
  instructions.forEach((x) => {
    sum += ExecMul(x)
  })
  return sum
}

const Part2Calc = (instructions) => {
  return Part1Calc(instructions)
}

const sData = fs.readFileSync(samplePath, "utf8")
const sDataP2 = fs.readFileSync(samplePathP2, "utf8")
const data = fs.readFileSync(inputPath, "utf8")

let sInstr = ParseCode(sData)
let s1 = Part1Calc(sInstr)

let sInstrP2 = ParseCodeV2(sDataP2)
let s2 = Part2Calc(sInstrP2)

let instr = ParseCode(data)
let p1 = Part1Calc(instr)

let instrP2 = ParseCodeV2(data)
let p2 = Part2Calc(instrP2)

module.exports = { s1, s2, p1, p2 }
