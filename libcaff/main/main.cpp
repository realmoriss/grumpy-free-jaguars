#include <iostream>
#include <main/main.h>

int main(int argc, char *argv[]) {
    kaitai::kstream ks(&std::cin);

    try {
        auto parsed = caff_t(&ks);
        return print_preview(parsed);
    } catch (...) {
        return 1;
    }
}

int print_preview(caff_t& parsed) {
    auto blocks = parsed.block();

    auto found_credits = false;
    ciff_t *image = nullptr;

    if (blocks) {
        for (auto block: *blocks) {
            auto id = block->id();
            switch (id) {
                case caff_t::caff_block_t::CAFF_BLOCK_ID_HEADER: {
                    auto header = static_cast<caff_t::caff_header_t*>(block->data());
                    if (header->num_anim() < 1) {
                        return 1;
                    }
                }
                break;

                case caff_t::caff_block_t::CAFF_BLOCK_ID_CREDITS: {
                    //auto credits = static_cast<caff_t::caff_credits_t*>(block->data());
                    found_credits = true;
                }
                break;

                case caff_t::caff_block_t::CAFF_BLOCK_ID_ANIMATION: {
                    if (!image) {
                        auto frame = static_cast<caff_t::caff_animation_t*>(block->data());
                        image = frame->ciff_data();
                    }
                }
                break;
            }
        }
    }

    if (found_credits && image) {
        auto ciff_header = image->header();
        auto width = ciff_header->fixed_size()->width();
        auto height = ciff_header->fixed_size()->height();
        std::cout << width << " " << height << std::endl;
        std::cout << image->pixel_data();
        return 0;
    }

    return 1;
}