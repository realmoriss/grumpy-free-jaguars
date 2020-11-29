#ifndef INCLUDE_CAFF_UTIL_H_
#define INCLUDE_CAFF_UTIL_H_

#include <libcaff/caff.h>

caff_t parse_caff(std::istream& input);
void pretty_print(caff_t& parsed);
int print_preview(caff_t& parsed);

#endif
