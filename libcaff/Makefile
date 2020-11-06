# TODO: code coverage, if we have the time
.PHONY: all library clean test

OBJ_DIR := ./obj
BUILD_DIR := ./build
SRC_DIR := ./src
INCLUDE_DIR := ./include
TEST_DIR := ./test
OBJ_DIR_TEST := ./test-obj

SRC_FILES := $(shell find "$(SRC_DIR)" -type f -iname '*.cpp' -print)
OBJ_FILES := $(patsubst $(SRC_DIR)/%.cpp,$(OBJ_DIR)/%.o,$(SRC_FILES))

SRC_FILES_TEST := $(shell find "$(TEST_DIR)" -type f -iname '*.cpp' -print)
OBJ_FILES_TEST := $(patsubst $(TEST_DIR)/%.cpp,$(OBJ_DIR_TEST)/%.o,$(SRC_FILES_TEST))

CXX := g++
CXXFLAGS := -O2 -Wall -Wextra -pedantic -c -std=c++17 -I$(INCLUDE_DIR)

LD := g++
LDFLAGS := -lc -lstdc++
LDFLAGS_TEST := -lgtest

# FIXME: hardcoded shared library extension
LIB := $(BUILD_DIR)/lib.so
TEST_RUNNER := $(BUILD_DIR)/test-runner

all: library test

library: $(LIB)

$(LIB): $(OBJ_FILES)
	@-mkdir -p $(BUILD_DIR)
	$(LD) $(LDFLAGS) -shared -o $(LIB) $^

test: $(TEST_RUNNER)
	$(TEST_RUNNER)

$(TEST_RUNNER): $(OBJ_FILES_TEST) $(LIB)
	$(LD) $(LDFLAGS) $(LDFLAGS_TEST) -o $(TEST_RUNNER) $^

$(OBJ_DIR)/%.o:: $(SRC_DIR)/%.cpp
	@-mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) -o $@ $<

$(OBJ_DIR_TEST)/%.o:: $(TEST_DIR)/%.cpp
	@-mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) -o $@ $<

clean:
	@-rm -r $(OBJ_DIR) $(OBJ_DIR_TEST) $(BUILD_DIR)
