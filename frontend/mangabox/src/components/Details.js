import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import axios from 'axios';

const Details = () => {
  const [manga, setManga] = useState(null);
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
      <p>{manga.synopsis}</p>
      {/* Add more manga details here as needed */}
    </div>
  );
};

export default Details;