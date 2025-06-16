import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import { useNavigate } from 'react-router-dom';
import api from '../services/api';
import { useFormValidation } from '../utility/userFormValidation';
import { 
    validateEmail, 
    validatePassword, 
    validateRepeatPassword, 
    validateRequired, 
    validateName, 
    validateWalletAddress, 
    validatePhone, 
    validateWebsite 
} from '../utility/validationUtils';

type UserType = 'student' | 'provider';

const studentInitialValues = {
    email: '',
    name: '',
    surname: '',
    walletAddress: '',
    password: '',
    repeatPassword: '',
};

const providerInitialValues = {
    institutionName: '',
    email: '',
    // walletAddress: '',
    address: '',
    phoneNumber: '',
    website: '',
    password: '',
    repeatPassword: '',
};

function Signup() {
    const [showPassword, setShowPassword] = useState(false);
    const [userType, setUserType] = useState<UserType>('student');
    const navigate = useNavigate();

    const studentForm = useFormValidation(studentInitialValues);
    const providerForm = useFormValidation(providerInitialValues);

    // Get current form based on user type
    const currentForm = userType === 'student' ? studentForm : providerForm;

    const handleUserTypeChange = (newUserType: UserType) => {
        setUserType(newUserType);
    };

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        let validations: Record<string, (value: string) => string>;

        if (userType === 'student') {
            validations = {
                email: validateEmail,
                name: validateName,
                surname: validateName,
                walletAddress: validateWalletAddress,
                password: validatePassword,
                repeatPassword: (value) => validateRepeatPassword(studentForm.values.password, value),
            };
        } else {
            validations = {
                institutionName: validateRequired,
                email: validateEmail,
                // walletAddress:    validateWalletAddress,
                address: validateRequired,
                phoneNumber: validatePhone,
                website: validateWebsite, // Optional
                password: validatePassword,
                repeatPassword: (value) => validateRepeatPassword(providerForm.values.password, value),
            };
        }

        const isValid = currentForm.validateAll(validations);

        if (isValid) {
            try {
                if (userType === 'student') {
                    // Register the student
                    const studentPayload = {
                        email: studentForm.values.email,
                        name: studentForm.values.name,
                        surname: studentForm.values.surname,
                        walletAddress: studentForm.values.walletAddress,
                        password: studentForm.values.password,
                    };
                    await api.post('/students', studentPayload); // to delete the students signup option later !!!

                } else {
                    // Register the provider
                    const providerPayload = {
                        email:           providerForm.values.email,
                        password:        providerForm.values.password,
                        // contractAddress: providerForm.values.walletAddress,
                        organizationInfo: {
                          name:        providerForm.values.institutionName,
                          address:     providerForm.values.address,
                          phoneNumber: providerForm.values.phoneNumber,
                          website:     providerForm.values.website,
                        }
                    };

                    await api.post('/providers', providerPayload);

                    // Login to get JWT
                    const { data } = await api.post<{ token: string }>('/providers/token', {
                        name: providerForm.values.email,
                        password: providerForm.values.password,
                    });
                    localStorage.setItem('authToken', data.token);
                    api.defaults.headers.common.Authorization = `Bearer ${data.token}`;

                    navigate('/registerStudents');
                }
            } catch (err: any) {
                const serverErrors = err.response?.data?.errors;
                if (serverErrors) {
                    Object.entries(serverErrors).forEach(([field, msg]) => {
                        currentForm.setError(field, msg as string);
                    });
                } else {
                    console.error(err);
                }
            }
        }
    };

    return (
        <div className="d-flex justify-content-center py-4">
            <Form className="d-flex flex-column align-items-start w-25" onSubmit={handleSubmit}>
                {/* User Type Selection */}
                <Form.Group className="mb-3 w-100">
                    <Form.Label>I am a:</Form.Label>
                    <div>
                        <Form.Check
                            type="radio"
                            name="userType"
                            id="student"
                            label="Student"
                            checked={userType === 'student'}
                            onChange={() => handleUserTypeChange('student')}
                        />
                        <Form.Check
                            type="radio"
                            name="userType"
                            id="provider"
                            label="Educational Provider"
                            checked={userType === 'provider'}
                            onChange={() => handleUserTypeChange('provider')}
                        />
                    </div>
                </Form.Group>

                {/* Student Fields */}
                {userType === 'student' && (
                    <>
                        <Form.Group className="mb-3 w-100" controlId="formBasicEmail">
                            <Form.Label>Email address</Form.Label>
                            <Form.Control
                                type="email"
                                placeholder="Enter email"
                                value={studentForm.values.email}
                                onChange={(e) => studentForm.handleChange('email', e.target.value, validateEmail)}
                            />
                            {studentForm.errors.email && <Form.Text className="text-danger">{studentForm.errors.email}</Form.Text>}
                        </Form.Group>

                        <Form.Group className="mb-3 w-100" controlId="formBasicName">
                            <Form.Label>First Name</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="Enter first name"
                                value={studentForm.values.name}
                                onChange={(e) => studentForm.handleChange('name', e.target.value, validateName)}
                            />
                            {studentForm.errors.name && <Form.Text className="text-danger">{studentForm.errors.name}</Form.Text>}
                        </Form.Group>

                        <Form.Group className="mb-3 w-100" controlId="formBasicSurname">
                            <Form.Label>Last Name</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="Enter last name"
                                value={studentForm.values.surname}
                                onChange={(e) => studentForm.handleChange('surname', e.target.value, validateName)}
                            />
                            {studentForm.errors.surname && <Form.Text className="text-danger">{studentForm.errors.surname}</Form.Text>}
                        </Form.Group>

                        <Form.Group className="mb-3 w-100" controlId="formBasicWallet">
                            <Form.Label>Wallet Address</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="0x..."
                                value={studentForm.values.walletAddress}
                                onChange={(e) => studentForm.handleChange('walletAddress', e.target.value, validateWalletAddress)}
                            />
                            {studentForm.errors.walletAddress && <Form.Text className="text-danger">{studentForm.errors.walletAddress}</Form.Text>}
                        </Form.Group>
                    </>
                )}

                {/* Provider Fields */}
                {userType === 'provider' && (
                    <>
                        <Form.Group className="mb-3 w-100" controlId="formBasicInstitution">
                            <Form.Label>Institution Name</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="Enter institution name"
                                value={providerForm.values.institutionName}
                                onChange={(e) => providerForm.handleChange('institutionName', e.target.value, validateRequired)}
                            />
                            {providerForm.errors.institutionName && <Form.Text className="text-danger">{providerForm.errors.institutionName}</Form.Text>}
                        </Form.Group>

                        <Form.Group className="mb-3 w-100" controlId="formBasicEmail">
                            <Form.Label>Email address</Form.Label>
                            <Form.Control
                                type="email"
                                placeholder="Enter email"
                                value={providerForm.values.email}
                                onChange={(e) => providerForm.handleChange('email', e.target.value, validateEmail)}
                            />
                            {providerForm.errors.email && <Form.Text className="text-danger">{providerForm.errors.email}</Form.Text>}
                        </Form.Group>
{/* 
                        <Form.Group className="mb-3 w-100" controlId="providerWallet">
                            <Form.Label>Wallet (Contract) Address</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="0x1234…"
                                value={providerForm.values.walletAddress}
                                isInvalid={!!providerForm.errors.walletAddress}
                                onChange={e => providerForm.handleChange(
                                'walletAddress',
                                e.target.value,
                                validateWalletAddress
                                )}
                            />
                            <Form.Control.Feedback type="invalid">
                                {providerForm.errors.walletAddress}
                            </Form.Control.Feedback>
                        </Form.Group> */}

                        <Form.Group className="mb-3 w-100" controlId="formBasicAddress">
                            <Form.Label>Address</Form.Label>
                            <Form.Control
                                as="textarea"
                                rows={3}
                                placeholder="Enter institution address"
                                value={providerForm.values.address}
                                onChange={(e) => providerForm.handleChange('address', e.target.value, validateRequired)}
                            />
                            {providerForm.errors.address && <Form.Text className="text-danger">{providerForm.errors.address}</Form.Text>}
                        </Form.Group>

                        <Form.Group className="mb-3 w-100" controlId="formBasicPhone">
                            <Form.Label>Phone Number</Form.Label>
                            <Form.Control
                                type="tel"
                                placeholder="Enter phone number"
                                value={providerForm.values.phoneNumber}
                                onChange={(e) => providerForm.handleChange('phoneNumber', e.target.value, validatePhone)}
                            />
                            {providerForm.errors.phoneNumber && <Form.Text className="text-danger">{providerForm.errors.phoneNumber}</Form.Text>}
                        </Form.Group>

                        <Form.Group className="mb-3 w-100" controlId="formBasicWebsite">
                            <Form.Label>Website (Optional)</Form.Label>
                            <Form.Control
                                type="url"
                                placeholder="https://example.com"
                                value={providerForm.values.website}
                                onChange={(e) => providerForm.handleChange('website', e.target.value, validateWebsite)}
                            />
                            {providerForm.errors.website && <Form.Text className="text-danger">{providerForm.errors.website}</Form.Text>}
                        </Form.Group>
                    </>
                )}

                {/* Common Password Fields */}
                <Form.Group className="mb-3 w-100" controlId="formBasicPassword">
                    <Form.Label>Password</Form.Label>
                    <Form.Control
                        type={showPassword ? 'text' : 'password'}
                        placeholder="Password"
                        value={currentForm.values.password}
                        onChange={(e) => currentForm.handleChange('password', e.target.value, validatePassword)}
                    />
                    {currentForm.errors.password && <Form.Text className="text-danger">{currentForm.errors.password}</Form.Text>}
                </Form.Group>

                <Form.Group className="mb-3 w-100" controlId="formBasicRepeatPassword">
                    <Form.Label>Repeat Password</Form.Label>
                    <Form.Control
                        type={showPassword ? 'text' : 'password'}
                        placeholder="Repeat Password"
                        value={currentForm.values.repeatPassword}
                        onChange={(e) => currentForm.handleChange('repeatPassword', e.target.value, (value) => validateRepeatPassword(currentForm.values.password, value))}
                    />
                    {currentForm.errors.repeatPassword && <Form.Text className="text-danger">{currentForm.errors.repeatPassword}</Form.Text>}
                </Form.Group>

                {/* Show Password Checkbox */}
                <Form.Group className="mb-3 w-100" controlId="formBasicCheckbox">
                    <Form.Check
                        type="checkbox"
                        label="Show password"
                        checked={showPassword}
                        onChange={(e) => setShowPassword(e.target.checked)}
                    />
                </Form.Group>

                {/* Submit Button */}
                <Button variant="secondary w-100" type="submit">
                    Sign Up as {userType === 'student' ? 'Student' : 'Provider'}
                </Button>
            </Form>
        </div>
    );
}

export default Signup;
