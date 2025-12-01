# Go Exercism Codebase Index

**Generated:** 2025-12-01  
**Workspace:** `c:\Users\povil\Exercism\go`  
**Total Exercises:** 25  
**Total Files:** 144

---

## 📋 Overview

This workspace contains Go programming exercises from Exercism. Each exercise directory follows a standard structure with source files, tests, and documentation.

### Standard Exercise Structure
```
exercise-name/
├── exercise_name.go          # Main implementation
├── exercise_name_test.go     # Test suite
├── go.mod                    # Module definition
├── README.md                 # Exercise description
├── HELP.md                   # Getting started guide
└── HINTS.md                  # Exercise hints
```

---

## 📦 Exercises by Package

### 1. **airport-robot** (`package airportrobot`)
- **Path:** `airport-robot/`
- **Files:** `airport_robot.go`, `airport_robot_test.go`
- **Key Types:** `Greeter` (interface), `Italian`, `Portuguese`
- **Key Functions:** `SayHello()`, `LanguageName()`, `Greet()`
- **Concepts:** Interfaces, methods, polymorphism

### 2. **animal-magic** (`package chance`)
- **Path:** `animal-magic/`
- **Files:** `animal_magic.go`, `animal_magic_test.go`
- **Concepts:** Random number generation, probability

### 3. **annalyns-infiltration** (`package annalyn`)
- **Path:** `annalyns-infiltration/`
- **Files:** `annalyns_infiltration.go`, `annalyns_infiltration_test.go`
- **Concepts:** Boolean logic, conditionals

### 4. **bird-watcher** (`package birdwatcher`)
- **Path:** `bird-watcher/`
- **Files:** `bird_watcher.go`, `bird_watcher_test.go`
- **Concepts:** Slices, loops, data manipulation

### 5. **blackjack** (`package blackjack`)
- **Path:** `blackjack/`
- **Files:** `blackjack.go`, `blackjack_test.go`
- **Concepts:** Switch statements, game logic

### 6. **card-tricks** (`package cards`)
- **Path:** `card-tricks/`
- **Files:** `card_tricks.go`, `card_tricks_test.go`
- **Concepts:** Slice operations, indexing

### 7. **cars-assemble** (`package cars`)
- **Path:** `cars-assemble/`
- **Files:** `cars_assemble.go`, `cars_assemble_test.go`
- **Concepts:** Floating-point arithmetic, conditionals

### 8. **chessboard** (`package chessboard`)
- **Path:** `chessboard/`
- **Files:** `chessboard.go`, `chessboard_test.go`
- **Key Types:** `File` ([]bool), `Chessboard` (map[string]File)
- **Concepts:** Maps, custom types, nested data structures

### 9. **election-day** (`package electionday`)
- **Path:** `election-day/`
- **Files:** `election_day.go`, `election_result.go`, `election_day_test.go`
- **Key Types:** `ElectionResult` (struct)
- **Concepts:** Structs, pointers, methods

### 10. **exc2** (`package fanin`)
- **Path:** `exc2/`
- **Files:** `exercise2.go`, `exercise2_test.go`
- **Concepts:** Concurrency, fan-in pattern, channels

### 11. **exc3** (`package tasks`)
- **Path:** `exc3/`
- **Files:** `exercise3.go`, `exercise3_test.go`
- **Key Types:** `Task` (struct)
- **Concepts:** Concurrency patterns, task management

### 12. **exc4** (`package concurpatterns`) 🔄 **CURRENTLY OPEN**
- **Path:** `exc4/`
- **Files:** `exercise4.go`, `exercise4_test.go`, `README_Exercise4 (2).md`
- **Key Types:** `Source` (func type)
- **Key Functions:** `FetchAll()` (TODO: not implemented)
- **Dependencies:** `context` package
- **Concepts:** Context, cancellation, WaitGroups, error propagation
- **Status:** ⚠️ Implementation pending

### 13. **expenses** (`package expenses`)
- **Path:** `expenses/`
- **Files:** `expenses.go`, `expenses_test.go`
- **Key Types:** `Record` (struct), `DaysPeriod` (struct)
- **Concepts:** Structs, data modeling

### 14. **gross-store** (`package gross`)
- **Path:** `gross-store/`
- **Files:** `gross_store.go`, `gross_store_test.go`
- **Concepts:** Maps, unit conversions

### 15. **jedliks-toys** (`package jedlik`)
- **Path:** `jedliks-toys/`
- **Files:** `jedliks_toys.go`, `car.go`, `jedliks_toys_test.go`
- **Key Types:** `Car` (struct)
- **Concepts:** Package organization, multiple files, structs

### 16. **lasagna** (`package lasagna`)
- **Path:** `lasagna/`
- **Files:** `lasagna.go`, `lasagna_test.go`
- **Key Constants:** `OvenTime = 40`
- **Key Functions:** 
  - `RemainingOvenTime(actualMinutesInOven int) int`
  - `PreparationTime(numberOfLayers int) int`
  - `ElapsedTime(numberOfLayers, actualMinutesInOven int) int`
- **Concepts:** Constants, functions, basic arithmetic

### 17. **lasagna-master** (`package lasagna`)
- **Path:** `lasagna-master/`
- **Files:** `lasagna_master.go`, `lasagna_master_test.go`
- **Concepts:** Variadic functions, advanced patterns

### 18. **need-for-speed** (`package speed`)
- **Path:** `need-for-speed/`
- **Files:** `need_for_speed.go`, `need_for_speed_test.go`
- **Key Types:** `Car` (struct), `Track` (struct)
- **Concepts:** Structs, methods, pointers

### 19. **party-robot** (`package partyrobot`)
- **Path:** `party-robot/`
- **Files:** `party_robot.go`, `party_robot_test.go`
- **Key Functions:**
  - `Welcome(name string) string`
  - `HappyBirthday(name string, age int) string`
  - `AssignTable(name string, table int, seatmate string, direction string, distance float64) string`
- **Concepts:** String formatting, multiple parameter types

### 20. **pingpong** (`package pingpong`)
- **Path:** `pingpong/`
- **Files:** `pingpong.go`, `pingpong_test.go`
- **Key Functions:**
  - `RunPingPong(count int) []string`
  - `sendPing(wg *sync.WaitGroup, ch chan<- string, times int)`
  - `sendPong(wg *sync.WaitGroup, ch chan<- string, times int)`
- **Dependencies:** `sync` package
- **Concepts:** Concurrency, channels, WaitGroups, goroutines

### 21. **the-farm** (`package thefarm`)
- **Path:** `the-farm/`
- **Files:** `the_farm.go`, `types.go`, `the_farm_test.go`
- **Key Types:** 
  - `FodderCalculator` (interface)
  - `InvalidCowsError` (struct - custom error type)
- **Key Functions:**
  - `DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error)`
  - `ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error)`
  - `ValidateNumberOfCows(cows int) error`
- **Concepts:** Interfaces, error handling, custom errors

### 22. **two-fer** (`package twofer`)
- **Path:** `two-fer/`
- **Files:** `two_fer.go`, `two_fer_test.go`
- **Key Functions:** `ShareWith(name string) string`
- **Concepts:** String formatting, conditionals

### 23. **vehicle-purchase** (`package purchase`)
- **Path:** `vehicle-purchase/`
- **Files:** `vehicle_purchase.go`, `vehicle_purchase_test.go`
- **Key Functions:**
  - `NeedsLicense(kind string) bool`
  - `ChooseVehicle(option1, option2 string) string`
  - `CalculateResellPrice(originalPrice, age float64) float64`
- **Concepts:** Conditionals, string comparison, calculations

### 24. **weather-forecast** (`package weather`)
- **Path:** `weather-forecast/`
- **Files:** `weather_forecast.go`, `weather_forecast_test.go`
- **Key Functions:** `Forecast(city, condition string) string`
- **Concepts:** Documentation comments, exported functions

### 25. **welcome-to-tech-palace** (`package techpalace`)
- **Path:** `welcome-to-tech-palace/`
- **Files:** `welcome_to_tech_palace.go`, `welcome_to_tech_palace_test.go`
- **Key Functions:**
  - `WelcomeMessage(customer string) string`
  - `AddBorder(welcomeMsg string, numStarsPerLine int) string`
  - `CleanupMessage(oldMsg string) string`
- **Concepts:** String manipulation, formatting

---

## 🏗️ Architecture Patterns

### Concurrency Exercises
1. **exc2** (fanin pattern - channels)
2. **exc3** (task management - concurrency patterns)
3. **exc4** (context, cancellation, WaitGroups) ⚠️ In progress
4. **pingpong** (channels, goroutines, WaitGroups)

### Interface & OOP Exercises
1. **airport-robot** (interfaces, methods)
2. **the-farm** (interfaces, custom errors)

### Data Structure Exercises
1. **chessboard** (maps, custom types)
2. **bird-watcher** (slices)
3. **card-tricks** (slice operations)

### Error Handling Exercises
1. **the-farm** (custom error types, error propagation)
2. **exc4** (error handling in concurrent code)

---

## 📊 Statistics

### File Distribution
- **Go Source Files:** 28 (non-test)
- **Test Files:** 25+
- **Module Files (go.mod):** 25
- **Documentation Files:** ~91 (README.md, HELP.md, HINTS.md)

### Complexity Levels
- **Beginner:** lasagna, two-fer, vehicle-purchase, weather-forecast
- **Intermediate:** airport-robot, bird-watcher, party-robot, the-farm
- **Advanced:** exc2, exc3, exc4, pingpong (concurrency exercises)

---

## 🔍 Quick Reference

### Finding Code by Concept

**Concurrency & Channels:**
- exc2, exc3, exc4, pingpong

**Structs & Methods:**
- election-day, expenses, jedliks-toys, need-for-speed

**Interfaces:**
- airport-robot, the-farm

**Error Handling:**
- the-farm, exc4

**String Manipulation:**
- party-robot, welcome-to-tech-palace, weather-forecast

**Maps:**
- chessboard, gross-store

**Slices:**
- bird-watcher, card-tricks

---

## 🎯 Current Focus

**Active Files:**
- `exc4/exercise4.go` - Context-based concurrent fetching (TODO)
- `exc4/exercise4_test.go` - Tests for FetchAll function
- `exc4/README_Exercise4 (2).md` - Exercise documentation

**Next Steps:**
- Implement `FetchAll()` function in exc4
- Handle context cancellation properly
- Implement error propagation from multiple sources
- Use WaitGroups for coordinating goroutines

---

## 🛠️ Development Patterns

### Common Imports
```go
"context"          // exc4 - context handling
"fmt"              // Most exercises - formatting
"sync"             // pingpong, exc2-4 - concurrency
"testing"          // All *_test.go files
```

### Testing Pattern
All exercises follow standard Go testing:
```go
func TestFunctionName(t *testing.T) {
    // Test cases
}
```

### Module Pattern
Each exercise is a standalone module with `go.mod`:
```
module <exercise-name>

go 1.XX
```

---

**Index End** | Last Updated: 2025-12-01
