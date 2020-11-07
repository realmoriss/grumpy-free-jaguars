#include <libcaff/caff.h>
#include <iostream>

caff_t parse_caff(std::istream &input) {
    kaitai::kstream ks(&input);
    return caff_t(&ks);
}

void pretty_print(caff_t &parsed) {
    auto blocks = parsed.block();
    if(blocks) {
        for(auto block: *blocks) {
            auto id = block->id();
            std::cerr << "-----------------------" << std::endl;
            std::cerr << "ID    : " << id << std::endl;

            switch(id) {
                case caff_t::caff_block_t::CAFF_BLOCK_ID_HEADER: {
                    auto header = static_cast<caff_t::caff_header_t*>(block->data());
                    std::cerr << "\tMagic bytes: " << header->magic() << std::endl;
                    std::cerr << "\tHeader size: " << header->header_size() << std::endl;
                    std::cerr << "\tNumber of frames: " << header->num_anim() << std::endl;
                }
                break;

                case caff_t::caff_block_t::CAFF_BLOCK_ID_CREDITS: {
                    auto credits = static_cast<caff_t::caff_credits_t*>(block->data());
                    std::cerr << "\tAuthor: " << credits->creator() << std::endl;
                    std::cerr << "\tAuthored date: "
                        << credits->year() << "-"
                        << credits->month() + 0 << "-"
                        << credits->day() + 0
                        << std::endl;
                }
                break;

                case caff_t::caff_block_t::CAFF_BLOCK_ID_ANIMATION: {
                    auto frame = static_cast<caff_t::caff_animation_t*>(block->data());
                    std::cerr << "\tDuration: " << frame->duration() << std::endl;

                    auto image = frame->ciff_data();
                    auto ciff_header = image->header();
                    auto width = ciff_header->fixed_size()->width();
                    auto height = ciff_header->fixed_size()->height();
                    std::cerr << "\tWidth: " << width << std::endl;
                    std::cerr << "\tHeight: " << height << std::endl;
                }
                break;
            }
        }
    }
}
