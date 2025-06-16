import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../services/api';

function Logout() {
	const navigate = useNavigate();

	useEffect(() => {
		// Clear the token from localStorage
		localStorage.removeItem('authToken');
		
		// Clear the Authorization header
		delete api.defaults.headers.common.Authorization;
		
		// Redirect to home page
		navigate('/home');
	}, [navigate]);

	return null; // This component doesn't render anything
}

export default Logout;