# This is a Makefile for learning makefile knowledge !!!
# Knowledge point:
# 1. What's the usage of ".PHONY" ?
# 2. What's the diffirences between "=", ":=", "?=", and "+=" ?

TEST_VARIABLE  = "123"
TEST_VARIABLE ?= "456"
TEST_VARIABLE += "789"

all: some_targets
# Simulate to create the lastest target file.
	@touch target_final_file && echo create target_final_file ...
	@echo Make all targets done !!!
	@echo ""
#	@echo "TEST_VARIABLE=$(TEST_VARIABLE)"
	@echo ""

TEST_VARIABLE := "abc"

some_targets:
# Simulate to create some target files.
	@touch target_file1 && echo create target_file1 ...
	@touch target_file2 && echo create target_file2 ...
	@touch target_file3 && echo create target_file3 ...
	@echo ""
#	@echo "TEST_VARIABLE=$(TEST_VARIABLE)"
	@echo ""

clean:
	@echo Clean builded project.
	@rm -rf target_file1 target_file2 target_file3 target_final_file

.PHONY: clean some_targets