import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../services/api';

function Logout() {
	const navigate = useNavigate();

	useEffect(() => {
		const userType = localStorage.getItem('userType');

		if (userType === 'provider') {
			// Clear the token from localStorage
			localStorage.removeItem('authToken');
			
			// Clear the Authorization header
			delete api.defaults.headers.common.Authorization;
		}

		// Clear user type and user-specific data
		localStorage.removeItem('userType');
		localStorage.removeItem('studentName');
		localStorage.removeItem('studentWalletId');
		
		// Redirect to home page
		navigate('/home');
	}, [navigate]);

	return null; // This component doesn't render anything
}

export default Logout;