// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "caff.h"
#include "kaitai/exceptions.h"

caff_t::caff_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent, caff_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = this;
    m_block = 0;

    try {
        _read();
    } catch(...) {
        _clean_up();
        throw;
    }
}

void caff_t::_read() {
    m_block = new std::vector<caff_block_t*>();
    {
        int i = 0;
        while (!m__io->is_eof()) {
            m_block->push_back(new caff_block_t(m__io, this, m__root));
            i++;
        }
    }
}

caff_t::~caff_t() {
    _clean_up();
}

void caff_t::_clean_up() {
    if (m_block) {
        for (std::vector<caff_block_t*>::iterator it = m_block->begin(); it != m_block->end(); ++it) {
            delete *it;
        }
        delete m_block; m_block = 0;
    }
}

caff_t::caff_block_t::caff_block_t(kaitai::kstream* p__io, caff_t* p__parent, caff_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;
    m__io__raw_data = 0;

    try {
        _read();
    } catch(...) {
        _clean_up();
        throw;
    }
}

void caff_t::caff_block_t::_read() {
    m_id = static_cast<caff_t::caff_block_t::caff_block_id_t>(m__io->read_u1());
    m_length = m__io->read_u8le();
    n_data = true;
    switch (id()) {
    case caff_t::caff_block_t::CAFF_BLOCK_ID_HEADER: {
        n_data = false;
        m__raw_data = m__io->read_bytes(length());
        m__io__raw_data = new kaitai::kstream(m__raw_data);
        m_data = new caff_header_t(m__io__raw_data, this, m__root);
        break;
    }
    case caff_t::caff_block_t::CAFF_BLOCK_ID_CREDITS: {
        n_data = false;
        m__raw_data = m__io->read_bytes(length());
        m__io__raw_data = new kaitai::kstream(m__raw_data);
        m_data = new caff_credits_t(m__io__raw_data, this, m__root);
        break;
    }
    case caff_t::caff_block_t::CAFF_BLOCK_ID_ANIMATION: {
        n_data = false;
        m__raw_data = m__io->read_bytes(length());
        m__io__raw_data = new kaitai::kstream(m__raw_data);
        m_data = new caff_animation_t(m__io__raw_data, this, m__root);
        break;
    }
    default: {
        m__raw_data = m__io->read_bytes(length());
        break;
    }
    }
}

caff_t::caff_block_t::~caff_block_t() {
    _clean_up();
}

void caff_t::caff_block_t::_clean_up() {
    if (!n_data) {
        if (m__io__raw_data) {
            delete m__io__raw_data; m__io__raw_data = 0;
        }
        if (m_data) {
            delete m_data; m_data = 0;
        }
    }
}

caff_t::caff_header_t::caff_header_t(kaitai::kstream* p__io, caff_t::caff_block_t* p__parent, caff_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;

    try {
        _read();
    } catch(...) {
        _clean_up();
        throw;
    }
}

void caff_t::caff_header_t::_read() {
    m_magic = m__io->read_bytes(4);
    if (!(magic() == std::string("\x43\x41\x46\x46", 4))) {
        throw kaitai::validation_not_equal_error<std::string>(std::string("\x43\x41\x46\x46", 4), magic(), _io(), std::string("/types/caff_header/seq/0"));
    }
    m_header_size = m__io->read_u8le();
    m_num_anim = m__io->read_u8le();
}

caff_t::caff_header_t::~caff_header_t() {
    _clean_up();
}

void caff_t::caff_header_t::_clean_up() {
}

caff_t::caff_credits_t::caff_credits_t(kaitai::kstream* p__io, caff_t::caff_block_t* p__parent, caff_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;

    try {
        _read();
    } catch(...) {
        _clean_up();
        throw;
    }
}

void caff_t::caff_credits_t::_read() {
    m_year = m__io->read_u2le();
    m_month = m__io->read_u1();
    m_day = m__io->read_u1();
    m_hour = m__io->read_u1();
    m_minute = m__io->read_u1();
    m_creator_len = m__io->read_u8le();
    m_creator = kaitai::kstream::bytes_to_str(m__io->read_bytes(creator_len()), std::string("ASCII"));
}

caff_t::caff_credits_t::~caff_credits_t() {
    _clean_up();
}

void caff_t::caff_credits_t::_clean_up() {
}

caff_t::caff_animation_t::caff_animation_t(kaitai::kstream* p__io, caff_t::caff_block_t* p__parent, caff_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;
    m_ciff_data = 0;
    m__io__raw_ciff_data = 0;

    try {
        _read();
    } catch(...) {
        _clean_up();
        throw;
    }
}

void caff_t::caff_animation_t::_read() {
    m_duration = m__io->read_u8le();
    m__raw_ciff_data = m__io->read_bytes_full();
    m__io__raw_ciff_data = new kaitai::kstream(m__raw_ciff_data);
    m_ciff_data = new ciff_t(m__io__raw_ciff_data);
}

caff_t::caff_animation_t::~caff_animation_t() {
    _clean_up();
}

void caff_t::caff_animation_t::_clean_up() {
    if (m__io__raw_ciff_data) {
        delete m__io__raw_ciff_data; m__io__raw_ciff_data = 0;
    }
    if (m_ciff_data) {
        delete m_ciff_data; m_ciff_data = 0;
    }
}
