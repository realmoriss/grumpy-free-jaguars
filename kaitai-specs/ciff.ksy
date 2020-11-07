meta:
  id: ciff
  file-extension: ciff
  endian: le

seq:
  - id: header
    type: header
    size: header_size
  - id: pixel_data
    size: header.fixed_size.content_size


instances:
  header_size:
    pos: 0x4
    type: u8

types:
  header:
    seq:
      - id: fixed_size
        type: fixed_header
      - id: variable_size
        type: variable_header

  fixed_header:
    seq:
      - id: magic
        contents: 'CIFF'
      - id: header_size
        type: u8
      - id: content_size
        type: u8
      - id: width
        type: u8
      - id: height
        type: u8
    instances:
      expected_content_size:
        value: width * height * 3

  variable_header:
    seq:
      - id: caption
        type: str
        encoding: 'ASCII'
        terminator: 10
      - id: tags
        type: str
        encoding: 'ASCII'
        terminator: 0
        repeat: eos
