import { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import { useNavigate } from 'react-router-dom';
import { useFormValidation } from '../utility/userFormValidation';
import { validateEmail, validatePassword, validateRepeatPassword } from '../utility/validationUtils';


function Login() {
	const[showPassword, setShowPassword] = useState(false);
	const[ validFormData, setValidFormData] = useState(false);
	const navigate = useNavigate();

	const { values, errors, handleChange, validateAll } = useFormValidation({
			email: '',
		});
	
	const handleSubmit = (event: React.FormEvent) => {
		event.preventDefault();

		const isValid = validateAll({
			email: validateEmail,
		});

		if (isValid) {
			navigate('/registerStudents');
			setValidFormData(true);
		} else {
			setValidFormData(false);
		}
	};

	return (
		<div className="d-flex justify-content-center py-4">
			<Form className='d-flex flex-column align-items-start w-25'>
				<Form.Group className="mb-3 w-100" controlId="formBasicEmail">
					<Form.Label>Email address</Form.Label>
					<Form.Control
                        type="email"
                        placeholder="Enter email"
                        value={values.email}
                        onChange={(e) => handleChange('email', e.target.value, validateEmail)}
                    />
                    {errors.email && <Form.Text className="text-danger">{errors.email}</Form.Text>}
				</Form.Group>

				<Form.Group className="mb-3 w-100" controlId="formBasicPassword">
					<Form.Label className='d-inline-flex'>Password</Form.Label>
					<Form.Control type={showPassword ? 'text' : 'password'} placeholder="Password" className="w-100"/>
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