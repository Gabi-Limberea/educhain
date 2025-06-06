export const validateEmail = (value: string): string => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(value) ? '' : 'Invalid email address';
};

export const validatePassword = (value: string): string => {
    return value.length >= 6 ? '' : 'Password must be at least 6 characters long';
};

export const validateRepeatPassword = (password: string, repeatPassword: string): string => {
    return password === repeatPassword ? '' : 'Passwords do not match';
};