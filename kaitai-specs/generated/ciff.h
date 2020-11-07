#ifndef CIFF_H_
#define CIFF_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>
#include <vector>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif

class ciff_t : public kaitai::kstruct {

public:
    class header_t;
    class fixed_header_t;
    class variable_header_t;

    ciff_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = 0, ciff_t* p__root = 0);

private:
    void _read();
    void _clean_up();

public:
    ~ciff_t();

    class header_t : public kaitai::kstruct {

    public:

        header_t(kaitai::kstream* p__io, ciff_t* p__parent = 0, ciff_t* p__root = 0);

    private:
        void _read();
        void _clean_up();

    public:
        ~header_t();

    private:
        fixed_header_t* m_fixed_size;
        variable_header_t* m_variable_size;
        ciff_t* m__root;
        ciff_t* m__parent;

    public:
        fixed_header_t* fixed_size() const { return m_fixed_size; }
        variable_header_t* variable_size() const { return m_variable_size; }
        ciff_t* _root() const { return m__root; }
        ciff_t* _parent() const { return m__parent; }
    };

    class fixed_header_t : public kaitai::kstruct {

    public:

        fixed_header_t(kaitai::kstream* p__io, ciff_t::header_t* p__parent = 0, ciff_t* p__root = 0);

    private:
        void _read();
        void _clean_up();

    public:
        ~fixed_header_t();

    private:
        bool f_expected_content_size;
        int32_t m_expected_content_size;

    public:
        int32_t expected_content_size();

    private:
        std::string m_magic;
        uint64_t m_header_size;
        uint64_t m_content_size;
        uint64_t m_width;
        uint64_t m_height;
        ciff_t* m__root;
        ciff_t::header_t* m__parent;

    public:
        std::string magic() const { return m_magic; }
        uint64_t header_size() const { return m_header_size; }
        uint64_t content_size() const { return m_content_size; }
        uint64_t width() const { return m_width; }
        uint64_t height() const { return m_height; }
        ciff_t* _root() const { return m__root; }
        ciff_t::header_t* _parent() const { return m__parent; }
    };

    class variable_header_t : public kaitai::kstruct {

    public:

        variable_header_t(kaitai::kstream* p__io, ciff_t::header_t* p__parent = 0, ciff_t* p__root = 0);

    private:
        void _read();
        void _clean_up();

    public:
        ~variable_header_t();

    private:
        std::string m_caption;
        std::vector<std::string>* m_tags;
        ciff_t* m__root;
        ciff_t::header_t* m__parent;

    public:
        std::string caption() const { return m_caption; }
        std::vector<std::string>* tags() const { return m_tags; }
        ciff_t* _root() const { return m__root; }
        ciff_t::header_t* _parent() const { return m__parent; }
    };

private:
    bool f_header_size;
    uint64_t m_header_size;

public:
    uint64_t header_size();

private:
    header_t* m_header;
    std::string m_pixel_data;
    ciff_t* m__root;
    kaitai::kstruct* m__parent;
    std::string m__raw_header;
    kaitai::kstream* m__io__raw_header;

public:
    header_t* header() const { return m_header; }
    std::string pixel_data() const { return m_pixel_data; }
    ciff_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
    std::string _raw_header() const { return m__raw_header; }
    kaitai::kstream* _io__raw_header() const { return m__io__raw_header; }
};

#endif  // CIFF_H_
