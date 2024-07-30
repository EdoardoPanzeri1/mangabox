import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';

const Search = () => {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);

  const handleSearch = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.get(`http://localhost:8080/search?q=${query}`);

      if (response.status === 200) {
        setResults(response.data);
      }
    } catch (error) {
      console.error('There was an error with the search!', error);
    }
  };

  return (
    <div style={{ backgroundColor: 'black', color: 'white', minHeight: '100vh', padding: '20px' }}>
      <form onSubmit={handleSearch}>
        <input
          type="text"
          placeholder="Search Manga"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          style={{ backgroundColor: 'black', color: 'white', border: '1px solid white' }}
        />
        <button type="submit">Search</button>
      </form>
      <div>
        {results.length > 0 && (
          <ul>
            {results.map((manga, index) => (
              <li key={index}>
                <h3>{manga.Title}</h3>
                <p>{manga.Author}</p>
                <img src={manga.ImageURL} alt={manga.Title} style={{ maxWidth: '100px'}} />
              </li>
            ))}
          </ul>
        )}
      </div>
      <nav>
        <ul>
          <li><Link to="/mangas" style={{ color: 'white' }}>My Mangas</Link></li>
          <li><Link to="/profile" style={{ color: 'white' }}>Profile</Link></li>
        </ul>
      </nav>
    </div>
  );
};

export default Search;