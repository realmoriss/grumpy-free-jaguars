#ifndef CAFF_H_
#define CAFF_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>
#include "ciff.h"
#include <vector>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif
class ciff_t;

class caff_t : public kaitai::kstruct {

public:
    class caff_block_t;
    class caff_header_t;
    class caff_credits_t;
    class caff_animation_t;

    caff_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = 0, caff_t* p__root = 0);

private:
    void _read();
    void _clean_up();

public:
    ~caff_t();

    class caff_block_t : public kaitai::kstruct {

    public:

        enum caff_block_id_t {
            CAFF_BLOCK_ID_HEADER = 1,
            CAFF_BLOCK_ID_CREDITS = 2,
            CAFF_BLOCK_ID_ANIMATION = 3
        };

        caff_block_t(kaitai::kstream* p__io, caff_t* p__parent = 0, caff_t* p__root = 0);

    private:
        void _read();
        void _clean_up();

    public:
        ~caff_block_t();

    private:
        caff_block_id_t m_id;
        uint64_t m_length;
        kaitai::kstruct* m_data;
        bool n_data;

    public:
        bool _is_null_data() { data(); return n_data; };

    private:
        caff_t* m__root;
        caff_t* m__parent;
        std::string m__raw_data;
        kaitai::kstream* m__io__raw_data;

    public:

        /**
         * Type of the CAFF block
         */
        caff_block_id_t id() const { return m_id; }

        /**
         * Length of the CAFF block data
         */
        uint64_t length() const { return m_length; }

        /**
         * CAFF block data
         */
        kaitai::kstruct* data() const { return m_data; }
        caff_t* _root() const { return m__root; }
        caff_t* _parent() const { return m__parent; }
        std::string _raw_data() const { return m__raw_data; }
        kaitai::kstream* _io__raw_data() const { return m__io__raw_data; }
    };

    class caff_header_t : public kaitai::kstruct {

    public:

        caff_header_t(kaitai::kstream* p__io, caff_t::caff_block_t* p__parent = 0, caff_t* p__root = 0);

    private:
        void _read();
        void _clean_up();

    public:
        ~caff_header_t();

    private:
        std::string m_magic;
        uint64_t m_header_size;
        uint64_t m_num_anim;
        caff_t* m__root;
        caff_t::caff_block_t* m__parent;

    public:
        std::string magic() const { return m_magic; }

        /**
         * Size of the header (all fields included)
         */
        uint64_t header_size() const { return m_header_size; }

        /**
         * Number of CIFF animation blocks
         */
        uint64_t num_anim() const { return m_num_anim; }
        caff_t* _root() const { return m__root; }
        caff_t::caff_block_t* _parent() const { return m__parent; }
    };

    /**
     * CAFF credits block which specifies the CAFF creation date, creation time and author
     */

    class caff_credits_t : public kaitai::kstruct {

    public:

        caff_credits_t(kaitai::kstream* p__io, caff_t::caff_block_t* p__parent = 0, caff_t* p__root = 0);

    private:
        void _read();
        void _clean_up();

    public:
        ~caff_credits_t();

    private:
        uint16_t m_year;
        uint8_t m_month;
        uint8_t m_day;
        uint8_t m_hour;
        uint8_t m_minute;
        uint64_t m_creator_len;
        std::string m_creator;
        caff_t* m__root;
        caff_t::caff_block_t* m__parent;

    public:
        uint16_t year() const { return m_year; }
        uint8_t month() const { return m_month; }
        uint8_t day() const { return m_day; }
        uint8_t hour() const { return m_hour; }
        uint8_t minute() const { return m_minute; }
        uint64_t creator_len() const { return m_creator_len; }

        /**
         * Creator of the CAFF file
         */
        std::string creator() const { return m_creator; }
        caff_t* _root() const { return m__root; }
        caff_t::caff_block_t* _parent() const { return m__parent; }
    };

    /**
     * CAFF animation block which contains a CIFF image to be animated
     */

    class caff_animation_t : public kaitai::kstruct {

    public:

        caff_animation_t(kaitai::kstream* p__io, caff_t::caff_block_t* p__parent = 0, caff_t* p__root = 0);

    private:
        void _read();
        void _clean_up();

    public:
        ~caff_animation_t();

    private:
        uint64_t m_duration;
        ciff_t* m_ciff_data;
        caff_t* m__root;
        caff_t::caff_block_t* m__parent;
        std::string m__raw_ciff_data;
        kaitai::kstream* m__io__raw_ciff_data;

    public:

        /**
         * The duration in miliseconds for which the CIFF image must be displayed during animation
         */
        uint64_t duration() const { return m_duration; }
        ciff_t* ciff_data() const { return m_ciff_data; }
        caff_t* _root() const { return m__root; }
        caff_t::caff_block_t* _parent() const { return m__parent; }
        std::string _raw_ciff_data() const { return m__raw_ciff_data; }
        kaitai::kstream* _io__raw_ciff_data() const { return m__io__raw_ciff_data; }
    };

private:
    std::vector<caff_block_t*>* m_block;
    caff_t* m__root;
    kaitai::kstruct* m__parent;

public:
    std::vector<caff_block_t*>* block() const { return m_block; }
    caff_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};

#endif  // CAFF_H_
