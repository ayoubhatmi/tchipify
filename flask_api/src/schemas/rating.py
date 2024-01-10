from marshmallow import Schema, fields, validates_schema, ValidationError

class RatingSchema(Schema):
    id = fields.String(description="ID")
    user_id = fields.String(description="User ID")
    song_id = fields.String(description="Song ID")
    rating = fields.Integer(description="Rating")
    comment = fields.String(description="Comment")
    rating_date = fields.String(description="Rating Date")

    @validates_schema
    def validate_rating(self, data, **kwargs):
        if 'rating' in data and not (0 <= data['rating'] <= 5):
            raise ValidationError("Rating must be between 0 and 5")

    # You can add additional validation logic as needed
