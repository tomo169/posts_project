import axios from  'axios';

export async function handleLogin(loginData) {
    try {
        const response = await axios.post('http://localhost:8080/login', loginData);
        return response.data;
    } catch (error) {
        throw new Error(error.response ? error.response.data.message : error.message);
    }
}

export async function handleRegister(registerData) {
    try {
      const response = await axios.post('http://localhost:8080/register', registerData);
      return response.data;
    } catch (error) {
      throw new Error(error.response ? error.response.data.message : error.message);
    }
  }