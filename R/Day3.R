library(dplyr)
library(purrr)
library(stringr)

ParseCode <- function(file){
  file |> 
    readLines() |> 
    str_extract_all("mul\\(\\d+,\\d+\\)") |> 
    unlist()
}

ParseCodeV2 <- function(file){
  
  txt <- file |> 
    readLines() |> 
    paste0(collapse = "")
  
  # determine when start of the mul regex falls between start of 
  # do regex and start of dont regex
  mult <- str_extract_all(txt, "mul\\(\\d+,\\d+\\)") |> unlist()
  loc <- str_locate_all(txt, "mul\\(\\d+,\\d+\\)")
  doLoc <- str_locate_all(txt, "do\\(\\)")
  dontLoc <- str_locate_all(txt, "don't\\(\\)")
  
  target <- loc[[1]][,1]
  start <- data.frame(
    # implicit do at start
    i = c(0, doLoc[[1]][,1]), include = TRUE
  ) 
  stop <- data.frame(
    # implicit dont at end
    i = c(dontLoc[[1]][,1], str_length(txt) + 1), include = FALSE
  ) 
  condLoc <- bind_rows(start, stop) |> 
    arrange(i)

  # find the first conditional that follows a mul instruction
  # look at previous conditional to determine if to include instruction
  keep <- map_lgl(target, \(x) condLoc$include[which.max(condLoc$i > x) - 1])
  
  mult[keep]
  
}

Part1Calc <- function(instructions) {
  instructions |> 
    str_extract_all("\\d+") |> 
    map_int(\(x) as.integer(x[1]) * as.integer(x[2])) |> 
    sum()
}

Part2Calc <- function(instructions) {
  Part1Calc(instructions)
}

sampleData <- ParseCode("../Data/Sample_Day3.txt") 
sampleDataP2 <- ParseCodeV2("../Data/Sample_Day3p2.txt") 
data <- ParseCode("../Data/Day3.txt")
dataP2 <- ParseCodeV2("../Data/Day3.txt")
  
s1 <- Part1Calc(sampleData) 
p1 <- Part1Calc(data)

s2 <- Part2Calc(sampleDataP2)
p2 <- Part2Calc(dataP2)