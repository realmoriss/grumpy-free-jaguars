#include <fstream>
#include <gtest/gtest.h>
#include <iostream>
#include <istream>
#include <map>
#include <stdexcept>

#include <libcaff/caff.h>
#include <libcaff/util.h>

namespace {
    std::fstream open_fixture(std::string fixture) {
        auto fixtures_path = "../../fixtures/inputs/";
        auto full_path = fixtures_path + fixture;
        std::fstream file;
        file.open(full_path, std::fstream::in);
        if(!file.is_open()) {
            throw std::runtime_error("can't find file " + full_path);
        }
        return file;
    }
}

namespace {
    TEST(DummySet,A) {
        ASSERT_EQ(42, 42);
    }

    TEST(DummySet,OpenFixture) {
        // TODO: close with RAAI? The test runner is going to terminate soon anyway.
        auto infile = open_fixture("1.caff");
        auto caff = parse_caff(infile);
        //pretty_print(caff);
    }
}
