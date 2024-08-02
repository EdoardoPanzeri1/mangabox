import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const Catalog = () => {
  const [mangas, setMangas] = useState([]);
  const [message, setMessage] = useState('');
  const userID = localStorage.getItem('user_id');
  const navigate = useNavigate(); // Hook for redirecting

  useEffect(() => {
    const userID = localStorage.getItem('user_id');
    if (!userID) {
      setMessage('You must be logged in to view your catalog');
      return;
    }

    const fetchCatalog = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/mangas?user_id=${userID}`);
        if (response.status === 200) {
          const data = response.data;
          console.log("Fetched data:", data)
          setMangas(response.data);
        } else {
          setMessage('Failed to retrieve catalog');
        }
      } catch (error) {
        console.error('Error fetching catalog:', error);
        setMessage('An error occurred while retrieving the catalog');
      }
    };

    fetchCatalog();
  }, []);

  const deleteManga = async (id) => {
    if (!userID) {
      setMessage('You must be logged in to delete mangas');
      return;
    }

    console.log('Deleting manga with ID:', id); // Debug log

    try {
      const response = await axios.delete(`http://localhost:8080/mangas/${id}`, {
        params: {
          user_id: userID
        }
      });
      if (response.status === 200) {
        setMessage('Manga deleted successfully');
        // Update the manga list after deletion
        setMangas(prevMangas => prevMangas.filter(manga => manga.id !== id));
        // Redirect to the catlog page
        navigate('/mangas');
      } else {
        setMessage('Failed to delete manga');
      }
    } catch (error) {
      console.error('Error deleting manga:', error);
      setMessage('An error occurred while deleting the manga');
    }
  };


  // Debug check if fetching is working
  console.log('Catalog:', mangas);

  if (message) {
    return <p>{message}</p>;
  }

  return (
    <div style={{ backgroundColor: 'black', color: 'white', minHeight: '100vh', padding: '20px' }}>
      <h1>My Mangas</h1>
      {mangas && mangas.length > 0 ? (
      <ul>
        {mangas.map((manga, index) => (
          <li key={index} style={{ marginBottom: '20px' }}>
            <h3>{manga.title}</h3>
            {manga.cover_art_url && (
              <img src={manga.cover_art_url} alt={manga.title} style={{ maxWidth: '200px' }} />
            )}
            <p>{Array.isArray(manga.authors) ? manga.authors.join(', ') : 'Unknown authors'}</p>
            <p>Status: {manga.status}</p>
            <button onClick={() => deleteManga(manga.id)}>Delete</button>
          </li>
        ))}
      </ul>
    ) : (
      <p>No mangas found</p>
    )}
  </div>
);
}

export default Catalog;