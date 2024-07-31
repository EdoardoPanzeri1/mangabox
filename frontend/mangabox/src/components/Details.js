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
    try {
      const response = await axios.post('http://localhost:8080/mangas', {
        ID: mangaId, 
        Title: manga.title,
        Authors: JSON.stringify(manga.authors.map(author => author.name)), 
        CoverArtUrl: manga.images.jpg.image_url,
        Synopsis: manga.synopsis,
        Type: manga.type,
        Chapters: manga.chapters,
        Volumes: manga.volumes,
        Status: 'bought', // Example status value, adjust as needed
        UserID: 1, // Replace with the actual user ID, you might need to dynamically set this
        IssueNumber: manga.issue_number || 0,
        PublicationDate: manga.publication_date || new Date().toISOString(),
        Storyline: 'some storyline', // Example storyline, adjust as needed
        UpdatedAt: new Date().toISOString(),
        Images: JSON.stringify(manga.images),
        Serializations: JSON.stringify([]), // Adjust as needed
        Genres: JSON.stringify([]), // Adjust as needed
        ExplicitGenres: JSON.stringify([]), // Adjust as needed
        Themes: JSON.stringify([]), // Adjust as needed
        Demographics: JSON.stringify([]), // Adjust as needed
        Score: 0, // Example score value, adjust as needed
        ScoredBy: 0, // Example scored_by value, adjust as needed
        Rank: 0, // Example rank value, adjust as needed
        Popularity: 0, // Example popularity value, adjust as needed
        Members: 0, // Example members value, adjust as needed
        Favorites: 0, // Example favorites value, adjust as needed
        Background: 'some background', // Example background, adjust as needed
        Relations: JSON.stringify([]), // Adjust as needed
        ExternalLinks: JSON.stringify([]), // Adjust as needed
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