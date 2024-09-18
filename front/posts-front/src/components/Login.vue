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

                this.$router.push('/create-post');
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
}
</style>