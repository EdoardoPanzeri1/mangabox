import React from 'react';
import { useNavigate } from 'react-router-dom';

const Feedback = ({ message, returnPath }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(returnPath); // Navigate to the path specified by returnPath prop
  };

  return (
    <div style={{ backgroundColor: 'black', color: 'white', minHeight: '100vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <h2>{message}</h2>
      <button onClick={handleClick} style={{ marginTop: '20px', padding: '10px 20px', backgroundColor: 'grey', border: 'none', color: 'white', cursor: 'pointer' }}>
        Back to Previous Page
      </button>
    </div>
  );
};

export default Feedback;