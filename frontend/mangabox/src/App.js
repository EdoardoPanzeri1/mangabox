import React, { useState, useEffect } from 'react';

function App() {
  const [mangas, setMangas] = useState([]);

  useEffect(() => {
    // Fetch manga data from the backend
    fetch('http://localhost:8080/mangas')
      .then(response => response.json())
      .then(data => setMangas(data))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  return (
    <div className="App">
      <h1>Manga List</h1>
      <ul>
        {mangas.map(manga => (
          <li key={manga.id}>{manga.title}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;