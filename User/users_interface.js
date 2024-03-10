// User Service Frontend Code

import axios from 'axios';

const API_URL = 'http://user-service-api.com'; // Replace with your actual User Service API URL

// Get all users
export const getUsers = async () => {
  try {
    const response = await axios.get(`${API_URL}/users`);
    return response.data;
  } catch (error) {
    console.error('Error getting users:', error);
    return null;
  }
};

// Get a single user by ID
export const getUserById = async (userId) => {
  try {
    const response = await axios.get(`${API_URL}/users/${userId}`);
    return response.data;
  } catch (error) {
    console.error(`Error getting user with ID ${userId}:`, error);
    return null;
  }
};

// Create a new user
export const createUser = async (user) => {
  try {
    const response = await axios.post(`${API_URL}/users`, user);
    return response.data;
  } catch (error) {
    console.error('Error creating user:', error);
    return null;
  }
};

// Update an existing user
export const updateUser = async (userId, user) => {
  try {
    const response = await axios.put(`${API_URL}/users/${userId}`, user);
    return response.data;
  } catch (error) {
    console.error(`Error updating user with ID ${userId}:`, error);
    return null;
  }
};

// Delete a user
export const deleteUser = async (userId) => {
  try {
    const response = await axios.delete(`${API_URL}/users/${userId}`);
    return response.data;
  } catch (error) {
    console.error(`Error deleting user with ID ${userId}:`, error);
    return null;
  }
};