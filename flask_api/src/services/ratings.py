import json
import requests

ratings_url = "https://ratings-mike.edu.forestier.re"

def get_ratings(song_id):
    response = requests.get(f"{ratings_url}/songs/{song_id}/ratings")
    response.raise_for_status()
    return response.json()

def add_rating(song_id, data):
    response = requests.post(f"{ratings_url}/songs/{song_id}/ratings", json=data)
    response.raise_for_status()
    return response.json()

def get_rating(song_id, rating_id):
    response = requests.get(f"{ratings_url}/songs/{song_id}/ratings/{rating_id}")
    response.raise_for_status()
    return response.json()

def update_rating(song_id, rating_id, data):
    response = requests.put(f"{ratings_url}/songs/{song_id}/ratings/{rating_id}", json=data)
    response.raise_for_status()
    return response.json()

def delete_rating(song_id, rating_id):
    response = requests.delete(f"{ratings_url}/songs/{song_id}/ratings/{rating_id}")
    response.raise_for_status()
    return response.json()
