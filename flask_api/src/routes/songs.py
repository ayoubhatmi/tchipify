from flask import Blueprint, jsonify, request
from marshmallow import Schema, fields, ValidationError

from src.models.http_exceptions import *
import src.services.songs as songs_service
from src.schemas.song import SongSchema

songs = Blueprint(name="songs", import_name=__name__)

@songs.route('/', methods=['GET'])
def get_songs():
    try:
        songs_data = songs_service.get_songs()
        return jsonify(SongSchema(many=True).dump(songs_data)), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 500


@songs.route('/', methods=['POST'])
def create_song():
    try:
        data = request.json
        SongSchema().load(data)
        new_song = songs_service.create_song(data)
        return jsonify(SongSchema().dump(new_song)), 201
    except ValidationError as e:
        return jsonify({"error": str(e)}), 400
    except Exception as e:
        return jsonify({"error": str(e)}), 500


@songs.route('/<int:song_id>', methods=['GET'])
def get_song(song_id):
    try:
        song_data = songs_service.get_song(song_id)
        if song_data:
            return jsonify(SongSchema().dump(song_data)), 200
        else:
            return jsonify({"error": "Song not found"}), 404
    except Exception as e:
        return jsonify({"error": str(e)}), 500
