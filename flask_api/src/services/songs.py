import json
import requests

songs_url = "http://localhost:8080/songs/"  # URL de l'API songs (golang)*

def get_songs():
    response = requests.get(songs_url)
    response.raise_for_status()  # Raise an HTTPError for bad responses
    return response.json()

def create_song(data):
    response = requests.post(songs_url, json=data)
    response.raise_for_status()
    return response.json()

def get_song(song_id):
    response = requests.get(f"{songs_url}{song_id}")
    response.raise_for_status()
    return response.json()
