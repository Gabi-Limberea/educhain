import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import { useNavigate } from 'react-router-dom';
import { useFormValidation } from '../utility/userFormValidation';
import { validateEmail, validatePassword, validateRepeatPassword } from '../utility/validationUtils';

function Signup() {
    const [showPassword, setShowPassword] = useState(false);
    const navigate = useNavigate();
    const { values, errors, handleChange, validateAll } = useFormValidation({
        email: '',
        password: '',
        repeatPassword: '',
    });

    const handleSubmit = (event: React.FormEvent) => {
        event.preventDefault();

        const isValid = validateAll({
            email: validateEmail,
            password: validatePassword,
            repeatPassword: (value) => validateRepeatPassword(values.password, value),
        });

        if (isValid) {
            navigate('/registerStudents');
        }
    };

    return (
        <div className="d-flex justify-content-center py-4">
            <Form className="d-flex flex-column align-items-start w-25" onSubmit={handleSubmit}>
                {/* introduce mail */}
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

                {/* introduce password */}
                <Form.Group className="mb-3 w-100" controlId="formBasicPassword">
                    <Form.Label>Password</Form.Label>
                    <Form.Control
                        type={showPassword ? 'text' : 'password'}
                        placeholder="Password"
                        value={values.password}
                        onChange={(e) => handleChange('password', e.target.value, validatePassword)}
                    />
                    {errors.password && <Form.Text className="text-danger">{errors.password}</Form.Text>}
                </Form.Group>

                {/* repeat password */}
                <Form.Group className="mb-3 w-100" controlId="formBasicRepeatPassword">
                    <Form.Label>Repeat Password</Form.Label>
                    <Form.Control
                        type={showPassword ? 'text' : 'password'}
                        placeholder="Repeat Password"
                        value={values.repeatPassword}
                        onChange={(e) =>
                            handleChange('repeatPassword', e.target.value, (value) =>
                                validateRepeatPassword(values.password, value)
                            )
                        }
                    />
                    {errors.repeatPassword && (
                        <Form.Text className="text-danger">{errors.repeatPassword}</Form.Text>
                    )}
                </Form.Group>

                {/* show both passwords */}
                <Form.Group className="mb-3 w-100" controlId="formBasicCheckbox">
                    <Form.Check
                        type="checkbox"
                        label="Show password"
                        checked={showPassword}
                        onChange={(e) => setShowPassword(e.target.checked)}
                    />
                </Form.Group>

                {/* submit button */}
                <Button variant="secondary w-100" type="submit">
                    Signup
                </Button>
            </Form>
        </div>
    );
}

export default Signup;