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
                if (!token) {
                    throw new Error("No token found");
                }

                const response = await axios.get('http://localhost:8080/profile', {
                    headers: { Authorization: `Bearer ${token}` },
                });

                const profileData = response.data;

                // Set user object and update fields conditionally to handle different casing
                setUser(profileData);
                setEmail(profileData.email || profileData.Email || '');
                console.log("Profile fetched successfully", response.data);
            } catch (error) {
                if (error.response?.status === 404) {
                    navigate('/create-profile');
                } else if (error.response?.status === 401) {
                    navigate('/login');
                } else {
                    console.error('There was an error fetching the profile!', error);
                }
            }
        };

        fetchProfile();
    }, [navigate]);
  const handleUpdate = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem('token');
      const updateData = { email };
      if (password) {
        updateData.password = password;
      }

      const response = await axios.put(
        'http://localhost:8080/profile',
        updateData,
        { 
          headers: { Authorization: `Bearer ${token}`}
         }
      );
      
      if (response.status === 200) {
        const updateProfile = response.data;
        setUser(updateProfile);
        setEmail(updateProfile.email || updateProfile.Email);
        setPassword('');
        setIsEditing(false);
      }
    } catch (error) {
      console.error('There was an error updating the profile!', error);
    }
  };

  return (
    <div style={{ backgroundColor: 'black', color: 'white', minHeight: '100vh', padding: '20px' }}>
      <h2>Profile</h2>
      {!isEditing ? (
        <div>
          <p>Username: {user.Username}</p>
          <p>Email: {user.Email}</p>
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