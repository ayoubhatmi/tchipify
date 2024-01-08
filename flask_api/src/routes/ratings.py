from flask import Blueprint, jsonify, request
from marshmallow import Schema, fields, validates_schema, ValidationError

from src.models.http_exceptions import *
import src.services.ratings as ratings_service

ratings = Blueprint(name="ratings", import_name=__name__)


class RatingSchema(Schema):
    id = fields.Int(description="Rating ID")
    value = fields.Int(description="Rating Value")

# Add any other fields relevant to your ratings model


@ratings.route('/songs/<int:song_id>/ratings', methods=['GET'])
def get_ratings(song_id):
    try:
        ratings_data = ratings_service.get_ratings(song_id)
        return jsonify(RatingSchema(many=True).dump(ratings_data)), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 500


@ratings.route('/songs/<int:song_id>/ratings', methods=['POST'])
def add_rating(song_id):
    try:
        data = request.json
        RatingSchema().load(data)
        new_rating = ratings_service.add_rating(song_id, data)
        return jsonify(RatingSchema().dump(new_rating)), 201
    except ValidationError as e:
        return jsonify({"error": str(e)}), 400
    except Exception as e:
        return jsonify({"error": str(e)}), 500


@ratings.route('/songs/<int:song_id>/ratings/<int:rating_id>', methods=['GET'])
def get_rating(song_id, rating_id):
    try:
        rating_data = ratings_service.get_rating(song_id, rating_id)
        if rating_data:
            return jsonify(RatingSchema().dump(rating_data)), 200
        else:
            return jsonify({"error": "Rating not found"}), 404
    except Exception as e:
        return jsonify({"error": str(e)}), 500


@ratings.route('/songs/<int:song_id>/ratings/<int:rating_id>', methods=['PUT'])
def update_rating(song_id, rating_id):
    try:
        data = request.json
        RatingSchema().load(data)
        updated_rating = ratings_service.update_rating(song_id, rating_id, data)
        return jsonify(RatingSchema().dump(updated_rating)), 200
    except ValidationError as e:
        return jsonify({"error": str(e)}), 400
    except Exception as e:
        return jsonify({"error": str(e)}), 500


@ratings.route('/songs/<int:song_id>/ratings/<int:rating_id>', methods=['DELETE'])
def delete_rating(song_id, rating_id):
    try:
        ratings_service.delete_rating(song_id, rating_id)
        return jsonify({"message": "Rating deleted successfully"}), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 500
