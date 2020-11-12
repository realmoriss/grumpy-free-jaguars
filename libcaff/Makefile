# TODO: code coverage, if we have the time
.PHONY: all library clean test

OBJ_DIR := ./obj
BUILD_DIR := ./build
SRC_DIR := ./src
INCLUDE_DIR := ./include
TEST_DIR := ./test
EXECUTABLE_DIR := ./main
OBJ_DIR_TEST := ./test-obj
OBJ_DIR_EXECUTABLE := ./main-obj
KSY_DIR := ./kaitai-specs

SRC_FILES := $(shell find "$(SRC_DIR)" -type f -iname '*.cpp' -print)
OBJ_FILES := $(patsubst $(SRC_DIR)/%.cpp,$(OBJ_DIR)/%.o,$(SRC_FILES))

SRC_FILES_TEST := $(shell find "$(TEST_DIR)" -type f -iname '*.cpp' -print)
OBJ_FILES_TEST := $(patsubst $(TEST_DIR)/%.cpp,$(OBJ_DIR_TEST)/%.o,$(SRC_FILES_TEST))

SRC_FILES_EXECUTABLE := $(shell find "$(EXECUTABLE_DIR)" -type f -iname '*.cpp' -print)
OBJ_FILES_EXECUTABLE := $(patsubst $(EXECUTABLE_DIR)/%.cpp,$(OBJ_DIR_EXECUTABLE)/%.o,$(SRC_FILES_EXECUTABLE))

KSY_FILES := $(shell find "$(KSY_DIR)" -type f -iname '*.ksy' -print)

CXX := g++
CXXFLAGS := \
	-O2 \
	-Wall -Wextra -pedantic \
	-c \
	-std=c++17 \
	-I$(INCLUDE_DIR) -I$(INCLUDE_DIR)/libcaff \
	-fPIC

LD := g++
LDFLAGS := -lc -lstdc++ -L$(BUILD_DIR) -lkaitai_struct_cpp_stl_runtime
LDFLAGS_TEST := -lgtest -Wl,-rpath,. -L$(BUILD_DIR) -lcaff
LDFLAGS_EXECUTABLE := -Wl,-rpath,. -L$(BUILD_DIR) -lcaff


# FIXME: hardcoded shared library extension
LIB := libcaff.so
TEST_RUNNER := $(BUILD_DIR)/test-runner
EXECUTABLE := $(BUILD_DIR)/caff

all: library test executable test-executable

library: $(BUILD_DIR)/$(LIB)

$(BUILD_DIR)/libkaitai_struct_cpp_stl_runtime.so:
	DOCKER_BUILDKIT=1 docker build -o $(BUILD_DIR) -f kaitai-runtime.docker .

$(BUILD_DIR)/$(LIB): $(OBJ_FILES) $(BUILD_DIR)/libkaitai_struct_cpp_stl_runtime.so
	@-mkdir -p $(BUILD_DIR)
	$(LD) $(LDFLAGS) -shared -o $@ $^

test: $(TEST_RUNNER)
	cd $(shell dirname $(TEST_RUNNER)) && ./$(shell basename $(TEST_RUNNER))

$(TEST_RUNNER): $(OBJ_FILES_TEST) $(BUILD_DIR)/$(LIB)
	$(LD) $(LDFLAGS) $(LDFLAGS_TEST) -o $(TEST_RUNNER) $(OBJ_FILES_TEST)

test-executable: $(EXECUTABLE)
	cd $(shell dirname $(EXECUTABLE)) && ./$(shell basename $(EXECUTABLE)) < "../../fixtures/inputs/1.caff" > /dev/null
.PHONY: test-executable

executable: $(EXECUTABLE)
.PHONY: executable

$(EXECUTABLE): $(OBJ_FILES_EXECUTABLE) $(BUILD_DIR)/$(LIB)
	$(LD) $(LDFLAGS) $(LDFLAGS_EXECUTABLE) -o $(EXECUTABLE) $(OBJ_FILES_EXECUTABLE)

$(SRC_DIR)/caff/caff.cpp $(SRC_DIR)/caff/ciff.cpp: $(KSY_FILES)
	DOCKER_BUILDKIT=1 docker build -o . -f kaitai-gen.docker .

$(OBJ_DIR)/%.o:: $(SRC_DIR)/%.cpp $(SRC_DIR)/caff/caff.cpp $(SRC_DIR)/caff/ciff.cpp
	@-mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) -o $@ $<

$(OBJ_DIR_TEST)/%.o:: $(TEST_DIR)/%.cpp
	@-mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) -o $@ $<

$(OBJ_DIR_EXECUTABLE)/%.o:: $(EXECUTABLE_DIR)/%.cpp
	@-mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) -o $@ $<

clean:
	@-rm -r $(OBJ_DIR) $(OBJ_DIR_TEST) $(OBJ_DIR_EXECUTABLE) $(BUILD_DIR)
