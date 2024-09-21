<template>
    <div class="register-form">
        <h2>Register</h2>
        <form @submit.prevent="handleRegisterForm" class="inputs">
            <div>
                <label for="name">Name</label>
                <input type="text" id="name" v-model="name" required />
            </div>
            <div>
                <label for="email">Email</label>
                <input type="email" id="email" v-model="email" required />
            </div>
            <div>
                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" required />
            </div>
            <div>
                <label for="confirmPassword">Confirm Password</label>
                <input type="password" id="confirmPassword" v-model="confirmPassword" required />
            </div>
            <button type="submit">Register</button>
        </form>
    </div>
</template>

<script>
import { handleRegister, handleLogin } from '../services/authService';

export default {
    data() {
        return {
            name:'',
            email:'',
            password:'',
            confirmPassword:'',
        };
    },
    methods: {
        async handleRegisterForm() {

            if (this.password !== this.confirmPassword) {
                alert("Password don't match!")
                return;
            }

            const registerData = {
                name: this.name,
                email: this.email,
                password: this.password,
            };

            try {
                const response = await handleRegister(registerData);
                console.log('Register successful:', response.data);

                await this.loginUser();
                this.$router.push('/create-post');

            } catch (error) {
                console.error('Eror logging in:', error.message);
            }
        },
        async loginUser() {
            const loginData = {
                email: this.email,
                password: this.password,
            };

            try {
                const response = await handleLogin(loginData);
                const token = response.token

                localStorage.setItem('authToken', token)

                console.log('login successful:', response.data);

            } catch (error) {
                console.error('Error loging in', error.message);
            }
        },

    },
};
</script>

<style scoped>
.register-form {
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

input[type="text"],
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