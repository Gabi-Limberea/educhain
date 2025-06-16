import { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import api from '../services/api';
import { useNavigate } from 'react-router-dom';
import { useFormValidation } from '../utility/userFormValidation';
import { validateName, validateWalletAddress, validateEmail, validatePassword, validateRepeatPassword } from '../utility/validationUtils';
const rootNode = document.getElementById('root')

function Login() {
	const [userType, setUserType] = useState<'student' | 'provider'>('provider');
	const [name, setName] = useState('');
	const [walletId, setWalletId] = useState('');
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const [showPassword, setShowPassword] = useState(false);
	const [validFormData, setValidFormData] = useState(false);
	const navigate = useNavigate();

	const { values, errors, handleChange, validateAll } = useFormValidation({
		name: '',
		walletId: '',
		email: '',
	});

	const handleSubmit = async (event: React.FormEvent) => {
		event.preventDefault();

		if (userType === 'student') {
			const isValid = validateAll({
				name: validateName,
				walletId: validateWalletAddress,
			});

			if (isValid) {
				// Store student info in localStorage and redirect
				localStorage.setItem('userType', 'student');
				localStorage.setItem('studentName', name);
				localStorage.setItem('studentWalletId', walletId);
				navigate('/account');
			}
		} else {
			const isValid = validateAll({
				email: validateEmail,
			});

			if (isValid) {
				const config = {
					auth: {
						"username": email,
						"password": password
					}
				};

				try {
					const { data } = await api.post<string>('/providers/token', "", config);
					if (!data) {
						throw new Error('No token received from server');
					}

					localStorage.setItem('authToken', data);
					localStorage.setItem('userType', 'provider');
					api.defaults.headers.common.Authorization = `Bearer ${data}`;
					navigate('/registerStudents');
				} catch (error) {
					console.error('Login failed:', error);
					alert("Failed login, please try again!");
				}
			}
		}

		setValidFormData(true);
	};

	return (
		<div className="d-flex justify-content-center py-4">
			<Form className='d-flex flex-column align-items-start w-25' onSubmit={handleSubmit}>
				<Form.Group className="mb-3 w-100">
					<Form.Label>I am a:</Form.Label>
					<div>
						<Form.Check
							type="radio"
							name="userType"
							id="student"
							label="Student"
							checked={userType === 'student'}
							onChange={() => setUserType('student')}
						/>
						<Form.Check
							type="radio"
							name="userType"
							id="provider"
							label="Educational Provider"
							checked={userType === 'provider'}
							onChange={() => setUserType('provider')}
						/>
					</div>
				</Form.Group>

				{userType === 'student' ? (
					<>
						<Form.Group className="mb-3 w-100" controlId="formBasicName">
							<Form.Label>Name</Form.Label>
							<Form.Control
								type="text"
								placeholder="Enter your name"
								value={values.name}
								onChange={e => {handleChange('name', e.target.value, validateName); setName(e.target.value)}}
							/>
							{errors.name && <Form.Text className="text-danger">{errors.name}</Form.Text>}
						</Form.Group>

						<Form.Group className="mb-3 w-100" controlId="formBasicWalletId">
							<Form.Label>Wallet Address</Form.Label>
							<Form.Control
								type="text"
								placeholder="Enter your wallet address"
								value={values.walletId}
								onChange={e => {handleChange('walletId', e.target.value, validateWalletAddress); setWalletId(e.target.value)}}
							/>
							{errors.walletId && <Form.Text className="text-danger">{errors.walletId}</Form.Text>}
						</Form.Group>
					</>
				) : (
					<>
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
							<Form.Label>Password</Form.Label>
							<Form.Control
								type="password"
								placeholder="Password"
								onChange={(e) => setPassword(e.target.value)}
							/>
						</Form.Group>
					</>
				)}

				<Button variant="secondary w-20" type="submit">
					Login
				</Button>
			</Form>
		</div>
	);
}

export default Login;