import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';

const Search = () => {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);

  const handleSearch = async (e) => {
    e.preventDefault();

    try {
      // Make request to go server  
      const response = await axios.get(`http://localhost:8080/search?q=${query}`);

      if (response.status === 200) {
        console.log(response.data) // Log the response data for debugging
        setResults(response.data);
      }
    } catch (error) {
      console.error('There was an error with the search!', error);
    }
  };

  return (
    <div style={{ backgroundColor: 'black', color: 'white', minHeight: '100vh', padding: '20px' }}>
      <div style={{ display: 'flex', alignItems: 'center', marginBottom: '20px' }}>
        <form onSubmit={handleSearch} style={{ display: 'flex', alignItems: 'center', marginRight: '20px' }}>
          <input
            type="text"
            placeholder="Search Manga"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            style={{ backgroundColor: 'black', color: 'white', border: '1px solid white', marginRight: '10px' }}
          />
          <button type="submit">Search</button>
        </form>
        <nav>
          <ul style={{ display: 'flex', listStyleType: 'none', padding: '0', margin: '0' }}>
            <li><Link to="/mangas" style={{ color: 'white', textDecoration: 'none', marginRight: '10px' }}>My Mangas</Link></li>
            <li><Link to="/profile" style={{ color: 'white', textDecoration: 'none' }}>Profile</Link></li>
          </ul>
        </nav>
      </div>
      <div>
        {results.length > 0 ? (
          <ul>
            {results.map((manga, index) => (
              <li key={index} style={{ marginBottom: '20px' }}>
                <Link to={`/details?id=${manga.id}`} style={{ color: 'white', textDecoration: 'none' }}>
                    <h3>{manga.title}</h3>
                </Link>
                <Link to={`/details?id=${manga.id}`} style={{ textDecoration: 'none' }}>
                  <img src={manga.image_url} alt={manga.title} style={{ maxWidth: '100px' }} />
                </Link>
                <p>{manga.author}</p>
              </li>
            ))}
          </ul>
        ) : (
          <p>No results found</p>
        )}
      </div>
    </div>
  );
};

export default Search;