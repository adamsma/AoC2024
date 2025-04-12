library(readr)
library(dplyr)
library(purrr)
library(glue)

ParseData <- function(file){
  file |> 
    read_fwf(
      fwf_empty(file, col_names = c("col1", "col2")),
      col_types = "dd"
    )
}

testData <- ParseData("../Data/Sample_Day1.txt")

data <- ParseData("../Data/Day1.txt")

CalcTotalDist <- function(data) {
  
  orders <- data |> 
    mutate_all(sort) |> 
    mutate(dist = abs(col1 - col2))
  
  sum(orders$dist)
  
}

s1 <- CalcTotalDist(testData)
glue("Part 1 Sample Answer: {s1}") |>  print()

a1 <- CalcTotalDist(data)
glue("Part 1 Answer: {a1}") |>  print()

CalcSimilarity <- function(data) {
  
  sims <- map_dbl(data$col1, \(x) sum(data$col2 == x))
  
  sum(sims * data$col1)
  
}

s2 <- CalcSimilarity(testData)
glue("Part 2 Sample Answer: {s2}") |>  print()

a2 <- CalcSimilarity(data)
glue("Part 2 Answer: {a2}") |>  print()
