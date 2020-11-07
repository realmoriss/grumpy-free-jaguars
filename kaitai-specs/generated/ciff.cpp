// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "ciff.h"
#include "kaitai/exceptions.h"

ciff_t::ciff_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent, ciff_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = this;
    m_header = 0;
    m__io__raw_header = 0;
    f_header_size = false;

    try {
        _read();
    } catch(...) {
        _clean_up();
        throw;
    }
}

void ciff_t::_read() {
    m__raw_header = m__io->read_bytes(header_size());
    m__io__raw_header = new kaitai::kstream(m__raw_header);
    m_header = new header_t(m__io__raw_header, this, m__root);
    m_pixel_data = m__io->read_bytes(header()->fixed_size()->content_size());
}

ciff_t::~ciff_t() {
    _clean_up();
}

void ciff_t::_clean_up() {
    if (m__io__raw_header) {
        delete m__io__raw_header; m__io__raw_header = 0;
    }
    if (m_header) {
        delete m_header; m_header = 0;
    }
    if (f_header_size) {
    }
}

ciff_t::header_t::header_t(kaitai::kstream* p__io, ciff_t* p__parent, ciff_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;
    m_fixed_size = 0;
    m_variable_size = 0;

    try {
        _read();
    } catch(...) {
        _clean_up();
        throw;
    }
}

void ciff_t::header_t::_read() {
    m_fixed_size = new fixed_header_t(m__io, this, m__root);
    m_variable_size = new variable_header_t(m__io, this, m__root);
}

ciff_t::header_t::~header_t() {
    _clean_up();
}

void ciff_t::header_t::_clean_up() {
    if (m_fixed_size) {
        delete m_fixed_size; m_fixed_size = 0;
    }
    if (m_variable_size) {
        delete m_variable_size; m_variable_size = 0;
    }
}

ciff_t::fixed_header_t::fixed_header_t(kaitai::kstream* p__io, ciff_t::header_t* p__parent, ciff_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;
    f_expected_content_size = false;

    try {
        _read();
    } catch(...) {
        _clean_up();
        throw;
    }
}

void ciff_t::fixed_header_t::_read() {
    m_magic = m__io->read_bytes(4);
    if (!(magic() == std::string("\x43\x49\x46\x46", 4))) {
        throw kaitai::validation_not_equal_error<std::string>(std::string("\x43\x49\x46\x46", 4), magic(), _io(), std::string("/types/fixed_header/seq/0"));
    }
    m_header_size = m__io->read_u8le();
    m_content_size = m__io->read_u8le();
    m_width = m__io->read_u8le();
    m_height = m__io->read_u8le();
}

ciff_t::fixed_header_t::~fixed_header_t() {
    _clean_up();
}

void ciff_t::fixed_header_t::_clean_up() {
}

int32_t ciff_t::fixed_header_t::expected_content_size() {
    if (f_expected_content_size)
        return m_expected_content_size;
    m_expected_content_size = ((width() * height()) * 3);
    f_expected_content_size = true;
    return m_expected_content_size;
}

ciff_t::variable_header_t::variable_header_t(kaitai::kstream* p__io, ciff_t::header_t* p__parent, ciff_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;
    m_tags = 0;

    try {
        _read();
    } catch(...) {
        _clean_up();
        throw;
    }
}

void ciff_t::variable_header_t::_read() {
    m_caption = kaitai::kstream::bytes_to_str(m__io->read_bytes_term(10, false, true, true), std::string("ASCII"));
    m_tags = new std::vector<std::string>();
    {
        int i = 0;
        while (!m__io->is_eof()) {
            m_tags->push_back(kaitai::kstream::bytes_to_str(m__io->read_bytes_term(0, false, true, true), std::string("ASCII")));
            i++;
        }
    }
}

ciff_t::variable_header_t::~variable_header_t() {
    _clean_up();
}

void ciff_t::variable_header_t::_clean_up() {
    if (m_tags) {
        delete m_tags; m_tags = 0;
    }
}

uint64_t ciff_t::header_size() {
    if (f_header_size)
        return m_header_size;
    std::streampos _pos = m__io->pos();
    m__io->seek(4);
    m_header_size = m__io->read_u8le();
    m__io->seek(_pos);
    f_header_size = true;
    return m_header_size;
}
