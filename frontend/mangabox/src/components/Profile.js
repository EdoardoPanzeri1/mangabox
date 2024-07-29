import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const Profile = () => {
  const [user, setUser] = useState({});
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isEditing, setIsEditing] = useState(false);
  const navigate = useNavigate();
  
  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/profile', {
          headers: { Authorization: `Bearer ${token}` },
        });
        setUser(response.data);
        setEmail(response.data.email);
      } catch (error) {
        console.error('There was an error fetching the profile!', error);
        navigate('/login');
      }
    };
    
    fetchProfile();
  }, [navigate]);

  const handleUpdate = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem('token');
      const response = await axios.put(
        'http://localhost:8080/profile',
        { email, password },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );
      
      if (response.status === 200) {
        setUser(response.data);
        setIsEditing(false);
      }
    } catch (error) {
      console.error('There was an error updating the profile!', error);
    }
  };

  return (
    <div>
      <h2>Profile</h2>
      {!isEditing ? (
        <div>
          <p>Username: {user.username}</p>
          <p>Email: {user.email}</p>
          <button onClick={() => setIsEditing(true)}>Edit Profile</button>
        </div>
      ) : (
        <form onSubmit={handleUpdate}>
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <input
            type="password"
            placeholder="Password (leave blank to keep current)"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <button type="submit">Update Profile</button>
          <button type="button" onClick={() => setIsEditing(false)}>
            Cancel
          </button>
        </form>
      )}
    </div>
  );
};

export default Profile;