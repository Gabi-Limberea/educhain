import { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import api from '../services/api';
import { useNavigate } from 'react-router-dom';
import { useFormValidation } from '../utility/userFormValidation';
import { validateEmail, validatePassword, validateRepeatPassword } from '../utility/validationUtils';
const rootNode = document.getElementById('root')

function Login() {
	const[showPassword, setShowPassword] = useState(false);
	const[ validFormData, setValidFormData] = useState(false);
	const navigate = useNavigate();
	const[password, setPassword] = useState('');
	const[email, setEmail] = useState('');

	// const ttl = {
	// 	""
	// }


	const { values, errors, handleChange, validateAll } = useFormValidation({
			email: '',
	});
	
	const handleSubmit = async (event: React.FormEvent) => {
		event.preventDefault();

		// console.log("emailul este", email);

		const isValid = validateAll({
			email: validateEmail,
		});

		const config = {
			auth: {
					"username": email,
					"password": password
			}
		};
			
		try {
			const { data } = await api.post<string>('/providers/token', "", config);
			console.log('Full response data:', data); // Debug log to see the exact response structure
			
			// The token is the response data itself
			if (!data) {
				console.error('No token in response');
				throw new Error('No token received from server');
			}

			// Store the token
			localStorage.setItem('authToken', data);
			console.log('Token stored in localStorage:', localStorage.getItem('authToken')); // Verify storage
			
			// Set the token in API headers
			api.defaults.headers.common.Authorization = `Bearer ${data}`;
			console.log('API headers after setting token:', api.defaults.headers.common); // Verify headers
			
			navigate('/registerStudents');
		} catch (error) {
			console.error('Login failed:', error);
			alert("Failed login, please try again!");
		}

		setValidFormData(true);
	};

	return (
		<div className="d-flex justify-content-center py-4">
			<Form className='d-flex flex-column align-items-start w-25' onSubmit={handleSubmit}>
				<Form.Group className="mb-3 w-100" controlId="formBasicEmail">
					<Form.Label>Email address</Form.Label>
					<Form.Control
                        type="email"
                        placeholder="Enter email"
                        value={values.email}
                        onChange={e => {handleChange('email', e.target.value, validateEmail); setEmail(e.target.value)}}
                    />
                    {errors.email && <Form.Text className="text-danger">{errors.email}</Form.Text>}
				</Form.Group>

				<Form.Group className="mb-3 w-100" controlId="formBasicPassword">
					<Form.Label className='d-inline-flex'>Password</Form.Label>
					<Form.Control type={showPassword ? 'text' : 'password'} placeholder="Password" className="w-100"
					onChange={(e) => setPassword(e.target.value)}/>
				</Form.Group>
				<Form.Group className="mb-3 w-100 d-flex align-items-center" controlId="formBasicCheckbox">
					<Form.Check type="checkbox" label="Show password" checked={showPassword} 
								onChange={(e) => setShowPassword(e.target.checked)} />
				</Form.Group>

				<Button variant="secondary w-20" type="submit">
					Login
				</Button>
            </Form>
		</div>
	)
}

export default Login;