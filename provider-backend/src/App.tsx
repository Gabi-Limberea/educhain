import React from 'react';
import './App.css';
import NavbarComponent from './components/Navbar';
import FooterComponent from './components/Footer';
import axios from 'axios';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import UploadDiplomas from './pages/UploadDiplomas';
import RegisterStudents from './pages/RegisterStudents';
import Login from './pages/Login';
import Logout from './pages/Logout';
import Signup from './pages/Signup';
import Home from './pages/Home';
import Account from './pages/Account';
import RegisterCertification from './pages/RegisterCertification';

// for testing only
const apiCall = () => {
	axios.get('http://localhost:8080').then((data) => {
	  console.log(data)
	})
}



const App: React.FC = () => {
	return (
		<BrowserRouter>
			<div className="d-flex flex-column min-vh-100"> 
				<div> <NavbarComponent/> </div>

				{/* for testing send request */}
				{/* <button onClick={apiCall}>Make API Call</button> */}

				<div className="flex-grow-1"> 
					<Routes>
						<Route path="/" element={<Navigate to="/home"/>} />

						<Route path="/uploadDiplomas" element={<UploadDiplomas />} />
						<Route path="/registerStudents" element={<RegisterStudents />} />
						<Route path="/login" element={<Login />} />
						<Route path="/signup" element={<Signup />} />
						<Route path="/logout" element={<Logout />} />
						<Route path="/home" element={<Home />}/>
						<Route path="/account" element={<Account/>}/>
						<Route path="/registerCertification" element={<RegisterCertification/>}/>

					</Routes>
				</div>
				<FooterComponent/>
			</div>
		</BrowserRouter>
	);
};
export default App;
