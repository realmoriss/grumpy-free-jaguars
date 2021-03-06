meta:
  id: caff
  file-extension: caff
  endian: le
  imports:
    - ciff

seq:
- id: block
  type: caff_block
  repeat: eos

types:
  caff_block:
    seq:
    - id: id
      type: u1
      enum: caff_block_id
      doc: 'Type of the CAFF block'
    - id: length
      type: u8
      doc: 'Length of the CAFF block data'
    - id: data
      size: length
      type:
        switch-on: id
        cases:
          'caff_block_id::header': caff_header
          'caff_block_id::credits': caff_credits
          'caff_block_id::animation': caff_animation
      doc: 'CAFF block data'
    enums:
      caff_block_id:
        1: header
        2: credits
        3: animation

  caff_header:
    seq:
    - id: magic
      contents: 'CAFF'
    - id: header_size
      doc: 'Size of the header (all fields included)'
      type: u8
    - id: num_anim
      doc: 'Number of CIFF animation blocks'
      type: u8

  caff_credits:
    seq:
    - id: year
      type: u2
    - id: month
      type: u1
    - id: day
      type: u1
    - id: hour
      type: u1
    - id: minute
      type: u1
    - id: creator_len
      type: u8
    - id: creator
      type: str
      size: creator_len
      encoding: ASCII
      doc: 'Creator of the CAFF file'
    doc: 'CAFF credits block which specifies the CAFF creation date, creation time and author'

  caff_animation:
    seq:
    - id: duration
      type: u8
      doc: 'The duration in miliseconds for which the CIFF image must be displayed during animation'
    - id: ciff_data
      type: ciff
      size-eos: true
    doc: 'CAFF animation block which contains a CIFF image to be animated'