from partsbuilder import PartsBuilder


builder = PartsBuilder()

builder.set_fm_file_tag()
builder.build()
builder.transfer_to_json()
