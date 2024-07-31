import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import axios from 'axios';

const Details = () => {
  const [manga, setManga] = useState(null);
  const [message, setMessage] = useState('');
  const location = useLocation();
  const queryParams = new URLSearchParams(location.search);
  const mangaId = queryParams.get('id');

  // Debugging: Log location and mangaId
  console.log('Location:', location);
  console.log('Manga ID:', mangaId);

  useEffect(() => {
    if (!mangaId) {
      console.error('Manga ID is undefined or null');
      return;
    }

    const fetchMangaDetails = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/details?id=${mangaId}`);
        if (response.status === 200) {
          setManga(response.data);
        } else {
          console.error('Error: Received status', response.status);
        }
      } catch (error) {
        console.error('Error fetching manga details:', error);
      }
    };

    fetchMangaDetails();
  }, [mangaId]);

  const addToDatabase = async () => {
    const userID = localStorage.getItem('user_id'); // Get the user id from localstorage

    if (!userID) {
      console.error('User ID is not available in localStorage');
      setMessage('You must be logged in to add manga to the catalog');
      return;
    }

    try {
      const response = await axios.post('http://localhost:8080/mangas', {
        id: mangaId, 
        title: manga.title,
        authors: JSON.stringify(manga.authors.map(author => author.name)), 
        cover_art_url: manga.images.jpg.image_url,
        synopsis: manga.synopsis,
        status: 'bought', 
        user_id: userID,
        issue_number: manga.issue_number || 0,
        publication_date: manga.publication_date || new Date().toISOString(),
        storyline: manga.storyline || 'placeholder storyline',
        images: JSON.stringify(manga.images),
        serializations: JSON.stringify(manga.serializations || []), 
        genres: JSON.stringify(manga.genres.map(genre => genre.name)), 
        explicit_genres: JSON.stringify(manga.explicit_genres || []), 
        themes: JSON.stringify(manga.themes || []), 
        demographics: JSON.stringify(manga.demographics || []),
        score: manga.score || 0, 
        scored_by: manga.scored_by || 0, 
        rank: manga.rank || 0, 
        popularity: manga.popularity || 0, 
        members: manga.members || 0, 
        favorites: manga.favorites || 0, 
        background: manga.background || 'placeholder background', 
        relations: JSON.stringify(manga.relations || []), 
        external_links: JSON.stringify(manga.external_links || []), 
      });

      if (response.status === 201) {
        setMessage('Manga successfully added to your database!');
      } else {
        setMessage('Failed to add manga to the database.');
      }
    } catch (error) {
      console.error('Error adding manga to the database:', error);
      setMessage('An error occurred while adding the manga.');
    }
  };

  if (!manga) {
    return <p>Loading...</p>;
  }

  return (
    <div style={{ backgroundColor: 'black', color: 'white', minHeight: '100vh', padding: '20px' }}>
      <h1>{manga.title}</h1>
      <img 
        src={manga.images.jpg.image_url} 
        alt={manga.title} 
        style={{ maxWidth: '200px' }} 
      />
      <p>Author: {manga.authors.map(author => author.name).join(', ')}</p>
      <p>Genres: {manga.genres.map(genres => genres.name).join(', ')}</p>
      <p>Chapter: {manga.chapters || 'N/A'}</p>
      <p>Volumes: {manga.volumes || 'N/A'}</p>
      <p>{manga.synopsis}</p>
      <button 
        onClick={addToDatabase} 
        style={{ margin: '10px 0', padding: '10px', backgroundColor: 'green', color: 'white', border: 'none', cursor: 'pointer' }}
      >
        Add to Database
      </button>
      {message && <p>{message}</p>}
    </div>
  );
};

export default Details;