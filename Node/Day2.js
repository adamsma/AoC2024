const fs = require("node:fs")

const samplePath = "../data/Sample_Day2.txt"
const inputPath = "../data/Day2.txt"

const ParseReports = (data) => {
  var reports = []
  const rows = data.trim().split("\r\n")
  rows.forEach((row) => {
    reports.push(row.split(" ").map((x) => parseInt(x)))
  })

  return reports
}

const IsSafe = (report) => {
  var delta = []
  for (i = 1; i < report.length; i++) {
    delta.push(report[i] - report[i - 1])
  }
  // determine if slope is all positive or negative
  if (delta[0] === 0) {
    return false
  }

  var lastSign = Math.sign(delta[0])
  for (i = 1; i < delta.length; i++) {
    if (delta[i] === 0) {
      return false
    }

    if (Math.sign(delta[i]) != lastSign) {
      return false
    }
  }

  // greatest change must be have abs less than 3
  return Math.max(...delta) <= 3 && Math.min(...delta) >= -3
}

const IsDampenSafe = (report) => {
  if (IsSafe(report)) {
    return true
  }

  // drop elements one by one to see if any make it safe
  for (var i = 0; i < report.length; i++) {
    if (IsSafe(report.toSpliced(i, 1))) {
      return true
    }
  }

  return false
}

const Part1Calc = (reports) => {
  var sum = 0
  reports.forEach((r) => (sum += IsSafe(r) ? 1 : 0))
  return sum
}

const Part2Calc = (reports) => {
  var sum = 0
  reports.forEach((r) => (sum += IsDampenSafe(r) ? 1 : 0))
  return sum
}

const sData = fs.readFileSync(samplePath, "utf8")
const data = fs.readFileSync(inputPath, "utf8")

var reports = ParseReports(sData)
var s1 = Part1Calc(reports)
var s2 = Part2Calc(reports)

var reports = ParseReports(data)
var p1 = Part1Calc(reports)
var p2 = Part2Calc(reports)

module.exports = { s1, s2, p1, p2 }
