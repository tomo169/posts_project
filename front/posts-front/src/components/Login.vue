<template>
    <div class="login-form">
        <h2>Login</h2>
        <form @submit.prevent="hendleLoginForm">
            <div>
                <label for="email">Email</label>
                <input type="email" id="email" v-model="email" required />
            </div>
            <div>
                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" required />
            </div>
            <button type="submit">Login</button>
        </form>
    </div>
</template>

<script>
import { handleLogin } from '../services/authService';

export default {
    data() {
        return {
            email:'',
            password:'',
        };
    },
    methods: {
       async hendleLoginForm() {
            const loginData = {
                email: this.email,
                password: this.password,
            };

            try {
                const response = await handleLogin(loginData);
                const token = response.token

                localStorage.setItem('authToken', token)

                console.log('login successful:', response.data);

                this.$router.push('/posts');
            } catch (error) {
                console.error('Error loging in', error.message);
            }
        },
    },
};
</script>

<style scoped>
.login-form {
  max-width: 400px;
  margin: auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

form {
  background-color: #333;
  padding: 30px;
  border-radius: 8px;
  border: 1px solid #6fbff5;
  width: 100%;
  max-width: 400px;
}

label {
  font-weight: bold;
  display: block;
  margin-bottom: 8px;
  margin-top: 8px;
  color: white;
}

input[type="password"],
input[type="email"] {
  width: 100%;
  padding: 12px;
  border: 1px solid #f56ff5;
  background-color: #242424;
  border-radius: 4px;
  font-size: 16px;
  box-sizing: border-box;
}

button {
    margin-top: 8px;
}
</style>