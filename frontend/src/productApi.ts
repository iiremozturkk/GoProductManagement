import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080',
    headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
    },
    timeout: 10000
});

// Request interceptor to add Authorization header
api.interceptors.request.use((config) => {
    console.log('API Request:', config.method?.toUpperCase(), config.url);
    const isLoggedIn = localStorage.getItem('isLoggedIn') === 'true';
    if (isLoggedIn) {
        config.headers['Authorization'] = 'Bearer dummy-token';
    }
    return config;
}, error => {
    console.error('Request Error:', error);
    return Promise.reject(error);
});

// Response interceptor to handle errors
api.interceptors.response.use(
    response => {
        console.log('API Response:', response.status, response.data);
        return response;
    },
    error => {
        if (error.response) {
            // Server yanıt verdi ama hata kodu döndü
            console.error('API Error:', {
                status: error.response.status,
                data: error.response.data,
                headers: error.response.headers
            });
        } else if (error.request) {
            // İstek yapıldı ama yanıt alınamadı
            console.error('API Error: No response received', error.request);
        } else {
            // İstek oluşturulurken bir hata oluştu
            console.error('API Error:', error.message);
        }
        return Promise.reject(error);
    }
);

interface LoginResponse {
    success: boolean;
    message?: string;
}

interface RegisterResponse {
    success: boolean;
    message?: string;
}

interface Product {
    id: number;
    name: string;
    description: string;
    price: number;
    stock: number;
    imageURL?: string;
}

interface User {
    name: string;
    email: string;
    password: string;
}

export async function login(email: string, password: string): Promise<LoginResponse> {
    try {
        console.log('Login attempt for:', email);
        const response = await api.post<LoginResponse>('/login', { email, password });
        if (response.data.success) {
            localStorage.setItem('isLoggedIn', 'true');
            console.log('Login successful');
        }
        return response.data;
    } catch (error) {
        console.error('Login error:', error);
        throw error;
    }
}

export async function register(user: User): Promise<RegisterResponse> {
    try {
        console.log('Registration attempt for:', user.email);
        const response = await api.post<RegisterResponse>('/register', user);
        console.log('Registration response:', response.data);
        return response.data;
    } catch (error) {
        console.error('Register error:', error);
        throw error;
    }
}

export async function getProducts(): Promise<Product[]> {
    try {
        const response = await api.get<Product[]>('/products');
        return response.data;
    } catch (error) {
        console.error('Get products error:', error);
        throw error;
    }
}

export async function getProductById(id: number): Promise<Product> {
    try {
        const response = await api.get<Product>(`/products/${id}`);
        return response.data;
    } catch (error) {
        console.error('Get product error:', error);
        throw error;
    }
}

export async function createProduct(product: Omit<Product, 'id'>): Promise<void> {
    try {
        await api.post('/products', product);
    } catch (error) {
        console.error('Create product error:', error);
        throw error;
    }
}

export async function updateProduct(id: number, product: Partial<Product>): Promise<void> {
    try {
        await api.put(`/products/${id}`, product);
    } catch (error) {
        console.error('Update product error:', error);
        throw error;
    }
}

export async function deleteProduct(id: number): Promise<void> {
    try {
        await api.delete(`/products/${id}`);
    } catch (error) {
        console.error('Delete product error:', error);
        throw error;
    }
}

export async function logout(): Promise<void> {
    localStorage.removeItem('isLoggedIn');
    console.log('Logged out');
}

export { api }; 
