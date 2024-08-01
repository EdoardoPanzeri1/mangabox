import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const Catalog = () => {
  const [mangas, setMangas] = useState([]);
  const [message, setMessage] = useState('');

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
            {manga.coverArtUrl && (
              <img src={manga.coverArtUrl} alt={manga.title} style={{ maxWidth: '200px' }} />
            )}
            <p>{Array.isArray(manga.authors) ? manga.authors.join(', ') : 'Unknown authors'}</p>
            <p>Status: {manga.status}</p>
            <p>Issue: {manga.issueNumber}</p>
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