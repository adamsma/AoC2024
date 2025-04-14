library(dplyr)
library(purrr)


ParseReports <- function(file){
  file |> 
    readLines() |> 
    map(\(x) strsplit(x, " ", fixed = TRUE) |>  unlist() |> as.integer())
}

IsSafe <- function(rpt) {
  
  delta <- diff(rpt)
  sameSign <- all(delta > 0) || all(delta < 0)
  safeMag <- max(abs(delta)) <= 3
  
  sameSign && safeMag
  
}

IsDampenSafe <- function(rpt) {
  
  # Check if safe as is
  if(IsSafe(rpt)) return(TRUE)
  
  # check if removing any single value makes it safe
  # could improve efficiency by short circuiting first safe
  # but reports are short enough to test all
  map_lgl(seq_along(rpt), \(i) IsSafe(rpt[-i])) |> 
    any()

}

Part1Calc <- function(rpts) {
  rpts |> 
    map_lgl(IsSafe) |> 
    sum()
}

Part2Calc <- function(rpts) {
  rpts |> 
    map_lgl(IsDampenSafe) |> 
    sum()
}

sampleData <- ParseReports("../Data/Sample_Day2.txt")
data <- ParseReports("../Data/Day2.txt")

s1 <- Part1Calc(sampleData) 
p1 <- Part1Calc(data)

s2 <- Part2Calc(sampleData)
p2 <- Part2Calc(data)