import { useState } from 'react';

export const useFormValidation = (initialState: Record<string, string>) => {
    const [values, setValues] = useState(initialState);
    const [errors, setErrors] = useState<Record<string, string>>({});

    const handleChange = (field: string, value: string, validateFn?: (value: string) => string) => {
        setValues({ ...values, [field]: value });

        if (validateFn) {
            const error = validateFn(value);
            setErrors({ ...errors, [field]: error });
        }
    };

    const validateField = (field: string, validateFn: (value: string) => string) => {
        const error = validateFn(values[field]);
        setErrors({ ...errors, [field]: error });
    };

    const validateAll = (validations: Record<string, (value: string) => string>) => {
        const newErrors: Record<string, string> = {};
        Object.keys(validations).forEach((field) => {
            const error = validations[field](values[field]);
            if (error) newErrors[field] = error;
        });
        setErrors(newErrors);
        return Object.keys(newErrors).length === 0; // Return true if no errors
    };

    return { values, errors, handleChange, validateField, validateAll };
};