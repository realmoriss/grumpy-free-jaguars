#include <libcaff/util.h>
#include <main/main.h>

#include <iostream>

int main(int argc, char *argv[]) {
  kaitai::kstream ks(&std::cin);

  try {
    auto parsed = caff_t(&ks);
    return print_preview(parsed);
  } catch (...) {
    return 1;
  }
}
