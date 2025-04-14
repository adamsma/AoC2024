library(readr)
library(dplyr)
library(purrr)

ParseData <- function(file){
  file |> 
    read_fwf(
      fwf_empty(file, col_names = c("col1", "col2")),
      col_types = "dd"
    )
}

testData <- ParseData("../Data/Sample_Day1.txt")

data <- ParseData("../Data/Day1.txt")

# Calculate total distances between lists
Part1Calc <- function(data) {
  
  orders <- data |> 
    mutate_all(sort) |> 
    mutate(dist = abs(col1 - col2))
  
  sum(orders$dist)
  
}

# Calculate similarity score
Part2Calc <- function(data) {
  
  sims <- map_dbl(data$col1, \(x) sum(data$col2 == x))
  
  sum(sims * data$col1)
  
}

s1 <- Part1Calc(testData)
p1 <- Part1Calc(data)


s2 <- Part2Calc(testData)
p2 <- Part2Calc(data)

