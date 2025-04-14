# clear potential conflicting variables
typical <- c("s1", "p1", "s2", "p2", "sampleData", "data", "Part1Calc", "Part2Calc")
rm(list = intersect(ls(), typical))

CURRENT_DAY = 2

PrintSolutions <- function(s1 = NULL, p1 = NULL, s2 = NULL, p2 = NULL) {
  glue::glue("Part 1 Sample Answer: {s1}") |>  print()
  glue::glue("Part 1 Answer: {p1}") |>  print()
  glue::glue("Part 2 Sample Answer: {s2}") |>  print()
  glue::glue("Part 2 Answer: {p2}") |>  print()
}

 # Will create at least s1, p1, s2, p2 variables in global environment
source(paste0("Day", CURRENT_DAY, ".R"))

PrintSolutions(s1, p1, s2, p2)
