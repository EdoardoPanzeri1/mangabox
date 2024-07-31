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
        const response = await axios.get(`http://localhost:8080/catalog?username=${userID}`);
        if (response.status === 200) {
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

  if (message) {
    return <p>{message}</p>;
  }

  return (
    <div style={{ backgroundColor: 'black', color: 'white', minHeight: '100vh', padding: '20px' }}>
      <h1>My Mangas</h1>
      {mangas.length === 0 ? (
        <p>No mangas found in your catalog.</p>
      ) : (
        <ul>
          {mangas.map(manga => (
            <li key={manga.ID}>
              <h2>{manga.Title}</h2>
              <img src={manga.CoverArtUrl} alt={manga.Title} style={{ maxWidth: '200px' }} />
              <p>Authors: {manga.Authors.join(', ')}</p>
              <p>Status: {manga.Status}</p>
              <p>Issue Number: {manga.IssueNumber}</p>
            </li>
          ))}
        </ul>
      )}
      <Link to="/search" style={{ color: 'white' }}>Back to Search</Link>
    </div>
  );
};

export default Catalog;