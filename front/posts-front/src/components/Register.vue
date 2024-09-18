<template>
    <div class="register-form">
        <h2>Register</h2>
        <form @submit.prevent="handleRegisterForm">
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
import { handleRegister } from '../services/authService';

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
            } catch (error) {
                console.error('Eror logging in:', error.message);
            }
        },
    },
};
</script>

<style scoped>
.register-form {
  max-width: 400px;
  margin: auto;
}
</style>