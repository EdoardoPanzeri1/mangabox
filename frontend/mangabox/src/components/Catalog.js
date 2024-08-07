import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';


const Catalog = () => {
  const [mangas, setMangas] = useState([]);
  const [message, setMessage] = useState('');
  const [feedbackMessage, setFeedbackMessage] = useState('');
  const userID = localStorage.getItem('user_id');
  const navigate = useNavigate();

  const fetchCatalog = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/mangas?user_id=${userID}`);
      if (response.status === 200) {
        const data = response.data;
        console.log("Fetched data:", data);
        setMangas(data);
      } else {
        setMessage('Failed to retrieve catalog');
      }
    } catch (error) {
      console.error('Error fetching catalog:', error);
      setMessage('An error occurred while retrieving the catalog');
    }
  };

  useEffect(() => {
    if (!userID) {
      setMessage('You must be logged in to view your catalog');
    } else {
      fetchCatalog();
    }
  }, [userID, fetchCatalog]);

  const deleteManga = async (id) => {
    if (!userID) {
      setMessage('You must be logged in to delete mangas');
      return;
    }
  
    try {
      const response = await axios.delete(`http://localhost:8080/mangas/${id}`, {
        params: {
          user_id: userID,
        },
      });
  
      if (response.status === 200) {
        setFeedbackMessage('Manga deleted successfully');
        fetchCatalog(); // Re-fetch the catalog to update the list of mangas
      } else {
        setFeedbackMessage('Failed to delete the manga');
      }
    } catch (error) {
      console.error('Error deleting manga:', error);
      setFeedbackMessage('An error occurred while deleting the manga');
    }
  };

  const updateMangaStatus = async (id, newStatus) => {
    if (!userID) {
      setMessage('You must be logged in to update manga status');
      return;
    }
  
    try {
      const payload = { status: newStatus, user_id: userID };
      console.log(`Updating manga ID: ${id} to status: ${newStatus}`);
      console.log('Payload being sent:', payload); // Debug log
  
      const response = await axios.put(
        `http://localhost:8080/mangas/${id}`,
        payload,
        {
          headers: {
            'Content-Type': 'application/json'
          }
        }
      );
  
      if (response.status === 200) {
        setFeedbackMessage('Manga status updated successfully');
        // Update the local state to reflect the status change
        setMangas(prevMangas =>
          prevMangas.map(manga => manga.id === id ? { ...manga, status: newStatus } : manga)
        );
      } else {
        setFeedbackMessage('Failed to update manga status');
      }
    } catch (error) {
      console.error('Error updating manga status:', error);
      setFeedbackMessage('An error occurred while updating the manga status');
    }
  };

  // Debug check if fetching is working
  console.log('Catalog:', mangas);


  return (
    <div style={{ backgroundColor: 'black', color: 'white', minHeight: '100vh', padding: '20px' }}>
    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
      <h1>My Mangas</h1>
      <button 
        onClick={() => navigate('/search')}
        style={{ backgroundColor: 'white', color: 'black', padding: '10px', borderRadius: '5px' }}
      >
        Back to Search
      </button>
    </div>
    {message && <p>{message}</p>}
    {feedbackMessage && <p>{feedbackMessage}</p>}
      
      <div style={{ display: 'flex', justifyContent: 'space-between' }}>
        <div style={{ flex: 1, margin: '10px' }}>
          <h2>Bought</h2>
          {mangas && mangas.filter(manga => manga.status === 'bought').length > 0 ? (
            <ul>
              {mangas.filter(manga => manga.status === 'bought').map((manga, index) => (
                <li key={index} style={{ marginBottom: '20px' }}>
                  <h3>{manga.title}</h3>
                  {manga.cover_art_url && (
                    <img src={manga.cover_art_url} alt={manga.title} style={{ maxWidth: '200px' }} />
                  )}
                  <p>{Array.isArray(manga.authors) ? manga.authors.join(', ') : 'Unknown authors'}</p>
                  <p>Status: {manga.status}</p>
                  <button onClick={() => deleteManga(manga.id)}>
                    Delete
                  </button>
                  <button onClick={() => updateMangaStatus(manga.id, 'read')}>
                    Mark as Read
                  </button>
                </li>
              ))}
            </ul>
          ) : (
            <p>No mangas found</p>
          )}
        </div>

        <div style={{ flex: 1, margin: '10px' }}>
        <h2>Read</h2>
        {mangas && mangas.filter(manga => manga.status === 'read').length > 0 ? (
          <ul>
            {mangas.filter(manga => manga.status === 'read').map((manga, index) => (
              <li key={index} style={{ marginBottom: '20px' }}>
                <h3>{manga.title}</h3>
                {manga.cover_art_url && (
                  <img src={manga.cover_art_url} alt={manga.title} style={{ maxWidth: '200px' }} />
                )}
                <p>{Array.isArray(manga.authors) ? manga.authors.join(', ') : 'Unknown authors'}</p>
                <p>Status: {manga.status}</p>
                <button onClick={() => deleteManga(manga.id)}>
                  Delete
                </button>
                <button onClick={() => updateMangaStatus(manga.id, 'bought')}>
                  Mark as Bought
                </button>
              </li>
            ))}
          </ul>
        ) : (
          <p>No mangas found</p>
        )}
      </div>
    </div>
  </div>
);
};

export default Catalog;