from marshmallow import Schema, fields, validates_schema, ValidationError

class SongSchema(Schema):
    id = fields.String(description="Song ID")
    artist = fields.String(description="Artist")
    file_name = fields.String(description="File Name")
    published_date = fields.String(description="Published Date")
    title = fields.String(description="Title")