#include <gtest/gtest.h>
#include <libcaff/caff.h>
#include <libcaff/util.h>

#include <fstream>
#include <iostream>
#include <istream>
#include <map>
#include <stdexcept>

namespace {
std::fstream open_fixture(std::string fixture) {
  auto fixtures_path = "../../fixtures/";
  auto full_path = fixtures_path + fixture;
  std::fstream file;
  file.open(full_path, std::fstream::in);
  if (!file.is_open()) {
    throw std::runtime_error("can't find file " + full_path);
  }
  return file;
}
}  // namespace

namespace {

TEST(UnitTestSet, OpenFixture) {
  // TODO: close with RAAI? The test runner is going to terminate soon anyway.
  auto infile = open_fixture("inputs/1.caff");
  ASSERT_NO_THROW(parse_caff(infile));
}

TEST(UnitTestSet, PrettyPrint) {
  std::string expected =
      "-----------------------\n"
      "ID    : 1\n"
      "\tMagic bytes: CAFF\n"
      "\tHeader size: 20\n"
      "\tNumber of frames: 2\n"
      "-----------------------\n"
      "ID    : 2\n"
      "\tAuthor: Test Creator\n"
      "\tAuthored date: 2020-7-2 14:50\n"
      "-----------------------\n"
      "ID    : 3\n"
      "\tDuration: 1000\n"
      "\tWidth: 1000\n"
      "\tHeight: 667\n"
      "-----------------------\n"
      "ID    : 3\n"
      "\tDuration: 1000\n"
      "\tWidth: 1000\n"
      "\tHeight: 667\n"
      "";

  auto infile = open_fixture("inputs/1.caff");
  auto caff = parse_caff(infile);
  testing::internal::CaptureStderr();
  pretty_print(caff);
  std::string output = testing::internal::GetCapturedStderr();
  ASSERT_EQ(expected, output);
}

TEST(UnitTestSet, CheckSizeParse) {
  std::string expWidth = "1000";
  std::string expHeight = "667";

  auto infile = open_fixture("inputs/1.caff");
  auto caff = parse_caff(infile);

  testing::internal::CaptureStdout();
  int ret = print_preview(caff);
  std::string output = testing::internal::GetCapturedStdout();

  ASSERT_EQ(0, ret);
  auto width = output.substr(0, output.find(" "));
  ASSERT_EQ(expWidth, width);
  auto height = output.substr(output.find(" ") + 1,
                              output.find("\n") - output.find(" ") - 1);
  ASSERT_EQ(expHeight, height);
}

TEST(UnitTestSet, CheckParsedData) {
  auto infile = open_fixture("inputs/1.caff");
  auto caff = parse_caff(infile);

  auto expfile = open_fixture("outputs/1.out.raw");
  std::string expContents((std::istreambuf_iterator<char>(expfile)),
                          (std::istreambuf_iterator<char>()));

  testing::internal::CaptureStdout();
  int ret = print_preview(caff);
  std::string output = testing::internal::GetCapturedStdout();

  std::string contents = output;

  ASSERT_EQ(expContents, contents);
}
}  // namespace
